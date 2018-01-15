// package for play by play event with methods for retrieving data and determining event type
package nba_pbp

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

const (
	PERSONAL_FOUL_TYPE_STRING            = "PersonalFouls"
	SHOOTING_FOUL_TYPE_STRING            = "ShootingFouls"
	LOOSE_BALL_FOUL_TYPE_STRING          = "LooseBallFouls"
	OFFENSIVE_FOUL_TYPE_STRING           = "OffensiveFouls"
	AWAY_FROM_PLAY_FOUL_TYPE_STRING      = "AwayFromPlayFouls"
	CHARGE_FOUL_TYPE_STRING              = "ChargeFouls"
	PERSONAL_BLOCK_TYPE_STRING           = "PersonalBlockFouls"
	PERSONAL_TAKE_TYPE_STRING            = "PersonalTakeFouls"
	SHOOTING_BLOCK_TYPE_STRING           = "ShootingBlockFouls"
	CLEAR_PATH_FOUL_TYPE_STRING          = "ClearPathFouls"
	DEFENSIVE_3_SECONDS_FOUL_TYPE_STRING = "Defensive3SecondsViolations"
	FLAGRANT_1_FOUL_TYPE_STRING          = "Flagrant1Fouls"
	FLAGRANT_2_FOUL_TYPE_STRING          = "Flagrant2Fouls"
	DOUBLE_FOUL_TYPE_STRING              = "DoubleFouls"
	INBOUND_FOUL_TYPE_STRING             = "InboundFouls"
)

type PbpEvent struct {
	EventNum              int64  `json:"evt"`
	ClockTime             string `json:"cl"`
	Description           string `json:"de"`
	LocX                  int64  `json:"locX"`
	LocY                  int64  `json:"locY"`
	Opt1                  int64  `json:"opt1"`
	Opt2                  int64  `json:"opt2"`
	Mtype                 int64  `json:"mtype"`
	Etype                 int64  `json:"etype"`
	OPlayerId             string `json:"opid"`
	TeamId                int64  `json:"tid"`
	PlayerId              int64  `json:"pid"`
	HomeScore             int64  `json:"hs"`
	VisitorScore          int64  `json:"vs"`
	EPlayerId             string `json:"epid"`
	OfTeamId              int64  `json:"oftid"`
	CurrentPlayers        map[int64][]int64
	SecondsSinceLastEvent float64
	SecondsToNextEvent    float64
}

func (event *PbpEvent) GetAllEventsAtEventTime(period_events []PbpEvent) ([]PbpEvent, error) {
	// gets all events that happened at the same time as event
	events := []PbpEvent{}
	seconds_remaining, err := event.GetSecondsRemaining()
	if err != nil {
		return events, err
	}
	for _, period_event := range period_events {
		period_event_seconds_remaining, err := period_event.GetSecondsRemaining()
		if err != nil {
			return events, err
		}
		if period_event_seconds_remaining == seconds_remaining {
			events = append(events, period_event)
		}
	}
	return events, nil
}

func (event *PbpEvent) GetOPlayerIdInt() (int64, error) {
	if event.OPlayerId == "" {
		return int64(0), nil
	} else {
		player_id_int, err := strconv.ParseInt(event.OPlayerId, 10, 64)
		if err != nil {
			return 0, err
		}
		return player_id_int, nil
	}
}

func (event *PbpEvent) GetEPlayerIdInt() (int64, error) {
	if event.EPlayerId == "" {
		return int64(0), nil
	} else {
		player_id_int, err := strconv.ParseInt(event.EPlayerId, 10, 64)
		if err != nil {
			return 0, err
		}
		return player_id_int, nil
	}
}

func (event *PbpEvent) IsMadeFG() bool {
	if event.Etype == 1 {
		return true
	}
	return false
}

func (event *PbpEvent) IsMissedFG() bool {
	if event.Etype == 2 {
		return true
	}
	return false
}

func (event *PbpEvent) IsAssistedShot() bool {
	if event.IsMadeFG() && strings.Contains(event.Description, " Assist: ") {
		return true
	}
	return false
}

func (event *PbpEvent) Is3PointShot() bool {
	if (event.IsMadeFG() || event.IsMissedFG()) && strings.Contains(event.Description, "3pt Shot") {
		return true
	}
	return false
}

func (event *PbpEvent) GetSecondsRemaining() (float64, error) {
	// returns seconds remaining in period from clock time formatted mm:ss
	split := strings.Split(event.ClockTime, ":")
	minutes_int, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		return 0, err
	}
	seconds_int, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
		return 0, err
	}
	seconds_remaining := minutes_int*60 + seconds_int
	return seconds_remaining, err
}

func (event *PbpEvent) IsMissedFT() bool {
	if event.Etype == 3 && strings.Contains(event.Description, " Missed") {
		return true
	}
	return false
}

func (event *PbpEvent) IsMadeFT() bool {
	if event.Etype == 3 {
		return !event.IsMissedFT()
	}
	return false
}

func (event *PbpEvent) IsTimeout() bool {
	if event.Etype == 9 {
		return true
	}
	return false
}

func (event *PbpEvent) IsSubstitution() bool {
	if event.Etype == 8 {
		return true
	}
	return false
}

func (event *PbpEvent) IsTurnover() bool {
	if event.Etype == 5 && !strings.Contains(event.Description, "No Turnover") {
		return true
	}
	return false
}

func (event *PbpEvent) IsSteal() bool {
	if event.IsTurnover() && strings.Contains(event.Description, " Steal:") {
		return true
	}
	return false
}

func (event *PbpEvent) IsFoul() bool {
	if event.Etype == 6 {
		return true
	}
	return false
}

func (event *PbpEvent) IsShotClockViolation() bool {
	if event.IsTurnover() && event.Mtype == 11 {
		return true
	}
	return false
}

func (event *PbpEvent) IsRebound() bool {
	// mtype is 1 on team rebounds after ft 1 of 2
	// sometimes mtype is 1 on player rebound, which is why pid != 0 check is there
	if event.Etype == 4 && (event.Mtype == 0 || event.PlayerId != 0) {
		return true
	}
	return false
}

func (event *PbpEvent) IsPutback(previous_event PbpEvent) (bool, error) {
	// assumes event is a made fg
	event_seconds_remaining, err := event.GetSecondsRemaining()
	if err != nil {
		return false, err
	}
	previous_event_seconds_remaining, err := event.GetSecondsRemaining()
	if err != nil {
		return false, err
	}
	if previous_event.IsRebound() && event.PlayerId == previous_event.PlayerId && event_seconds_remaining-previous_event_seconds_remaining <= 2 {
		return true, nil
	}
	return false, nil
}

func (event *PbpEvent) IsFt1of1() bool {
	if event.Etype == 3 && event.Mtype == 10 {
		return true
	}
	return false
}

func (event *PbpEvent) IsFt1of2() bool {
	if event.Etype == 3 && event.Mtype == 11 {
		return true
	}
	return false
}

func (event *PbpEvent) IsFt2of2() bool {
	if event.Etype == 3 && event.Mtype == 12 {
		return true
	}
	return false
}

func (event *PbpEvent) IsFt1of3() bool {
	if event.Etype == 3 && event.Mtype == 13 {
		return true
	}
	return false
}

func (event *PbpEvent) IsFt3of3() bool {
	if event.Etype == 3 && event.Mtype == 15 {
		return true
	}
	return false
}

func (event *PbpEvent) IsFirstFT() bool {
	if event.Etype == 3 && strings.Contains(event.Description, "1 of ") {
		return true
	}
	return false
}

func (event *PbpEvent) IsTechnicalFT() bool {
	if event.Etype == 3 && (event.Mtype == 16 || strings.Contains(event.Description, "Technical")) {
		return true
	}
	return false
}

func (event *PbpEvent) IsTechnicalFoul() bool {
	technical_foul_mtypes := []int64{11, 12, 13, 18, 19, 25, 30}
	if event.Etype == 6 && IntInSlice(event.Mtype, technical_foul_mtypes) {
		return true
	}
	return false
}
func (event *PbpEvent) IsDoubleTechnical() bool {
	if event.Etype == 6 && event.Mtype == 16 {
		return true
	}
	return false
}
func (event *PbpEvent) IsEjection() bool {
	if event.Etype == 11 {
		return true
	}
	return false
}

func (event *PbpEvent) GetFoulType() string {
	/*
		mtype:
				*1 - Personal
				*2 - Shooting
				*3 - Loose Ball
				*4 - Offensive
				*5 - Inbound foul (1 FTA)
				*6 - Away from play
				8 - Punch foul (Technical)
				*9 - Clear Path
				*10 - Double Foul
				11 - Technical
				12 - Non-Unsportsmanlike (Technical)
				13 - Hanging (Technical)
				*14 - Flagrant 1
				*15 - Flagrant 2
				16 - Double Technical
				*17 - Defensive 3 seconds (Technical)
				18 - Delay of game
				19 - Taunting (Technical)
				25 - Excess Timeout (Technical)
				*26 - Charge
				*27 - Personal Block
				*28 - Personal Take
				*29 - Shooting Block
				30 - Too many players (Technical)
	*/
	if event.Etype == 6 {
		if event.Mtype == 1 {
			return PERSONAL_FOUL_TYPE_STRING
		}
		if event.Mtype == 2 {
			return SHOOTING_FOUL_TYPE_STRING
		}
		if event.Mtype == 3 {
			return LOOSE_BALL_FOUL_TYPE_STRING
		}
		if event.Mtype == 4 {
			return OFFENSIVE_FOUL_TYPE_STRING
		}
		if event.Mtype == 5 {
			return INBOUND_FOUL_TYPE_STRING
		}
		if event.Mtype == 6 {
			return AWAY_FROM_PLAY_FOUL_TYPE_STRING
		}
		if event.Mtype == 9 {
			return CLEAR_PATH_FOUL_TYPE_STRING
		}
		if event.Mtype == 10 {
			return DOUBLE_FOUL_TYPE_STRING
		}
		if event.Mtype == 14 {
			return FLAGRANT_1_FOUL_TYPE_STRING
		}
		if event.Mtype == 15 {
			return FLAGRANT_2_FOUL_TYPE_STRING
		}
		if event.Mtype == 17 {
			return DEFENSIVE_3_SECONDS_FOUL_TYPE_STRING
		}
		if event.Mtype == 26 {
			return CHARGE_FOUL_TYPE_STRING
		}
		if event.Mtype == 27 {
			return PERSONAL_BLOCK_TYPE_STRING
		}
		if event.Mtype == 28 {
			return PERSONAL_TAKE_TYPE_STRING
		}
		if event.Mtype == 29 {
			return SHOOTING_BLOCK_TYPE_STRING
		}
	}
	return ""
}

func (event *PbpEvent) GetFouledFgm(period_events []PbpEvent) (PbpEvent, error) {
	// gets shot that was an and 1
	// event is the foul event
	seconds_remaining, err := event.GetSecondsRemaining()
	if err != nil {
		return PbpEvent{}, err
	}
	for _, period_event := range period_events {
		if event.EventNum == period_event.EventNum {
			// period_events is ordered sequentially, if we reach foul event we can assume no and1
			return PbpEvent{}, nil
		} else if period_event.IsMadeFG() {
			period_event_seconds_remaining, err := period_event.GetSecondsRemaining()
			if err != nil {
				return PbpEvent{}, err
			}
			if period_event_seconds_remaining == seconds_remaining && event.TeamId != period_event.TeamId {
				// shot before foul at same time as foul by opposing team
				return period_event, nil
			}
		}
	}
	return PbpEvent{}, nil
}

func (event *PbpEvent) GetFoulThatResultedInFt(period_events []PbpEvent) (PbpEvent, error) {
	// gets foul that led to free throws
	// event is FT
	possible_events := []PbpEvent{}
	events_at_time_of_ft, err := event.GetAllEventsAtEventTime(period_events)
	if err != nil {
		return PbpEvent{}, err
	}
	for _, period_event := range events_at_time_of_ft {
		if event.EventNum == period_event.EventNum {
			// once we get to ft event, we can return the foul event
			// if no foul event yet, need to keep going since sometimes pbp is out of order and foul is after ft
			if len(possible_events) != 0 {
				return possible_events[0], nil
			}
		} else if period_event.IsFoul() && !period_event.IsTechnicalFoul() && !period_event.IsDoubleTechnical() && event.TeamId != period_event.TeamId {
			// foul at same time as ft by opposing team
			possible_events = append(possible_events, period_event)
		}
	}
	if len(possible_events) != 0 {
		// bug in pbp where foul is after FT, that will get returned here
		return possible_events[0], nil
	} else {
		return PbpEvent{}, nil
	}
}

func (event *PbpEvent) GetLastFoul(period_events []PbpEvent) (PbpEvent, error) {
	// this is the same as GetFoulThatResultedInFt without technical foul check, TODO: combine these
	// event is FT
	possible_events := []PbpEvent{}
	events_at_time_of_ft, err := event.GetAllEventsAtEventTime(period_events)
	if err != nil {
		return PbpEvent{}, err
	}
	for _, period_event := range events_at_time_of_ft {
		if event.EventNum == period_event.EventNum {
			// once we get to ft event, we can return the foul event
			// if no foul event yet, need to keep going since sometimes pbp is out of order and foul is after ft
			if len(possible_events) != 0 {
				return possible_events[0], nil
			}
		} else if period_event.IsFoul() {
			// foul at same time as ft by opposing team
			possible_events = append(possible_events, period_event)
		}
	}
	if len(possible_events) != 0 {
		// bug in pbp where foul is after FT, that will get returned here
		return possible_events[0], nil
	} else {
		return PbpEvent{}, nil
	}
}

func (event *PbpEvent) GetNumberOfFtaForFoul() int64 {
	// assumes an event is foul
	if strings.Contains(event.Description, "(1 FTA)") {
		return int64(1)
	} else if strings.Contains(event.Description, "(2 FTA)") {
		return int64(2)
	} else if strings.Contains(event.Description, "(3 FTA)") {
		return int64(3)
	}
	return int64(0)
}

func (event *PbpEvent) IsEndOfPeriod() bool {
	if event.Etype == 13 && event.Mtype == 0 {
		return true
	}
	return false
}

func (event *PbpEvent) IsDelayOfGame() bool {
	if event.Etype == 6 && event.Mtype == 18 {
		return true
	}
	return false
}

func (event *PbpEvent) IsJumpBall() bool {
	if event.Etype == 10 {
		return true
	}
	return false
}

func (event *PbpEvent) IsBlockedShot() bool {
	// blocked shot is missed shot with non blank opid
	if event.IsMissedFG() && event.OPlayerId != "" {
		return true
	}
	return false
}

func (event *PbpEvent) IsFtFromInboundFoul(period_events []PbpEvent) (bool, error) {
	if event.IsFt1of1() {
		events_at_time_of_ft, err := event.GetAllEventsAtEventTime(period_events)
		if err != nil {
			return false, err
		}
		for _, period_event := range events_at_time_of_ft {
			if period_event.GetFoulType() == INBOUND_FOUL_TYPE_STRING {
				return true, nil
			}
		}
	}
	return false, nil
}

func (event *PbpEvent) IsFtFromAwayFromPlayFoul(period_events []PbpEvent) (bool, error) {
	if event.IsFt1of1() {
		events_at_time_of_ft, err := event.GetAllEventsAtEventTime(period_events)
		if err != nil {
			return false, err
		}
		away_from_play_fouls := []PbpEvent{}
		for _, period_event := range events_at_time_of_ft {
			if period_event.GetFoulType() == AWAY_FROM_PLAY_FOUL_TYPE_STRING {
				away_from_play_fouls = append(away_from_play_fouls, period_event)
			}
		}
		if len(away_from_play_fouls) == 1 {
			made_shots_at_time_of_foul := []PbpEvent{}
			for _, period_event := range events_at_time_of_ft {
				if period_event.IsMadeFG() {
					made_shots_at_time_of_foul = append(made_shots_at_time_of_foul, period_event)
				}
			}
			if len(made_shots_at_time_of_foul) == 0 {
				return true, nil
			} else if made_shots_at_time_of_foul[0].TeamId == away_from_play_fouls[0].TeamId && event.TeamId != made_shots_at_time_of_foul[0].TeamId {
				// team that made shot is team that committed foul at time of shot
				return true, nil
			}
		}
	}
	return false, nil
}

func (event *PbpEvent) IsAnd1Shot(period_events []PbpEvent) (bool, error) {
	// assumes event is made fg
	events_at_time_of_shot, err := event.GetAllEventsAtEventTime(period_events)
	if err != nil {
		return false, err
	}
	// check for ft 1 of 1 and foul drawn by player at same time as event
	player_drew_foul := false
	player_shot_ft_1_of_1 := false
	for _, period_event := range events_at_time_of_shot {
		if period_event.IsFt1of1() && period_event.PlayerId == event.PlayerId {
			player_shot_ft_1_of_1 = true
		}
		if period_event.IsFoul() {
			foul_drawn_player_id, err := period_event.GetOPlayerIdInt()
			if err != nil {
				return false, err
			}
			if foul_drawn_player_id == event.PlayerId {
				player_drew_foul = true
			}
		}
	}
	if player_drew_foul && player_shot_ft_1_of_1 {
		return true, nil
	}
	return false, nil
}

func (event *PbpEvent) GetReboundedShot(period_events []PbpEvent) (PbpEvent, error) {
	// period_events should be ordered from start of period to end
	if event.ClockTime == "00:00.0" && event.PlayerId != 0 {
		// Ignore team rebounds at end of period since they aren't actual rebounds, they are just placeholder events
		return PbpEvent{}, nil
	} else {
		seconds_remaining, err := event.GetSecondsRemaining()
		if err != nil {
			return PbpEvent{}, err
		}
		last_miss_before_rebound := PbpEvent{}
		found_rebound := false
		for _, period_event := range period_events {
			if event.EventNum == period_event.EventNum {
				found_rebound = true
			}
			if !found_rebound && (period_event.IsMissedFT() || period_event.IsMissedFG()) {
				last_miss_before_rebound = period_event
			}else if period_event.IsShotClockViolation() && event.IsTeamRebound(){
				// check for shot clock violation at time of rebound, we can ignore these rebounds
				shot_clock_seconds_remaining, err := period_event.GetSecondsRemaining()
				if err != nil {
					return PbpEvent{}, err
				}
				if shot_clock_seconds_remaining == seconds_remaining {
					return PbpEvent{}, nil
				}
			}
		}
		return last_miss_before_rebound, nil
	}
}

func (event *PbpEvent) IsDefensiveRebound(missed_shot PbpEvent) bool {
	if event.TeamId != missed_shot.TeamId {
		return true
	}
	return false
}

func (event *PbpEvent) IsTeamRebound() bool {
	if event.PlayerId == 0 {
		return true
	}
	return false
}

func (event *PbpEvent) IsSelfRebound(missed_shot PbpEvent) bool {
	if event.PlayerId == missed_shot.PlayerId {
		return true
	}
	return false
}

func (event *PbpEvent) IsTrackedEvent() bool {
	// returns true if event is an event that should trigger a counting stat getting incremented
	if event.IsTurnover() {
		return true
	}
	if event.IsFoul() {
		return true
	}
	if event.IsMadeFT() {
		return true
	}
	if event.IsMissedFT() {
		return true
	}
	if event.IsMadeFG() {
		return true
	}
	if event.IsMissedFG() {
		return true
	}
	if event.IsRebound() {
		return true
	}

	return false
}

func (event *PbpEvent) GenerateLineupIds() map[int64]string {
	// lineup id is '-' separated player ids(ordered as strings)
	lineup_ids := make(map[int64]string)
	for team_id, team_players := range event.CurrentPlayers {
		player_ids_as_strings := []string{}
		for _, player_id := range team_players {
			player_ids_as_strings = append(player_ids_as_strings, strconv.Itoa(int(player_id)))
		}
		sort.Strings(player_ids_as_strings)
		lineup_id := strings.Join(player_ids_as_strings, "-")
		lineup_ids[team_id] = lineup_id
	}
	return lineup_ids
}

func (event *PbpEvent) GetShotDistance() float64 {
	// get shot distance from net for shot, center of rim coordinates are 0,0
	x_squared := float64(event.LocX * event.LocX)
	y_squared := float64(event.LocY * event.LocY)
	shot_distance := math.Sqrt(x_squared+y_squared) / 10 // unit for distance is off by factor of 10, divide by 10 to convert to feet
	return shot_distance
}

func (event *PbpEvent) IsCorner3() bool {
	if math.Abs(float64(event.LocX)) >= 220 && event.LocY <= 87 {
		return true
	}
	return false
}
