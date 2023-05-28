package server

import (
	"net"
	"sftp-server/pkg/auth"
	"sftp-server/pkg/config"
	"sftp-server/pkg/sftp"
)

func Run(configFilePath string) error {
	authHandler, err := auth.NewHandler()
	if err != nil {
		return err
	}

	configHandler, err := config.NewConfig(configFilePath)
	if err != nil {
		return err
	}

	server, err := sftp.NewServer(authHandler, configHandler)
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", "0.0.0.0:55555")
	if err != nil {
		return err
	}

	return server.Serve(listener)
}
