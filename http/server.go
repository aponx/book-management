package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/aponx/book-management/common"

	"github.com/rs/zerolog/log"
)

type IServer interface {
	Start()
}

type server struct {
	*common.Server
	http.Handler
}

func NewServer(svr *common.Server, handler http.Handler) IServer {
	return &server{
		svr,
		handler,
	}
}

func (s *server) Start() {
	var srv http.Server
	idleConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Info().Msg("Server is shutting down")

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Error().Err(err).Msg("Server failed to shut down")
		}
		close(idleConnectionClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%s", s.Host, s.Port)

	srv.Handler = s.Handler

	log.Info().Msg("Server is starting at " + srv.Addr)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Error().Err(err).Msg("Server failed to start")
	}

	<-idleConnectionClosed
	log.Info().Msg("Server shutted down")
}
