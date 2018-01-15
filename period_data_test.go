package nba_pbp

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func helperLoadPeriodData(t *testing.T) PbpDataResponse {
	bytes, err := ioutil.ReadFile("testdata/period_data.json")
	if err != nil {
		t.Fatal(err)
	}

	var period_data PbpDataResponse

	json.Unmarshal(bytes, &period_data)
	return period_data
}

func TestEvents(t *testing.T) {
	data := helperLoadPeriodData(t)

	expected_period_number := int64(1)
	if data.Data.PeriodNumber != expected_period_number {
		t.Errorf("Wrong Period Number. Expected %d. Got %d.", expected_period_number, data.Data.PeriodNumber)
	}

	expected_events_length := 129
	if len(data.Data.Events) != expected_events_length {
		t.Errorf("Incorrect events length. Expected %d. Got %d.", expected_events_length, len(data.Data.Events))
	}
}

func TestParsingPbp(t *testing.T) {
	period_data := helperLoadPeriodData(t)
	game_details := helperLoadGameDetails(t)
	period_events := period_data.Data
	// get starters
	starters, err := period_events.GetStarters(game_details)
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	expected_starters := map[int64][]int64{
		1610612744: []int64{2585, 201939, 2738, 202691, 203110},
		1610612746: []int64{201599, 203710, 1628393, 201933, 1626155},
	}
	match := reflect.DeepEqual(starters, expected_starters)
	if !match {
		t.Errorf("Expected %#v, got %#v", expected_starters, starters)
	}

	// add lineups
	err = period_events.AddCurrentPlayersAndTimeElapsed(starters)
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	for _, event := range period_events.Events {
		if event.IsSubstitution() {
			coming_in_player_id, err := event.GetEPlayerIdInt()
			if err != nil {
				t.Fatal(err)
			}
			going_out_player_id := event.PlayerId
			team_id := event.TeamId
			if IntInSlice(going_out_player_id, event.CurrentPlayers[team_id]) {
				t.Errorf("PlayerId %d. In Event %s", going_out_player_id, event.Description)
			}
			if !IntInSlice(coming_in_player_id, event.CurrentPlayers[team_id]) {
				t.Errorf("PlayerId %d. Not In Event %s", coming_in_player_id, event.Description)
			}
		}
	}
	expected_seconds_to_next_event := float64(14)
	if period_events.Events[1].SecondsToNextEvent != expected_seconds_to_next_event {
		t.Errorf("Expected %d. Got %d", expected_seconds_to_next_event, period_events.Events[1].SecondsToNextEvent)
	}
	expected_seconds_since_previous_event := float64(6)
	if period_events.Events[1].SecondsSinceLastEvent != expected_seconds_since_previous_event {
		t.Errorf("Expected %d. Got %d", expected_seconds_since_previous_event, period_events.Events[1].SecondsSinceLastEvent)
	}

	// parse possessions
	possession_details, err := period_events.ParsePossessions(game_details)
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	expected_number_of_possessions := 51
	if len(possession_details) != expected_number_of_possessions {
		t.Errorf("PlayerId %d. In Event %d", expected_number_of_possessions, len(possession_details))
	}

	// getting events for possession
	possession_events, err := possession_details[0].GetAllEventsForPossession(period_events.Events)
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	expected_number_of_events := 5
	if len(possession_events) != expected_number_of_events {
		t.Errorf("Expected %d. Got %d", expected_number_of_events, len(possession_events))
	}

	for i, possession := range possession_details {
		if i != 0 {
			if possession.OffenseTeamId != possession_details[i-1].DefenseTeamId {
				t.Errorf("Back to Back Possessions. TeamId %d. Possession Number %s", possession.OffenseTeamId, i+1)
			}
			possession_details[i].AddPossessionStats(period_events.Events, possession_details[i-1])
		} else {
			possession_details[i].AddPossessionStats(period_events.Events, PossessionDetails{})
		}
		possession_details[i].AddPlayerStatsForPossession()
	}

	// check possesion stats
	event_index := 0
	expected_possession_start_type := OFF_DEADBALL_STRING
	if possession_details[event_index].PossessionStartType != expected_possession_start_type {
		t.Errorf("Expected %s. Got %s.", expected_possession_start_type, possession_details[event_index].PossessionStartType)
	}

	// team stats
	team_id := int64(1610612744)
	lineup_id := "201939-202691-203110-2585-2738"
	opponent_lineup_id := "1626155-1628393-201599-201933-203710"

	player_id := int64(2585)
	stat_key := "AtRimMisses"
	expected_stat_value := int64(1)
	stat_value := possession_details[event_index].PlayerStats[team_id][lineup_id][opponent_lineup_id][player_id][stat_key]
	if stat_value != expected_stat_value {
		t.Errorf("Expected %d. Got %d. For %d %s on possession number %d", expected_stat_value, stat_value, player_id, stat_key, possession_details[event_index].PossessionNumber)
	}

	player_id = int64(0)
	stat_key = "OffAtRimRebound"
	expected_stat_value = int64(1)
	stat_value = possession_details[event_index].PlayerStats[team_id][lineup_id][opponent_lineup_id][player_id][stat_key]
	if stat_value != expected_stat_value {
		t.Errorf("Expected %d. Got %d. For %d %s on possession number %d", expected_stat_value, stat_value, player_id, stat_key, possession_details[event_index].PossessionNumber)
	}

	player_id = int64(2738)
	stat_key = "AtRimMisses"
	expected_stat_value = int64(1)
	stat_value = possession_details[event_index].PlayerStats[team_id][lineup_id][opponent_lineup_id][player_id][stat_key]
	if stat_value != expected_stat_value {
		t.Errorf("Expected %d. Got %d. For %d %s on possession number %d", expected_stat_value, stat_value, player_id, stat_key, possession_details[event_index].PossessionNumber)
	}

	// opponent stats
	player_id = int64(201599)
	team_id = int64(1610612746)
	opponent_lineup_id = "201939-202691-203110-2585-2738"
	lineup_id = "1626155-1628393-201599-201933-203710"

	stat_key = "DefAtRimRebound"
	player_id = int64(201599)
	expected_stat_value = int64(1)
	stat_value = possession_details[event_index].PlayerStats[team_id][lineup_id][opponent_lineup_id][player_id][stat_key]
	if stat_value != expected_stat_value {
		t.Errorf("Expected %d. Got %d. For %d %s on possession number %d", expected_stat_value, stat_value, player_id, stat_key, possession_details[event_index].PossessionNumber)
	}

	// new possession
	event_index = 1
	expected_possession_start_type = OFF_AT_RIM_MISS_STRING
	if possession_details[event_index].PossessionStartType != expected_possession_start_type {
		t.Errorf("Expected %s. Got %s.", expected_possession_start_type, possession_details[event_index].PossessionStartType)
	}
	expected_previous_possession_shooter := int64(2585)
	if possession_details[event_index].PreviousPossessionEndShooterPlayerId != expected_previous_possession_shooter {
		t.Errorf("Expected %d. Got %d.", expected_previous_possession_shooter, possession_details[event_index].PreviousPossessionEndShooterPlayerId)
	}

	expected_previous_possession_rebounder := int64(201599)
	if possession_details[event_index].PreviousPossessionEndReboundPlayerId != expected_previous_possession_rebounder {
		t.Errorf("Expected %d. Got %d.", expected_previous_possession_rebounder, possession_details[event_index].PreviousPossessionEndReboundPlayerId)
	}

	expected_number_of_events = 2
	if len(possession_details[event_index].Events) != expected_number_of_events {
		t.Errorf("Expected %d. Got %d.", expected_number_of_events, len(possession_details[event_index].Events))
	}

	// new possession
	event_index = 14
	expected_possession_start_type = OFF_LIVE_BALL_TURNOVER_STRING
	if possession_details[event_index].PossessionStartType != expected_possession_start_type {
		t.Errorf("Expected %s. Got %s.", expected_possession_start_type, possession_details[event_index].PossessionStartType)
	}
	expected_previous_possession_turnover := int64(201933)
	if possession_details[event_index].PreviousPossessionEndTurnoverPlayerId != expected_previous_possession_turnover {
		t.Errorf("Expected %d. Got %d.", expected_previous_possession_turnover, possession_details[event_index].PreviousPossessionEndTurnoverPlayerId)
	}

	expected_previous_possession_steal := int64(203110)
	if possession_details[event_index].PreviousPossessionEndStealPlayerId != expected_previous_possession_steal {
		t.Errorf("Expected %d. Got %d.", expected_previous_possession_steal, possession_details[event_index].PreviousPossessionEndStealPlayerId)
	}
	// offensive rebounds
	event_index = 7
	expected_offensive_rebounds := int64(1)
	if possession_details[event_index].OffensiveRebounds != expected_offensive_rebounds {
		t.Errorf("Expected %d. Got %d.", expected_offensive_rebounds, possession_details[event_index].OffensiveRebounds)
	}

	expected_second_chance_time := float64(7)
	if possession_details[event_index].SecondChanceTime != expected_second_chance_time {
		t.Errorf("Expected %f. Got %f.", expected_second_chance_time, possession_details[event_index].SecondChanceTime)
	}

	// score differential
	event_index = 3 // first possession with a score - should be 0
	expected_score_diff := int64(0)
	if possession_details[event_index].PossessionStartScoreDifferential != expected_score_diff {
		t.Errorf("Expected %d. Got %d.", expected_score_diff, possession_details[event_index].PossessionStartScoreDifferential)
	}

	event_index = 4 // possession following first possession with a score
	expected_score_diff = int64(-2)
	if possession_details[event_index].PossessionStartScoreDifferential != expected_score_diff {
		t.Errorf("Expected %d. Got %d.", expected_score_diff, possession_details[event_index].PossessionStartScoreDifferential)
	}

	// sum up stats
	team_stats, opponent_stats, player_stats, lineup_stats, lineup_opponent_stats, err := SumPossessionStats(possession_details, int64(1610612744))
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}

	// check rebound chances mirror each other
	for key, value := range team_stats {
		if strings.HasSuffix(key, "Chances") {
			if strings.HasPrefix(key, "Off") {
				if opponent_stats[strings.Replace(key, "Off", "Def", 1)] != value {
					t.Errorf("Expected %d. Got %d.", value, opponent_stats[strings.Replace(key, "Off", "Def", 1)])
				}
			} else if strings.HasPrefix(key, "Def") {
				if opponent_stats[strings.Replace(key, "Def", "Off", 1)] != value {
					t.Errorf("Expected %d. Got %d.", value, opponent_stats[strings.Replace(key, "Def", "Off", 1)])
				}
			}
		}
	}

	// team stats
	expected_team_stats := map[string]int64{
		"2ptAnd1s":                             int64(1),
		"2ptShootingFoulFreeThrowTrips":        int64(1),
		"3ptShootingFoulFreeThrowTrips":        int64(1),
		"Arc3Assists":                          int64(2),
		"Arc3Makes":                            int64(4),
		"Arc3Misses":                           int64(3),
		"AssistedArc3":                         int64(2),
		"AssistedAtRim":                        int64(2),
		"AssistedCorner3":                      int64(1),
		"AssistedLongMidRange":                 int64(1),
		"AtRimAssists":                         int64(2),
		"AtRimBlocks":                          int64(2),
		"AtRimMakes":                           int64(3),
		"AtRimMisses":                          int64(6),
		"BlockedAtRim":                         int64(3),
		"BlockedLongMidRange":                  int64(1),
		"Corner3Assists":                       int64(1),
		"Corner3Makes":                         int64(1),
		"Corner3Misses":                        int64(4),
		"DeadballTurnovers":                    int64(1),
		"DefArc3Rebound":                       int64(2),
		"DefArc3ReboundChances":                int64(3),
		"DefAtRimRebound":                      int64(3),
		"DefAtRimReboundChances":               int64(3),
		"DefBlockedAtRimRebound":               int64(1),
		"DefBlockedAtRimReboundChances":        int64(2),
		"DefensivePossessions":                 int64(25),
		"DefFTRebound":                         int64(1),
		"DefFTReboundChances":                  int64(1),
		"DefLongMidRangeRebound":               int64(1),
		"DefLongMidRangeReboundChances":        int64(1),
		"DefShortMidRangeRebound":              int64(2),
		"DefShortMidRangeReboundChances":       int64(2),
		"FTMakes":                              int64(8),
		"LiveballTurnovers":                    int64(1),
		"LongMidRangeAssists":                  int64(1),
		"LongMidRangeMakes":                    int64(1),
		"LongMidRangeMisses":                   int64(3),
		"OffArc3Rebound":                       int64(1),
		"OffArc3ReboundChances":                int64(2),
		"OffAtRimRebound":                      int64(2),
		"OffAtRimReboundChances":               int64(3),
		"OffBlockedAtRimReboundChances":        int64(3),
		"OffBlockedLongMidRangeRebound":        int64(1),
		"OffBlockedLongMidRangeReboundChances": int64(1),
		"OffCorner3Rebound":                    int64(1),
		"OffCorner3ReboundChances":             int64(4),
		"OffensivePossessions":                 int64(26),
		"OffLongMidRangeRebound":               int64(1),
		"OffLongMidRangeReboundChances":        int64(2),
		"OffShortMidRangeReboundChances":       int64(2),
		"PenaltyFreeThrowTrips":                int64(1),
		"PersonalFouls":                        int64(2),
		"PersonalFoulsDrawn":                   int64(2),
		"Seconds":                              int64(719),
		"ShootingBlockFouls":                   int64(1),
		"ShootingFouls":                        int64(2),
		"ShootingFoulsDrawn":                   int64(3),
		"ShortMidRangeMisses":                  int64(2),
		"Steals":                               int64(3),
	}
	match = reflect.DeepEqual(team_stats, expected_team_stats)
	if !match {
		for key, value := range expected_team_stats{
			if value != team_stats[key]{
				t.Errorf("Expected %d, got %d for %s", value, team_stats[key], key)
			}
		}
	}
	// opponent stats
	expected_opponent_stats := map[string]int64{
		"2ptShootingFoulFreeThrowTrips":        int64(3),
		"Arc3Makes":                            int64(2),
		"Arc3Misses":                           int64(3),
		"AssistedLongMidRange":                 int64(3),
		"AtRimBlocks":                          int64(3),
		"AtRimMakes":                           int64(4),
		"AtRimMisses":                          int64(5),
		"BlockedAtRim":                         int64(2),
		"DefArc3Rebound":                       int64(1),
		"DefArc3ReboundChances":                int64(2),
		"DefAtRimRebound":                      int64(1),
		"DefAtRimReboundChances":               int64(3),
		"DefBlockedAtRimRebound":               int64(3),
		"DefBlockedAtRimReboundChances":        int64(3),
		"DefBlockedLongMidRangeReboundChances": int64(1),
		"DefCorner3Rebound":                    int64(3),
		"DefCorner3ReboundChances":             int64(4),
		"DefensivePossessions":                 int64(26),
		"DefLongMidRangeRebound":               int64(1),
		"DefLongMidRangeReboundChances":        int64(2),
		"DefShortMidRangeRebound":              int64(2),
		"DefShortMidRangeReboundChances":       int64(2),
		"FTMakes":                              int64(6),
		"FTMisses":                             int64(1),
		"LiveballTurnovers":                    int64(3),
		"LongMidRangeAssists":                  int64(3),
		"LongMidRangeBlocks":                   int64(1),
		"LongMidRangeMakes":                    int64(3),
		"LongMidRangeMisses":                   int64(1),
		"OffArc3Rebound":                       int64(1),
		"OffArc3ReboundChances":                int64(3),
		"OffAtRimReboundChances":               int64(3),
		"OffBlockedAtRimRebound":               int64(1),
		"OffBlockedAtRimReboundChances":        int64(2),
		"OffensivePossessions":                 int64(25),
		"OffFTReboundChances":                  int64(1),
		"OffLongMidRangeReboundChances":        int64(1),
		"OffShortMidRangeReboundChances":       int64(2),
		"PersonalFouls":                        int64(2),
		"PersonalFoulsDrawn":                   int64(2),
		"Seconds":                              int64(719),
		"ShootingBlockFoulsDrawn":              int64(1),
		"ShootingFouls":                        int64(3),
		"ShootingFoulsDrawn":                   int64(2),
		"ShortMidRangeMakes":                   int64(1),
		"ShortMidRangeMisses":                  int64(2),
		"Steals":                               int64(1),
		"TechnicalFreeThrows":                  int64(1),
	}
	match = reflect.DeepEqual(opponent_stats, expected_opponent_stats)
	if !match {
		for key, value := range expected_opponent_stats{
			if value != opponent_stats[key]{
				t.Errorf("Expected %d, got %d for %s", value, opponent_stats[key], key)
			}
		}
	}

	//player stats
	player_id = int64(0) // for team rebounds/ shot clock violations
	team_id = int64(1610612744)
	expected_stats := map[string]int64{
		"OffAtRimRebound":   int64(1),
		"DefArc3Rebound":    int64(1),
		"OffCorner3Rebound": int64(1),
	}
	match = reflect.DeepEqual(player_stats[team_id][player_id], expected_stats)
	if !match {
		t.Errorf("Expected %#v, got %#v", expected_stats, player_stats[team_id][player_id])
	}

	player_id = int64(201939)
	expected_stats = map[string]int64{
		"2ptShootingFoulFreeThrowTrips":        int64(1),
		"3ptShootingFoulFreeThrowTrips":        int64(1),
		"Arc3Assists":                          int64(1),
		"Arc3Makes":                            int64(3),
		"Arc3Misses":                           int64(2),
		"AssistedArc3":                         int64(1),
		"AssistedCorner3":                      int64(1),
		"Corner3Makes":                         int64(1),
		"Corner3Misses":                        int64(1),
		"DefArc3ReboundChances":                int64(3),
		"DefAtRimReboundChances":               int64(3),
		"DefBlockedAtRimReboundChances":        int64(2),
		"DefensivePossessions":                 int64(25),
		"DefFTReboundChances":                  int64(1),
		"DefLongMidRangeRebound":               int64(1),
		"DefLongMidRangeReboundChances":        int64(1),
		"DefShortMidRangeReboundChances":       int64(2),
		"FTMakes":                              int64(5),
		"LiveballTurnovers":                    int64(1),
		"OffArc3ReboundChances":                int64(2),
		"OffAtRimReboundChances":               int64(3),
		"OffBlockedAtRimReboundChances":        int64(3),
		"OffBlockedLongMidRangeReboundChances": int64(1),
		"OffCorner3ReboundChances":             int64(4),
		"OffensivePossessions":                 int64(26),
		"OffLongMidRangeRebound":               int64(1),
		"OffLongMidRangeReboundChances":        int64(2),
		"OffShortMidRangeReboundChances":       int64(2),
		"Seconds":                              int64(719),
		"ShootingFoulsDrawn":                   int64(2),
		"Steals":                               int64(1),
	}
	match = reflect.DeepEqual(player_stats[team_id][player_id], expected_stats)
	if !match {
		t.Errorf("Expected %#v, got %#v", expected_stats, player_stats[team_id][player_id])
	}

	team_id = int64(1610612746)
	player_id = int64(201933)
	expected_stats = map[string]int64{
		"2ptShootingFoulFreeThrowTrips":        int64(2),
		"Arc3Misses":                           int64(1),
		"AssistedLongMidRange":                 int64(1),
		"AtRimBlocks":                          int64(1),
		"DefArc3Rebound":                       int64(1),
		"DefArc3ReboundChances":                int64(2),
		"DefAtRimReboundChances":               int64(3),
		"DefBlockedAtRimReboundChances":        int64(3),
		"DefBlockedLongMidRangeReboundChances": int64(1),
		"DefCorner3ReboundChances":             int64(1),
		"DefensivePossessions":                 int64(21),
		"DefLongMidRangeReboundChances":        int64(2),
		"DefShortMidRangeRebound":              int64(1),
		"DefShortMidRangeReboundChances":       int64(2),
		"FTMakes":                              int64(3),
		"FTMisses":                             int64(1),
		"LiveballTurnovers":                    int64(2),
		"LongMidRangeAssists":                  int64(1),
		"LongMidRangeMakes":                    int64(1),
		"OffArc3ReboundChances":                int64(3),
		"OffAtRimReboundChances":               int64(2),
		"OffBlockedAtRimReboundChances":        int64(2),
		"OffensivePossessions":                 int64(20),
		"OffFTReboundChances":                  int64(1),
		"OffLongMidRangeReboundChances":        int64(1),
		"OffShortMidRangeReboundChances":       int64(2),
		"Seconds":                              int64(589),
		"ShootingBlockFoulsDrawn":              int64(1),
		"ShootingFouls":                        int64(1),
		"ShootingFoulsDrawn":                   int64(1),
	}
	match = reflect.DeepEqual(player_stats[team_id][player_id], expected_stats)
	if !match {
		t.Errorf("Expected %#v, got %#v", expected_stats, player_stats[team_id][player_id])
	}

	// lineup stats
	lineup_id = "201939-202691-203110-2585-2738"
	expected_stats = map[string]int64{
		"2ptShootingFoulFreeThrowTrips":        int64(1),
		"Arc3Assists":                          int64(1),
		"Arc3Makes":                            int64(2),
		"Arc3Misses":                           int64(1),
		"AssistedArc3":                         int64(1),
		"AssistedAtRim":                        int64(2),
		"AssistedLongMidRange":                 int64(1),
		"AtRimAssists":                         int64(2),
		"AtRimBlocks":                          int64(2),
		"AtRimMakes":                           int64(2),
		"AtRimMisses":                          int64(4),
		"BlockedAtRim":                         int64(2),
		"BlockedLongMidRange":                  int64(1),
		"Corner3Misses":                        int64(1),
		"DeadballTurnovers":                    int64(1),
		"DefArc3Rebound":                       int64(1),
		"DefArc3ReboundChances":                int64(2),
		"DefAtRimRebound":                      int64(1),
		"DefAtRimReboundChances":               int64(1),
		"DefBlockedAtRimRebound":               int64(1),
		"DefBlockedAtRimReboundChances":        int64(2),
		"DefensivePossessions":                 int64(12),
		"DefFTRebound":                         int64(1),
		"DefFTReboundChances":                  int64(1),
		"DefLongMidRangeRebound":               int64(1),
		"DefLongMidRangeReboundChances":        int64(1),
		"DefShortMidRangeRebound":              int64(2),
		"DefShortMidRangeReboundChances":       int64(2),
		"FTMakes":                              int64(2),
		"LongMidRangeAssists":                  int64(1),
		"LongMidRangeMakes":                    int64(1),
		"LongMidRangeMisses":                   int64(1),
		"OffAtRimRebound":                      int64(1),
		"OffAtRimReboundChances":               int64(2),
		"OffBlockedLongMidRangeRebound":        int64(1),
		"OffBlockedLongMidRangeReboundChances": int64(1),
		"OffensivePossessions":                 int64(13),
		"PersonalFouls":                        int64(1),
		"PersonalFoulsDrawn":                   int64(1),
		"Seconds":                              int64(393),
		"ShootingFouls":                        int64(1),
		"ShootingFoulsDrawn":                   int64(1),
		"ShortMidRangeMisses":                  int64(1),
		"Steals":                               int64(1),
	}
	match = reflect.DeepEqual(lineup_stats[lineup_id], expected_stats)
	if !match {
		for key, value := range expected_stats{
			if value != lineup_stats[lineup_id][key]{
				t.Errorf("Expected %d, got %d for %s", value, lineup_stats[lineup_id][key], key)
			}
		}
	}

	// lineup opponent stats
	expected_stats = map[string]int64{
		"2ptShootingFoulFreeThrowTrips":  int64(1),
		"Arc3Misses":                     int64(2),
		"AssistedLongMidRange":           int64(3),
		"AtRimBlocks":                    int64(2),
		"AtRimMakes":                     int64(1),
		"AtRimMisses":                    int64(3),
		"BlockedAtRim":                   int64(2),
		"DefArc3Rebound":                 int64(1),
		"DefArc3ReboundChances":          int64(1),
		"DefAtRimRebound":                int64(1),
		"DefAtRimReboundChances":         int64(2),
		"DefBlockedAtRimRebound":         int64(2),
		"DefBlockedAtRimReboundChances":  int64(2),
		"DefCorner3Rebound":              int64(1),
		"DefCorner3ReboundChances":       int64(1),
		"DefensivePossessions":           int64(13),
		"DefShortMidRangeRebound":        int64(1),
		"DefShortMidRangeReboundChances": int64(1),
		"FTMakes":                        int64(1),
		"FTMisses":                       int64(1),
		"LiveballTurnovers":              int64(1),
		"LongMidRangeAssists":            int64(3),
		"LongMidRangeBlocks":             int64(1),
		"LongMidRangeMakes":              int64(3),
		"LongMidRangeMisses":             int64(1),
		"OffArc3Rebound":                 int64(1),
		"OffArc3ReboundChances":          int64(2),
		"OffBlockedAtRimRebound":         int64(1),
		"OffBlockedAtRimReboundChances":  int64(2),
		"OffensivePossessions":           int64(12),
		"PersonalFouls":                  int64(1),
		"PersonalFoulsDrawn":             int64(1),
		"Seconds":                        int64(393),
		"ShootingFouls":                  int64(1),
		"ShootingFoulsDrawn":             int64(1),
		"ShortMidRangeMisses":            int64(2),
	}
	match = reflect.DeepEqual(lineup_opponent_stats[lineup_id], expected_stats)
	if !match {
		for key, value := range expected_stats{
			if value != lineup_opponent_stats[lineup_id][key]{
				t.Errorf("Expected %d, got %d for %s", value, lineup_opponent_stats[lineup_id][key], key)
			}
		}
	}
}
