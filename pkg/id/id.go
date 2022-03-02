package id

import (
	"math/rand"
	"time"
)

type ID string

// Sum of degit symbols and upper/lower case latin symbols
const idCharsCount = 10 + 26 + 26
const idDefaultLen = 6

var idChars []rune

func init() {
	rand.Seed(time.Now().UnixNano())

	// Append smallcase and uppercase symbols to idChars
	small, big := 97, 65
	for i := 0; i < 26; i++ {
		idChars = append(idChars, rune(i+small), rune(i+big))
	}

	// Append degit symbols to idChars
	num := 48
	for i := 0; i < 10; i++ {
		idChars = append(idChars, rune(i+num))
	}
}

func randomSym() rune {
	return idChars[rand.Intn(idCharsCount)]
}

func GenerateId() ID {
	var id ID
	for i := 0; i < idDefaultLen; i++ {
		id += ID(randomSym())
	}

	return id
}
