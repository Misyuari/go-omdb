package main

import (
	"fmt"
	"strings"
	"testing"
)

var stringTestBracket = "test(bracket)s"

func TestFindFirstStringInBracket(t *testing.T) {
	findFirstStringInBracket := func(str string) string {
		if len(str) > 0 {
			indexFirstBracketFound := strings.Index(str, "(")
			if indexFirstBracketFound >= 0 {
				runes := []rune(str)
				wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
				indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
				if indexClosingBracketFound >= 0 {
					runes := []rune(wordsAfterFirstBracket)
					return string(runes[1 : indexClosingBracketFound-1])
				} else {
					return ""
				}
			} else {
				return ""
			}
		} else {
			return ""
		}
		return ""
	}
	fmt.Println(findFirstStringInBracket(stringTestBracket)) // output = bracke
}

func TestFindFirstStringInBracketNew(t *testing.T) {
	findFirstStringInBracket := func(str string) string {
		indexFirstBracket := strings.IndexByte(str, '(')
		if indexFirstBracket < 0 {
			return ""
		}
		indexFirstBracket++
		indexClosingBracket := strings.IndexByte(str[indexFirstBracket:], ')')
		if indexClosingBracket < 0 {
			return ""
		}
		//return str[indexFirstBracket : indexFirstBracket+indexClosingBracket-1]  // output = bracke
		return str[indexFirstBracket : indexFirstBracket+indexClosingBracket] // output = bracket
	}

	fmt.Println(findFirstStringInBracket(stringTestBracket))
}

func TestAnagram(t *testing.T) {
	groupAnagrams := func(strs []string) [][]string {
		res := [][]string{}
		tmp := map[[26]int][]string{} // 26 is number of alphabet
		for _, s := range strs {
			chars := [26]int{}
			for _, c := range s {
				chars[c-'a']++
			}
			tmp[chars] = append(tmp[chars], s)
		}
		for _, v := range tmp {
			res = append(res, v)
		}
		return res
	}
	fmt.Println(groupAnagrams([]string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}))
}
