package handler

import (
	"errors"
	"github.com/spf13/cobra"
	"sdkman-go/internal/mock"
)

func ListHandler() *Handler {
	return &Handler{
		ExecuteFunc: func(cmd *cobra.Command, args []string) error {
			javaRegistry := mock.NewJavaRegistry()
			if len(javaRegistry.Registry) == 0 {
				return errors.New("FAIL:1") // No items in the registry
			}

			for key := range javaRegistry.Registry {
				cmd.Println(key)
			}

			if shouldRefresh() {
				return errors.New("SUCCESS_REFRESH:100") // List success and refresh needed
			}

			return errors.New("SUCCESS:0") // List success
		},
	}
}
