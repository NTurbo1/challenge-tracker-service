package util

import (
	"os"
	"fmt"
)

// Function FileExists returns true if the file exists.
func FileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}

	return true
}

// Function WriteHeaderLineTo writes a given header line to a given file.
// Returns an error if there's a write fail. The returned error is nil in case of success. 
func WriteHeaderLineTo(file *os.File, headerLine string) error {
	n, err := file.WriteString(headerLine + "\n")
	if err != nil {
		return err
	}
	if n < len(headerLine) {
		// TODO: Implement multiple write tries until success or max number of tries has been reached.
		// Can be ignored for now.
		return fmt.Errorf("Incomplete header write to file %s.", file)
	}

	return nil
}
