package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "alpha alpha alpha alpha alpha"
	str = replaceNth(str, "alpha", "beta", 3)
	fmt.Println(str)

}

func replaceNth(s, old, new string, n int) string {
	// index
	i := 0

	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			// we did not find it.
			break
		}

		// you have found it.

		i = i + x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}

		i += len(old)
	}

	return s
}
