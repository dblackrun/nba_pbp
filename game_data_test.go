package nba_pbp

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func helperLoadGameDetails(t *testing.T) GameDetails {
	bytes, err := ioutil.ReadFile("testdata/game_details.json")
	if err != nil {
		t.Fatal(err)
	}

	var game_details GameDetails

	json.Unmarshal(bytes, &game_details)
	return game_details
}

func TestDidPlayerPlayInGame(t *testing.T) {
	game_details := helperLoadGameDetails(t)
	played, team_id := DidPlayerPlayInGame(int64(203110), game_details)
	if played != true {
		t.Errorf("Expected true. Got %t.", played)
	}
	expected_team_id := int64(1610612744)
	if team_id != expected_team_id {
		t.Errorf("Expected %d. Got %d.", expected_team_id, team_id)
	}

	played, team_id = DidPlayerPlayInGame(int64(101150), game_details)
	if played != true {
		t.Errorf("Expected true. Got %t.", played)
	}
	expected_team_id = int64(1610612746)
	if team_id != expected_team_id {
		t.Errorf("Expected %d. Got %d.", expected_team_id, team_id)
	}

	played, team_id = DidPlayerPlayInGame(int64(1628449), game_details)
	if played != false {
		t.Errorf("Expected false. Got %t.", played)
	}
	expected_team_id = int64(0)
	if team_id != expected_team_id {
		t.Errorf("Expected %d. Got %d.", expected_team_id, team_id)
	}
}

func TestSwapTeamIdForGame(t *testing.T) {
	game_details := helperLoadGameDetails(t)
	team_id := SwapTeamIdForGame(game_details.HomeTeamData.TeamId, game_details)
	if team_id != game_details.VisitorTeamData.TeamId {
		t.Errorf("Expected %d. Got %d.", game_details.VisitorTeamData.TeamId, team_id)
	}
}
