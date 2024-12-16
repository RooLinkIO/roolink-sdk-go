package roolink

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	bazadebezolkohpepadrPattern = regexp.MustCompile(`bazadebezolkohpepadr="([^"]+)"`)
	scriptURLPattern            = regexp.MustCompile(`<script type="text/javascript"\s+(?:nonce=".*?")?\s+src="([a-z\d/\-_]+)"></script>`)
)

func GetBazadebezolkohpepadr(html string) (int, error) {
	match := bazadebezolkohpepadrPattern.FindStringSubmatch(html)
	if len(match) < 2 {
		return -1, errors.New("failed to parse bazadebezolkohpepadr")
	}
	return strconv.Atoi(match[1])
}

func ParseScriptURL(html string) (string, error) {
	match := scriptURLPattern.FindStringSubmatch(html)
	if len(match) < 2 {
		return "", errors.New("failed to parse script URL")
	}
	return match[1], nil
}
