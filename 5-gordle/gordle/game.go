package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// New returns a Game, which can be used to play
func New(playerInput io.Reader, corpus []string, maxAttempts int) (*Game, error) {
	if len(corpus) == 0 {
		return nil, ErrorCorpusIsEmpty
	}
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(pickWord(corpus)),
		maxAttempts: maxAttempts,
	}
	return g, nil
}

// Play runs the game
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")
	for attempt := 1; attempt <= g.maxAttempts; attempt++ {
		guess := g.ask()
		if slices.Equal(guess, g.solution) {
			fmt.Printf("👏 Congratulations! You guessed the word in %d attempts!\n", attempt)
			return
		}
		fmt.Println(computeFeedback(guess, g.solution).String())
	}
	fmt.Printf(
		"😞 You failed to guess the word in %d attempts. The solution was: %s\n",
		g.maxAttempts,
		string(g.solution),
	)
}

const solutionLength = 5

// ask reads input until a valid suggestion is mage (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)
	for {
		input, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}
		guess := splitToUppercaseCharacters(string(input))
		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(
				os.Stderr,
				"Your attempt is invalid with Grodle's solution: %s\n",
				err.Error(),
			)
		} else {
			return guess
		}
	}
}

// errInvalidWordLength is returned when
// the guess has the wrong number of characters.
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the " +
	"same number of characters as the solution")

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf(
			"expected %d, got %d, %w",
			solutionLength,
			len(guess),
			errInvalidWordLength,
		)
	}
	return nil
}

// splitToUppercaseCharacters is a naive implementation to turn a string
// into a list of characters
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

func computeFeedback(guess []rune, solution []rune) feedback {
	result := make(feedback, len(guess))
	if len(guess) != len(solution) {
		_,_=fmt.Fprintf(
			os.Stderr, 
			"Guess length (%d) is't equal to solution's length (%d)",
			len(guess), 
			len(solution),
		)
		return result
	}
	used := make([]bool, len(guess))
	for posInGuest,character := range guess {
		if character != solution[posInGuest] {
			continue
		}
		result[posInGuest] = correctPosition
		used[posInGuest] = true
	}
	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			continue
		}
		for posInSolution, target := range solution {
			if used[posInSolution] {
				continue
			}
			if target == character {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				break
			}
		}
	}
	return result
}
