package main

import (
	"bytes"   // package for working with byte buffers
	"flag"    // package for command line flag parsing
	"fmt"     // package for console output
	"os"      // package for operating system functionality
	"strings" // package for working with strings
)

// morseLetters is a map that maps Morse code strings to their corresponding letters or symbols.
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

func main() {
	// output is a buffer that will be used to store the decoded message.
	var output bytes.Buffer

	// morse is the input Morse code to decode.
	var morse string

	// flag.StringVar parses the command line argument for the Morse code to decode.
	flag.StringVar(&morse, "morse", "", "Morse code to decode.\nWARNING: DO NOT REMOVE THE QUOTES.")
	flag.Parse()

	// If the Morse code input is empty, output an error message and exit the program with an error code.
	if len(morse) == 0 {
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
				output.WriteString("?")
			}
		}
	}

	// Output the decoded message and a note about invalid codes.
	fmt.Printf("Result: %s\nNOTE: ? means that you put invalid code.\n", output.String())
}