## Features

* Adds lineups on floor for each play by play event
* Adds detailed data for each possession including start time, end time, score margin, how the previous possession ended, second chance time, offensive rebounds
* Sums up team, opponent, player and lineup stats for list of possessions
* Shooting, assist and rebound stats broken down by shot zone


## Install

	go get github.com/dblackrun/nba_pbp

## Usage
```
package main

import (
  "fmt"
  "github.com/dblackrun/nba_pbp"
)

func main() {
  game_id := "0021700001"

  if game_details, err := nba_pbp.GetGameDetailResponseData(game_id); err != nil {
    panic(err)
  } else {
    period := int64(1)
    period_data, err := nba_pbp.GetPbpResponseData(game_id, period)
    if err != nil {
      panic(err)
    }
    starters, err := period_data.Data.GetStarters(game_details.Data)
    if err != nil {
      panic(err)
    }
    err = period_data.Data.AddCurrentPlayersAndTimeElapsed(starters)
    if err != nil {
      panic(err)
    }
    possession_details, err := period_data.Data.ParsePossessions(game_details.Data)
    if err != nil {
      panic(err)
    }
    for i, _ := range possession_details {
      if i != 0 {
        possession_details[i].AddPossessionStats(period_data.Data.Events, possession_details[i-1])
      } else {
        possession_details[i].AddPossessionStats(period_data.Data.Events, nba_pbp.PossessionDetails{})
      }
      possession_details[i].AddPlayerStatsForPossession()
    }
    team_stats, opponent_stats, player_stats, lineup_stats, lineup_opponent_stats, err := nba_pbp.SumPossessionStats(possession_details, game_details.Data.VisitorTeamData.TeamId)
    if err != nil {
      panic(err)
    }
    fmt.Println(team_stats)
    fmt.Println(opponent_stats)
    fmt.Println(player_stats)
    fmt.Println(lineup_stats)
    fmt.Println(lineup_opponent_stats)
  }
}
```
