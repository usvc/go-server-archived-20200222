package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

// fileExists checks if :pathToFile leads to a file, returning
// an error if it doesn't exist or is a directory
func fileExists(pathToFile string) error {
	if fileInfo, err := os.Lstat(pathToFile); err != nil {
		return NewError("FILE_NO_EXISTS", err.Error())
	} else if fileInfo.IsDir() {
		return NewError("FILE_NO_EXISTS", fmt.Sprintf("path '%s' leads to a directory", pathToFile))
	}
	return nil
}

// stringDefined checks if a string has been set
func stringDefined(someString string) error {
	if len(someString) == 0 {
		return NewError("STRING_UNDEFINED", fmt.Sprint("provided string was empty"))
	}
	return nil
}

// tlsCertKeyMatches attempts to load the certificate and key, returning
// an error if they don't match
func tlsCertKeyMatches(pathToCert, pathToKey string) error {
	if _, err := tls.LoadX509KeyPair(pathToCert, pathToKey); err != nil {
		return NewError("CERT_KEY_MISMATCH", fmt.Sprintf("cert at '%s' does not match key at '%s'", pathToCert, pathToKey))
	}
	return nil
}
