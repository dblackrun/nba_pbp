package nba_pbp

type GameDetails struct {
	Mid             int64       `json:"mid"`
	GameId          string      `json:"gid"`
	GameCode        string      `json:"gcode"`
	GameDate        string      `json:"gdte"`
	Status          string      `json:"stt"`
	Period          int64       `json:"p"`
	VisitorTeamData TeamDetails `json:"vls"`
	HomeTeamData    TeamDetails `json:"hls"`
}

type TeamDetails struct {
	TeamId           int64        `json:"tid"`
	TeamAbbreviation string       `json:"ta"`
	Score            int64        `json:"s"`
	Players          []PlayerData `json:"pstsg"`
}

type PlayerData struct {
	PlayerId      int64  `json:"pid"`
	FirstName     string `json:"fn"`
	LastName      string `json:"ln"`
	SecondsPlayed int64  `json:"totsec"`
}

func DidPlayerPlayInGame(player_id int64, game_details GameDetails) (bool, int64) {
	// checks if player played in game, also returns team id
	for _, player := range game_details.VisitorTeamData.Players {
		if player_id == player.PlayerId && player.SecondsPlayed > int64(0) {
			return true, game_details.VisitorTeamData.TeamId
		}
	}

	for _, player := range game_details.HomeTeamData.Players {
		if player_id == player.PlayerId && player.SecondsPlayed > int64(0) {
			return true, game_details.HomeTeamData.TeamId
		}
	}

	return false, int64(0)
}

func SwapTeamIdForGame(team_id int64, game_details GameDetails) int64 {
	if team_id == game_details.HomeTeamData.TeamId {
		return game_details.VisitorTeamData.TeamId
	}
	if team_id == game_details.VisitorTeamData.TeamId {
		return game_details.HomeTeamData.TeamId
	}
	return int64(0)
}
