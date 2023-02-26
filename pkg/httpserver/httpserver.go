package httpserver

import (
	"NoteKeeper/pkg/common"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net"
	"net/http"
)

type Server struct {
	logger *zap.Logger
	config *common.Config
	server *fasthttp.Server
}

func New(logger *zap.Logger, routes *router.Router) *Server {
	server := &fasthttp.Server{
		Name:               fmt.Sprintf("note_service fasthttp"),
		Handler:            routes.Handler,
		MaxConnsPerIP:      MAX_CONNECTIONS_PER_IP,
		MaxRequestsPerConn: MAX_REQUESTS_PER_CONNECTION,
		MaxRequestBodySize: MAX_REQUEST_BODY_SIZE * 1024 * 1024,
		WriteTimeout:       WRITE_TIMEOUT,
		ReadTimeout:        READ_TIMEOUT,
		IdleTimeout:        IDLE_TIMEOUT,
		TCPKeepalive:       TCP_KEEPALIVE,
	}

	return &Server{
		logger: logger.Named("httpserver"),
		server: server,
	}
}

// Run starts HTTP server
func (srv *Server) Run(ln net.Listener) {
	srv.logger.Info("Running HTTP server")
	if err := srv.server.Serve(ln); err != nil && err != http.ErrServerClosed {
		srv.logger.Fatal(err.Error())
	}
}

func (srv *Server) Shutdown() {
	if err := srv.server.Shutdown(); err != nil {
		srv.logger.Fatal(err.Error())
	}
}
