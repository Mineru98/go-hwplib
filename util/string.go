package util

import (
	"errors"
	"strings"
)

// UTF-16LE 문자열의 바이트 크기를 반환합니다.
func GetUTF16LEStringSize(s string) int {
	if s == "" {
		return 2
	}
	return 2 + len(s)*2
}

// 바이트 배열을 16진수 문자열로 변환합니다.
func BytesToHex(bytes []byte) string {
	hexArray := "0123456789ABCDEF"
	hexChars := make([]byte, len(bytes)*3)
	for j, b := range bytes {
		v := b & 0xFF
		hexChars[j*3] = hexArray[v>>4]
		hexChars[j*3+1] = hexArray[v&0x0F]
		hexChars[j*3+2] = ' '
	}
	return string(hexChars)
}

// 두 문자열이 같은지 확인합니다.
func Equals(str1, str2 string) bool {
	return str1 == str2
}

// 문자열을 검색 및 대체합니다.
func ReplaceEach(text string, searchList, replacementList []string) (string, error) {
	return replaceEach(text, searchList, replacementList, false, 0)
}

func replaceEach(text string, searchList, replacementList []string, repeat bool, timeToLive int) (string, error) {
	if timeToLive < 0 {
		searchSet := make(map[string]struct{}, len(searchList))
		replacementSet := make(map[string]struct{}, len(replacementList))
		for _, s := range searchList {
			searchSet[s] = struct{}{}
		}
		for _, r := range replacementList {
			replacementSet[r] = struct{}{}
		}
		for s := range searchSet {
			if _, found := replacementSet[s]; found {
				return "", errors.New("aborting to protect against StackOverflowError")
			}
		}
	}

	if isEmpty(text) || len(searchList) == 0 || len(replacementList) == 0 || (len(searchList) > 0 && timeToLive == -1) {
		return text, nil
	}

	searchLength := len(searchList)
	replacementLength := len(replacementList)

	if searchLength != replacementLength {
		return "", errors.New("search and replace array lengths don't match")
	}

	noMoreMatchesForReplIndex := make([]bool, searchLength)
	textIndex := -1
	replaceIndex := -1
	tempIndex := -1

	for i := 0; i < searchLength; i++ {
		if noMoreMatchesForReplIndex[i] || isEmpty(searchList[i]) || replacementList[i] == "" {
			continue
		}
		tempIndex = strings.Index(text, searchList[i])

		if tempIndex == -1 {
			noMoreMatchesForReplIndex[i] = true
		} else if textIndex == -1 || tempIndex < textIndex {
			textIndex = tempIndex
			replaceIndex = i
		}
	}

	if textIndex == -1 {
		return text, nil
	}

	start := 0
	increase := 0

	for i := 0; i < searchLength; i++ {
		if searchList[i] == "" || replacementList[i] == "" {
			continue
		}
		greater := len(replacementList[i]) - len(searchList[i])
		if greater > 0 {
			increase += 3 * greater
		}
	}
	increase = min(increase, len(text)/5)

	var buf strings.Builder
	buf.Grow(len(text) + increase)

	for textIndex != -1 {
		buf.WriteString(text[start:textIndex])
		buf.WriteString(replacementList[replaceIndex])

		start = textIndex + len(searchList[replaceIndex])

		textIndex = -1
		replaceIndex = -1

		for i := 0; i < searchLength; i++ {
			if noMoreMatchesForReplIndex[i] || searchList[i] == "" || replacementList[i] == "" {
				continue
			}
			tempIndex = strings.Index(text[start:], searchList[i])
			if tempIndex == -1 {
				noMoreMatchesForReplIndex[i] = true
			} else if textIndex == -1 || tempIndex < textIndex {
				textIndex = tempIndex + start
				replaceIndex = i
			}
		}
	}

	buf.WriteString(text[start:])
	result := buf.String()
	if !repeat {
		return result, nil
	}

	return replaceEach(result, searchList, replacementList, repeat, timeToLive-1)
}

func isEmpty(cs string) bool {
	return cs == ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
