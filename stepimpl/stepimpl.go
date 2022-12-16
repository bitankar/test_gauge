package stepImpl

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/getgauge-contrib/gauge-go/gauge"
	m "github.com/getgauge-contrib/gauge-go/models"
	. "github.com/getgauge-contrib/gauge-go/testsuit"
)

var vowels map[rune]bool

var _ = gauge.Step("Vowels in English language are <vowels>.", func(vowelString string) {
	vowels = make(map[rune]bool, 0)
	for _, ch := range vowelString {
		vowels[ch] = true
	}
})

var _ = gauge.Step("Almost all words have vowels <table>", func(tbl *m.Table) {
	for _, row := range tbl.Rows {
		word := row.Cells[0]
		expectedCount, err := strconv.Atoi(row.Cells[1])
		if err != nil {
			T.Fail(fmt.Errorf("Failed to parse string %s to integer", row.Cells[1]))
		}
		actualCount := countVowels(word)
		if actualCount != expectedCount {
			T.Fail(fmt.Errorf("Vowel count in word %s - got: %d, want: %d", word, actualCount, expectedCount))
		}
	}
})

var _ = gauge.Step("The word <word> has <expectedCount> vowels.", func(word string, expected string) {
	actualCount := countVowels(word)
	expectedCount, err := strconv.Atoi(expected)
	if err != nil {
		T.Fail(fmt.Errorf("Failed to parse string %s to integer", expected))
	}
	if actualCount != expectedCount {
		T.Fail(fmt.Errorf("got: %d, want: %d", actualCount, expectedCount))
	}
})

func countVowels(word string) int {
	vowelCount := 0
	for _, ch := range word {
		if _, ok := vowels[ch]; ok {
			vowelCount++
		}
	}
	return vowelCount
}

var _ = gauge.Step("All tkn version should be as below <table>", func(tbl *m.Table) {
	out, err := exec.Command("tkn", "version").Output()
	if err != nil {
		T.Fail(fmt.Errorf("Error getting tkn version : %s", err))
	} else {
		op := string(out)
		fmt.Println(op)
	}
	for _, row := range tbl.Rows {
		tool := row.Cells[0]
		expectedVersion := row.Cells[1]
		fmt.Println(tool, expectedVersion)
		// if err != nil {
		// 	T.Fail(fmt.Errorf("Failed to parse string %s to integer", row.Cells[1]))
		// }
		// actualCount := countVowels(word)
		// if actualCount != expectedCount {
		// 	T.Fail(fmt.Errorf("Vowel count in word %s - got: %d, want: %d", word, actualCount, expectedCount))
		// }
	}
})
