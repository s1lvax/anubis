package main

import (
	"anubis/utils"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("added task: ", cmd.Args().First())
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println("completed task: ", cmd.Args().First())
					return nil
				},
			},
			{
				Name:    "password",
				Aliases: []string{"pwd"},
				Usage:   "options for passwords",
				Commands: []*cli.Command{
					{
						Name:  "gen",
						Usage: "Generate a secure password",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							arg := cmd.Args().First()

							// Check if the argument is a valid integer
							number, err := strconv.Atoi(arg)
							if err != nil {
								// Handle the error if the argument is not a number
								return cli.Exit("Error: Argument must be a valid integer.", 1)
							}

							// Pass the number to generatePassword
							password, err := utils.GeneratePassword(number)
							if err != nil {
								fmt.Println("Error:", err)
								return err // Or handle the error in a way that's appropriate for your application
							}
							fmt.Println("Generated Password:", password)

							return nil
						},
					},
					{
						Name:  "rate",
						Usage: "Rate the strength of an existing password",
						Action: func(ctx context.Context, cmd *cli.Command) error {
							arg := cmd.Args().First()

							// Check if a password was provided
							if arg == "" {
								return cli.Exit("Error: Please provide a password to rate.", 1)
							}

							// Evaluate the password strength
							score, entropy, rating, feedback := utils.EvaluatePassword(arg)

							// Display results
							fmt.Printf("Password: %s\n", arg)
							fmt.Printf("Score: %d/10\n", score)
							fmt.Printf("Entropy: %.2f bits\n", entropy)
							fmt.Printf("Rating: %s\n", rating)
							fmt.Println("Feedback:")
							for _, tip := range feedback {
								fmt.Printf("- %s\n", tip)
							}

							return nil
						},
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
