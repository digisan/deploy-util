package main

import (
	"regexp"
	"strings"
	"time"

	. "github.com/digisan/go-generics/v2"
	gio "github.com/digisan/gotk/io"
)

type symbol map[string]string

var (
	m = symbol{
		"Updated@": "Updated@ " + time.Now().Format(time.RFC3339), // 2022-09-24T18:55:19+10:00
	}
	rm = map[string]*regexp.Regexp{
		"Updated@": regexp.MustCompile(`Updated@ \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+\d{2}:\d{2}`),
	}
)

func replace(line string) string {
	for k, v := range m {
		if rm[k].MatchString(line) {
			line = rm[k].ReplaceAllString(line, v)
		} else {
			line = strings.ReplaceAll(line, k, v)
		}
	}
	return line
}

func isCommentary(line string) bool {
	okSrc := strings.HasPrefix(line, "//")
	return okSrc
}

func ReplaceSymbol(onlyCmt bool, fPaths ...string) error {
	for _, fPath := range fPaths {
		if _, err := gio.FileLineScan(fPath, func(line string) (bool, string) {
			switch {
			case onlyCmt:
				line = IF(isCommentary(line), replace(line), line)
			default:
				line = replace(line)
			}
			return true, line
		}, fPath); err != nil {
			return err
		}
	}
	return nil
}
