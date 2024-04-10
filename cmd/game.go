package cmd

import (
	"fmt"
	"goker/internals/plays"
	"goker/pkg/log"
)

func Execute() {
	var answer string
	rageQuit := false
	p0, err := plays.NewPlayer("Skrib", 0)
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
	log.Debugf("%v is sitting on the bench", p2.Name)
	//players := []*plays.Player{p0, p1, p2}
	players := []*plays.Player{p0, p1}

	for !rageQuit {

		p, err := plays.NewPlay(players)
		if err != nil {
			log.Fatalf("%w", err)
		}
		log.Infof("BlindInfos : %v", p.Blinds)

		for p.Round != 4 {
			currentRount := p.Round
			if err := p.BeginRound(); err != nil {
				log.Fatal("round could not begin")
			}
			for id := range p.Players {
				if !p.Players[id].StillPlays {
					continue
				}
				if p.Round == 0 && len((*p.Bets)[0]) == 2 {
					if len(p.Players) > 2 && id < 2 {
						log.Debug(">2 is a crowd, skip 2 players due to blinds")
						continue
					} else if id == 0 {
						log.Debug("begining new play in heads up, skip small blind player")
						continue
					}
				}

				log.Info("#####################################")
				log.Infof("######### P%v: %v's turn  ############", id, p.Players[id].Name)
				log.Info("#####################################")

				isAutonomous := false
				if p.Players[id].Strategy == 1 {
					log.Info("Player is autonomous")
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
					log.Infof("%v has: %v          board: %v", p.Players[id].Name, *cards, *board)
					log.Infof("Bets for this round are: %v", (*p.Bets)[p.Round])
				}

				canCheck, err := p.CanCheck(id)
				if err != nil {
					log.Fatalf("got error during CanCheck %w", err)
				}
				if isAutonomous {
					if err := p.Call(id); err != nil {
						log.Fatalf("%w", err)
					}
				} else {
					if canCheck {
						log.Infof("%v, What do you want to do ? (raise/check/fold)", p.Players[id].Name)
					} else {
						log.Infof("%v, What do you want to do ? (call/raise/fold)", p.Players[id].Name)
					}

				QUESTION:
					//read input from user
					fmt.Scanln(&answer)
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
							log.Info("You can't do that")
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
						log.Info("I didn't understand")
						goto QUESTION
					}
					answer = ""
				}

				if err := p.EndRound(); err != nil {
					log.Debugf("%w", err) //don't exit, just continue playing
				}

				if p.Round == currentRount+1 {
					//Manages case where player 1 raised, player 0 had to call, we want to start player loop again
					//goto NEXTROUND
					break
				}
			}
		}
		for i := range p.Players {
			p.Players[i].StillPlays = true
			log.Infof("Stack of %v is: $%v", p.Players[i].Name, p.Players[i].Stack)
		}

		log.Info("Continue playing (Y/n)?")
		fmt.Scanln(&answer)
		if answer == "n" {
			log.Info("Thanks for playing")
			rageQuit = true
		}
		answer = ""

		for id := range p.Players {
			p.Players[id].ID = (id + 1) % len(p.Players) //move the Dealer button
		}
		players = []*plays.Player{p1, p0}
	}
}
