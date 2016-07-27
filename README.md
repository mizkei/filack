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
regexp_list = [
  "fail",
  "error",
]
[filters.channel]
id = "CHANNEL ID"
name = "CHANNEL NAME"


[[filters]]
regexp_list = [
  "check",
]
[filters.channel]
id = "CHANNEL ID2"
name = "CHANNEL NAME2"
```

# usage

- linux

```shell-session
$ ./filack | while read text; do notify-send "$text"; done
```

- mac

```shell-session
$ ./filack | while read text; do osascript  -e 'display notification "'"$text"'" with title "slack"'; done
```
