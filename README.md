# fsync

fsync 目前对于跨平台不太好，我主要是做一些文件拷贝(例如拷贝到磁盘中)，因此需要一个跨平台工具.

# 使用

一、下载

前往 release 下载即可

二、介绍

```shell
➜  fsync git:(dev) ✗ fsync            
The File Sync CLI 1.0.0

Usage:
  fsync [command]

Available Commands:
  copy        For copying directories
  help        Help about any command

Flags:
  -h, --help               help for fsync
      --log-level string   set the log level(debug|info|notice|warn|error) (default "info")
  -v, --version            version for fsync

Use "fsync [command] --help" for more information about a command.
```

# 复制场景

```shell
➜  fsync git:(dev) ✗ fsync copy --help
For copying directories

Usage:
  fsync copy [--from DIR] [--to DIR] [--exclude-from FILE] [--exclude PATTERN] [flags]

Flags:
      --exclude strings       exclude files matching PATTERN
      --exclude-from string   read exclude patterns from FILE
  -f, --from string           the from dir
  -h, --help                  help for copy
  -t, --to string             the to dir

Global Flags:
      --log-level string   set the log level(debug|info|notice|warn|error) (default "info")
```

## Mac

```shell
cd ~/Pictures/Photos\ Library.photoslibrary/originals
bin/darwin_amd64/mac_tools -from '.' -to '/Volumes/fanaodong/photos/originals'
```

## Windows

```shell
C:\Users\bella\Desktop\mac_tools\bin\windows_amd64\mac_tools.exe copy --from .  --to E:\photos\windows\huawei
```

## Linux

同 mac

# 缺陷

1. linux 文件是没有create time属性，mac/windows都是存在的，但是Go没有提供，这里需要采用system call的方式去修改了，后续支持
2. 支持文件协议提供远程copy(参考rsync)
