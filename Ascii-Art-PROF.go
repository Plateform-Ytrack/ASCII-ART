// @title  - ASCII ART main module
// @author - Cedric OBEJERO <cedric.obejero@tanooki.fr>
// @date   - Dec. 21st, 2020
// Release - V1R5
//
// Change Log
// V1R6 - Add output file management to save ascii art
// V1R5 - Refactoring code for maintenance : split into files
// V1R4 - Add align management for center|right|left|jusity
// V1R3 - Add multiple banner files management
// V1R2 - Add color handling feature for basic ANSI Colors
// V1R1 - Refactoring cli arguments parsing with FLAG package
// V1R0 - Initial release to print input text in ASCII Art with standard.txt banner

package main

import (
	"fmt"
	"os"
	"strings"
)

// MAIN ENTRY POINT
func main() {
	var validArgs ASCIIArtOptions
	var char2print []rune
	var bannerLines []string
	var padding int
	var outputFile *os.File
	var err error

	// Validate command line inputs - Format & values
	validArgs = ParseArguments(os.Args)

	// Check banner file & get content by line in array
	bannerLines = GetBanner(&validArgs)

	if validArgs.IsReady {
		// Convert string to rune array
		char2print = []rune(validArgs.Text)

		// Convert ascii value of characters to first line position in banner file
		for i := 0; i < len(char2print); i++ {
			char2print[i] -= 32 // as we start from Space character in banner file
			char2print[i] *= 9  // as each banner letter uses 8 lines plus one empty line
			char2print[i]++     // as banner file start with an empty line
		}

		// To prepare text alignment
		// Define number of cols to print the text in ascii art
		asciiLenght := 0
		spaceCount := 0
		for index := 0; index < len(char2print); index++ {
			tmp := int(char2print[index])
			// IF we have to print SPACE character count it
			if tmp == 1 {
				spaceCount++
			}
			asciiLenght += len(bannerLines[tmp])
		}

		// Prepare data for text alignment
		if strings.Compare(validArgs.Align, "center") == 0 {
			padding = validArgs.Width - asciiLenght
			if padding < 0 {
				padding = 0
			} else {
				padding /= 2
			}
		} else if strings.Compare(validArgs.Align, "right") == 0 {
			padding = validArgs.Width - asciiLenght
			if padding < 0 {
				padding = 0
			}
		} else if strings.Compare(validArgs.Align, "justify") == 0 {
			padding = validArgs.Width - asciiLenght
			if padding < 0 {
				padding = 0
			} else {
				padding /= spaceCount
			}
		} else {
			// Align to the left by default
			padding = 0
		}

		// IF output is requested, prepare output file
		outputFile = nil
		if len(validArgs.OutFile) > 0 {
			outputFile, err = os.Open(validArgs.OutFile)
			if err != nil {
				fmt.Println("ERROR - ", err.Error())
				os.Exit(ErrorFileNotFound)
			}

		}

		// Print in ASCII ART = using 8 lines
		for i := 0; i < 8; i++ {
			if strings.Compare(validArgs.Align, "center") == 0 || strings.Compare(validArgs.Align, "right") == 0 {
				for i := 0; i < padding; i++ {
					PrintColor(validArgs.Color, " ", outputFile)
				}
			}

			for index := 0; index < len(char2print); index++ {
				lineNumber := int(char2print[index]) + i
				PrintColor(validArgs.Color, bannerLines[lineNumber], outputFile)
				if int(char2print[index]) == 1 && strings.Compare(validArgs.Align, "justify") == 0 {
					for i := 0; i < padding; i++ {
						PrintColor(validArgs.Color, " ", outputFile)
					}
				}
			}
			PrintColor(validArgs.Color, "\n", outputFile)
		}

		// IF output was requested, close file
		if outputFile != nil {
			outputFile.Close()
		}

	}
}