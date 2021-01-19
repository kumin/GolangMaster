package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func StringConcatenation(prefix string, seed ...string) string {
	key := strings.Join(seed, ":")
	return fmt.Sprintf("%s:%s", prefix, key)
}

func GenUUID() {
	for i := 0; i < 100; i++ {
		fmt.Println(uuid.New())
	}
}

func main() {
	fmt.Println(StringConcatenation("experiment", "1"))
	GenUUID()
}
