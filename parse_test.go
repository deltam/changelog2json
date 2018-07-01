package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	testdata := []struct {
		s    string
		logs []changelog
	}{
		{`2018-07-01  Nanashi Gombeh2  <deltam+2@gmail.com>

        * tag3: title3
        hogehoge

        * tag2: title2
        fugafuga
        hogehoge

2018-06-29  Nanashi Gombeh  <deltam@gmail.com>

        * tag1: title1
        hogefuga
`, []changelog{
			{
				Date:  "2018-07-01",
				Name:  "Nanashi Gombeh2",
				Mail:  "deltam+2@gmail.com",
				Tags:  "tag3",
				Title: "title3",
				Body:  "        hogehoge\n\n",
				Line:  3,
			},
			{
				Date:  "2018-07-01",
				Name:  "Nanashi Gombeh2",
				Mail:  "deltam+2@gmail.com",
				Tags:  "tag2",
				Title: "title2",
				Body:  "        fugafuga\n        hogehoge\n\n",
				Line:  6,
			},
			{
				Date:  "2018-06-29",
				Name:  "Nanashi Gombeh",
				Mail:  "deltam@gmail.com",
				Tags:  "tag1",
				Title: "title1",
				Body:  "        hogefuga\n",
				Line:  12,
			},
		},
		},
	}

	for _, td := range testdata {
		ret, _ := parse(td.s)
		for i, log := range td.logs {
			if log != ret[i] {
				t.Errorf("parse() is %v, want %v", ret[i], log)
			}
		}
	}
}
