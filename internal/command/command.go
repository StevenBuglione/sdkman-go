package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sdkman-go/internal/exitcode"
	"sdkman-go/internal/handler"
)

type Command interface {
	Init(use, short, long string, handler *handler.Handler)
	Register(*cobra.Command)
}

type BaseCommand struct {
	Cmd *cobra.Command
}

func (b *BaseCommand) Init(use, short, long string, h *handler.Handler) {
	b.Cmd = &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := h.Execute(cmd, args)
			// Handle error and capture exit code
			if err != nil {
				errParts := err.Error()
				exitCode := exitcode.ParseExitCode(errParts)

				// Use exit codes as per the operation result
				switch exitCode {
				case 100: // Success refresh code
					fmt.Println(exitCode) // Print the message
					// Call your refresh function here
				default:
					fmt.Println(exitCode) // Print the message
				}
				fmt.Println(exitCode) // Print the message
				os.Exit(exitCode)
			}
			return nil
		},
	}
}

func (b *BaseCommand) Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(b.Cmd)
}
