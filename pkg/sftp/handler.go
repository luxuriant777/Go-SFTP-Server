package sftp

import (
	"errors"
	"github.com/pkg/sftp"
	"sftp-server/pkg/auth"
	"sftp-server/pkg/config"
)

type Handler struct {
	config      *config.Config
	authHandler *auth.AuthHandler
}

func NewHandler(config *config.Config, authHandler *auth.AuthHandler) *Handler {
	return &Handler{
		config:      config,
		authHandler: authHandler,
	}
}

func (h *Handler) HandleSFTPRequest(req *sftp.Request) error {
	switch req.Method {
	case "List":
		return h.handleListRequest(req)
	case "Open":
		return h.handleOpenRequest(req)
	case "Write":
		return h.handleWriteRequest(req)
	case "Read":
		return h.handleReadRequest(req)
	case "Close":
		return h.handleCloseRequest(req)
	case "Remove":
		return h.handleRemoveRequest(req)
	case "Rename":
		return h.handleRenameRequest(req)
	default:
		return errors.New("unsupported request method")
	}
}

func (h *Handler) handleListRequest(req *sftp.Request) error {
	// handle listing request here
	// use req.RespondWithFileList([]os.FileInfo)
	return nil
}

func (h *Handler) handleOpenRequest(req *sftp.Request) error {
	// handle open request here
	// use req.RespondWithHandle(string)
	return nil
}

func (h *Handler) handleWriteRequest(req *sftp.Request) error {
	// handle write request here
	// use req.Respond() after write operation
	return nil
}

func (h *Handler) handleReadRequest(req *sftp.Request) error {
	// handle read request here
	// use req.RespondWithData([]byte)
	return nil
}

func (h *Handler) handleCloseRequest(req *sftp.Request) error {
	// handle close request here
	// use req.Respond() after close operation
	return nil
}

func (h *Handler) handleRemoveRequest(req *sftp.Request) error {
	// handle remove request here
	// use req.Respond() after remove operation
	return nil
}

func (h *Handler) handleRenameRequest(req *sftp.Request) error {
	// handle rename request here
	// use req.Respond() after rename operation
	return nil
}
