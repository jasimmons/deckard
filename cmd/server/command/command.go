package command

import (
	"github.com/spf13/cobra"
)

type Command struct {
	Root *cobra.Command
}

func New() *Command {
	c := &Command{
		Root: &cobra.Command{
			Use: "deckard",
		},
	}

	c.Root.AddCommand(
		c.NewServeCommand(),
	)

	c.setupFlags()
	return c
}

func (c *Command) run(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func (c *Command) Execute() error {
	return c.Root.Execute()
}

func (c *Command) setupFlags() {
}
