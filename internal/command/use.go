package command

import "sdkman-go/internal/handler"

type UseCommand struct {
	*Command
}

func NewUseCommand() ICommand {
	c := &UseCommand{Command: &Command{}}
	h := handler.NewUseHandler()
	c.Command.Init(
		"use",
		"Use Command",
		"This is the Use Command",
		h,
	)
	return c
}
