package models

import (
	"encoding/base64"
	"strings"
	"time"
)

type Episode struct {
	Title       string
	Description string
	Url         string
	PubDate     *time.Time
	Season      *Season
}

// GetRealUrl. From API we will get an encoded URL. We need to get the real URL to the video
func (e *Episode) GetRealUrl() (string, error) {
	outSeq := []string{
		"u", "5", "Y", "I", "E", "4", "7", "D", "6", "G", "j", "n", "2",
		"g", "T", "b", "L", "v", "S", "F", "X", "h", "1", "q", "=", "Z",
	}
	inSeq := []string{
		"W", "x", "m", "M", "3", "r", "t", "0", "8", "z", "U", "A", "B",
		"d", "P", "y", "K", "O", "i", "V", "N", "w", "9", "l", "R", "C",
	}
	str := e.Url
	for i, _ := range e.Url {
		strings.Replace(str, inSeq[i], "__")
		strings.Replace(str, outSeq[i], inSeq[i])
		strings.Replace(str, "__", outSeq[i])
	}
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("Error in decoding url #{str}")
		return "", err
	}
	return data, nil
}
