package dice

import (
	"crypto/rand"
	"math/big"
)

var dice = [6]int{1, 2, 3, 4, 5, 6}
var bigDiceMax = big.NewInt(6)

func Roll() int {
	br, err := rand.Int(rand.Reader, bigDiceMax)
	if err != nil {
		panic(err)
	}
	return dice[br.Int64()]
}
