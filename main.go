package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joewalnes/websocketd/libwebsocketd"
)

func log(l *libwebsocketd.LogScope, level libwebsocketd.LogLevel, levelName string, category string, msg string, args ...interface{}) {
	if level < l.MinLevel {
		return
	}
	fullMsg := fmt.Sprintf(msg, args...)

	assocDump := ""
	for index, pair := range l.Associated {
		if index > 0 {
			assocDump += " "
		}
		assocDump += fmt.Sprintf("%s:'%s'", pair.Key, pair.Value)
	}

	l.Mutex.Lock()
	fmt.Printf("%s | %-6s | %-10s | %s | %s\n", libwebsocketd.Timestamp(), levelName, category, assocDump, fullMsg)
	l.Mutex.Unlock()
}

func main() {
	flag.Usage = PrintHelp
	config := parseCommandLine()

	log := libwebsocketd.RootLogScope(config.LogLevel, log)

	if config.DevConsole {
		if config.StaticDir != "" {
			log.Fatal("server", "Invalid parameters: --devconsole cannot be used with --staticdir. Pick one.")
			os.Exit(4)
		}
		if config.CgiDir != "" {
			log.Fatal("server", "Invalid parameters: --devconsole cannot be used with --cgidir. Pick one.")
			os.Exit(4)
		}
	}

	http.Handle(config.BasePath, libwebsocketd.HttpWsMuxHandler{
		Config: config.Config,
		Log:    log})

	log.Info("server", "Starting WebSocket server   : ws://%s%s", config.Addr, config.BasePath)
	if config.DevConsole {
		log.Info("server", "Developer console enable  d : http://%s/", config.Addr)
	}
	if config.UsingScriptDir {
		log.Info("server", "Serving from directory      : %s", config.ScriptDir)
	} else if config.CommandName != "" {
		log.Info("server", "Serving using application   : %s %s", config.CommandName, strings.Join(config.CommandArgs, " "))
	}
	if config.StaticDir != "" {
		log.Info("server", "Serving static content from : %s", config.StaticDir)
	}
	if config.CgiDir != "" {
		log.Info("server", "Serving CGI scripts from    : %s", config.CgiDir)
	}

	err := http.ListenAndServe(config.Addr, nil)
	if err != nil {
		log.Fatal("server", "Could start server: %s", err)
		os.Exit(3)
	}
}
