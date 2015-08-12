package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net", "io"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func removeWhites(text string) string {
	var newText []rune
	for _, r := range text {
		if unicode.IsSpace(r) {
			continue
		}
		// appending rune to the slice
		newText = append(newText, r)
	}
	return string(newText)
}
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// converting text to lower case (obviously)
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			// check whether whether rune is in allowed characters string
			if !strings.ContainsRune(allowedChars, r) {
				// if rune not allowed - skipping it and moving forward
				continue
			}
			// appending rune to the slice
			newText = append(newText, r)
		}
		// converting []rune slice to a string and outputing
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
