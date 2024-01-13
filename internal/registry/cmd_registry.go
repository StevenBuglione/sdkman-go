package registry

import (
	"github.com/spf13/cobra"
	"sdkman-go/internal/command"
)

var CommandRegistry = []command.ICommand{
	//command.NewInstallCommand(),
	//command.NewCurrentCommand(),
	//command.NewListCommand(),
	command.NewUseCommand(),
}

func RegisterCommands(rootCmd *cobra.Command) {
	for _, cmd := range CommandRegistry {
		cmd.Register(rootCmd)
	}
}
