package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Command struct {
	Root *cobra.Command
	*viper.Viper
}

func New() *Command {
	c := &Command{
		Root: &cobra.Command{
			Use: "deckard",
		},
		Viper: viper.New(),
	}

	c.Root.AddCommand(
		c.NewServeCommand(),
	)

	c.SetEnvPrefix("DECKARD")
	c.AutomaticEnv()

	return c
}

func (c *Command) run(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func (c *Command) Execute() error {
	return c.Root.Execute()
}
