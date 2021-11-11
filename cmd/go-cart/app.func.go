package goCart

import (
	b64 "encoding/base64"
	"os"
	"fmt"

	"github.com/gosimple/slug"
	"gopkg.in/yaml.v2"
)

// slugger slugs a string
func slugger(textToSlug string) string {
	return slug.Make(textToSlug)
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	checkAndFail(err)
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	readConfig = config

	return config, nil
}

// PreflightSetup just makes sure the stage is set
func PreflightSetup() {
	logStdOut("Preflight complete!")
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// rmStrFromStrSlice removes a string from a string slice
func rmStrFromStrSlice(r string, s []string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// B64EncodeBytesToStr converts a byte slice to a Base64 Encoded String
func B64EncodeBytesToStr(input []byte) string {
	return b64.StdEncoding.EncodeToString(input)
}

// B64DecodeBytesToBytes converts a Base64 byte slice to a Base64 Decoded Byte slice
func B64DecodeBytesToBytes(input []byte) ([]byte, error) {
	return B64DecodeStrToBytes(string(input))
}

// B64DecodeStrToBytes converts a Base64 string to a Base64 Decoded Byte slice
func B64DecodeStrToBytes(input string) ([]byte, error) {
	return b64.StdEncoding.DecodeString(input)
}