package nba_pbp

import (
	"fmt"
)

type PeriodData struct {
	GameId       string     `json:"gid"`
	PeriodNumber int64      `json:"p"`
	Events       []PbpEvent `json:"pla"`
}

func (period *PeriodData) GetStarters(game_details GameDetails) (map[int64][]int64, error) {
	starters := map[int64][]int64{}
	subbed_in_players := map[int64][]int64{}
	for _, event := range period.Events {
		player_id := event.PlayerId
		team_id := event.TeamId
		if team_id != 0 && player_id != 0 {
			if _, exists := starters[team_id]; !exists {
				starters[team_id] = []int64{}
				subbed_in_players[team_id] = []int64{}
			}
			secondary_player_id, err := event.GetEPlayerIdInt()
			if err != nil {
				return starters, err
			}
			if event.IsSubstitution() {
				// player_id is player going out, secondary_player_id is playing coming in
				if !IntInSlice(secondary_player_id, starters[team_id]) && !IntInSlice(secondary_player_id, subbed_in_players[team_id]) {
					// track players who have been subbed in
					subbed_in_players[team_id] = append(subbed_in_players[team_id], secondary_player_id)
				}
				if !IntInSlice(player_id, starters[team_id]) && !IntInSlice(player_id, subbed_in_players[team_id]) {
					// if player is subbed out without having subbed in this period, they started period
					starters[team_id] = append(starters[team_id], player_id)
				}
			} else {
				player_in_starters := IntInSlice(player_id, starters[team_id])
				player_in_subs := IntInSlice(player_id, subbed_in_players[team_id])
				// ignore technicals because player can get technical when not in game
				is_technical_or_ejection := event.IsEjection() || event.IsTechnicalFoul() || event.IsDoubleTechnical()
				// check that player_id played > 0 seconds, avoids including team id and other unknown ids
				player_played, _ := DidPlayerPlayInGame(player_id, game_details)
				if !player_in_starters && !player_in_subs && !is_technical_or_ejection && player_played {
					starters[team_id] = append(starters[team_id], player_id)
				}
				secondary_player_played, secondary_player_team_id := DidPlayerPlayInGame(secondary_player_id, game_details)
				secondary_player_in_starters := IntInSlice(secondary_player_id, starters[team_id])
				secondary_player_in_subs := IntInSlice(secondary_player_id, subbed_in_players[team_id])
				if !secondary_player_in_starters && !secondary_player_in_subs && !is_technical_or_ejection && secondary_player_played {
					starters[secondary_player_team_id] = append(starters[secondary_player_team_id], secondary_player_id)
				}
			}
		}
	}
	// make sure each team has 5 starters, if not need to fix pbp
	for team_id, team_starters := range starters {
		if len(team_starters) != 5 {
			return starters, fmt.Errorf("TeamId: %s has %d starters for period %d. Should have 5", team_id, len(team_starters), period.PeriodNumber)
		}
	}
	return starters, nil
}

func (period *PeriodData) AddCurrentPlayersAndTimeElapsed(starters map[int64][]int64) error {
	// adds current players to each pbp event
	current_players := starters
	for i, event := range period.Events {
		if event.IsSubstitution() {
			coming_in_player_id, err := event.GetEPlayerIdInt()
			if err != nil {
				return err
			}
			going_out_player_id := event.PlayerId
			team_id := event.TeamId
			updated_lineup := make([]int64, 5)
			for i, player_id := range current_players[team_id] {
				if player_id == going_out_player_id {
					updated_lineup[i] = coming_in_player_id
				} else {
					updated_lineup[i] = player_id
				}
			}
			current_players[team_id] = updated_lineup
		}
		period.Events[i].CurrentPlayers = make(map[int64][]int64)
		for team_id, lineup := range current_players {
			period.Events[i].CurrentPlayers[team_id] = lineup
		}

		// add time
		seconds_remaining, err := event.GetSecondsRemaining()
		if err != nil {
			return err
		}
		if i < len(period.Events)-1 {
			next_event_seconds_remaining, err := period.Events[i+1].GetSecondsRemaining()
			if err != nil {
				return err
			}
			period.Events[i].SecondsToNextEvent = seconds_remaining - next_event_seconds_remaining
		}
		if i > 0 {
			previous_event_seconds_remaining, err := period.Events[i-1].GetSecondsRemaining()
			if err != nil {
				return err
			}
			period.Events[i].SecondsSinceLastEvent = previous_event_seconds_remaining - seconds_remaining
		}
	}
	return nil
}

func (period *PeriodData) ParsePossessions(game_details GameDetails) ([]PossessionDetails, error) {
	period_possessions := []PossessionDetails{}
	possession_number := int64(1)
	var team_starting_possession_with_ball int64
	var possession_start_seconds_remaining float64
	var previous_possession_end_event_num int64
	var score_differential int64 // from perspective of team on offense
	for i, event := range period.Events {
		seconds_remaining, err := event.GetSecondsRemaining()
		if err != nil {
			return []PossessionDetails{}, err
		}
		if i == 0 {
			possession_start_seconds_remaining = seconds_remaining
			previous_possession_end_event_num = event.EventNum
			score_differential = event.HomeScore - event.VisitorScore
		}
		// determine if possession has changed
		possession_changing_event := false
		// defensive rebound
		if event.IsRebound() {
			rebounded_shot, err := event.GetReboundedShot(period.Events)
			if err != nil {
				return []PossessionDetails{}, err
			}
			if rebounded_shot.EventNum > 0 {
				def_reb := event.IsDefensiveRebound(rebounded_shot)
				team_reb := event.IsTeamRebound()
				if def_reb && !(seconds_remaining == 0 && team_reb) {
					// don't include team rebounds with 0 seconds left - they appear in pbp following missed shots at the buzzer
					possession_changing_event = true
					team_starting_possession_with_ball = SwapTeamIdForGame(event.TeamId, game_details)
				}
			}
		}
		// turnovers
		if event.IsTurnover() {
			possession_changing_event = true
			team_starting_possession_with_ball = event.TeamId
		}
		// made FGs
		if event.IsMadeFG() {
			and1, err := event.IsAnd1Shot(period.Events)
			if err != nil {
				return []PossessionDetails{}, err
			}
			if !and1 {
				// team that started possession with ball is team that made the shot
				possession_changing_event = true
				team_starting_possession_with_ball = event.TeamId
			}
		}
		// made final FTA
		if event.IsMadeFT() && (event.IsFt1of1() || event.IsFt2of2() || event.IsFt3of3()) {
			// Ignore FT 1 of 1 on away from play fouls and inbound fouls
			away_from_play, err := event.IsFtFromAwayFromPlayFoul(period.Events)
			if err != nil {
				return []PossessionDetails{}, err
			}
			inbound_foul, err := event.IsFtFromInboundFoul(period.Events)
			if err != nil {
				return []PossessionDetails{}, err
			}
			if !away_from_play && !inbound_foul {
				possession_changing_event = true
				team_starting_possession_with_ball = event.TeamId
			}
		}
		// change of possession on jump ball
		if event.IsJumpBall() && possession_number != 1 {
			if team_starting_possession_with_ball != event.OfTeamId {
				if !(period.Events[i+1].IsRebound() || period.Events[i+1].IsJumpBall()) {
					// ignore jump ball if next event is a rebound or jump ball since that will trigger possession change
					possession_changing_event = true
				}
				if next_event_seconds_remaining, err := period.Events[i+1].GetSecondsRemaining(); err == nil {
					if period.Events[i+1].IsTurnover() && next_event_seconds_remaining == seconds_remaining {
						// if next event is steal at same time of pbp don't need to change possession since steal takes care of it
						possession_changing_event = false
					}
				} else {
					return []PossessionDetails{}, err
				}
				if previous_event_seconds_remaining, err := period.Events[i+1].GetSecondsRemaining(); err == nil {
					if period.Events[i-1].IsTurnover() && previous_event_seconds_remaining == seconds_remaining {
						// if previous event is steal at same time of pbp don't need to change possession since steal takes care of it
						possession_changing_event = false
					}
				} else {
					return []PossessionDetails{}, err
				}
			}
		} else {
			team_starting_possession_with_ball = event.OfTeamId
		}

		if i == len(period.Events)-1 {
			possession_changing_event = true
		}

		if possession_changing_event {
			if team_starting_possession_with_ball != game_details.HomeTeamData.TeamId {
				// score_differential is from home team perspective, change to from offensive team perspective
				score_differential = -1 * score_differential
			}
			lineup_ids := event.GenerateLineupIds()
			defense_team_id := SwapTeamIdForGame(team_starting_possession_with_ball, game_details)
			possession_details := PossessionDetails{
				Period:                           period.PeriodNumber,
				PossessionNumber:                 possession_number,
				OffenseTeamId:                    team_starting_possession_with_ball,
				DefenseTeamId:                    defense_team_id,
				PossessionStartTime:              possession_start_seconds_remaining,
				PossessionEndTime:                seconds_remaining,
				PreviousPossessionEndEventNum:    previous_possession_end_event_num,
				PossessionEndEventNum:            event.EventNum,
				PossessionStartScoreDifferential: score_differential,
				OffenseLineupId:                  lineup_ids[team_starting_possession_with_ball],
				DefenseLineupId:                  lineup_ids[defense_team_id],
			}
			period_possessions = append(period_possessions, possession_details)
			possession_number += 1
			possession_start_seconds_remaining = seconds_remaining
			previous_possession_end_event_num = event.EventNum
			score_differential = event.HomeScore - event.VisitorScore
		}
	}
	return period_possessions, nil
}
