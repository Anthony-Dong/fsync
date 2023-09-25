package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	ignore "github.com/sabhiram/go-gitignore"

	"github.com/spf13/cobra"

	"github.com/anthony-dong/fsync/logs"
	"github.com/anthony-dong/fsync/utils"
)

func NewCopyCmd() (*cobra.Command, error) {
	helper := CopyHelper{
		done:    make(chan bool),
		Exclude: make([]string, 0),
	}
	cmd := &cobra.Command{
		Use:   "copy [--from DIR] [--to DIR] [--exclude-from FILE] [--exclude PATTERN]",
		Short: `For copying directories`,
		RunE: func(cmd *cobra.Command, args []string) error {
			RegisterDeferTask(func() {
				helper.Close()
			})
			return helper.Run()
		},
	}
	cmd.Flags().StringVarP(&helper.From, "from", "f", "", "the from dir")
	cmd.Flags().StringVarP(&helper.To, "to", "t", "", "the to dir")
	cmd.Flags().StringVar(&helper.ExcludeFrom, "exclude-from", "", "read exclude patterns from FILE")
	cmd.Flags().StringSliceVar(&helper.Exclude, "exclude", []string{}, "exclude files matching PATTERN")
	return cmd, nil
}

type CopyHelper struct {
	From        string
	To          string
	ExcludeFrom string   // read exclude patterns from FILE
	Exclude     []string // exclude patterns
	ignore      *ignore.GitIgnore
	done        chan bool
	doneOnce    sync.Once
}

func (c *CopyHelper) Close() {
	c.doneOnce.Do(func() {
		close(c.done)
	})
}

func (c *CopyHelper) init() error {
	var err error
	if c.From, err = filepath.Abs(c.From); err != nil {
		return err
	}
	if c.To, err = filepath.Abs(c.To); err != nil {
		return err
	}
	logs.Info("copy %s -> %s", utils.FormatFile(c.From), utils.FormatFile(c.To))
	if !utils.Exist(c.From) {
		return fmt.Errorf(`not exist file %s`, c.From)
	}
	if !utils.Exist(c.To) {
		return fmt.Errorf(`not exist file %s`, c.To)
	}

	if c.ExcludeFrom != "" {
		if c.ExcludeFrom, err = filepath.Abs(c.ExcludeFrom); err != nil {
			return err
		}
		if !utils.Exist(c.ExcludeFrom) {
			return fmt.Errorf(`not exist exclude file %s`, c.To)
		}
		if c.ignore, err = ignore.CompileIgnoreFileAndLines(c.ExcludeFrom, c.Exclude...); err != nil {
			return err
		}
	}
	if len(c.Exclude) > 0 && c.ignore == nil {
		c.ignore = ignore.CompileIgnoreLines(c.Exclude...)
	}
	return nil
}

func (c *CopyHelper) Run() error {
	if err := c.init(); err != nil {
		return err
	}
	return utils.WalkDirFiles(c.From, func(base string, from string, fs os.FileInfo) error {
		if strings.HasPrefix(from, c.To) {
			return nil
		}
		relPath, err := filepath.Rel(base, from)
		if err != nil {
			return err
		}
		if c.ignore != nil && c.ignore.MatchesPath(relPath) {
			logs.Debug("C(Ignore) %s", relPath)
			return nil
		}
		select {
		case <-c.done:
			return ErrorDone
		default:
		}

		to := filepath.Join(c.To, relPath)
		return utils.CopyFile(from, to)
	})
}
