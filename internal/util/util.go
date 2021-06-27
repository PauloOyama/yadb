// Package util contains utilitary functions.
package util

import (
	"fmt"
	"os"
)

// GetEnv returns the given environment variable or ends the program with an error message indicating that the variable
// was not found
func GetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		fmt.Fprintln(os.Stderr, "The", key, "environment variable is not set")
		os.Exit(1)
	}

	return val
}
