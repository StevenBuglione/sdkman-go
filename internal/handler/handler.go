package handler

import "github.com/spf13/cobra"

type Handler struct {
	ExecuteFunc func(cmd *cobra.Command, args []string)
}

func (h *Handler) Execute(cmd *cobra.Command, args []string) {
	h.ExecuteFunc(cmd, args)
}
