package shapes

import (
	"fmt"
	"math"
	"strconv"

	"github.com/urfave/cli/v2"
)

// TriangleArea return given triangle area
func TriangleArea(ctx *cli.Context) error {
	a, _ := strconv.ParseFloat(ctx.Args().Get(0), 64)
	b, _ := strconv.ParseFloat(ctx.Args().Get(1), 64)
	c, _ := strconv.ParseFloat(ctx.Args().Get(2), 64)

	p := (a + b + c) / 2
	area := math.Sqrt(p * (p - a) * (p - b) * (p - c))

	return cli.Exit(fmt.Sprintf("Triangle area is: %f", area), 0)
}
