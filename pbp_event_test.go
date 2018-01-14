package nba_pbp

import (
	"testing"
)

// start test events
var MISSED_2 = PbpEvent{
	EventNum:     11,
	ClockTime:    "11:16",
	Description:  "[LAC] Jordan Layup Shot: Missed",
	LocX:         -10,
	LocY:         12,
	Opt1:         2,
	Opt2:         0,
	Mtype:        5,
	Etype:        2,
	OPlayerId:    "",
	TeamId:       1610612746,
	PlayerId:     201599,
	HomeScore:    0,
	VisitorScore: 0,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var ASSISTED_MADE_2 = PbpEvent{
	EventNum:     17,
	ClockTime:    "10:39",
	Description:  "[LAC 2-0] Williams Pullup Jump shot: Made (2 PTS) Assist: Evans (1 AST)",
	LocX:         177,
	LocY:         77,
	Opt1:         2,
	Opt2:         0,
	Mtype:        79,
	Etype:        1,
	OPlayerId:    "",
	TeamId:       1610612746,
	PlayerId:     203710,
	HomeScore:    2,
	VisitorScore: 0,
	EPlayerId:    "1628393",
	OfTeamId:     1610612746,
}
var UNASSISTED_MADE_3 = PbpEvent{
	EventNum:     92,
	ClockTime:    "04:57",
	Description:  "[LAC 14-17] Williams 3pt Shot: Made (3 PTS)",
	LocX:         -21,
	LocY:         259,
	Opt1:         3,
	Opt2:         0,
	Mtype:        79,
	Etype:        1,
	OPlayerId:    "",
	TeamId:       1610612746,
	PlayerId:     101150,
	HomeScore:    14,
	VisitorScore: 17,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var MISSED_3 = PbpEvent{
	EventNum:     534,
	ClockTime:    "03:04",
	Description:  "[GSW] Curry 3pt Shot: Missed",
	LocX:         29,
	LocY:         276,
	Opt1:         3,
	Opt2:         0,
	Mtype:        1,
	Etype:        2,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     201939,
	HomeScore:    71,
	VisitorScore: 93,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var MISSED_CORNER_3 = PbpEvent{
	EventNum:     53,
	ClockTime:    "07:23",
	Description:  "[GSW] Curry 3pt Shot: Missed",
	LocX:         -228,
	LocY:         -22,
	Opt1:         3,
	Opt2:         0,
	Mtype:        1,
	Etype:        2,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     201939,
	HomeScore:    6,
	VisitorScore: 6,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}

var MADE_FT = PbpEvent{
	EventNum:     45,
	ClockTime:    "08:06",
	Description:  "[GSW 3-6] Curry Free Throw 1 of 2 (1 PTS)",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         0,
	Mtype:        11,
	Etype:        3,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     201939,
	HomeScore:    6,
	VisitorScore: 3,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var MISSED_FT = PbpEvent{
	EventNum:     73,
	ClockTime:    "06:13",
	Description:  "[LAC] Griffin Free Throw 2 of 2 Missed",
	LocX:         0,
	LocY:         -80,
	Opt1:         2,
	Opt2:         0,
	Mtype:        12,
	Etype:        3,
	OPlayerId:    "",
	TeamId:       1610612746,
	PlayerId:     201933,
	HomeScore:    9,
	VisitorScore: 11,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var TIMEOUT_EVENT = PbpEvent{
	EventNum:     109,
	ClockTime:    "03:15",
	Description:  "[GSW] Team Timeout : Regular",
	LocX:         0,
	LocY:         -80,
	Opt1:         0,
	Opt2:         0,
	Mtype:        1,
	Etype:        9,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     1610612744,
	HomeScore:    18,
	VisitorScore: 26,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var SUBSTITUTION_EVENT = PbpEvent{
	EventNum:     113,
	ClockTime:    "03:15",
	Description:  "[GSW] Thompson Substitution replaced by McCaw",
	LocX:         0,
	LocY:         -80,
	Opt1:         0,
	Opt2:         0,
	Mtype:        0,
	Etype:        8,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     202691,
	HomeScore:    18,
	VisitorScore: 26,
	EPlayerId:    "1627775",
	OfTeamId:     1610612744,
}
var SHOOTING_FOUL_EVENT = PbpEvent{
	EventNum:     127,
	ClockTime:    "02:42",
	Description:  "[LAC] Johnson Foul: Shooting (1 PF) (1 FTA)",
	LocX:         -35,
	LocY:         756,
	Opt1:         0,
	Opt2:         0,
	Mtype:        2,
	Etype:        6,
	OPlayerId:    "201580",
	TeamId:       1610612746,
	PlayerId:     202325,
	HomeScore:    20,
	VisitorScore: 28,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var AND1_SHOT_EVENT = PbpEvent{
	EventNum:     126,
	ClockTime:    "02:42",
	Description:  "[GSW 28-20] McGee Tip Layup Shot: Made (2 PTS)",
	LocX:         0,
	LocY:         -5,
	Opt1:         2,
	Opt2:         0,
	Mtype:        97,
	Etype:        1,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     201580,
	HomeScore:    20,
	VisitorScore: 28,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var AND1_FT_EVENT = PbpEvent{
	EventNum:     132,
	ClockTime:    "02:42",
	Description:  "[GSW 29-20] McGee Free Throw 1 of 1 (3 PTS)",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         0,
	Mtype:        10,
	Etype:        3,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     201580,
	HomeScore:    20,
	VisitorScore: 29,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var STEAL_EVENT = PbpEvent{
	EventNum:     130,
	ClockTime:    "02:23",
	Description:  "[LAC] Griffin Turnover : Bad Pass (2 TO) Steal:Young (1 ST)",
	LocX:         -23,
	LocY:         74,
	Opt1:         1,
	Opt2:         0,
	Mtype:        1,
	Etype:        5,
	OPlayerId:    "201156",
	TeamId:       1610612746,
	PlayerId:     201933,
	HomeScore:    20,
	VisitorScore: 29,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var DEADBALL_TURNOVER_EVENT = PbpEvent{
	EventNum:     195,
	ClockTime:    "11:11",
	Description:  "[GSW] Thompson Turnover : Out of Bounds - Bad Pass Turnover (2 TO)",
	LocX:         -185,
	LocY:         33,
	Opt1:         0,
	Opt2:         0,
	Mtype:        45,
	Etype:        5,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     202691,
	HomeScore:    30,
	VisitorScore: 33,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var SHOT_CLOCK_VIOLATION = PbpEvent{
	EventNum:     629,
	ClockTime:    "09:11",
	Description:  "[GSW] Team Turnover : Shot Clock Turnover",
	LocX:         170,
	LocY:         26,
	Opt1:         0,
	Opt2:         0,
	Mtype:        11,
	Etype:        5,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     0,
	HomeScore:    80,
	VisitorScore: 108,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var PLAYER_DEFENSIVE_REBOUND = PbpEvent{
	EventNum:     631,
	ClockTime:    "08:55",
	Description:  "[GSW] McCaw Rebound (Off:0 Def:3)",
	LocX:         15,
	LocY:         10,
	Opt1:         0,
	Opt2:         0,
	Mtype:        0,
	Etype:        4,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     1627775,
	HomeScore:    80,
	VisitorScore: 108,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var PLAYER_OFFENSIVE_REBOUND = PbpEvent{
	EventNum:     670,
	ClockTime:    "05:49",
	Description:  "[LAC] Evans Rebound (Off:2 Def:2)",
	LocX:         58,
	LocY:         142,
	Opt1:         1,
	Opt2:         0,
	Mtype:        0,
	Etype:        4,
	OPlayerId:    "",
	TeamId:       1610612746,
	PlayerId:     1628393,
	HomeScore:    85,
	VisitorScore: 113,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var TEAM_OFFENSIVE_REBOUND = PbpEvent{
	EventNum:     8,
	ClockTime:    "11:40",
	Description:  "[GSW] Team Rebound",
	LocX:         -14,
	LocY:         11,
	Opt1:         1,
	Opt2:         0,
	Mtype:        0,
	Etype:        4,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     0,
	HomeScore:    0,
	VisitorScore: 0,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var TEAM_REBOUND_TO_IGNORE = PbpEvent{
	EventNum:     317,
	ClockTime:    "04:19",
	Description:  "[GSW] Team Rebound",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         0,
	Mtype:        1,
	Etype:        4,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     0,
	HomeScore:    48,
	VisitorScore: 49,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var REBOUND_EVENT_BEFORE_PUTBACK = PbpEvent{
	EventNum:     446,
	ClockTime:    "07:36",
	Description:  "[LAC] Jordan Rebound (Off:3 Def:5)",
	LocX:         0,
	LocY:         11,
	Opt1:         1,
	Opt2:         0,
	Mtype:        0,
	Etype:        4,
	OPlayerId:    "",
	TeamId:       1610612746,
	PlayerId:     201599,
	HomeScore:    59,
	VisitorScore: 73,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var PUTBACK_EVENT = PbpEvent{
	EventNum:     447,
	ClockTime:    "07:36",
	Description:  "[LAC 61-73] Jordan Putback Layup Shot: Made (11 PTS)",
	LocX:         0,
	LocY:         -5,
	Opt1:         2,
	Opt2:         0,
	Mtype:        72,
	Etype:        1,
	OPlayerId:    "",
	TeamId:       1610612746,
	PlayerId:     201599,
	HomeScore:    61,
	VisitorScore: 73,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var TECHNICAL_FT_EVENT = PbpEvent{
	EventNum:     140,
	ClockTime:    "02:11",
	Description:  "[LAC 21-29] Williams Free Throw Technical (8 PTS)",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         0,
	Mtype:        16,
	Etype:        3,
	OPlayerId:    "",
	TeamId:       1610612746,
	PlayerId:     101150,
	HomeScore:    21,
	VisitorScore: 29,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var FT_1_OF_2_EVENT = PbpEvent{
	EventNum:     160,
	ClockTime:    "01:18",
	Description:  "[GSW 30-21] Casspi Free Throw 1 of 2 (1 PTS)",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         0,
	Mtype:        11,
	Etype:        3,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     201956,
	HomeScore:    21,
	VisitorScore: 30,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var FT_2_OF_2_EVENT = PbpEvent{
	EventNum:     161,
	ClockTime:    "01:18",
	Description:  "[GSW 31-21] Casspi Free Throw 2 of 2 (2 PTS)",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         0,
	Mtype:        12,
	Etype:        3,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     201956,
	HomeScore:    21,
	VisitorScore: 31,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var FT_1_OF_3_EVENT = PbpEvent{
	EventNum:     472,
	ClockTime:    "05:49",
	Description:  "[GSW 79-63] Curry Free Throw 1 of 3 (35 PTS)",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         1,
	Mtype:        13,
	Etype:        3,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     201939,
	HomeScore:    63,
	VisitorScore: 79,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var FT_3_OF_3_EVENT = PbpEvent{
	EventNum:     474,
	ClockTime:    "05:49",
	Description:  "[GSW 81-63] Curry Free Throw 3 of 3 (37 PTS)",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         1,
	Mtype:        15,
	Etype:        3,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     201939,
	HomeScore:    63,
	VisitorScore: 81,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var INBOUND_FOUL_EVENT = PbpEvent{
	EventNum:     575,
	ClockTime:    "00:21.2",
	Description:  "Inbound Foul",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         1,
	Mtype:        5,
	Etype:        6,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     923,
	HomeScore:    76,
	VisitorScore: 102,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var AWAY_FROM_PLAY_FOUL_EVENT = PbpEvent{
	EventNum:     575,
	ClockTime:    "00:21.2",
	Description:  "Away From Play Foul",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         1,
	Mtype:        6,
	Etype:        6,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     923,
	HomeScore:    76,
	VisitorScore: 102,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var FT_1_OF_1_EVENT = PbpEvent{
	EventNum:     576,
	ClockTime:    "00:21.2",
	Description:  "[LAC 76-103] Evans Free Throw 1 of 1 (5 PTS)",
	LocX:         0,
	LocY:         -80,
	Opt1:         1,
	Opt2:         1,
	Mtype:        10,
	Etype:        3,
	OPlayerId:    "",
	TeamId:       1610612746,
	PlayerId:     1628393,
	HomeScore:    76,
	VisitorScore: 103,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var TECHNICAL_FOUL = PbpEvent{
	EventNum:     138,
	ClockTime:    "02:11",
	Description:  "[GSW] Green Technical (1 FTA) (T Ford)",
	LocX:         0,
	LocY:         -80,
	Opt1:         0,
	Opt2:         0,
	Mtype:        11,
	Etype:        6,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     203110,
	HomeScore:    20,
	VisitorScore: 29,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var DOUBLE_TECHNICAL_FOUL = PbpEvent{
	EventNum:     138,
	ClockTime:    "02:11",
	Description:  "",
	LocX:         0,
	LocY:         -80,
	Opt1:         0,
	Opt2:         0,
	Mtype:        16,
	Etype:        6,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     203110,
	HomeScore:    20,
	VisitorScore: 29,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var EJECTION_EVENT = PbpEvent{
	EventNum:     138,
	ClockTime:    "02:11",
	Description:  "",
	LocX:         0,
	LocY:         -80,
	Opt1:         0,
	Opt2:         0,
	Mtype:        16,
	Etype:        11,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     203110,
	HomeScore:    20,
	VisitorScore: 29,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}
var BLOCKED_SHOT_EVENT = PbpEvent{
	EventNum:     564,
	ClockTime:    "01:06",
	Description:  "[GSW] Curry Running Reverse Layup Shot: Missed",
	LocX:         14,
	LocY:         15,
	Opt1:         2,
	Opt2:         1,
	Mtype:        74,
	Etype:        2,
	OPlayerId:    "1626149",
	TeamId:       1610612744,
	PlayerId:     201939,
	HomeScore:    73,
	VisitorScore: 101,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var JUMP_BALL_EVENT = PbpEvent{
	EventNum:     4,
	ClockTime:    "11:54",
	Description:  "Jump Ball Pachulia vs Jordan (Curry gains possession)",
	LocX:         0,
	LocY:         -80,
	Opt1:         0,
	Opt2:         0,
	Mtype:        0,
	Etype:        10,
	OPlayerId:    "201599",
	TeamId:       1610612744,
	PlayerId:     2585,
	HomeScore:    0,
	VisitorScore: 0,
	EPlayerId:    "201939",
	OfTeamId:     1610612744,
}
var END_OF_PERIOD_EVENT = PbpEvent{
	EventNum:     176,
	ClockTime:    "00:00.0",
	Description:  "End Period",
	LocX:         0,
	LocY:         -80,
	Opt1:         0,
	Opt2:         0,
	Mtype:        0,
	Etype:        13,
	OPlayerId:    "",
	TeamId:       0,
	PlayerId:     0,
	HomeScore:    28,
	VisitorScore: 31,
	EPlayerId:    "",
	OfTeamId:     1610612744,
}
var DELAY_OF_GAME_EVENT = PbpEvent{
	EventNum:     138,
	ClockTime:    "02:11",
	Description:  "",
	LocX:         0,
	LocY:         -80,
	Opt1:         0,
	Opt2:         0,
	Mtype:        18,
	Etype:        6,
	OPlayerId:    "",
	TeamId:       1610612744,
	PlayerId:     203110,
	HomeScore:    20,
	VisitorScore: 29,
	EPlayerId:    "",
	OfTeamId:     1610612746,
}

// end test events

func TestGetAllEventsAtEventTime(t *testing.T) {
	event_options := []PbpEvent{SHOOTING_FOUL_EVENT, AND1_SHOT_EVENT, AND1_FT_EVENT, DELAY_OF_GAME_EVENT}
	events, err := SHOOTING_FOUL_EVENT.GetAllEventsAtEventTime(event_options)
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	expected_length := 3
	if len(events) != expected_length {
		t.Errorf("Expected %d. Got %d.", expected_length, len(events))
	}
}

func TestGetEPlayerIdInt(t *testing.T) {
	player_id, err := SUBSTITUTION_EVENT.GetEPlayerIdInt()
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	expected_player_id := int64(1627775)
	if player_id != expected_player_id {
		t.Errorf("Expected %d. Got %d.", expected_player_id, player_id)
	}

	// when epid is empty string
	player_id, err = FT_1_OF_1_EVENT.GetEPlayerIdInt()
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	expected_player_id = int64(0)
	if player_id != expected_player_id {
		t.Errorf("Expected %d. Got %d.", expected_player_id, player_id)
	}
}

func TestGetOPlayerIdInt(t *testing.T) {
	player_id, err := STEAL_EVENT.GetOPlayerIdInt()
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	expected_player_id := int64(201156)
	if player_id != expected_player_id {
		t.Errorf("Expected %d. Got %d.", expected_player_id, player_id)
	}

	// when epid is empty string
	player_id, err = FT_1_OF_1_EVENT.GetOPlayerIdInt()
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	expected_player_id = int64(0)
	if player_id != expected_player_id {
		t.Errorf("Expected %d. Got %d.", expected_player_id, player_id)
	}
}

func TestMadeFG(t *testing.T) {
	missed_2 := MISSED_2.IsMadeFG()
	if missed_2 != false {
		t.Errorf("Expected false. Got %t.", missed_2)
	}

	made_2 := ASSISTED_MADE_2.IsMadeFG()
	if made_2 != true {
		t.Errorf("Expected true. Got %t.", made_2)
	}
}

func TestMissedFG(t *testing.T) {
	missed_2 := MISSED_2.IsMissedFG()
	if missed_2 != true {
		t.Errorf("Expected true. Got %t.", missed_2)
	}

	made_2 := ASSISTED_MADE_2.IsMissedFG()
	if made_2 != false {
		t.Errorf("Expected false. Got %t.", made_2)
	}
}

func TestAssistedShot(t *testing.T) {
	missed_2 := MISSED_2.IsAssistedShot()
	if missed_2 != false {
		t.Errorf("Expected false. Got %t.", missed_2)
	}

	assisted_made_2 := ASSISTED_MADE_2.IsAssistedShot()
	if assisted_made_2 != true {
		t.Errorf("Expected true. Got %t.", assisted_made_2)
	}

	unassisted_made_3 := UNASSISTED_MADE_3.IsAssistedShot()
	if unassisted_made_3 != false {
		t.Errorf("Expected false. Got %t.", unassisted_made_3)
	}
}

func Test3PointShot(t *testing.T) {
	missed_3 := MISSED_3.Is3PointShot()
	if missed_3 != true {
		t.Errorf("Expected true. Got %t.", missed_3)
	}

	assisted_made_2 := ASSISTED_MADE_2.Is3PointShot()
	if assisted_made_2 != false {
		t.Errorf("Expected false. Got %t.", assisted_made_2)
	}

	unassisted_made_3 := UNASSISTED_MADE_3.Is3PointShot()
	if unassisted_made_3 != true {
		t.Errorf("Expected true. Got %t.", unassisted_made_3)
	}
}

func TestSecondsRemaining(t *testing.T) {
	seconds_remaining, err := MISSED_3.GetSecondsRemaining()
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	expected_seconds_remaining := float64(184)
	if seconds_remaining != expected_seconds_remaining {
		t.Errorf("Expected %d. Got %d.", expected_seconds_remaining, seconds_remaining)
	}
}

func TestMadeFT(t *testing.T) {
	made_ft := MADE_FT.IsMadeFT()
	if made_ft != true {
		t.Errorf("Expected true. Got %t.", made_ft)
	}

	missed_ft := MISSED_FT.IsMadeFT()
	if missed_ft != false {
		t.Errorf("Expected false. Got %t.", missed_ft)
	}
}

func TestMissedFT(t *testing.T) {
	made_ft := MADE_FT.IsMissedFT()
	if made_ft != false {
		t.Errorf("Expected false. Got %t.", made_ft)
	}

	missed_ft := MISSED_FT.IsMissedFT()
	if missed_ft != true {
		t.Errorf("Expected true. Got %t.", missed_ft)
	}
}

func TestTimeout(t *testing.T) {
	timeout := TIMEOUT_EVENT.IsTimeout()
	if timeout != true {
		t.Errorf("Expected true. Got %t.", timeout)
	}

	missed_ft := MISSED_FT.IsTimeout()
	if missed_ft != false {
		t.Errorf("Expected false. Got %t.", missed_ft)
	}
}

func TestSubstitution(t *testing.T) {
	substitution := SUBSTITUTION_EVENT.IsSubstitution()
	if substitution != true {
		t.Errorf("Expected true. Got %t.", substitution)
	}

	missed_ft := MISSED_FT.IsSubstitution()
	if missed_ft != false {
		t.Errorf("Expected false. Got %t.", missed_ft)
	}
}

func TestTurnover(t *testing.T) {
	steal := STEAL_EVENT.IsTurnover()
	if steal != true {
		t.Errorf("Expected true. Got %t.", steal)
	}

	deadball := DEADBALL_TURNOVER_EVENT.IsTurnover()
	if deadball != true {
		t.Errorf("Expected true. Got %t.", deadball)
	}

	shot_clock_violation := SHOT_CLOCK_VIOLATION.IsTurnover()
	if shot_clock_violation != true {
		t.Errorf("Expected true. Got %t.", shot_clock_violation)
	}

	made_ft := MADE_FT.IsTurnover()
	if made_ft != false {
		t.Errorf("Expected false. Got %t.", made_ft)
	}
}

func TestSteal(t *testing.T) {
	steal := STEAL_EVENT.IsSteal()
	if steal != true {
		t.Errorf("Expected true. Got %t.", steal)
	}

	deadball := DEADBALL_TURNOVER_EVENT.IsSteal()
	if deadball != false {
		t.Errorf("Expected false. Got %t.", deadball)
	}

	made_ft := MADE_FT.IsSteal()
	if made_ft != false {
		t.Errorf("Expected false. Got %t.", made_ft)
	}
}

func TestFoul(t *testing.T) {
	shooting_foul := SHOOTING_FOUL_EVENT.IsFoul()
	if shooting_foul != true {
		t.Errorf("Expected true. Got %t.", shooting_foul)
	}

	made_ft := MADE_FT.IsFoul()
	if made_ft != false {
		t.Errorf("Expected false. Got %t.", made_ft)
	}
}

func TestShotClockViolation(t *testing.T) {
	deadball := DEADBALL_TURNOVER_EVENT.IsShotClockViolation()
	if deadball != false {
		t.Errorf("Expected false. Got %t.", deadball)
	}

	shot_clock_violation := SHOT_CLOCK_VIOLATION.IsShotClockViolation()
	if shot_clock_violation != true {
		t.Errorf("Expected true. Got %t.", shot_clock_violation)
	}

	made_ft := MADE_FT.IsShotClockViolation()
	if made_ft != false {
		t.Errorf("Expected false. Got %t.", made_ft)
	}
}

func TestRebound(t *testing.T) {
	def_reb := PLAYER_DEFENSIVE_REBOUND.IsRebound()
	if def_reb != true {
		t.Errorf("Expected true. Got %t.", def_reb)
	}

	off_reb := PLAYER_OFFENSIVE_REBOUND.IsRebound()
	if off_reb != true {
		t.Errorf("Expected true. Got %t.", off_reb)
	}

	team_reb := TEAM_OFFENSIVE_REBOUND.IsRebound()
	if team_reb != true {
		t.Errorf("Expected true. Got %t.", team_reb)
	}

	team_reb_to_ignore := TEAM_REBOUND_TO_IGNORE.IsRebound()
	if team_reb_to_ignore != false {
		t.Errorf("Expected false. Got %t.", team_reb_to_ignore)
	}
}

func TestPutback(t *testing.T) {
	putback, err := PUTBACK_EVENT.IsPutback(REBOUND_EVENT_BEFORE_PUTBACK)
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if putback != true {
		t.Errorf("Expected true. Got %t.", putback)
	}

	putback, err = ASSISTED_MADE_2.IsPutback(PLAYER_OFFENSIVE_REBOUND)
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if putback != false {
		t.Errorf("Expected false. Got %t.", putback)
	}

}

func TestFt1of1(t *testing.T) {
	ft_1_of_1 := FT_1_OF_1_EVENT.IsFt1of1()
	if ft_1_of_1 != true {
		t.Errorf("Expected true. Got %t.", ft_1_of_1)
	}

	ft_1_of_2 := FT_1_OF_2_EVENT.IsFt1of1()
	if ft_1_of_2 != false {
		t.Errorf("Expected false. Got %t.", ft_1_of_2)
	}

	tech_ft := TECHNICAL_FT_EVENT.IsFt1of1()
	if tech_ft != false {
		t.Errorf("Expected false. Got %t.", tech_ft)
	}
}

func TestFt1of2(t *testing.T) {
	ft_1_of_1 := FT_1_OF_1_EVENT.IsFt1of2()
	if ft_1_of_1 != false {
		t.Errorf("Expected false. Got %t.", ft_1_of_1)
	}

	ft_1_of_2 := FT_1_OF_2_EVENT.IsFt1of2()
	if ft_1_of_2 != true {
		t.Errorf("Expected true. Got %t.", ft_1_of_2)
	}
}

func TestFt2of2(t *testing.T) {
	ft_1_of_1 := FT_1_OF_1_EVENT.IsFt2of2()
	if ft_1_of_1 != false {
		t.Errorf("Expected false. Got %t.", ft_1_of_1)
	}

	ft_1_of_2 := FT_1_OF_2_EVENT.IsFt2of2()
	if ft_1_of_2 != false {
		t.Errorf("Expected false. Got %t.", ft_1_of_2)
	}

	ft_2_of_2 := FT_2_OF_2_EVENT.IsFt2of2()
	if ft_2_of_2 != true {
		t.Errorf("Expected true. Got %t.", ft_2_of_2)
	}
}

func TestFt1of3(t *testing.T) {
	ft_1_of_2 := FT_1_OF_2_EVENT.IsFt1of3()
	if ft_1_of_2 != false {
		t.Errorf("Expected false. Got %t.", ft_1_of_2)
	}

	ft_1_of_3 := FT_1_OF_3_EVENT.IsFt1of3()
	if ft_1_of_3 != true {
		t.Errorf("Expected true. Got %t.", ft_1_of_3)
	}
}

func TestFt3of3(t *testing.T) {
	ft_1_of_3 := FT_1_OF_3_EVENT.IsFt3of3()
	if ft_1_of_3 != false {
		t.Errorf("Expected false. Got %t.", ft_1_of_3)
	}

	ft_3_of_3 := FT_3_OF_3_EVENT.IsFt3of3()
	if ft_3_of_3 != true {
		t.Errorf("Expected true. Got %t.", ft_3_of_3)
	}
}

func TestFirstFT(t *testing.T) {
	ft_1_of_3 := FT_1_OF_3_EVENT.IsFirstFT()
	if ft_1_of_3 != true {
		t.Errorf("Expected true. Got %t.", ft_1_of_3)
	}

	ft_1_of_2 := FT_1_OF_2_EVENT.IsFirstFT()
	if ft_1_of_2 != true {
		t.Errorf("Expected true. Got %t.", ft_1_of_2)
	}

	ft_1_of_1 := FT_1_OF_1_EVENT.IsFirstFT()
	if ft_1_of_1 != true {
		t.Errorf("Expected true. Got %t.", ft_1_of_1)
	}

	ft_3_of_3 := FT_3_OF_3_EVENT.IsFirstFT()
	if ft_3_of_3 != false {
		t.Errorf("Expected false. Got %t.", ft_3_of_3)
	}

	ft_2_of_2 := FT_2_OF_2_EVENT.IsFirstFT()
	if ft_2_of_2 != false {
		t.Errorf("Expected false. Got %t.", ft_2_of_2)
	}

	tech_ft := TECHNICAL_FT_EVENT.IsFirstFT()
	if tech_ft != false {
		t.Errorf("Expected false. Got %t.", tech_ft)
	}
}

func TestTechnicalFT(t *testing.T) {
	ft_1_of_3 := FT_1_OF_3_EVENT.IsTechnicalFT()
	if ft_1_of_3 != false {
		t.Errorf("Expected false. Got %t.", ft_1_of_3)
	}

	ft_1_of_2 := FT_1_OF_2_EVENT.IsTechnicalFT()
	if ft_1_of_2 != false {
		t.Errorf("Expected false. Got %t.", ft_1_of_2)
	}

	ft_1_of_1 := FT_1_OF_1_EVENT.IsTechnicalFT()
	if ft_1_of_1 != false {
		t.Errorf("Expected false. Got %t.", ft_1_of_1)
	}

	ft_3_of_3 := FT_3_OF_3_EVENT.IsTechnicalFT()
	if ft_3_of_3 != false {
		t.Errorf("Expected false. Got %t.", ft_3_of_3)
	}

	ft_2_of_2 := FT_2_OF_2_EVENT.IsTechnicalFT()
	if ft_2_of_2 != false {
		t.Errorf("Expected false. Got %t.", ft_2_of_2)
	}

	tech_ft := TECHNICAL_FT_EVENT.IsTechnicalFT()
	if tech_ft != true {
		t.Errorf("Expected true. Got %t.", tech_ft)
	}
}

func TestTechnicalFoul(t *testing.T) {
	tech := TECHNICAL_FOUL.IsTechnicalFoul()
	if tech != true {
		t.Errorf("Expected true. Got %t.", tech)
	}

	ft_1_of_2 := FT_1_OF_2_EVENT.IsTechnicalFoul()
	if ft_1_of_2 != false {
		t.Errorf("Expected false. Got %t.", ft_1_of_2)
	}
}

func TestDoubleTechnical(t *testing.T) {
	tech := TECHNICAL_FOUL.IsDoubleTechnical()
	if tech != false {
		t.Errorf("Expected false. Got %t.", tech)
	}

	double_tech := DOUBLE_TECHNICAL_FOUL.IsDoubleTechnical()
	if double_tech != true {
		t.Errorf("Expected true. Got %t.", double_tech)
	}
}

func TestEjection(t *testing.T) {
	tech := TECHNICAL_FOUL.IsEjection()
	if tech != false {
		t.Errorf("Expected false. Got %t.", tech)
	}

	ejection := EJECTION_EVENT.IsEjection()
	if ejection != true {
		t.Errorf("Expected true. Got %t.", ejection)
	}
}

func TestGetFoulType(t *testing.T) {
	foul_type := SHOOTING_FOUL_EVENT.GetFoulType()
	expected_foul_type := SHOOTING_FOUL_TYPE_STRING
	if foul_type != expected_foul_type {
		t.Errorf("Expected %s. Got %s.", expected_foul_type, foul_type)
	}

	foul_type = EJECTION_EVENT.GetFoulType()
	expected_foul_type = ""
	if foul_type != expected_foul_type {
		t.Errorf("Expected %s. Got %s.", expected_foul_type, foul_type)
	}
}

func TestGetFouledFgm(t *testing.T) {
	and1_shot, err := SHOOTING_FOUL_EVENT.GetFouledFgm([]PbpEvent{AND1_SHOT_EVENT, SHOOTING_FOUL_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if and1_shot.EventNum != AND1_SHOT_EVENT.EventNum {
		t.Errorf("Expected Eventnum %s. Got %s.", AND1_SHOT_EVENT.EventNum, and1_shot.EventNum)
	}

	and1_shot, err = SHOOTING_FOUL_EVENT.GetFouledFgm([]PbpEvent{SHOOTING_FOUL_EVENT, AND1_SHOT_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if and1_shot.EventNum != 0 {
		t.Errorf("Expected Eventnum %d. Got %d.", 0, and1_shot.EventNum)
	}

	and1_shot, err = SHOOTING_FOUL_EVENT.GetFouledFgm([]PbpEvent{ASSISTED_MADE_2, SHOOTING_FOUL_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if and1_shot.EventNum != 0 {
		t.Errorf("Expected Eventnum %d. Got %d.", 0, and1_shot.EventNum)
	}
}

func TestGetFoulThatResultedInFt(t *testing.T) {
	foul, err := AND1_FT_EVENT.GetFoulThatResultedInFt([]PbpEvent{SHOOTING_FOUL_EVENT, AND1_FT_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if foul.EventNum != SHOOTING_FOUL_EVENT.EventNum {
		t.Errorf("Expected Eventnum %s. Got %s.", SHOOTING_FOUL_EVENT.EventNum, foul.EventNum)
	}

	foul, err = AND1_FT_EVENT.GetFoulThatResultedInFt([]PbpEvent{AND1_FT_EVENT, SHOOTING_FOUL_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if foul.EventNum != SHOOTING_FOUL_EVENT.EventNum {
		t.Errorf("Expected Eventnum %d. Got %d.", SHOOTING_FOUL_EVENT.EventNum, foul.EventNum)
	}

	foul, err = TECHNICAL_FT_EVENT.GetFoulThatResultedInFt([]PbpEvent{TECHNICAL_FT_EVENT, TECHNICAL_FOUL})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if foul.EventNum != 0 {
		t.Errorf("Expected Eventnum %d. Got %d.", 0, foul.EventNum)
	}

	foul, err = FT_1_OF_2_EVENT.GetFoulThatResultedInFt([]PbpEvent{FT_1_OF_2_EVENT, SHOOTING_FOUL_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if foul.EventNum != 0 {
		t.Errorf("Expected Eventnum %d. Got %d.", 0, foul.EventNum)
	}
}

func TestGetLastFoul(t *testing.T) {
	foul, err := AND1_FT_EVENT.GetLastFoul([]PbpEvent{SHOOTING_FOUL_EVENT, AND1_FT_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if foul.EventNum != SHOOTING_FOUL_EVENT.EventNum {
		t.Errorf("Expected Eventnum %s. Got %s.", SHOOTING_FOUL_EVENT.EventNum, foul.EventNum)
	}

	foul, err = AND1_FT_EVENT.GetLastFoul([]PbpEvent{AND1_FT_EVENT, SHOOTING_FOUL_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if foul.EventNum != SHOOTING_FOUL_EVENT.EventNum {
		t.Errorf("Expected Eventnum %d. Got %d.", SHOOTING_FOUL_EVENT.EventNum, foul.EventNum)
	}

	foul, err = TECHNICAL_FT_EVENT.GetLastFoul([]PbpEvent{TECHNICAL_FT_EVENT, TECHNICAL_FOUL})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if foul.EventNum != TECHNICAL_FOUL.EventNum {
		t.Errorf("Expected Eventnum %d. Got %d.", TECHNICAL_FOUL.EventNum, foul.EventNum)
	}

	foul, err = FT_1_OF_2_EVENT.GetLastFoul([]PbpEvent{FT_1_OF_2_EVENT, SHOOTING_FOUL_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if foul.EventNum != 0 {
		t.Errorf("Expected Eventnum %d. Got %d.", 0, foul.EventNum)
	}
}

func TestGetNumberOfFtaForFoul(t *testing.T) {
	fta := SHOOTING_FOUL_EVENT.GetNumberOfFtaForFoul()
	expected_fta := int64(1)
	if fta != expected_fta {
		t.Errorf("Expected %d. Got %d.", expected_fta, fta)
	}

	fta = DOUBLE_TECHNICAL_FOUL.GetNumberOfFtaForFoul()
	expected_fta = int64(0)
	if fta != expected_fta {
		t.Errorf("Expected %d. Got %d.", expected_fta, fta)
	}
}

func TestEndPeriod(t *testing.T) {
	tech := TECHNICAL_FOUL.IsEndOfPeriod()
	if tech != false {
		t.Errorf("Expected false. Got %t.", tech)
	}

	end_period := END_OF_PERIOD_EVENT.IsEndOfPeriod()
	if end_period != true {
		t.Errorf("Expected true. Got %t.", end_period)
	}
}

func TestDelayOfGame(t *testing.T) {
	tech := TECHNICAL_FOUL.IsDelayOfGame()
	if tech != false {
		t.Errorf("Expected false. Got %t.", tech)
	}

	delay_of_game := DELAY_OF_GAME_EVENT.IsDelayOfGame()
	if delay_of_game != true {
		t.Errorf("Expected true. Got %t.", delay_of_game)
	}
}

func TestJumpBall(t *testing.T) {
	tech := TECHNICAL_FOUL.IsJumpBall()
	if tech != false {
		t.Errorf("Expected false. Got %t.", tech)
	}

	jump_ball := JUMP_BALL_EVENT.IsJumpBall()
	if jump_ball != true {
		t.Errorf("Expected true. Got %t.", jump_ball)
	}
}

func TestBlockedShot(t *testing.T) {
	tech := TECHNICAL_FOUL.IsBlockedShot()
	if tech != false {
		t.Errorf("Expected false. Got %t.", tech)
	}

	blocked_shot := BLOCKED_SHOT_EVENT.IsBlockedShot()
	if blocked_shot != true {
		t.Errorf("Expected true. Got %t.", blocked_shot)
	}
}

func TestIsFtFromInboundFoul(t *testing.T) {
	inbound_foul, err := FT_1_OF_1_EVENT.IsFtFromInboundFoul([]PbpEvent{INBOUND_FOUL_EVENT, FT_1_OF_1_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if inbound_foul != true {
		t.Errorf("Expected true. Got %t.", inbound_foul)
	}

	inbound_foul, err = AND1_FT_EVENT.IsFtFromInboundFoul([]PbpEvent{INBOUND_FOUL_EVENT, AND1_FT_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if inbound_foul != false {
		t.Errorf("Expected false. Got %t.", inbound_foul)
	}
}

func TestIsFtFromAwayFromPlayFoul(t *testing.T) {
	away_from_play, err := FT_1_OF_1_EVENT.IsFtFromAwayFromPlayFoul([]PbpEvent{INBOUND_FOUL_EVENT, FT_1_OF_1_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if away_from_play != false {
		t.Errorf("Expected false. Got %t.", away_from_play)
	}

	away_from_play, err = AND1_FT_EVENT.IsFtFromAwayFromPlayFoul([]PbpEvent{INBOUND_FOUL_EVENT, AND1_FT_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if away_from_play != false {
		t.Errorf("Expected false. Got %t.", away_from_play)
	}

	away_from_play, err = FT_1_OF_1_EVENT.IsFtFromAwayFromPlayFoul([]PbpEvent{AWAY_FROM_PLAY_FOUL_EVENT, FT_1_OF_1_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if away_from_play != true {
		t.Errorf("Expected true. Got %t.", away_from_play)
	}
}

func TestIsAnd1Shot(t *testing.T) {
	and1, err := AND1_SHOT_EVENT.IsAnd1Shot([]PbpEvent{SHOOTING_FOUL_EVENT, AND1_SHOT_EVENT, AND1_FT_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if and1 != true {
		t.Errorf("Expected true. Got %t.", and1)
	}

	and1, err = AND1_SHOT_EVENT.IsAnd1Shot([]PbpEvent{SHOOTING_FOUL_EVENT, AND1_SHOT_EVENT, FT_1_OF_1_EVENT})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if and1 != false {
		t.Errorf("Expected false. Got %t.", and1)
	}
}

func TestGetReboundedShot(t *testing.T) {
	rebounded_shot, err := PLAYER_DEFENSIVE_REBOUND.GetReboundedShot([]PbpEvent{MISSED_2, PLAYER_DEFENSIVE_REBOUND})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if rebounded_shot.EventNum != MISSED_2.EventNum {
		t.Errorf("Expected Eventnum %d. Got %d.", MISSED_2.EventNum, rebounded_shot.EventNum)
	}

	rebounded_shot, err = PLAYER_DEFENSIVE_REBOUND.GetReboundedShot([]PbpEvent{PLAYER_DEFENSIVE_REBOUND, MISSED_2})
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if rebounded_shot.EventNum != 0 {
		t.Errorf("Expected Eventnum %d. Got %d.", 0, rebounded_shot.EventNum)
	}
}

func TestIsDefensiveRebound(t *testing.T) {
	def_reb := PLAYER_DEFENSIVE_REBOUND.IsDefensiveRebound(MISSED_2)
	if def_reb != true {
		t.Errorf("Expected true. Got %t.", def_reb)
	}

	def_reb = PLAYER_OFFENSIVE_REBOUND.IsDefensiveRebound(MISSED_2)
	if def_reb != false {
		t.Errorf("Expected false. Got %t.", def_reb)
	}
}

func TestIsTeamRebound(t *testing.T) {
	team_reb := PLAYER_DEFENSIVE_REBOUND.IsTeamRebound()
	if team_reb != false {
		t.Errorf("Expected false. Got %t.", team_reb)
	}

	team_reb = TEAM_OFFENSIVE_REBOUND.IsTeamRebound()
	if team_reb != true {
		t.Errorf("Expected true. Got %t.", team_reb)
	}
}

func TestIsSelfRebound(t *testing.T) {
	self_reb := PLAYER_DEFENSIVE_REBOUND.IsSelfRebound(MISSED_2)
	if self_reb != false {
		t.Errorf("Expected false. Got %t.", self_reb)
	}
	self_reb = REBOUND_EVENT_BEFORE_PUTBACK.IsSelfRebound(MISSED_2)
	if self_reb != true {
		t.Errorf("Expected true. Got %t.", self_reb)
	}
}

func TestIsTrackedEvent(t *testing.T) {
	tracked := PLAYER_DEFENSIVE_REBOUND.IsTrackedEvent()
	if tracked != true {
		t.Errorf("Expected true. Got %t.", tracked)
	}
	tracked = STEAL_EVENT.IsTrackedEvent()
	if tracked != true {
		t.Errorf("Expected true. Got %t.", tracked)
	}
	tracked = MISSED_2.IsTrackedEvent()
	if tracked != true {
		t.Errorf("Expected true. Got %t.", tracked)
	}
	tracked = FT_1_OF_1_EVENT.IsTrackedEvent()
	if tracked != true {
		t.Errorf("Expected true. Got %t.", tracked)
	}
	tracked = SUBSTITUTION_EVENT.IsTrackedEvent()
	if tracked != false {
		t.Errorf("Expected false. Got %t.", tracked)
	}
}

func TestGenerateLineupIds(t *testing.T) {
	event := PbpEvent{
		CurrentPlayers: map[int64][]int64{1: []int64{2, 45, 67, 909, 14}, 2: []int64{87, 65, 45, 121, 32}},
	}
	lineup_ids := event.GenerateLineupIds()
	expected_lineup_id := "14-2-45-67-909"
	if lineup_ids[int64(1)] != expected_lineup_id {
		t.Errorf("Expected %s. Got %s.", expected_lineup_id, lineup_ids[int64(1)])
	}
	expected_lineup_id = "121-32-45-65-87"
	if lineup_ids[int64(2)] != expected_lineup_id {
		t.Errorf("Expected %s. Got %s.", expected_lineup_id, lineup_ids[int64(2)])
	}

	empty_event := PbpEvent{}
	lineup_ids = empty_event.GenerateLineupIds()
	if len(lineup_ids) != 0 {
		t.Errorf("Expected 0. Got %d.", len(lineup_ids))
	}
}

func TestGetShotDistance(t *testing.T) {
	distance := MISSED_2.GetShotDistance()
	expected_distance := int64(1)
	if int64(distance) != expected_distance {
		t.Errorf("Expected %d. Got %d.", expected_distance, int64(distance))
	}

	distance = MISSED_3.GetShotDistance()
	expected_distance = int64(27)
	if int64(distance) != expected_distance {
		t.Errorf("Expected %d. Got %d.", expected_distance, int64(distance))
	}
}

func TestIsCorner3(t *testing.T) {
	corner3 := MISSED_3.IsCorner3()
	if corner3 != false {
		t.Errorf("Expected false. Got %t.", corner3)
	}

	corner3 = MISSED_CORNER_3.IsCorner3()
	if corner3 != true {
		t.Errorf("Expected true. Got %t.", corner3)
	}
}
