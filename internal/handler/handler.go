package handler

import (
	"github.com/spf13/cobra"
	"sdkman-go/internal/exitcode"
)

type IHandler interface {
	Execute(cmd *cobra.Command, args []string) exitcode.ExitCode
	Refresh(exitCode exitcode.ExitCode) exitcode.ExitCode
}
