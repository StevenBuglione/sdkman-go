package command

import (
	"github.com/spf13/cobra"
	"os"
	"sdkman-go/internal/handler"
)

type ICommand interface {
	Init(use, short, long string, handler handler.IHandler)
	Register(*cobra.Command)
}

type Command struct {
	Cmd *cobra.Command
}

func (b *Command) Init(use, short, long string, h handler.IHandler) {
	b.Cmd = &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		RunE: func(cmd *cobra.Command, args []string) error {
			resp := h.Execute(cmd, args)
			exitCode := h.Refresh(resp)
			os.Exit(exitCode.Value())
			return nil
		},
	}
}

func (b *Command) Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(b.Cmd)
}
