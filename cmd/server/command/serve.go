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
	root *Command

	listenAddr   string
	publicAddr   string
	writeTimeout time.Duration
	readTimeout  time.Duration
}

func (c *Command) NewServeCommand() *cobra.Command {
	s := &ServeCommand{root: c}
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "serve a deckard HTTP server",
		RunE:  s.run,
	}

	serveCmd.Flags().StringVarP(&s.listenAddr, "addr", "a", defaultListenAddr, "address for deckard server to listen on")
	serveCmd.Flags().DurationVarP(&s.writeTimeout, "timeout-write", "", defaultWriteTimeout, "timeout duration for write requests")
	serveCmd.Flags().DurationVarP(&s.readTimeout, "timeout-read", "", defaultReadTimeout, "timeout duration for read requests")

	return serveCmd
}

func (c *ServeCommand) run(cmd *cobra.Command, args []string) error {
	dsrv := server.New(
		server.ListenAddr(c.listenAddr),
		server.WriteTimeout(c.writeTimeout),
		server.ReadTimeout(c.readTimeout),
	)

	log.Printf("serving deckard on %s", c.listenAddr)
	log.Fatal(dsrv.ListenAndServe())
	return nil
}
