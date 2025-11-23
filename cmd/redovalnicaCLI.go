package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func RedovalnicaCLI() {
	cmd := &cli.Command{
		Name:  "Redovalnica",
		Usage: "Redovalnica shows student grades",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Least amout of grades to pass",
				Value: 9,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Lowest possible grade",
				Value: 5,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Highest possible grade",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

			fmt.Printf("%d, %d, %d\n", stOcen, minOcena, maxOcena)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Printf("something went wrong\n")
	}
}
