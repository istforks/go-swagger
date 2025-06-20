package commands

import "github.com/istforks/go-swagger/cmd/swagger/commands/initcmd"

// InitCmd is a command namespace for initializing things like a swagger spec.
type InitCmd struct {
	Model *initcmd.Spec `command:"spec"`
}

// Execute provides default empty implementation
func (i *InitCmd) Execute(args []string) error {
	return nil
}
