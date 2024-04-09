package cmd

import (
	"fmt"
	"goker/internals/plays"
	"goker/pkg/log"
)

/*
func deal(players int) []cards.Hand {
	hands := make([]cards.Hand, players, players)
	for p := 0; p < players; p++ {
		hands[p].Deal(2)
	}
	return hands
}
*/

func Execute() {
	p0, err := plays.NewPlayer("Skrib")
	if err != nil {
		log.Fatalf("could not instantiate Player1 : %w", err)
	}

	p1, err := plays.NewPlayer("Computer1")
	if err != nil {
		log.Fatalf("could not instantiate Player2 : %w", err)
	}

	players := []*plays.Player{p0, p1}

	p, err := plays.NewPlay(players)
	if err != nil {
		log.Fatalf("%w", err)
	}
	log.Infof("BlindInfos : %v", p.Blinds)

	var answer string
	rageQuit := false
	for !rageQuit {
		log.Info("#####################################")
		log.Info("############ P0's turn  #############")
		log.Info("#####################################")

		//time.Sleep(500 * time.Millisecond)
		for p.Round != 4 {
			if err := p.BeginRound(); err != nil {
				log.Fatal("round could not begin")
			}

			cards, err := p.GetPlayerCards(p0.ID)
			if err != nil {
				log.Fatalf("%w", err)
			}
			log.Infof("Player %v has: %v", p0.Name, *cards)

			log.Infof("Bets for this round are: %v", (*p.Bets)[p.Round])

			canCheck, err := p.CanCheck(p0.ID)
			if err != nil {
				log.Fatalf("got error during CanCheck %w", err)
			}
			if canCheck {
				log.Info("Player0, What do you want to do ? (bet/check/fold)")
			} else {
				log.Info("Player0, What do you want to do ? (bet/fold)")
			}

			//read input from user
			fmt.Scanln(&answer)
			switch answer {
			case "bet":
				if err := p.PutBet(p0.ID, p.Blinds.SmallBlind); err != nil {
					log.Fatalf("%w", err)
				}
			case "check":
				if err := p.PutBet(p0.ID, 0); err != nil {
					log.Fatalf("%w", err)
				}
			case "fold":
				if err := p.Fold(p0.ID); err != nil {
					log.Fatalf("%w", err)
				}
			default:
				log.Info("I didn't understand")
				continue
			}
			answer = ""

			if p.Round == 4 {
				break
			}

			log.Info("#######################")
			log.Info("##### AI's turn #######")
			log.Info("#######################")
			//time.Sleep(500 * time.Millisecond)

			cards, err = p.GetPlayerCards(p1.ID)
			if err != nil {
				log.Fatalf("%w", err)
			}
			log.Infof("AI has: %v", *cards)
			log.Infof("Bets for this round are: %v", (*p.Bets)[p.Round])

			canCheck, err = p.CanCheck(p1.ID)
			if err != nil {
				log.Fatalf("got error during CanCheck %w", err)
			}
			if canCheck {
				log.Info("AI, What do you want to do ? (bet/check/fold)")
			} else {
				log.Info("AI, What do you want to do ? (bet/fold)")
			}

			//read input from user
			fmt.Scanln(&answer)
			switch answer {
			case "bet":
				if err := p.PutBet(p1.ID, p.Blinds.SmallBlind); err != nil {
					log.Fatalf("%w", err)
				}
			case "check":
				if err := p.PutBet(p1.ID, 0); err != nil {
					log.Fatalf("%w", err)
				}
			case "fold":
				if err := p.Fold(p1.ID); err != nil {
					log.Fatalf("%w", err)
				}
			default:
				log.Info("I didn't understand")
				continue
			}

			answer = ""

			if err := p.EndRound(); err != nil {
				log.Fatalf("%w", err)
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

		players := []*plays.Player{p1, p0}

		p, err = plays.NewPlay(players)
		if err != nil {
			log.Fatalf("%w", err)
		}
		log.Infof("BlindInfos : %v", p.Blinds)

	}
}
