# filter slack message

#install

```shell-session
$ go get github.com/mizkei/filack
```

# config

create 'conf.toml'

```
token = "YOUR TOKEN"


[[filters]]
  [[filters.regexp_list]]
    query = "fail"
    flags = "im"
  [[filters.regexp_list]]
    query = "error"
    flags = "i"

  [filters.channel]
    id = "CHANNEL1 ID"
    name = "CHANNEL1 NAME"


[[filters]]
  [[filters.regexp_list]]
    query = "test\d"
    flags = "im"

  [filters.channel]
    id = "CHANNEL2 ID"
    name = "CHANNEL2 NAME"
```

default config file path is `$HOME/.config/filack/conf.toml`

command line options

```shell-session
$ filack --config=config/file/path
```

# usage

- linux

```shell-session
$ filack | while read text; do notify-send "$text"; done
```

- mac

```shell-session
$ filack | while read text; do osascript  -e 'display notification "'"$text"'" with title "slack"'; done
```
