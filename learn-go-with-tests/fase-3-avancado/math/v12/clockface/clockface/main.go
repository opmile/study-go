package main

import (
	"os"
	"time"

	"github.com/quii/learn-go-with-tests/fase-3-avancado/math/v12/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
