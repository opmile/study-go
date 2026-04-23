// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"os"
	"time"

	"github.com/quii/learn-go-with-tests/fase-3-avancado/math/vFinal/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
