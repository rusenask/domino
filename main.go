package main

import (
	"bufio"
	"flag"
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
	// concatenating tlds into one string
	defaultTlds := strings.Join(tlds, ", ")
	// supplying default names to topLevelDomains
	var topLevelDomains = flag.String("tlds", defaultTlds, "Top level domains")
	flag.Parse() // parse the flag

	newTlds := strings.Split(*topLevelDomains, ",")
	// Seeding randomizer
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
		fmt.Println(removeWhites(string(newText) + "." + newTlds[rand.Intn(len(newTlds))]))
	}
}
