package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewCmdApp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Manage launcher scripts in the apps directory",
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "add <name>...",
		Short: "Write apps/<name> for each binary found on PATH",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			for _, name := range args {
				if err := addLauncherApp(name); err != nil {
					return err
				}
				fmt.Fprintf(cmd.ErrOrStderr(), "added %s\n", name)
			}
			return nil
		},
	})
	return cmd
}

func addLauncherApp(name string) error {
	if !isSafeAppID(name) {
		return fmt.Errorf("invalid app id %q", name)
	}
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	return writeAgentApp(name)
}
