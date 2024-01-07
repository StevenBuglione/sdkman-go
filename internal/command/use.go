package command

import "sdkman-go/internal/handler"

type UseCommand struct {
	*BaseCommand
}

func NewUseCommand() *UseCommand {
	c := &UseCommand{BaseCommand: &BaseCommand{}}
	h := handler.UseHandler()
	c.BaseCommand.Init(
		"use",
		"Use Command",
		"This is the Use Command",
		h,
	)
	return c
}
