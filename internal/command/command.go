package command

import (
	"github.com/spf13/cobra"
	"sdkman-go/internal/handler"
)

type Command interface {
	Init(use, short, long string, handler *handler.Handler)
	Register(*cobra.Command)
}

type BaseCommand struct {
	Cmd *cobra.Command
}

func (b *BaseCommand) Init(use, short, long string, handler *handler.Handler) {
	b.Cmd = &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run:   handler.Execute,
	}
}

func (b *BaseCommand) Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(b.Cmd)
}
