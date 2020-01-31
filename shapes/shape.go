package shapes

import "github.com/urfave/cli/v2"

func checkArgs(ctx *cli.Context) error {
	if !ctx.Args().Present() {
		return cli.NewExitError("no lengths provided", 1)
	}

	return nil
}
