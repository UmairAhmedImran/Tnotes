package note

import "github.com/spf13/cobra"

func Cmd() []*cobra.Command {
	return []*cobra.Command{addCmd, listCmd}
}
