package main

import (
	"line/interview/utils"
	"testing"
)

type isUrlTestCase struct {
	url      string
	expected int
}

var isUrlTestCases = []isUrlTestCase{
	{"http://www.google.com", 0},
	{"http://www.google.com/search", 0},
	{"https://www.youtube.com", 0},
	{"https://www.facebook.com", 0},
	{"https://www.wikipedia.org", 0},
	{"https://www.instag777ram.com", 1},
	{"www.sanook.com", 1},
	{"www.twitter.com", 1},
	{"https://www.youtube.com", 0},
}

func TestIsUp(t *testing.T) {
	for _, testCase := range isUrlTestCases {
		actual := utils.IsUpOrDown([]string{testCase.url})
		if actual.CountFail != testCase.expected {
			t.Errorf("isUrl(%s) = %v, expected %v", testCase.url, actual.CountFail, testCase.expected)
		}
	}
}
