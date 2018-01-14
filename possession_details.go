package nba_pbp

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	AT_RIM_STRING                       = "AtRim"
	SHORT_MID_RANGE_STRING              = "ShortMidRange"
	LONG_MID_RANGE_STRING               = "LongMidRange"
	CORNER_3_STRING                     = "Corner3"
	ARC_3_STRING                        = "Arc3"
	FREE_THROW_STRING                   = "FT"
	OFF_AT_RIM_MAKE_STRING              = "Off" + AT_RIM_STRING + "Make"
	OFF_AT_RIM_MISS_STRING              = "Off" + AT_RIM_STRING + "Miss"
	OFF_AT_RIM_BLOCK_STRING             = "Off" + AT_RIM_STRING + "Block"
	OFF_SHORT_MID_RANGE_MAKE_STRING     = "Off" + SHORT_MID_RANGE_STRING + "Make"
	OFF_SHORT_MID_RANGE_MISS_STRING     = "Off" + SHORT_MID_RANGE_STRING + "Miss"
	OFF_SHORT_MID_RANGE_BLOCK_STRING    = "Off" + SHORT_MID_RANGE_STRING + "Block"
	OFF_LONG_MID_RANGE_MAKE_STRING      = "Off" + LONG_MID_RANGE_STRING + "Make"
	OFF_LONG_MID_RANGE_MISS_STRING      = "Off" + LONG_MID_RANGE_STRING + "Miss"
	OFF_LONG_MID_RANGE_BLOCK_STRING     = "Off" + LONG_MID_RANGE_STRING + "Block"
	OFF_ARC_3_MAKE_STRING               = "Off" + ARC_3_STRING + "Make"
	OFF_ARC_3_MISS_STRING               = "Off" + ARC_3_STRING + "Miss"
	OFF_ARC_3_BLOCK_STRING              = "Off" + ARC_3_STRING + "Block"
	OFF_CORNER_3_MAKE_STRING            = "Off" + CORNER_3_STRING + "Make"
	OFF_CORNER_3_MISS_STRING            = "Off" + CORNER_3_STRING + "Miss"
	OFF_CORNER_3_BLOCK_STRING           = "Off" + CORNER_3_STRING + "Block"
	OFF_MADE_FT_STRING                  = "Off" + FREE_THROW_STRING + "Make"
	OFF_MISSED_FT_STRING                = "Off" + FREE_THROW_STRING + "Miss"
	OFF_LIVE_BALL_TURNOVER_STRING       = "OffLiveBallTurnover"
	OFF_DEADBALL_STRING                 = "OffDeadball"
	OFF_TIMEOUT_STRING                  = "OffTimeout"
	AT_RIM_CUTOFF                       = float64(4)
	SHORT_MID_RANGE_CUTOFF              = float64(14)
	MISSES_KEY                          = "Misses"
	BLOCKS_KEY                          = "Blocks"
	BLOCKED_KEY                         = "Blocked"
	MAKES_KEY                           = "Makes"
	ASSISTED_KEY                        = "Assisted"
	ASSIST_KEY                          = "Assists"
	REBOUND_KEY                         = "Rebound"
	DEADBALL_TURNOVER_KEY               = "DeadballTurnovers"
	LIVEBALL_TURNOVER_KEY               = "LiveballTurnovers"
	STEALS_KEY                          = "Steals"
	TECHNICAL_FT_KEY                    = "TechnicalFreeThrows"
	THREE_POINT_AND1_KEY                = "3ptAnd1s"
	TWO_POINT_AND1_KEY                  = "2ptAnd1s"
	PENALTY_FREE_THROW_KEY              = "PenaltyFreeThrowTrips"
	THREE_POINT_SHOOTING_FOUL_TRIPS_KEY = "3ptShootingFoulFreeThrowTrips"
	TWO_POINT_SHOOTING_FOUL_TRIPS_KEY   = "2ptShootingFoulFreeThrowTrips"
	AWAY_FROM_PLAY_FOUL_TRIPS_KEY       = "1ShotAwayFromPlayFreeThrowTrips"
	SECONDS_KEY                         = "Seconds"
	CHANCES_STRING                      = "Chances"
	OFFENSIVE_REBOUND_PREIX             = "Off"
	DEFENSIVE_REBOUND_PREIX             = "Def"
	OFFENSIVE_POSSESSIONS_KEY           = "OffensivePossessions"
	DEFENSIVE_POSSESSIONS_KEY           = "DefensivePossessions"
)

type PossessionDetails struct {
	GameId                                string
	Period                                int64
	PossessionNumber                      int64
	OffenseTeamId                         int64
	DefenseTeamId                         int64
	OffenseLineupId                       string
	DefenseLineupId                       string
	PossessionStartTime                   float64
	PossessionEndTime                     float64
	PreviousPossessionEndEventNum         int64
	PossessionEndEventNum                 int64
	PossessionStartType                   string
	PossessionStartScoreDifferential      int64
	OffensiveRebounds                     int64
	SecondChanceTime                      float64
	PreviousPossessionEndShooterPlayerId  int64 // for both makes and misses
	PreviousPossessionEndReboundPlayerId  int64 // only for misses
	PreviousPossessionEndTurnoverPlayerId int64 // only for live ball turnovers
	PreviousPossessionEndStealPlayerId    int64 // only for live ball turnovers
	PlayerStats                           map[int64]map[string]map[string]map[int64]map[string]int64
	Events                                []PbpEvent
	PreviousPossessionEvents              []PbpEvent
}

func (possession *PossessionDetails) GetAllEventsForPossession(period_events []PbpEvent) ([]PbpEvent, error) {
	var possession_events []PbpEvent
	add_events := false
	for _, event := range period_events {
		if add_events {
			possession_events = append(possession_events, event)
		}
		if event.EventNum == possession.PreviousPossessionEndEventNum {
			add_events = true
		}
		if event.EventNum == possession.PossessionEndEventNum {
			return possession_events, nil
		}
	}
	return possession_events, fmt.Errorf("Never got to end event num for GameId: %d Period: %d Possession Number: %d", possession.GameId, possession.Period, possession.PossessionNumber)
}

func (possession *PossessionDetails) AddPossessionStats(period_events []PbpEvent, previous_possession PossessionDetails) error {
	/*
			  Adds the following possession stats to possession details:
			    - PossessionStartType
			    - OffensiveRebounds
		      - SecondChanceTime
			    - PreviousPossessionEndShooterPlayerId
			    - PreviousPossessionEndReboundPlayerId
			    - PreviousPossessionEndTurnoverPlayerId
			    - PreviousPossessionEndStealPlayerId
			    - Events
	*/
	possession_events, err := possession.GetAllEventsForPossession(period_events)
	if err != nil {
		return err
	}

	possession.Events = possession_events

	has_timeout := false
	for _, event := range possession_events {
		if event.IsTimeout() {
			has_timeout = true
		}
		if event.IsRebound() {
			rebounded_shot, err := event.GetReboundedShot(period_events)
			if err != nil {
				return err
			}
			if !event.IsDefensiveRebound(rebounded_shot) {
				possession.OffensiveRebounds += 1
				if possession.OffensiveRebounds == 1 {
					// add second chance time
					if seconds_remaining, err := event.GetSecondsRemaining(); err != nil {
						return err
					} else {
						possession.SecondChanceTime = seconds_remaining - possession.PossessionEndTime
					}
				}
			}
		}
	}

	if possession.PossessionNumber == 1 {
		possession.PossessionStartType = OFF_DEADBALL_STRING
	} else {
		previous_possession_events, err := previous_possession.GetAllEventsForPossession(period_events)
		if err != nil {
			return err
		}
		possession.PreviousPossessionEvents = previous_possession_events
		previous_possession_ending_event := previous_possession_events[len(previous_possession_events)-1]

		if previous_possession_ending_event.IsMadeFG() {
			if previous_possession_ending_event.Is3PointShot() {
				if previous_possession_ending_event.IsCorner3() {
					possession.PossessionStartType = OFF_CORNER_3_MAKE_STRING
				} else {
					possession.PossessionStartType = OFF_ARC_3_MAKE_STRING
				}
			} else {
				shot_distance := previous_possession_ending_event.GetShotDistance()
				if shot_distance <= AT_RIM_CUTOFF {
					possession.PossessionStartType = OFF_AT_RIM_MAKE_STRING
				} else if shot_distance <= SHORT_MID_RANGE_CUTOFF {
					possession.PossessionStartType = OFF_SHORT_MID_RANGE_MAKE_STRING
				} else {
					possession.PossessionStartType = OFF_LONG_MID_RANGE_MAKE_STRING
				}
			}
			if !has_timeout {
				possession.PreviousPossessionEndShooterPlayerId = previous_possession_ending_event.PlayerId
			}
		} else if previous_possession_ending_event.IsMadeFT() {
			possession.PossessionStartType = OFF_MADE_FT_STRING
			if !has_timeout {
				possession.PreviousPossessionEndShooterPlayerId = previous_possession_ending_event.PlayerId
			}
		} else if previous_possession_ending_event.IsSteal() {
			possession.PossessionStartType = OFF_LIVE_BALL_TURNOVER_STRING
			if !has_timeout {
				possession.PreviousPossessionEndTurnoverPlayerId = previous_possession_ending_event.PlayerId
			}
			steal_player_id, err := previous_possession_ending_event.GetOPlayerIdInt()
			if err != nil {
				return err
			}
			possession.PreviousPossessionEndStealPlayerId = steal_player_id
		} else if previous_possession_ending_event.IsTurnover() {
			possession.PossessionStartType = OFF_DEADBALL_STRING
		} else if previous_possession_ending_event.IsRebound() {
			if previous_possession_ending_event.IsTeamRebound() {
				possession.PossessionStartType = OFF_DEADBALL_STRING
			} else {
				rebounded_shot, err := previous_possession_ending_event.GetReboundedShot(period_events)
				if err != nil {
					return err
				}
				if rebounded_shot.IsMissedFT() {
					possession.PossessionStartType = OFF_MISSED_FT_STRING
				} else if rebounded_shot.Is3PointShot() {
					if rebounded_shot.IsCorner3() {
						if rebounded_shot.IsBlockedShot() {
							possession.PossessionStartType = OFF_CORNER_3_BLOCK_STRING
						} else {
							possession.PossessionStartType = OFF_CORNER_3_MISS_STRING
						}
					} else {
						if rebounded_shot.IsBlockedShot() {
							possession.PossessionStartType = OFF_ARC_3_BLOCK_STRING
						} else {
							possession.PossessionStartType = OFF_ARC_3_MISS_STRING
						}
					}
				} else {
					shot_distance := rebounded_shot.GetShotDistance()
					if shot_distance <= AT_RIM_CUTOFF {
						if rebounded_shot.IsBlockedShot() {
							possession.PossessionStartType = OFF_AT_RIM_BLOCK_STRING
						} else {
							possession.PossessionStartType = OFF_AT_RIM_MISS_STRING
						}
					} else if shot_distance <= SHORT_MID_RANGE_CUTOFF {
						if rebounded_shot.IsBlockedShot() {
							possession.PossessionStartType = OFF_SHORT_MID_RANGE_BLOCK_STRING
						} else {
							possession.PossessionStartType = OFF_SHORT_MID_RANGE_MISS_STRING
						}
					} else {
						if rebounded_shot.IsBlockedShot() {
							possession.PossessionStartType = OFF_LONG_MID_RANGE_BLOCK_STRING
						} else {
							possession.PossessionStartType = OFF_LONG_MID_RANGE_MISS_STRING
						}
					}
				}
				if !has_timeout {
					possession.PreviousPossessionEndShooterPlayerId = rebounded_shot.PlayerId
					possession.PreviousPossessionEndReboundPlayerId = previous_possession_ending_event.PlayerId
				}
			}
		} else if previous_possession_ending_event.IsJumpBall() {
			// jump balls tipped out of bounds have no epid and should be off deadball
			if epid, err := previous_possession_ending_event.GetEPlayerIdInt(); err != nil {
				return nil
			} else if epid != 0 {
				possession.PossessionStartType = OFF_LIVE_BALL_TURNOVER_STRING
			} else {
				possession.PossessionStartType = OFF_DEADBALL_STRING
			}
		}
	}
	if has_timeout {
		possession.PossessionStartType = OFF_TIMEOUT_STRING
	}
	return nil
}

func (possession *PossessionDetails) AddPlayerStatsForPossession() error {
	/*
	  stats nested map format: {TeamId:{LineupId:{OpponentLineupId:{PlayerId:{StatKey:StatValue}}}}}
	*/
	stats := make(map[int64]map[string]map[string]map[int64]map[string]int64)
	o_team_id := possession.OffenseTeamId
	d_team_id := possession.DefenseTeamId
	stats[o_team_id] = make(map[string]map[string]map[int64]map[string]int64)
	stats[d_team_id] = make(map[string]map[string]map[int64]map[string]int64)
	for i, event := range possession.Events {
		// add keys to map if they are needed
		lineup_ids := event.GenerateLineupIds()
		o_lineup_id := lineup_ids[o_team_id]
		d_lineup_id := lineup_ids[d_team_id]
		if _, exists := stats[o_team_id][o_lineup_id]; !exists {
			stats[o_team_id][o_lineup_id] = make(map[string]map[int64]map[string]int64)
		}
		if _, exists := stats[o_team_id][o_lineup_id][d_lineup_id]; !exists {
			stats[o_team_id][o_lineup_id][d_lineup_id] = make(map[int64]map[string]int64)
		}
		if _, exists := stats[d_team_id][d_lineup_id]; !exists {
			stats[d_team_id][d_lineup_id] = make(map[string]map[int64]map[string]int64)
		}
		if _, exists := stats[d_team_id][d_lineup_id][o_lineup_id]; !exists {
			stats[d_team_id][d_lineup_id][o_lineup_id] = make(map[int64]map[string]int64)
		}
		// seconds played is SecondsToNextEvent
		for team_id, players := range event.CurrentPlayers {
			var lineup_id, opponent_lineup_id string
			if team_id == o_team_id {
				lineup_id = o_lineup_id
				opponent_lineup_id = d_lineup_id
			} else {
				lineup_id = d_lineup_id
				opponent_lineup_id = o_lineup_id
			}
			for _, player_id := range players {
				if _, exists := stats[team_id][lineup_id][opponent_lineup_id][player_id]; !exists {
					stats[team_id][lineup_id][opponent_lineup_id][player_id] = make(map[string]int64)
				}
				// add 0.5 since int64 conversion rounds down
				stats[team_id][lineup_id][opponent_lineup_id][player_id][SECONDS_KEY] += int64(event.SecondsToNextEvent + 0.5)
				if i == 0 && possession.PossessionNumber == 1 {
					// if first possession need to also add time since start of period
					stats[team_id][lineup_id][opponent_lineup_id][player_id][SECONDS_KEY] += int64(event.SecondsSinceLastEvent + 0.5)
				}
			}
		}
		if event.IsTrackedEvent() {
			if event.IsMissedFG() || event.IsMadeFG() {
				player_id := event.PlayerId
				if _, exists := stats[o_team_id][o_lineup_id][d_lineup_id][player_id]; !exists {
					stats[o_team_id][o_lineup_id][d_lineup_id][player_id] = make(map[string]int64)
				}

				var shot_range string
				if event.Is3PointShot() {
					if event.IsCorner3() {
						shot_range = CORNER_3_STRING
					} else {
						shot_range = ARC_3_STRING
					}
				} else {
					shot_distance := event.GetShotDistance()
					if shot_distance <= AT_RIM_CUTOFF {
						shot_range = AT_RIM_STRING
					} else if shot_distance <= SHORT_MID_RANGE_CUTOFF {
						shot_range = SHORT_MID_RANGE_STRING
					} else {
						shot_range = LONG_MID_RANGE_STRING
					}
				}

				if event.IsMissedFG() {
					stats[o_team_id][o_lineup_id][d_lineup_id][player_id][shot_range+MISSES_KEY] += 1

					if event.IsBlockedShot() {
						if block_player_id, err := event.GetOPlayerIdInt(); err != nil {
							return err
						} else {
							if _, exists := stats[d_team_id][d_lineup_id][o_lineup_id][block_player_id]; !exists {
								stats[d_team_id][d_lineup_id][o_lineup_id][block_player_id] = make(map[string]int64)
							}
							stats[d_team_id][d_lineup_id][o_lineup_id][block_player_id][shot_range+BLOCKS_KEY] += 1
							stats[o_team_id][o_lineup_id][d_lineup_id][player_id][BLOCKED_KEY+shot_range] += 1
						}
					}
				} else {
					stats[o_team_id][o_lineup_id][d_lineup_id][player_id][shot_range+MAKES_KEY] += 1

					if event.IsAssistedShot() {
						if assist_player_id, err := event.GetEPlayerIdInt(); err != nil {
							return err
						} else {
							if _, exists := stats[o_team_id][o_lineup_id][d_lineup_id][assist_player_id]; !exists {
								stats[o_team_id][o_lineup_id][d_lineup_id][assist_player_id] = make(map[string]int64)
							}
							stats[o_team_id][o_lineup_id][d_lineup_id][assist_player_id][shot_range+ASSIST_KEY] += 1
							stats[o_team_id][o_lineup_id][d_lineup_id][player_id][ASSISTED_KEY+shot_range] += 1
						}
					}
				}
			} else if event.IsTurnover() {
				player_id := event.PlayerId
				if _, exists := stats[o_team_id][o_lineup_id][d_lineup_id][player_id]; !exists {
					stats[o_team_id][o_lineup_id][d_lineup_id][player_id] = make(map[string]int64)
				}
				if event.IsSteal() {
					if steal_player_id, err := event.GetOPlayerIdInt(); err != nil {
						return err
					} else {
						if _, exists := stats[d_team_id][d_lineup_id][o_lineup_id][steal_player_id]; !exists {
							stats[d_team_id][d_lineup_id][o_lineup_id][steal_player_id] = make(map[string]int64)
						}
						stats[d_team_id][d_lineup_id][o_lineup_id][steal_player_id][STEALS_KEY] += 1
						stats[o_team_id][o_lineup_id][d_lineup_id][player_id][LIVEBALL_TURNOVER_KEY] += 1
					}
				} else {
					stats[o_team_id][o_lineup_id][d_lineup_id][player_id][DEADBALL_TURNOVER_KEY] += 1
				}
			} else if event.IsFoul() {
				foul_type := event.GetFoulType()
				if foul_type != "" && !event.IsTechnicalFoul() && !event.IsDoubleTechnical() {
					player_id := event.PlayerId
					team_id := event.TeamId

					// foul may be committed by either offensive team or defensive team so need to get team ids right
					var lineup_id, opponent_lineup_id string
					var opponent_team_id int64
					if team_id == o_team_id {
						opponent_team_id = d_team_id
						lineup_id = o_lineup_id
						opponent_lineup_id = d_lineup_id
					} else {
						opponent_team_id = o_team_id
						lineup_id = d_lineup_id
						opponent_lineup_id = o_lineup_id
					}

					if _, exists := stats[team_id][lineup_id][opponent_lineup_id][player_id]; !exists {
						stats[team_id][lineup_id][opponent_lineup_id][player_id] = make(map[string]int64)
					}
					stats[team_id][lineup_id][opponent_lineup_id][player_id][foul_type] += 1

					// opid is player who drew foul, for double foul they are commiting a foul
					var o_player_id_stat_key string
					if foul_type == DOUBLE_FOUL_TYPE_STRING {
						o_player_id_stat_key = foul_type
					} else {
						o_player_id_stat_key = foul_type + "Drawn"
					}

					// opid is player who drew foul, for double foul they are commiting a foul
					if o_player_id, err := event.GetOPlayerIdInt(); err != nil {
						return err
					} else if o_player_id != 0 {
						if _, exists := stats[opponent_team_id][opponent_lineup_id][lineup_id][o_player_id]; !exists {
							stats[opponent_team_id][opponent_lineup_id][lineup_id][o_player_id] = make(map[string]int64)
						}
						stats[opponent_team_id][opponent_lineup_id][lineup_id][o_player_id][o_player_id_stat_key] += 1
					}
				}
			} else if event.IsMadeFT() || event.IsMissedFT() {
				// check for foul that resulted in FTs and use lineups on floor for foul
				if last_foul, err := event.GetLastFoul(append(possession.Events, possession.PreviousPossessionEvents...)); err != nil {
					return err
				} else {
					foul_lineup_ids := last_foul.GenerateLineupIds()
					foul_o_lineup_id := foul_lineup_ids[o_team_id]
					foul_d_lineup_id := foul_lineup_ids[d_team_id]
					// if FT is last event for possession, update lineup ids
					if i == len(possession.Events)-1 {
						possession.OffenseLineupId = foul_o_lineup_id
						possession.DefenseLineupId = foul_d_lineup_id
					}
					if _, exists := stats[o_team_id][foul_o_lineup_id]; !exists {
						stats[o_team_id][foul_o_lineup_id] = make(map[string]map[int64]map[string]int64)
					}
					if _, exists := stats[o_team_id][foul_o_lineup_id][foul_d_lineup_id]; !exists {
						stats[o_team_id][foul_o_lineup_id][foul_d_lineup_id] = make(map[int64]map[string]int64)
					}
					if _, exists := stats[d_team_id][foul_d_lineup_id]; !exists {
						stats[d_team_id][foul_d_lineup_id] = make(map[string]map[int64]map[string]int64)
					}
					if _, exists := stats[d_team_id][foul_d_lineup_id][foul_o_lineup_id]; !exists {
						stats[d_team_id][foul_d_lineup_id][foul_o_lineup_id] = make(map[int64]map[string]int64)
					}
					player_id := event.PlayerId
					if _, exists := stats[o_team_id][foul_o_lineup_id][foul_d_lineup_id][player_id]; !exists {
						stats[o_team_id][foul_o_lineup_id][foul_d_lineup_id][player_id] = make(map[string]int64)
					}
					if event.IsMadeFT() {
						stats[o_team_id][foul_o_lineup_id][foul_d_lineup_id][player_id][FREE_THROW_STRING+MAKES_KEY] += 1
					} else if event.IsMissedFT() {
						stats[o_team_id][foul_o_lineup_id][foul_d_lineup_id][player_id][FREE_THROW_STRING+MISSES_KEY] += 1
					}
				}
			} else if event.IsRebound() {
				seconds_remaining, err := event.GetSecondsRemaining()
				if err != nil {
					return err
				}
				if !(seconds_remaining <= 0.1 && event.IsTeamRebound()) {
					if rebounded_shot, err := event.GetReboundedShot(possession.Events); err != nil {
						return err
					} else {
						var missed_shot_type string
						if rebounded_shot.IsMissedFT() {
							missed_shot_type = FREE_THROW_STRING
						} else if rebounded_shot.Is3PointShot() {
							if rebounded_shot.IsCorner3() {
								missed_shot_type = CORNER_3_STRING
							} else {
								missed_shot_type = ARC_3_STRING
							}
						} else {
							shot_distance := rebounded_shot.GetShotDistance()
							if shot_distance <= AT_RIM_CUTOFF {
								missed_shot_type = AT_RIM_STRING
							} else if shot_distance <= SHORT_MID_RANGE_CUTOFF {
								missed_shot_type = SHORT_MID_RANGE_STRING
							} else {
								missed_shot_type = LONG_MID_RANGE_STRING
							}
						}

						if rebounded_shot.IsBlockedShot() {
							missed_shot_type = "Blocked" + missed_shot_type
						}

						def_reb := event.IsDefensiveRebound(rebounded_shot)
						var rebound_team_id int64
						var rebound_lineup_id, rebound_opponent_lineup_id string
						if def_reb {
							rebound_team_id = d_team_id
							rebound_lineup_id = d_lineup_id
							rebound_opponent_lineup_id = o_lineup_id
						} else {
							rebound_team_id = o_team_id
							rebound_lineup_id = o_lineup_id
							rebound_opponent_lineup_id = d_lineup_id
						}

						player_id := event.PlayerId
						if _, exists := stats[rebound_team_id][rebound_lineup_id][rebound_opponent_lineup_id][player_id]; !exists {
							stats[rebound_team_id][rebound_lineup_id][rebound_opponent_lineup_id][player_id] = make(map[string]int64)
						}
						stats[rebound_team_id][rebound_lineup_id][rebound_opponent_lineup_id][player_id][missed_shot_type+REBOUND_KEY] += 1
					}
				}
			}

			if event.IsFirstFT() {
				// for determining source of fts, makes/misses already added
				if foul_event, err := event.GetFoulThatResultedInFt(possession.Events); err != nil {
					return err
				} else {
					player_id := event.PlayerId
					if _, exists := stats[o_team_id][o_lineup_id][d_lineup_id][player_id]; !exists {
						stats[o_team_id][o_lineup_id][d_lineup_id][player_id] = make(map[string]int64)
					}
					if foul_event.EventNum == 0 {
						// free throws to begin quarter, assume they are technical fts from technical between quarters
						stats[o_team_id][o_lineup_id][d_lineup_id][player_id][TECHNICAL_FT_KEY] += 1
					} else {
						foul_type := foul_event.GetFoulType()
						num_fts := foul_event.GetNumberOfFtaForFoul()
						var and1_shot PbpEvent
						var free_throw_type string
						if foul_type == SHOOTING_FOUL_TYPE_STRING || foul_type == SHOOTING_BLOCK_TYPE_STRING {
							// check if 1, 2 or 3 shots
							if num_fts == 1 {
								// and 1
								if and1_shot, err = foul_event.GetFouledFgm(possession.Events); err != nil {
									return err
								} else {
									if and1_shot.EventNum == 0 {
										// bug - not an actual shooting foul - away from play foul
										free_throw_type = AWAY_FROM_PLAY_FOUL_TRIPS_KEY
									} else if and1_shot.Is3PointShot() {
										free_throw_type = THREE_POINT_AND1_KEY
									} else {
										free_throw_type = TWO_POINT_AND1_KEY
									}
								}
							} else {
								free_throw_type = fmt.Sprintf("%dptShootingFoulFreeThrowTrips", num_fts)
							}
						} else if foul_type == FLAGRANT_1_FOUL_TYPE_STRING || foul_type == FLAGRANT_2_FOUL_TYPE_STRING {
							if num_fts == 0 {
								// assume 2 shot flagrant if num_fts is 0
								num_fts = 2
							}
							free_throw_type = fmt.Sprintf("%dptShotFlagrantFreeThrowTrips", num_fts)
						} else if foul_type == AWAY_FROM_PLAY_FOUL_TYPE_STRING {
							free_throw_type = fmt.Sprintf("%dptShotAwayFromPlayFreeThrowTrips", num_fts)
						} else if foul_type == INBOUND_FOUL_TYPE_STRING {
							free_throw_type = fmt.Sprintf("%dptShotInboundFoulFreeThrowTrips", num_fts)
						} else {
							// penalty
							// check for 3 shots since sometimes shooting fouls are counted as personal fouls - can't really fix for 2 shot fouls but can for 3
							if num_fts == 3 {
								free_throw_type = THREE_POINT_SHOOTING_FOUL_TRIPS_KEY
							} else if num_fts == 1 {
								if and1_shot, err = foul_event.GetFouledFgm(possession.Events); err != nil {
									return err
								} else if and1_shot.EventNum == 0 {
									if player_id == and1_shot.PlayerId {
										if and1_shot.Is3PointShot() {
											free_throw_type = THREE_POINT_AND1_KEY
										} else {
											free_throw_type = TWO_POINT_AND1_KEY
										}
									} else {
										free_throw_type = AWAY_FROM_PLAY_FOUL_TRIPS_KEY
									}
								} else {
									free_throw_type = AWAY_FROM_PLAY_FOUL_TRIPS_KEY
								}
							} else {
								free_throw_type = PENALTY_FREE_THROW_KEY
							}
						}
						stats[o_team_id][o_lineup_id][d_lineup_id][player_id][free_throw_type] += 1
					}
				}
			} else if event.IsTechnicalFT() {
				player_id := event.PlayerId
				if _, exists := stats[o_team_id][o_lineup_id][d_lineup_id][player_id]; !exists {
					stats[o_team_id][o_lineup_id][d_lineup_id][player_id] = make(map[string]int64)
				}
				stats[o_team_id][o_lineup_id][d_lineup_id][player_id][TECHNICAL_FT_KEY] += 1
			}
		}
	}
	possession.PlayerStats = stats
	return nil
}

func SumPossessionStats(possessions []PossessionDetails, team_id int64) (map[string]int64, map[string]int64, map[int64]map[int64]map[string]int64, map[string]map[string]int64, map[string]map[string]int64, error) {
	// note that for all non player stats, seconds played will be 5 times actual value since sums up seconds for each player in lineup
	team_stats := make(map[string]int64)
	opponent_stats := make(map[string]int64)
	player_stats := make(map[int64]map[int64]map[string]int64)
	lineup_stats := make(map[string]map[string]int64)
	lineup_opponent_stats := make(map[string]map[string]int64)
	for _, possession := range possessions {
		for stats_team_id, possession_team_stats := range possession.PlayerStats {
			if _, exists := player_stats[stats_team_id]; !exists {
				player_stats[stats_team_id] = make(map[int64]map[string]int64)
			}
			var opponent_team_id int64
			if stats_team_id == possession.OffenseTeamId {
				opponent_team_id = possession.DefenseTeamId
			} else {
				opponent_team_id = possession.OffenseTeamId
			}
			if _, exists := player_stats[opponent_team_id]; !exists {
				player_stats[opponent_team_id] = make(map[int64]map[string]int64)
			}
			for lineup_id, all_lineup_stats := range possession_team_stats {
				if _, exists := lineup_stats[lineup_id]; !exists {
					lineup_stats[lineup_id] = make(map[string]int64)
				}
				for opponent_lineup_id, possession_lineup_stats := range all_lineup_stats {
					if _, exists := lineup_opponent_stats[opponent_lineup_id]; !exists {
						lineup_opponent_stats[opponent_lineup_id] = make(map[string]int64)
					}
					for player_id, possession_player_stats := range possession_lineup_stats {
						for stat_key, stat_value := range possession_player_stats {
							if strings.Contains(stat_key, REBOUND_KEY) {
								// add Off/Def to rebound key
								var rebound_team_chances_key, rebound_opponent_chances_key string
								if stats_team_id == possession.OffenseTeamId {
									rebound_team_chances_key = OFFENSIVE_REBOUND_PREIX + stat_key + CHANCES_STRING
									rebound_opponent_chances_key = DEFENSIVE_REBOUND_PREIX + stat_key + CHANCES_STRING
									stat_key = OFFENSIVE_REBOUND_PREIX + stat_key
								} else {
									rebound_team_chances_key = DEFENSIVE_REBOUND_PREIX + stat_key + CHANCES_STRING
									rebound_opponent_chances_key = OFFENSIVE_REBOUND_PREIX + stat_key + CHANCES_STRING
									stat_key = DEFENSIVE_REBOUND_PREIX + stat_key
								}
								for _, player_id_string := range strings.Split(lineup_id, "-") {
									player_id_int, err := strconv.ParseInt(player_id_string, 10, 64)
									if err != nil {
										return team_stats, opponent_stats, player_stats, lineup_stats, lineup_opponent_stats, err
									}
									if _, exists := player_stats[stats_team_id][player_id_int]; !exists {
										player_stats[stats_team_id][player_id_int] = make(map[string]int64)
									}
									player_stats[stats_team_id][player_id_int][rebound_team_chances_key] += 1
								}
								for _, player_id_string := range strings.Split(opponent_lineup_id, "-") {
									player_id_int, err := strconv.ParseInt(player_id_string, 10, 64)
									if err != nil {
										return team_stats, opponent_stats, player_stats, lineup_stats, lineup_opponent_stats, err
									}
									if _, exists := player_stats[opponent_team_id][player_id_int]; !exists {
										player_stats[opponent_team_id][player_id_int] = make(map[string]int64)
									}
									player_stats[opponent_team_id][player_id_int][rebound_opponent_chances_key] += 1
								}

								if stats_team_id == team_id {
									team_stats[rebound_team_chances_key] += 1
									opponent_stats[rebound_opponent_chances_key] += 1
									lineup_stats[lineup_id][rebound_team_chances_key] += 1
									lineup_opponent_stats[opponent_lineup_id][rebound_opponent_chances_key] += 1
								} else {
									team_stats[rebound_opponent_chances_key] += 1
									opponent_stats[rebound_team_chances_key] += 1
									lineup_stats[lineup_id][rebound_opponent_chances_key] += 1
									lineup_opponent_stats[opponent_lineup_id][rebound_team_chances_key] += 1
								}
							}
							if _, exists := player_stats[stats_team_id][player_id]; !exists {
								player_stats[stats_team_id][player_id] = make(map[string]int64)
							}
							player_stats[stats_team_id][player_id][stat_key] += stat_value
							if team_id == stats_team_id {
								team_stats[stat_key] += stat_value
								lineup_stats[lineup_id][stat_key] += stat_value
							} else {
								opponent_stats[stat_key] += stat_value
								lineup_opponent_stats[opponent_lineup_id][stat_key] += stat_value
							}
						}
					}
				}
			}
		}
		// add to possession counts - count possession for players who finished possession on the floor
		for _, player_id_string := range strings.Split(possession.OffenseLineupId, "-") {
			player_id_int, err := strconv.ParseInt(player_id_string, 10, 64)
			if err != nil {
				return team_stats, opponent_stats, player_stats, lineup_stats, lineup_opponent_stats, err
			}
			if _, exists := player_stats[possession.OffenseTeamId][player_id_int]; !exists {
				player_stats[possession.OffenseTeamId][player_id_int] = make(map[string]int64)
			}
			player_stats[possession.OffenseTeamId][player_id_int][OFFENSIVE_POSSESSIONS_KEY] += 1
		}
		for _, player_id_string := range strings.Split(possession.DefenseLineupId, "-") {
			player_id_int, err := strconv.ParseInt(player_id_string, 10, 64)
			if err != nil {
				return team_stats, opponent_stats, player_stats, lineup_stats, lineup_opponent_stats, err
			}
			if _, exists := player_stats[possession.DefenseTeamId][player_id_int]; !exists {
				player_stats[possession.DefenseTeamId][player_id_int] = make(map[string]int64)
			}
			player_stats[possession.DefenseTeamId][player_id_int][DEFENSIVE_POSSESSIONS_KEY] += 1
		}
		if team_id == possession.OffenseTeamId {
			team_stats[OFFENSIVE_POSSESSIONS_KEY] += 1
			opponent_stats[DEFENSIVE_POSSESSIONS_KEY] += 1
			lineup_stats[possession.OffenseLineupId][OFFENSIVE_POSSESSIONS_KEY] += 1
			lineup_opponent_stats[possession.OffenseLineupId][DEFENSIVE_POSSESSIONS_KEY] += 1
		} else {
			team_stats[DEFENSIVE_POSSESSIONS_KEY] += 1
			opponent_stats[OFFENSIVE_POSSESSIONS_KEY] += 1
			lineup_stats[possession.DefenseLineupId][DEFENSIVE_POSSESSIONS_KEY] += 1
			lineup_opponent_stats[possession.DefenseLineupId][OFFENSIVE_POSSESSIONS_KEY] += 1
		}
	}

	return team_stats, opponent_stats, player_stats, lineup_stats, lineup_opponent_stats, nil
}
