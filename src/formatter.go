package main

import (
	"errors"
	"slices"
)

const F_TOP string = "[__________________________________]"
const F_BOTTOM string = "[__________________________________]"
const F_LEN int = 36
const F_PAD = 2

func FormatCenter(s string) (string, error) {
	flen := F_LEN
	content_len := len(s)
	// Ensure that the content is less than 32
	if len(s) > 28 {
		return "", errors.New("Format: More Than Allowed")
	}
	// Do the padding
	padding := (flen - content_len)
	if padding > 0 {
		padding = padding / 2
	}

	padding += 4
	var paddingStr string = "| "
	for range padding {
		paddingStr = string(append([]byte(paddingStr), byte(' ')))
	}
	// append the content
	paddingStr = string(append([]byte(paddingStr), []byte(s)...))
	paddingStr = string(append([]byte(paddingStr), []byte(" |")...))
	return paddingStr, nil
}

func FormatStart(s string) (string, error) {
	if len(s) > 32 {
		return "", errors.New("Format: More Than Allowed")
	}
	content := "| "
	diff := 32 - len(s)
	// append the content
	content = string(append([]byte(content), []byte(s)...))
	// give the diff back
	for range diff {
		content = string(append([]byte(content), byte(' ')))
	}
	content = string(append([]byte(content), []byte(" |")...))
	return content, nil
}

func FormatEnd(s string) (string, error) {
	if len(s) > 32 {
		return "", errors.New("Format: More Than Allowed")
	}
	content := "|_ "
	diff := 32 - len(s)
	reverse := s
	slices.Reverse([]byte(reverse))
	// append the content
	content = string(append([]byte(content), []byte(reverse)...))
	// give the diff back
	for range diff {
		content = string(append([]byte(content), byte(' ')))
	}
	content = string(append([]byte(content), []byte(" _|")...))
	copy := []byte(content)
	slices.Reverse(copy)
	content = string(copy)
	return content, nil
}

func FHeader(str string) (string, error) {
	buf := []byte(F_TOP)
	buf = append(buf, '\n')
	buf = append(buf, []byte(str)...)
	cntered, err := FormatCenter(str)
	if err != nil {
		return cntered, err
	}
	buf = append(buf, []byte(cntered)...)
	buf = append(buf, []byte("|__________________________________|\n")...)
	return string(buf), nil
}

type stringed string
