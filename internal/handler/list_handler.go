package handler

import (
	"errors"
	"github.com/spf13/cobra"
	"sdkman-go/internal/exitcode"
	"sdkman-go/internal/mock"
)

func ListHandler() *Handler {
	return &Handler{
		ExecuteFunc: func(cmd *cobra.Command, args []string) error {
			javaRegistry := mock.NewJavaRegistry()
			if len(javaRegistry.Registry) == 0 {
				return errors.New(exitcode.IsFailure()) // No items in the registry
			}

			for key := range javaRegistry.Registry {
				cmd.Println(key)
			}

			if shouldRefresh() {
				return errors.New(exitcode.IsSuccessRefresh()) // List success and refresh needed
			}

			return errors.New(exitcode.IsSuccess()) // List success
		},
	}
}
