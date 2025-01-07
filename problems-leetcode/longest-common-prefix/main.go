package main

import "fmt"

func recursionNext(indexRange int, strsCopy []string, prefix string) bool {
	for _, str := range strsCopy {
		/* first condition: sub string with size smaller that prefix */
		/* second condition: prefix has in sub string? */
		if indexRange >= len(str) || str[:indexRange+1] != prefix {
			return false
		}
	}

	return true
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		/* handle empty arr of strings */
		return ""
	}

	/* inicialize variables (prefix and range of index) */
	firstLetter := string(strs[0])
	currentLength := len(firstLetter)
	indexRange := 0

	/* loop in array if handle with any empty sub string inside array*/
	for currentLength > 0 && indexRange < currentLength {
		prefix := firstLetter[:indexRange+1]
		if recursionNext(indexRange, strs[1:], prefix) {
			indexRange += 1
		} else {
			break
		}

	}

	if currentLength > 0 {
		/* its have prefix */
		return firstLetter[:indexRange]
	}

	return ""
}

func main() {
	strs := []string{"flower", "flow", "flight"}

	result := longestCommonPrefix(strs)

	fmt.Println(result)
}
