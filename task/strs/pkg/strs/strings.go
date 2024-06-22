package strs

import (
	"github.com/Grishun/problems/task/math/pkg/dectobin"
	"strings"
)

func RightString(s string) bool {
	words := strings.Fields(s)
	return words[0] == "Здравствуйте" && strings.Contains(s, "пожалуйста") && s[len(s)-1] == '.'
}

func ToAnotherCase(s string, p bool) string {

	newS := s[:0]
	for _, char := range s {
		if p {
			newS += strings.ToUpper(string(char))
		} else {
			newS += strings.ToLower(string(char))
		}
	}
	s = newS
	return s
}

func ChangeVowels(s string) string {
	vowels := "aeiou"

	for _, vowel := range vowels {
		s = strings.ReplaceAll(s, string(vowel), "*")
	}

	return s
}

func CountVowels(s string) (res int) {
	vowels := "AEIOUaeiou"

	for _, vowel := range vowels {
		res += strings.Count(s, string(vowel))
	}
	return
}

func CeasarCyfer(s string, shift int) (cyfered string) {

	words := strings.Fields(s)

	shiftBitMask := func(str string, shift int32) []int32 {

		var charMask int32
		charMasks := make([]int32, 0, len(str))

		for _, char := range str {
			if char-'a'+shift > 25 {
				charMask = 1 << (char - 'a' + shift - 26)
			} else {
				charMask = 1 << (char - 'a' + shift)
			}
			charMasks = append(charMasks, charMask)
		}

		return charMasks
	}

	rebuildWord := func(masks []int32) string {
		letters := ""
		for _, mask := range masks {
			letters += string(strings.Count(dectobin.DecToBin(uint64(mask)), "0") + 'a')
		}
		return letters
	}

	for _, word := range words {
		cyfered += rebuildWord(shiftBitMask(word, int32(shift))) + " "
	}

	return

}
