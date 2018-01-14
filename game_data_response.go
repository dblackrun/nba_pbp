package nba_pbp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PbpDataResponse struct {
	Data PeriodData `json:"g"`
}

func GetPbpResponseData(game_id string, period int64) (PbpDataResponse, error) {
	var data PbpDataResponse
	var client = &http.Client{Timeout: 10 * time.Second}
	year := GetYearForGame(game_id)
	url := fmt.Sprintf("https://data.nba.com/data/10s/v2015/json/mobile_teams/nba/%s/scores/pbp/%s_%d_pbp.json", year, game_id, period)
	r, err := client.Get(url)
	if err != nil {
		return data, err
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&data)
	if err != nil {
		return data, err
	}

	return data, err
}

type GameDetailResponse struct {
	Data GameDetails `json:"g"`
}

func GetGameDetailResponseData(game_id string) (GameDetailResponse, error) {
	var data GameDetailResponse
	var client = &http.Client{Timeout: 10 * time.Second}
	year := GetYearForGame(game_id)
	url := fmt.Sprintf("http://data.nba.com/data/10s/v2015/json/mobile_teams/nba/%s/scores/gamedetail/%s_gamedetail.json", year, game_id)
	r, err := client.Get(url)
	if err != nil {
		return data, err
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&data)
	if err != nil {
		return data, err
	}

	return data, err
}
