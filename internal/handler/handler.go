package handler

import "github.com/spf13/cobra"

type Handler struct {
	ExecuteFunc func(cmd *cobra.Command, args []string) error
}

func (h *Handler) Execute(cmd *cobra.Command, args []string) error {
	err := h.ExecuteFunc(cmd, args)
	return err
}
