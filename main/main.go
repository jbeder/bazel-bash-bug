package main

import (
	"log"

	fb "github.com/jbeder/bazel-bash-bug"
)

func main() {
	if err := fb.Run(); err != nil {
		log.Printf("%v", err)
	}
}
