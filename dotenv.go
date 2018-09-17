/*
Package dotenvconfig is an implementation of the Ruby dotenv library and inspired by the Laravel environment loader.
The purpose of the library is to load variables from a file into the environment based on the current environment (dev, staging, production...)
*/
package dotenvconfig

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Load the variables into the environment. By default, default.env will be loaded and can be overwritten by any .env file (dev.env, staging.env...)
func Load(files ...string) {
	files = append([]string{"default.env"}, files...)

	for _, file := range files {
		file, err := os.Open(file)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			splited := strings.SplitN(scanner.Text(), "=", 2)
			// Check if the value is defined AND that the line is not a comment
			if len(splited) > 1 && scanner.Text()[:1] != "#" {
				key := splited[0]
				value := splited[1]
				os.Setenv(key, value)
			}
		}
		checkForError(scanner.Err())
	}

	return
}

func checkForError(e error) {
	if e != nil {
		log.Print(e)
	}
}
