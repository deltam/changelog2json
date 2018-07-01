package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type changelog struct {
	Date  string `json:"date"`
	Name  string `json:"name"`
	Mail  string `json:"mail"`
	Tags  string `json:"tags"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Line  int    `json:"line"`
}

func parse(text string) ([]changelog, error) {
	clogs := make([]changelog, 0, int(len(text)/500))

	sc := bufio.NewScanner(strings.NewReader(text))
	i := 0
	eof := false
	nextLine := func() bool {
		eof = !sc.Scan()
		i++
		return !eof
	}
	skipBlank := func() {
		for nextLine() {
			if sc.Text() != "" {
				break
			}
		}
	}

	skipBlank()
	for !eof {
		h, err := parseHeader(sc.Text())
		if err != nil {
			return nil, fmt.Errorf("[line %d] %v", i, err)
		}

		skipBlank()
		for isTitle(sc.Text()) {
			m, err := parseTitle(sc.Text(), *h)
			if err != nil {
				return nil, fmt.Errorf("[line %d] %v", i, err)
			}
			m.Line = i

			for nextLine() && isBody(sc.Text()) {
				m.Body += sc.Text() + "\n"
			}

			clogs = append(clogs, *m)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("[line %d] Scan failed:, %v", i, err)
	}

	return clogs, nil
}

var headerRegex = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})  (.+)  <(.+)>$`)

func isHeader(line string) bool {
	return headerRegex.MatchString(line)
}

func parseHeader(line string) (*changelog, error) {
	h := headerRegex.FindStringSubmatch(line)
	if len(h) != 4 {
		return nil, fmt.Errorf("parse header failed: %s", line)
	}
	return &changelog{Date: h[1], Name: h[2], Mail: h[3]}, nil
}

var titleRegex = regexp.MustCompile(`^ +\* (.+): (.+)$`)

func isTitle(line string) bool {
	return titleRegex.MatchString(line)
}

func parseTitle(line string, m changelog) (*changelog, error) {
	t := titleRegex.FindStringSubmatch(line)
	if len(t) != 3 {
		return nil, fmt.Errorf("parse title failed: %s", line)
	}
	m.Tags = strings.Trim(t[1], " ")
	m.Title = strings.Trim(t[2], " ")
	return &m, nil
}

func isBody(line string) bool {
	return !isHeader(line) && !isTitle(line)
}
