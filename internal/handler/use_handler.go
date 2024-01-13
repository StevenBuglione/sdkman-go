package handler

import (
	"fmt"
	"github.com/spf13/cobra"
	"sdkman-go/internal/exitcode"
	"sdkman-go/internal/service"
)

type UseHandler struct {
	UseHandlerService service.IUseHandlerService
}

func NewUseHandler() IHandler {
	uh := &UseHandler{
		UseHandlerService: service.NewUseHandlerService(),
	}
	// Check interface implementation
	var _ IHandler = uh
	return uh
}

func (h *UseHandler) Execute(cmd *cobra.Command, args []string) exitcode.ExitCode {
	err := h.UseHandlerService.Use(args)
	if err != nil {
		fmt.Println("Error setting Java path:", err)
		return exitcode.Failure
	}
	return exitcode.Success
}

func (h *UseHandler) Refresh(exitCode exitcode.ExitCode) exitcode.ExitCode {
	if exitCode == exitcode.Success {
		return exitcode.SuccessRefresh
	}
	return exitcode.Failure
}
