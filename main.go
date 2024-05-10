package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	englishWords := []string{
		"apple", "banana", "chair", "table", "computer", "desk", "book", "pencil", "tree", "cat",
		"dog", "house", "car", "bike", "road", "river", "mountain", "beach", "ocean", "sun",
		"moon", "star", "cloud", "rain", "snow", "wind", "flower", "grass", "bird", "fish",
		"tiger", "lion", "elephant", "giraffe", "bear", "butterfly", "bee", "ant", "spider",
		"snake", "frog", "turtle", "dolphin", "whale", "shark", "octopus", "jellyfish",
		"coral", "iceberg", "volcano",
	}

	word := englishWords[rand.Intn(len(englishWords))]
	lives := 2 * len(word)

	// 2. generate the word blanks  "php" -> p_p
	blanks := []string{}
	for range word {
		blanks = append(blanks, "_")
	}
	for {
		// 3. show the word blanks and ask for letters
		fmt.Printf("❤️ %d, Word: %s Letter: ", lives, strings.Join(blanks, ""))

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		// 4. check provided letters
		for _, inputLetter := range input {
			correctGuess := false

			for i, wordLetter := range word {
				if inputLetter == wordLetter {
					blanks[i] = string(inputLetter)
					correctGuess = true
				}
			}

			if !correctGuess {
				lives--
			}
		}

		// 5. if no more lives, you lost
		if lives <= 0 {
			fmt.Printf("❤️ 0, Word: %s - sorry, you lost!\n", word)
			break
		}
		// 6. if word is guessed, you won
		if word == strings.Join(blanks, "") {
			fmt.Printf("❤️ %d, Word: %s - you won, congrats!\n", lives, word)
			break
		}
	}
	// repeat steps 3-6 until win or loss.

}
