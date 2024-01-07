package command

import "sdkman-go/internal/handler"

type ListCommand struct {
	*BaseCommand
}

func NewListCommand() *ListCommand {
	c := &ListCommand{BaseCommand: &BaseCommand{}}
	h := handler.ListHandler()
	c.BaseCommand.Init(
		"list",
		"List Command",
		"This is the List Command",
		h,
	)
	return c
}
