package utils

import (
	"io"
	"os"

	"github.com/anthony-dong/fsync/logs"
)

func CopyFile(from string, to string) (rErr error) {
	fromFileStat, err := os.Stat(from)
	if err != nil {
		return err
	}
	if Exist(to) {
		toFileStat, err := os.Stat(to)
		if err != nil {
			return err
		}
		if fromFileStat.ModTime().Unix() == toFileStat.ModTime().Unix() {
			logs.Debug("M(Cached) %s", FormatFile(to))
			return nil
		}
		logs.Info("M %s", FormatFile(to))
	} else {
		logs.Info("C %s", FormatFile(to))
	}
	defer func() {
		if rErr != nil {
			return
		}
		if err := os.Chtimes(to, fromFileStat.ModTime(), fromFileStat.ModTime()); err != nil {
			rErr = err
		}
	}()
	if fromFileStat.IsDir() {
		if Exist(to) {
			return nil
		}
		return os.Mkdir(to, fromFileStat.Mode())
	}
	fromFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromFile.Close()
	toFile, err := os.OpenFile(to, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fromFileStat.Mode())
	if err != nil {
		logs.Error("create file find err: %v, file: %s", err, FormatFile(to))
		return err
	}
	defer func() {
		if err := toFile.Close(); err != nil {
			rErr = err
		}
	}()
	if _, err := io.Copy(toFile, fromFile); err != nil {
		return err
	}
	return nil
}
