package main

import (
	"bytes"
	"log"
	"poker/cards"
)

func deal(players int) []cards.Hand {
	hands := make([]cards.Hand, players, players)
	for p := 0; p < players; p++ {
		hands[p].DealRandom(2)
	}
	return hands
}

func isDead(h cards.Hand) bool {
	return h.Score() > 21
}

func main() {
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
	var buf bytes.Buffer
	var logger = log.New(&buf, "logger: ", log.Lshortfile)

	logger.Print("Hello, log file!")

}
