package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	help = `
{{binary}} ({{version}})

{{binary}} is a command line tool that will allow any executable program
that accepts input on stdin and produces output on stdout to be turned into
a WebSocket server.

Usage:

  Export a single executable program a WebSocket server:

    {{binary}} [options] program [program args]

  Or, export an entire directory of executables as WebSocket endpoints:

    {{binary}} [options] --dir=SOMEDIR

Options:

  --port=PORT                    HTTP port to listen on.

  --address=ADDRESS              Address to bind to. 
                                 Default: 0.0.0.0 (all)

  --basepath=PATH                Base path in URLs to serve from. 
                                 Default: / (root of domain)

  --reverselookup={true,false}   Perform DNS reverse lookups on remote clients.
                                 Default: true

  --dir=DIR                      Allow all scripts in the local directory
                                 to be accessed as WebSockets. If using this,
                                 option, then the standard program and args
                                 options should not be specified.

  --staticdir=DIR                Serve static files in this directory over HTTP.

  --cgidir=DIR                   Serve CGI scripts in this directory over HTTP.

  --help                         Print help and exit.

  --version                      Print version and exit.

  --license                      Print license and exit.

  --devconsole                   Enable interactive development console.
                                 This enables you to access the websocketd
                                 server with a web-browser and use a
                                 user interface to quickly test WebSocket
                                 endpoints. For example, to test an
                                 endpoint at ws://[host]/foo, you can
                                 visit http://[host]/foo in your browser.
                                 This flag cannot be used in conjunction
                                 with --staticdir or --cgidir.

  --loglevel={debug,             Log level to use (from most to least
              trace,             verbose).
              access,            Default: access
              info,
              error,
              fatal}


` )

func PrintHelp() {
	msg := strings.Trim(help, " \n")
	msg = strings.Replace(msg, "{{binary}}", filepath.Base(os.Args[0]), -1)
	msg = strings.Replace(msg, "{{version}}", Version(), -1)
	fmt.Fprintf(os.Stderr, "%s\n", msg)
}
