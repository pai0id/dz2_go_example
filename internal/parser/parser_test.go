package parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type TestCase_t struct {
	filename string
	want     int
}

func getFileText(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func getTestCases() ([]TestCase_t, error) {
	testCases := []TestCase_t{}
	var testCase TestCase_t

	files, err := ioutil.ReadDir("../../data")
	if err != nil {
		return nil, err
	}

	pattern := regexp.MustCompile(`^.*_[0-9]+_out\.txt$`)

	for _, file := range files {
		if !file.IsDir() && pattern.MatchString(file.Name()) {
			fileName := fmt.Sprintf("../../data/%s", file.Name())

			content, err := getFileText(fileName)
			if err != nil {
				return nil, err
			}

			num, err := strconv.Atoi(string(content))
			if err != nil {
				return nil, err
			}

			testCase.filename = strings.ReplaceAll(fileName, "out", "in")
			testCase.want = num
			testCases = append(testCases, testCase)
		}
	}

	return testCases, nil
}

func TestCountPartsBasic(t *testing.T) {
	emailText, err := getFileText("../../data/pos_1_in.txt")
	if err != nil {
		t.Error(err)
	}

	cnt := CountParts(emailText)

	if cnt != 2 {
		t.Errorf("pos_1.txt = %v; Правильно: 2", cnt)
	}
}

func TestCountPartsTableDriven(t *testing.T) {
	var tests, err = getTestCases()
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v", tt.filename)
		t.Run(testname, func(t *testing.T) {
			emailText, err := getFileText(tt.filename)
			if err != nil {
				t.Error(err)
			}

			cnt := CountParts(emailText)
			if cnt != tt.want {
				t.Errorf("Получено: %v, Правильно: %v", cnt, tt.want)
			}
		})
	}
}

func areAllSubstrings(substrings []string, mainString string) bool {
	for _, substr := range substrings {
		if !strings.Contains(mainString, substr) {
			return false
		}
	}
	return true
}

func TestParseEmailBasic(t *testing.T) {
	emailText, err := getFileText("../../data/pos_1_in.txt")
	if err != nil {
		t.Error(err)
	}

	cnt, parts := ParseEmail(emailText)

	if cnt != 2 {
		t.Errorf("pos_1.txt = %v; Правильно: 2", cnt)
	} else if !areAllSubstrings(parts, emailText) {
		t.Error("pos_1.txt некорректные парты")
	}
}

func TestParseEmailDriven(t *testing.T) {
	var tests, err = getTestCases()
	if err != nil {
		t.Error(err)
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v", tt.filename)
		t.Run(testname, func(t *testing.T) {
			emailText, err := getFileText(tt.filename)
			if err != nil {
				t.Error(err)
			}

			cnt, parts := ParseEmail(emailText)
			if cnt != tt.want {
				t.Errorf("Получено: %v, Правильно: %v", cnt, tt.want)
			} else if !areAllSubstrings(parts, emailText) {
				t.Errorf("%v некорректные парты", tt.filename)
			}
		})
	}
}
