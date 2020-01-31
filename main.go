package main

import  (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/Jerome1337/sac/shapes"
)

var cmds []*cli.Command

func init() {
	cmds = []*cli.Command{
		{
			Name: "area",
			Usage: "Find shapes area!",
			Subcommands: []*cli.Command{
				{
					Name: "triangle",
					Aliases: []string{"t"},
					Action: shapes.TriangleArea,
				},
				{
					Name: "rectangle",
					Aliases: []string{"r"},
					Action: shapes.RectangleArea,
				},
			},
		},
	}
}

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Name: "sac",
		Usage: "Geometric shapes helper command!",
		Commands: cmds,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
