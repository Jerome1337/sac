package shapes

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
)

// RectangleArea return given rectangle area
func RectangleArea(ctx *cli.Context) error {
	l, _ := strconv.ParseFloat(ctx.Args().Get(0), 64)
	w, _ := strconv.ParseFloat(ctx.Args().Get(1), 64)

	return cli.Exit(fmt.Sprintf("Retangle area is: %f", l * w), 0)
}
