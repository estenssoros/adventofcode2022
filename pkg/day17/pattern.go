package day17

import (
	"errors"
	"fmt"
	"strings"
)

func findPattern(data string) string {
	fmt.Println("finding patterns")
	// start with two elements grow window
	// try to find repeating of window
	var largestMatch string
	// matches := map[string]int{}
	for i := 0; i < len(data); i++ {
		for j := i + 10; j-i < len(data)-j+1; j++ {
			match, err := kmp(data[j:], data[i:j])
			if err != nil {
				continue
			}
			// matches[match]++
			if len(match) > len(largestMatch) {
				largestMatch = match
			}
			// fmt.Println(i, j, data[i:j], data[j:])
		}
	}
	// common, count := mostCommonString(matches)
	// fmt.Println("most common", count)
	// fmt.Println(common)
	// if isStringRepeated(largestMatch, common) {
	// 	return common
	// }
	return largestMatch
}

func isRepeated(match string, matches map[string]int) bool {
	for k := range matches {
		if isStringRepeated(match, k) {
			fmt.Println("repeated!", k)
			return true
		}
	}
	return false
}

func isStringRepeated(s, sub string) bool {
	for strings.HasPrefix(s, sub) {
		s = strings.TrimRight(s, sub)
	}
	return len(s) == 0
}

func mostCommonString(l map[string]int) (string, int) {
	var s string
	var i int
	for k, v := range l {
		if v > i {
			i = v
			s = k
		}
	}
	return s, i
}

type Match struct {
	match string
}

func (m Match) String() string {
	// b := strings.Builder{}
	// for _, d := range m.match {
	// 	b.WriteString(fmt.Sprint(d))
	// }
	// return b.String()
	return m.match
}

func kmp(all, sub string) (string, error) {
	table := buildTable(sub)
	return doesMatch(all, sub, table)
}

func buildTable(sub string) []int {
	table := make([]int, len(sub))
	for i := 0; i < len(sub); i++ {
		table[i] = -1
	}
	var i, j = 1, 0
	for i < len(sub) {
		if sub[i] == sub[j] {
			table[i] = j
			i++
			j++
		} else if j > 0 {
			j = table[j-1] + 1
		} else {
			i++
		}
	}
	return table
}

var ErrNotMatch = errors.New("no match")

func doesMatch(all, sub string, table []int) (string, error) {
	var i, j int
	for i+len(sub)-j <= len(sub) {
		if all[i] == sub[j] {
			if j == len(sub)-1 {
				return sub, nil
			}
			i++
			j++
		} else if j > 0 {
			j = table[j-1] + 1
		} else {
			i++
		}
	}
	return "", ErrNotMatch
}
