package nba_pbp

func IntInSlice(a int64, list []int64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetYearForGame(game_id string) string {
	/*
		gets season from game id
		4th and 5th characters in game id represent season year
		ex. game id 0021700579 returns 2017
	*/
	if string(game_id[3]) == "9" {
		if string(game_id[4]) == "9" {
			return "1999"
		} else {
			return "19" + string(game_id[3]) + string(game_id[4])
		}
	}
	return "20" + string(game_id[3]) + string(game_id[4])
}
