package nba_pbp

import (
	"testing"
)

func TestGetPbpResponseData(t *testing.T) {
	game_id := "0021700578"
	period := int64(1)
	data, err := GetPbpResponseData(game_id, period)
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if data.Data.GameId != game_id {
		t.Errorf("Wrong Game Id. Expected %s. Got %s.", game_id, data.Data.GameId)
	}
}

func TestGetGameDetailResponseData(t *testing.T) {
	game_id := "0021700578"
	data, err := GetGameDetailResponseData(game_id)
	if err != nil {
		t.Errorf("Expected err=nil, got %#v", err)
	}
	if data.Data.GameId != game_id {
		t.Errorf("Wrong Game Id. Expected %s. Got %s.", game_id, data.Data.GameId)
	}
}
