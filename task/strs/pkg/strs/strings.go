package strs

import (
	"github.com/Grishun/problems/task/math/pkg/dectobin"
	"math/rand"
	"strings"
	"unicode"
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

func CeasarCyfer(s string, shift int) string {

	cyfered := strings.Builder{}

	for shift < 0 || shift >= 26 {
		if shift < 0 {
			shift += 26
		} else if shift >= 26 {
			shift -= 26
		}
	}

	words := strings.Fields(s)

	shiftBitMask := func(str string, shift int32) []int32 {

		var charMask int32
		charMasks := make([]int32, 0, len(str))

		for _, char := range str {
			startRune := 'a'
			var flagUpper, flagLower int32 = 0, 0

			if unicode.IsUpper(char) {
				startRune = 'A'
				flagUpper = 1
			} else {
				flagLower = 1
			}

			if char-startRune+shift > 25 {
				charMask = 1<<(char-startRune+shift-26+flagLower) - flagUpper
			} else {
				charMask = 1<<(char-startRune+shift+flagLower) - flagUpper
			}

			charMasks = append(charMasks, charMask)
		}

		return charMasks
	}

	rebuildWord := func(masks []int32) string {
		letters := strings.Builder{}
		for _, mask := range masks {
			if strings.Contains(dectobin.DecToBin(uint64(mask)), "0") && mask != 0 {
				letters.WriteRune(rune(strings.Count(dectobin.DecToBin(uint64(mask)), "0") + 'a' - 1))
			} else {
				letters.WriteRune(rune(strings.Count(dectobin.DecToBin(uint64(mask)), "1") + 'A'))
			}
		}

		return letters.String()
	}

	for _, word := range words {
		cyfered.WriteString(rebuildWord(shiftBitMask(word, int32(shift))))
	}

	return cyfered.String()
}

// CeasarCyferV2 this version is written by ChatGpt
func CeasarCyferV2(s string, shift int) string {
	shift = shift % 26
	if shift < 0 {
		shift += 26
	}
	ciphered := strings.Builder{}
	for _, char := range s {
		if unicode.IsLetter(char) {
			base := 'a'
			if unicode.IsUpper(char) {
				base = 'A'
			}
			char = (char-base+rune(shift))%26 + base
		}
		ciphered.WriteRune(char)
	}
	return ciphered.String()
}

func RandText(n int) (text string) {

	for i := 0; i < n; i++ {
		word := ""
		for j := 0; j < rand.Intn(10); j++ {
			flag := rand.Intn(2)
			if flag == 1 {
				word += string(rune('A' + rand.Intn(26)))
			} else {
				word += string(rune('a' + rand.Intn(26)))
			}
		}
		text += word + " "
	}

	return
}
