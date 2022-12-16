package stepImpl

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

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
	tkn_version_map := map[string]string{}
	if err != nil {
		T.Fail(fmt.Errorf("Error getting tkn version : %s", err))
	} else {
		temp_string_list := []string{}
		re := regexp.MustCompile(`([a-zA-Z ]*)\s*:\s*v*([0-9\.]*)`)
		op_temp := re.FindAllString(strings.ToLower(string(out)), -1)
		for _, s1 := range op_temp {
			temp_string_list = strings.Split(s1, `:`)
			tkn_version_map[strings.Trim(strings.Split(temp_string_list[0], ` `)[0], ` `)] = strings.Trim(strings.ReplaceAll(temp_string_list[1], `v`, ``), ` `)
		}
	}
	for _, row := range tbl.Rows {
		tool := strings.ToLower(row.Cells[0])
		expectedVersion := strings.ReplaceAll(strings.ToLower(row.Cells[1]), `v`, ``)
		// fmt.Printf("checking for tool :  %v - got: %v, want: %v", tool, x[tool], expectedVersion)
		if tkn_version_map[tool] != expectedVersion {
			T.Fail(fmt.Errorf("version mismatch for tool :  %v - got: %v, want: %v", tool, x[tool], expectedVersion))
		}
	}
})
