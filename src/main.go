package main

import (
	"bytes"   // package for working with byte buffers
	"flag"    // package for command line flag parsing
	"fmt"     // package for console output
	"os"      // package for operating system functionality
	"strings" // package for working with strings
)

// Morse->letter map
var morseLetters = map[string]string{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
	"-----": "0",
	"/":     " ",
}

func decodeMorse(morse string, warn bool) {
	// output is a buffer that will be used to store the decoded message.
	var output bytes.Buffer

	// If the Morse code input is empty, output an error message and exit the program with an error code.
	if morse == "" {
		fmt.Fprintln(os.Stderr, "Error: Morse is empty")
		os.Exit(1)
	}

	// Loop through each word in the input Morse code.
	for _, w := range strings.Split(morse, "  ") {
		// Loop through each character in the word and decode it.
		for _, char := range strings.Split(w, " ") {
			// If the morse code is valid, write the corresponding letter to the output buffer.
			// Otherwise, write a question mark to indicate an invalid code.
			if result, ok := morseLetters[char]; ok {
				output.WriteString(result)
			} else {
				warn = true
				output.WriteString("?")
			}
		}
	}

	if warn {
		fmt.Println("WARNING: ? means that you put invalid code")
	}

	// Output the decoded message and a note about invalid codes.
	fmt.Printf("Result: %s", output.String())
}

func main() {

	/* Command line flags for warn boolean (if it should warn the user about invalid codes)
	and morseInput string which is the target Morse code to decode */

	var warn bool
	flag.BoolVar(&warn, "warn", false, "Toggle warnings")

	var morseInput string
	flag.StringVar(&morseInput, "morse", "", "Morse code to decode.\nWARNING: DO NOT REMOVE THE QUOTES.")

	flag.Parse()

	decodeMorse(morseInput, warn)
}
