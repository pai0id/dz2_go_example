package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func getEmailText(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func TestCountPartsBasic(t *testing.T) {
	emailText, err := getEmailText("../../data/pos_1.txt")
	if err != nil {
		t.Error(err)
	}

	cnt := CountParts(emailText)

	if cnt != 2 {
		t.Errorf("pos_1.txt = %v; Правильно: 2", cnt)
	}
}

func TestCountPartsTableDriven(t *testing.T) {
	var tests = []struct {
		filename string
		want     int
	}{
		{"pos_1.txt", 2},
		{"pos_2.txt", 1},
		{"pos_3.txt", 1},
		{"neg_1.txt", 0},
		{"neg_2.txt", 0},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%v", tt.filename)
		t.Run(testname, func(t *testing.T) {
			emailText, err := getEmailText(fmt.Sprintf("../../data/%s", tt.filename))
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
