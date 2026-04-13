package main

import (
	"os"
	"time"

	"gowithtests.aliusman.net/math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
