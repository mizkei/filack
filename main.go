package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/BurntSushi/toml"
	"github.com/nlopes/slack"
)

type Conf struct {
	Token   string   `toml:"token"`
	Filters []Filter `toml:"filters"`
}

type Filter struct {
	Channel    Channel  `toml:"channel"`
	RegexpList []string `toml:"regexp_list"`
	ReList     []*regexp.Regexp
}

type Channel struct {
	ID   string `toml:"id"`
	Name string `toml:"name"`
}

func (f *Filter) CheckText(msg string) bool {
	for _, re := range f.ReList {
		if re.MatchString(msg) {
			return true
		}
	}

	return false
}

func main() {
	var confPath string

	flag.StringVar(&confPath, "config", "$HOME/.config/filack/conf.toml", "config file path")
	flag.Parse()

	var conf Conf
	if _, err := toml.DecodeFile(os.ExpandEnv(confPath), &conf); err != nil {
		panic(err)
	}

	// check the regexp syntax
	for i, f := range conf.Filters {
		for _, s := range f.RegexpList {
			r, err := regexp.Compile(s)
			if err != nil {
				panic(err)
			}

			conf.Filters[i].ReList = append(conf.Filters[i].ReList, r)
		}
	}

	api := slack.New(conf.Token)
	rtm := api.NewRTM()

	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				for _, f := range conf.Filters {
					if f.Channel.ID != ev.Channel {
						continue
					}

					if f.CheckText(ev.Text) {
						fmt.Printf("#%s: %s\n", f.Channel.Name, ev.Text)
					}
				}
			case *slack.InvalidAuthEvent:
				log.Println("error: Invalid Auth")
				break Loop
			default:
			}
		}
	}
}
