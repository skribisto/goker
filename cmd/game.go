package cmd

import (
	"fmt"

	"github.com/skribisto/goker/internals/common"
	"github.com/skribisto/goker/internals/plays"
	"github.com/skribisto/goker/pkg/log"
)

func Execute() {
	var answer string
	rageQuit := false

	log.GLogf("Current Version of Goker: %v", common.GetVersion())

	p0, err := plays.NewPlayer("Player", 0)
	if err != nil {
		log.Fatalf("could not instantiate Player0 : %w", err)
	}

	p1, err := plays.NewPlayer("CPU1", 1)
	if err != nil {
		log.Fatalf("could not instantiate Player1 : %w", err)
	}

	p2, err := plays.NewPlayer("CPU2", 1)
	if err != nil {
		log.Fatalf("could not instantiate Player1 : %w", err)
	}
	//log.Debugf("%v is sitting on the bench", p0.Name)
	players := []*plays.Player{p1, p2, p0}
	//players := []*plays.Player{p1, p2}
	//players := []*plays.Player{p0, p1}

	for !rageQuit {
		p, err := plays.NewPlay(players)
		if err != nil {
			log.Fatalf("%w", err)
		}
		log.GLogf("BlindInfos : %v", p.Blinds)

		blindPlayersSkipped := 0

		for p.Round != 4 {
			currentRount := p.Round
			if err := p.BeginRound(); err != nil {
				log.Fatal("round could not begin")
			}
			for id := range p.Players {
				if p.Round == 0 && len((*p.Bets)[0]) == 2 {
					if len(p.Players) > 2 && id < 2 && blindPlayersSkipped < 2 {
						log.Debug(">2 is a crowd, skip 2 players due to blinds")
						//BUG, if only 2 players left pre-flop, and they havec blind, we'll skip them in loop
						blindPlayersSkipped++
						continue
					} else if len(p.Players) == 2 && id == 0 {
						log.Debug("begining new play in heads up, skip small blind player")
						blindPlayersSkipped++
						continue
					}
				}
				if !p.Players[id].StillPlays {
					continue
				}

				log.GLog("#####################################")
				log.GLogf("######### P%v: %v's turn  ############", id, p.Players[id].Name)
				log.GLog("#####################################")

				isAutonomous := false
				if p.Players[id].Strategy == 1 {
					log.Debug("Player is autonomous")
					isAutonomous = true
				}

				cards, err := p.GetPlayerCards(id)
				if err != nil {
					log.Fatalf("%w", err)
				}

				board, err := p.GetBoard()
				if err != nil {
					log.Fatalf("%w", err)
				}
				if !isAutonomous {
					log.GLogf("%v has: %v          board: %v", p.Players[id].Name, *cards, *board)
					log.GLogf("Bets for this round are: %v", (*p.Bets)[p.Round])
				} else {
					log.GLogf("board: %v", *board)
					log.GLogf("Bets for this round are: %v", (*p.Bets)[p.Round])
				}

				canCheck, err := p.CanCheck(id)
				if err != nil {
					log.Fatalf("got error during CanCheck %w", err)
				}

			QUESTION:
				if isAutonomous {
					answer, err = p.ComputePlayerDecision(id)
					if err != nil {
						log.Fatalf("%w", err)
					}
				} else {
					if canCheck {
						log.GLogf("%v, What do you want to do ? (check/raise/fold)", p.Players[id].Name)
					} else {
						log.GLogf("%v, What do you want to do ? (call/raise/fold)", p.Players[id].Name)
					}

					//read input from user
					fmt.Scanln(&answer)
				}

				switch answer {
				case "call":
					if err := p.Call(id); err != nil {
						log.Fatalf("%w", err)
					}
				case "raise":
					if err := p.Raise(id, p.Blinds.BigBlind); err != nil {
						log.Fatalf("%w", err)
					}
				case "check":
					if !canCheck {
						log.GLog("You can't do that")
						if isAutonomous {
							log.Fatal("error computing decision for CPU")
						}
						goto QUESTION
					}
					if err := p.Call(id); err != nil {
						log.Fatalf("%w", err)
					}
				case "fold":
					if err := p.Fold(id); err != nil {
						log.Fatalf("%w", err)
					}
				default:
					log.Warn("I didn't understand")
					if isAutonomous {
						log.Fatal("error computing decision for CPU")
					}
					goto QUESTION
				}
				answer = ""

				if err := p.EndRound(); err != nil {
					log.Debugf("%w", err) //don't exit, just continue playing
				}

				if p.Round != currentRount {
					//if EndRound changed the current round, stop asking players for this round
					//e.g. EndRound when everyone else fold, p.Round is 4 and we need to start new game
					break
				}
			}
		}
		onlyAutonomousPlayers := true

		for i := range p.Players {
			p.Players[i].StillPlays = true
			log.GLogf("Stack of %v is: $%v", p.Players[i].Name, p.Players[i].Stack)
			if p.Players[i].Strategy != 1 {
				onlyAutonomousPlayers = false
			}
		}

		if !onlyAutonomousPlayers {
			log.GLog("Continue playing (Y/n)?")
			fmt.Scanln(&answer)
			if answer == "n" {
				log.GLog("Thanks for playing")
				rageQuit = true
			}
			answer = ""
		}

		players = make([]*plays.Player, len(p.Players))

		for id := range p.Players {
			if p.Players[id].Stack > (p.Blinds.BigBlind) {
				newId := (id + 1) % len(p.Players) //move the Dealer button
				players[newId] = p.Players[id]
			}
		}

		if len(players) == 1 {
			log.GLogf("TOURNAMENT WINNER: %v", players[0].Name)
			rageQuit = true
		}
	}
}
