package handler

import (
	"github.com/spf13/cobra"
	"sdkman-go/internal/mock"
)

func ListHandler() *Handler {
	return &Handler{
		ExecuteFunc: func(cmd *cobra.Command, args []string) {
			javaRegistry := mock.NewJavaRegistry()
			for key := range javaRegistry.Registry {
				cmd.Println(key)
			}
		},
	}
}
