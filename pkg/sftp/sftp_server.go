package sftp

import (
	"net"
	"sftp-server/pkg/auth"
	"sftp-server/pkg/config"
)

type Server struct {
	Handler *Handler
}

func NewServer(authHandler *auth.AuthHandler, config *config.Config) (*Server, error) {
	handler := NewHandler(config, authHandler)
	return &Server{
		Handler: handler,
	}, nil
}

func (s *Server) Serve(listener net.Listener) error {
	for {
		_, err := listener.Accept()
		if err != nil {
			return err
		}

		go func() {
			// Implement your connection handling logic here...
		}()
	}
}
