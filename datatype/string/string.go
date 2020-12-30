package main

import (
	"fmt"
	"strings"
)

func StringConcatenation(prefix string, seed ...string) string {
	key := strings.Join(seed, ":")
	return fmt.Sprintf("%s:%s", prefix, key)
}

func main() {
	fmt.Println(StringConcatenation("experiment", "1"))
}
