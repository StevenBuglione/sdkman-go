package handler

import (
	"errors"
	"github.com/spf13/cobra"
	"sdkman-go/internal/exitcode"
)

func InstallHandler() *Handler {
	return &Handler{
		ExecuteFunc: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New(exitcode.IsFailure()) // No args provided
			}

			// Your logic for install here...

			cmd.Println("Install Command: " + args[0])

			if shouldRefresh() {
				return errors.New(exitcode.IsSuccessRefresh()) // Install success and refresh needed
			}

			return errors.New(exitcode.IsSuccess()) // Install success
		},
	}
}
