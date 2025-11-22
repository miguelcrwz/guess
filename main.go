package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("\033[35;1m??? GUESSING GAME ???\033[0m")
	fmt.Println("\n\033[36;1mA RANDOM NUMBER WILL BE GENERATED. TRY TO ANSWER CORRECTLY. THE NUMBER MUST BE BETWEEN 0 AND 100.\033[0m")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses{
	fmt.Printf("\n\033[33;1mWHAT'S YOUR GUESS? \033[0m")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)

		if err != nil{
			fmt.Println("\n\033[31;1m✘ YOUR GUESS MUST BE A NON-NEGATIVE INTEGER.\033[0m")
			continue
		}

		switch{
		case guessInt < x:
			fmt.Println("\n\033[31;1m✘ WRONG — THE NUMBER IS HIGHER.\033[0m")
		case guessInt > x:
			fmt.Println("\n\033[31;1m✘ WRONG — THE NUMBER IS LOWER.\033[0m")
		case guessInt == x:
			fmt.Printf("\n\033[32;1m✔ CORRECT! YOU GUESSED THE NUMBER IN %d TRIES. YOUR TRIES WERE: %v\033[0m", i+1, guesses[:i])
			return
		}

		guesses[i] = guessInt
	}

	fmt.Printf("\n\033[37;1mUNFORTUNATELY, YOU DIDN'T GUESS THE NUMBER %d. YOUR TRIES WERE: %v\033[0m", x, guesses)
}
