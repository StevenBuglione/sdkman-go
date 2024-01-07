package handler

import "github.com/spf13/cobra"

func InstallHandler() *Handler {
	return &Handler{
		ExecuteFunc: func(cmd *cobra.Command, args []string) {
			cmd.Println("Install Command" + args[0])
		},
	}
}
