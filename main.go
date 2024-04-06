package main

import (
	"goker/cards"
	"goker/scores"
	"log"
	"os"
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

func main() {
	logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logFile)

	//log.Println("#####################################")
	log.Println("############# New Game   ############")
	//log.Println("#####################################")

	d := cards.NewDeck()

	var board0 cards.Board
	var board1 cards.Board
	var board2 cards.Board

	// Deal cards from the deck
	for i := 0; i < 5; i++ {
		card, err := d.Deal()
		if err != nil {
			log.Fatalln(err)
		}
		board0 = append(board0, card)
	}
	log.Println("Dealt board cards: ", board0)
	for i := 0; i < 2; i++ {
		card, err := d.Deal()
		if err != nil {
			log.Fatalln(err)
		}
		board1 = append(board1, card)
	}
	log.Println("Dealt p1 cards: ", board1)
	for i := 0; i < 2; i++ {
		card, err := d.Deal()
		if err != nil {
			log.Fatalln(err)
		}
		board2 = append(board2, card)
	}
	log.Println("Dealt p2 card: ", board2)
	board1 = append(board1, board0...)
	board2 = append(board2, board0...)

	sc1, err := scores.Score(board1)
	if err != nil {
		log.Fatalln(err)
	}
	sc2, err := scores.Score(board2)
	if err != nil {
		log.Fatalln(err)
	}

	i, err := scores.CompareScoreCards(sc1, sc2)
	if err != nil {
		log.Fatalln(err)
	}
	if i == -1 {
		log.Println("player 2 won")
	} else if i == 1 {
		log.Println("player 1 won")
	} else {
		log.Println("ex aequo")
	}

	/*
		var answer string
		rageQuit := false
		for !rageQuit {
			players := 4
			isGameEnded := false

			hands := deal(players)
			playerHand := &hands[0]

			fmt.Println()
			fmt.Println("#####################################")
			fmt.Println("##### New Game (" + strconv.Itoa(players) + " players)  #########")
			fmt.Println("#####################################")
			fmt.Println()
			time.Sleep(1000 * time.Millisecond)
			for !isGameEnded {
				fmt.Println("Your hands is composed of: " + playerHand.String() + " \t(Total: " + strconv.Itoa(playerHand.Score()) + ")")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Do you want to hit that ? (y/N)")

				//read input from user
				fmt.Scanln(&answer)
				if answer == "y" || answer == "h" {
					playerHand.DealRandom(1)

					if isDead(*playerHand) {
						fmt.Println("Your hands is composed of: " + playerHand.String() + " \t(Total: " + strconv.Itoa(playerHand.Score()) + ")")
						fmt.Println("You sadly overshot this one... More luck next time !")
						isGameEnded = true
						time.Sleep(1000 * time.Millisecond)
					}
				} else {
					isGameEnded = true
				}
				answer = ""
			}

			fmt.Println("#######################")
			fmt.Println("##### AI's turn #######")
			fmt.Println("#######################")
			time.Sleep(500 * time.Millisecond)

			bestPlayer := 0
			bestScore := 0

			for player, hand := range hands {
				if player != 0 {
					for !isDead(hand) && hand.Score() < 15 { //High level AI right there
						hand.DealRandom(1)
					}
				}

				if isDead(hand) {
					fmt.Println("Player" + strconv.Itoa(player) + " died")
				} else {

					if hand.Score() > bestScore {
						bestPlayer = player
						bestScore = hand.Score()
					}

					fmt.Println("Player" + strconv.Itoa(player) + " has \t" + hand.String() + " \t(Total: " + strconv.Itoa(hand.Score()) + ")")
				}
				time.Sleep(500 * time.Millisecond)
			}

			if bestPlayer == 0 {
				fmt.Println("Congrats !! you won !")
			} else {
				fmt.Println("You lost! Too bad ... Player" + strconv.Itoa(bestPlayer) + " won with a score of " + strconv.Itoa(bestScore))
			}
			time.Sleep(1000 * time.Millisecond)

			fmt.Println()

			fmt.Println("Continue playing (Y/n)?")
			fmt.Scanln(&answer)
			if answer == "n" {
				fmt.Println("Thanks for playing")
				rageQuit = true
			}
		}
	*/

}
