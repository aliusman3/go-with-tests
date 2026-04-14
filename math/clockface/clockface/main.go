package main

import (
	"os"
	"time"

	. "gowithtests.aliusman.net/math/clockface/svg"
)

func main() {
	t := time.Now()
	SVGWriter(os.Stdout, t)
}
