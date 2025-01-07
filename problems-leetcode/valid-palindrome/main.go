package main

/* (,],{,}) */

func isValid(s string) bool {
	options := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}"}
	stack := []string{}
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		/* GET OPEN PARENTHESES */
		if closing, ok := options[char]; ok {
			stack = append(stack, closing)
		} else {
			/* close parentheses case */
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			/* the last item is the char value current */
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {

}
