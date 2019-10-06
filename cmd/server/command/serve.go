package command

import (
	"log"
	"time"

	"github.com/spf13/cobra"

	"github.com/jasimmons/deckard/server"
)

const (
	defaultListenAddr   = "0.0.0.0:1998"
	defaultWriteTimeout = 10 * time.Second
	defaultReadTimeout  = 5 * time.Second
)

type ServeCommand struct {
	// Command is the parent command, which is where the Viper
	// configuration is. Without an embedded Command, we would need some
	// extra work to have Viper-bound flags and envvars.
	*Command

	listenAddr   string
	publicAddr   string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func (c *Command) NewServeCommand() *cobra.Command {
	s := &ServeCommand{Command: c}
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "serve a deckard HTTP server",
		RunE:  s.run,
	}

	serveCmd.Flags().String("addr", defaultListenAddr, "address for deckard server to listen on")
	serveCmd.Flags().Duration("timeout-read", defaultReadTimeout, "timeout duration for read requests")
	serveCmd.Flags().Duration("timeout-write", defaultWriteTimeout, "timeout duration for write requests")
	s.BindPFlags(serveCmd.Flags())

	return serveCmd
}

func (c *ServeCommand) run(cmd *cobra.Command, args []string) error {
	dsrv := server.New(
		server.ListenAddr(c.GetString("addr")),
		server.WriteTimeout(c.GetDuration("timeout-write")),
		server.ReadTimeout(c.GetDuration("timeout-read")),
	)

	log.Printf("serving deckard on %s", c.GetString("addr"))
	log.Fatal(dsrv.ListenAndServe())
	return nil
}
