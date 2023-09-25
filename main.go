package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"github.com/anthony-dong/fsync/command"
	"github.com/anthony-dong/fsync/logs"
)

func runCommand() error {
	logLevel := ""
	cmd := &cobra.Command{
		Use:                   command.App,
		Version:               command.AppVersion,
		Short:                 fmt.Sprintf(`The File Sync CLI %s`, command.AppVersion),
		SilenceUsage:          true, // 禁止失败打印 --help
		SilenceErrors:         true, // 禁止框架打印异常
		DisableFlagsInUseLine: true,
		CompletionOptions:     cobra.CompletionOptions{DisableDefaultCmd: true},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			logs.SelLevelString(logLevel)
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "set the log level(debug|info|notice|warn|error)")
	if err := command.AddCmd(cmd, command.NewCopyCmd); err != nil {
		return err
	}
	return cmd.ExecuteContext(context.Background())
}

func main() {
	command.InitDeferTask()
	wg := errgroup.Group{}
	wg.Go(func() error {
		command.RunDeferTask()
		return nil
	})
	wg.Go(func() error {
		defer command.CloseDeferTask()
		if err := runCommand(); err != nil && err != command.ErrorDone {
			return err
		}
		return nil
	})
	if err := wg.Wait(); err != nil {
		panic(err)
	}
}
