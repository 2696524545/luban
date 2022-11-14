package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func main() {

	fmt.Println(RandAllString(6))
}

var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func RandAllString(lenNum int) string {
	str := strings.Builder{}
	for i := 0; i < lenNum; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(36))
		if err != nil {
			return ""
		}
		l := CHARS[n.Int64()]
		str.WriteString(l)
	}
	return str.String()
}
