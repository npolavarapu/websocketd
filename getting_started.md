This is a quick tutorial to get started with `websocketd`. By the end of it, you'll have:
*   a WebSocket server up and running
*   an endpoint that slowly counts to 10 on the server
*   a web-page that displays the current count from the server

This tutorial assumes you're familiar with the basics of web development (HTML, JavaScript) and the concept of [WebSockets](http://www.html5rocks.com/en/tutorials/websockets/basics/).

Remember, `websocketd` allows WebSocket end points to be written in virtually any programming language. In this tutorial, we're going to use a shell script written in bash to keep things simple.

Before you get started, make sure you have downloaded and installed `websocketd`.

Let's go!

## Understanding the basics

A WebSocket endpoint is a program that you expose over a WebSocket URL (e.g. `ws://someserver/my-program`). Everytime a browser connects to that URL, the websocketd server will start a new instance of your process. When the browser disconnects, the process will be stopped.

If there are 10 browser connected to your server, there will be 10 independent instances of your program running. `websocketd` takes care of listening for WebSocket connections and starting/stopping your program processes.

If your process needs to send a WebSocket message it should print the message contents to STDOUT, followed by a new line. `websocketd` will then send the message to the browser.

As you'd expect, if the browser sends a WebSocket message, `websocketd` will send it to STDIN of your process, followed by a new line.

## A simple endpoint

This simple shell script will count from 1 to 10 and then terminate, pausing for 1 second between each count.

__count.sh__:

```bash
#!/bin/bash
for COUNT in $(seq 1 10); do
  echo $COUNT
  sleep 1
done
```

If you don't like bash use another programming language.

Before turning it into a WebSocket server, let's test it from the command line. The beauty of `websocketd` is servers work equally well on the command line as they do in the server - with no modifications required.

```bash
$ chmod +x count.sh
$ ./count.sh
1
2
3
4
5
6
7
8
9
10
```

## Starting the WebSocket server

Now we have our program, it's time to turn it into a server. We start `websocketd`:

```bash
$ websocketd --port=8080 --devconsole ./count.sh
```

The parameters we supplied are:

    --port=8080    Port to listen on.
    --devconsole   Enable the development console (see below).
    ./count.sh     The program to run when a WebSocket connection is establish.

There is now a WebSocket end point, available at `ws://localhost:8080/`.

## Testing with the dev console

See that `--devconsole` flag? That enables a console built into `websocketd` to manually interact with the WebSocket endpoints.

Point your browser to `http://localhost:8080/`, and you'll see the console. Press the checkbox to connect.


## Building a web-page that connects to the WebSocket

Finally, let's create a web-page that to test it.

In JavaScript, you can connect to a WebSocket like this:

```javascript
var ws = new WebSocket('ws://localhost:8080/');
ws.onopen    = function() { ... do something on connect ... };
ws.onclose   = function() { ... do something on disconnect ... };
ws.onmessage = function(event) { ... do something with event.data ... };
```

Here is an example that ties it all together:
*   Open a WebSocket connection
*   Display the current count from the server
*   Indicate connection status by changing the background color

__count.html__:

```html
<!DOCTYPE html>
<html>
  <head>
    <title>websocketd count example</title>
    <style>
      #count {
      font: bold 150px arial;
      margin: auto;
      padding: 10px;
      text-align: center;
      }
    </style>
  </head>
  <body>
    
    <div id="count"></div>
    
    <script>
      var ws = new WebSocket('ws://localhost:8080/');
      ws.onopen = function() {
      document.body.style.backgroundColor = '#cfc';
      };
      ws.onclose = function() {
      document.body.style.backgroundColor = null;
      };
      ws.onmessage = function(event) {
      document.getElementById('count').innerText = event.data;
      };
    </script>
    
  </body>
</html>
```

Open this page in your web-browser. It will even work if you open it directly
from disk using a `file://` URL.

------------------------------------------------------------------------------------------------------------------------
To help with development, debugging and testing, websocketd includes a developer console, making it easy to interact with your WebSocket endpoints without having to write JavaScript.

To enable the console, add `--devconsole` to the websocketd command line options.

e.g.

    websocketd --port=1234 --dir=examples --devconsole

To test a WebSocket endpoint, open a web-browser and browse to its URL, replacing the `ws://` URL prefix with `http://`.

The console allows you to connect to WebSocket endpoint, send and receive test messages.

------------------------------------------------------------------------------------------------------------------------
WebSocket server programs often need additional data about the connection, beyond just the raw stream of messages - such as remote host, which web-page they're on, URL query parameters, cookies, etc.

To expose these to programs, they are set as environment variables for each request.

The variables conform to the 
[Common Gateway Interface (CGI) 1.1 Specification](http://tools.ietf.org/html/rfc3875).

You can read these environment variables directly, or use a programming language specific CGI library that can make them easier to work with (e.g. [Perl](http://perldoc.perl.org/CGI.html), [Python](http://docs.python.org/2/library/cgi.html), [Ruby](http://www.ruby-doc.org/stdlib-1.9.3/libdoc/cgi/rdoc/CGI.html)).


***


## AUTH_TYPE

Empty. Not supported in WebSockets.

[RFC 3875 - section 4.1.1](http://tools.ietf.org/html/rfc3875#section-4.1.1)


## CONTENT_LENGTH

Empty. Not supported in WebSockets.

[RFC 3875 - section 4.1.2](http://tools.ietf.org/html/rfc3875#section-4.1.2)


## CONTENT_TYPE

Empty. Not supported in WebSockets.

[RFC 3875 - section 4.1.3](http://tools.ietf.org/html/rfc3875#section-4.1.3)


## GATEWAY_INTERFACE

The string `websocketd-CGI/0.1`.

Future enhancements may change the version.

Example:

    GATEWAY_INTERFACE=websocketd-CGI/0.1

[RFC 3875 - section 4.1.4](http://tools.ietf.org/html/rfc3875#section-4.1.4)


## HTTP_COOKIE

HTTP Cookies, if present.

Example:

    HTTP_COOKIE=session=1234; prefs=foo`

[RFC 6265 - section 4.2](https://tools.ietf.org/html/rfc6265#section-4.2)


## HTTP_ORIGIN

The base URL of the origin site that initiated the WebSocket.

Example:

    HTTP_ORIGIN=https://github.com/


[RFC 6454 - section 4.2](http://tools.ietf.org/html/rfc6454)


## PATH_TRANSLATED

TODO

[RFC 3875 - section 4.1.6](http://tools.ietf.org/html/rfc3875#section-4.1.6)


## PATH_INFO

TODO

[RFC 3875 - section 4.1.5](http://tools.ietf.org/html/rfc3875#section-4.1.5)


## QUERY_STRING

Everything after the `?` in the URL.

Example: If the requested WebSocket URL is `ws://localhost:1234/somedir/myscript.py?name=me&msg=hello%20world`, then:

    QUERY_STRING=name=me&msg=hello%20world

[RFC 3875 - section 4.1.7](http://tools.ietf.org/html/rfc3875#section-4.1.7)


## REMOTE_ADDR

IP address of remote WebSocket client.

Example:

    REMOTE_ADDR=123.123.123.123

[RFC 3875 - section 4.1.8](http://tools.ietf.org/html/rfc3875#section-4.1.8)


## REMOTE_HOST

Reverse DNS lookup of remote WebSocket client.

Example:

    REMOTE_HOST=somemachine.someisp.com

If the reverse DNS lookup fails, or if `--reverselookup=false` is specified on the `websocketd` command line, this value will fallback to the same value of `REMOTE_ADDR`.

[RFC 3875 - section 4.1.9](http://tools.ietf.org/html/rfc3875#section-4.1.9)


## REMOTE_IDENT

Empty. Not supported in WebSockets.

[RFC 3875 - section 4.1.10](http://tools.ietf.org/html/rfc3875#section-4.1.10)


## REMOTE_PORT

Source port of remote WebSocket client.

Example:

    REMOTE_PORT=52696

(Non standard)


## REMOTE_USER

Empty. Not supported in WebSockets.

[RFC 3875 - section 4.1.11](http://tools.ietf.org/html/rfc3875#section-4.1.11)


## REQUEST_METHOD

The HTTP request method. For WebSockets, this is always `GET`.

Example:

    REQUEST_METHOD=GET

[RFC 3875 - section 4.1.12](http://tools.ietf.org/html/rfc3875#section-4.1.12)


## REQUEST_URI

The original request URI, as sent by the WebSocket client. This does not include scheme, host or port.

Example: If the requested WebSocket URL is `ws://localhost:1234/somedir/myscript.py?name=me&msg=hello%20world`, then:

    REQUEST_URI=/somedir/myscript.py?name=me&msg=hello%20world

(Non standard)


## SCRIPT_NAME

TODO

[RFC 3875 - section 4.1.13](http://tools.ietf.org/html/rfc3875#section-4.1.13)


## SERVER_NAME

The hostname of the server, as specified in the URL.

Example:

    SERVER_NAME=www.example.com

[RFC 3875 - section 4.1.14](http://tools.ietf.org/html/rfc3875#section-4.1.14)


## SERVER_PORT

The listening port of the server.

Example:

    SERVER_PORT=8080

[RFC 3875 - section 4.1.15](http://tools.ietf.org/html/rfc3875#section-4.1.15)


## SERVER_PROTOCOL

The HTTP protocol, as specified by the client.

Example:

    SERVER_PROTOCOL=HTTP/1.1

[RFC 3875 - section 4.1.16](http://tools.ietf.org/html/rfc3875#section-4.1.16)


## SERVER_SOFTWARE

The string `websocketd/x.x.x.x`, where `x.x.x.x` is the version of the `websocketd` program.

Example:

    SERVER_SOFTWARE=websocketd/1.0.0.0

[RFC 3875 - section 4.1.17](http://tools.ietf.org/html/rfc3875#section-4.1.17)


## UNIQUE_ID

A unique string associated with each WebSocket connection. It can be used for logging and debugging purposes.

The ID should be treated as an opaque payload. It must be treated as a string, not a number. The implementation may change over time.

The ID is only guaranteed to unique within a single `websocketd` process.

Example:

    UNIQUE_ID=45462465645449101442

(Non standard)


***

# HTTP Headers
Per the CGI specification, all HTTP headers will be set as environment variables. The header name will be transformed:

*   String converted to upper-case
*   Dashes `-` converted to underscores `_`
*   Prefixed with `HTTP_`

Example: If the HTTP header `Sec-WebSocket-Version: 13` is present, then:

    HTTP_SEC_WEBSOCKET_VERSION=13

[RFC 3875 - section 4.1.18](http://tools.ietf.org/html/rfc3875#section-4.1.18)
------------------------------------------------------------------------------------------------------------------------
It is possible to embed `websocketd` directly inside an existing go project by registering a handler. 

Here is an example:

```go
package main

import (
    "fmt"
    wsd "github.com/joewalnes/websocketd/libwebsocketd"
    "net/http"
)

func main() {
    // A log scope allows you to customize the logging that websocketd performs. 
    //You can provide your own log scope with a log func.
    logScope := wsd.RootLogScope(wsd.LogAccess, func(l *wsd.LogScope, 
        level wsd.LogLevel, levelName string,
        category string, msg string, args ...interface{}) {
        fmt.Println(args...)
    })
    // Configuration options tell websocketd where to look for programs to
    // run as WebSockets.
    config := &wsd.Config{
        ScriptDir:      "./ws-bin",
        UsingScriptDir: true,
        StartupTime:    time.Now(),
        DevConsole:     true,
    }
    // Register your route and handler.
    http.HandleFunc("/ws-bin/", func(rw http.ResponseWriter, req *http.Request) {
        handler := http.StripPrefix("/ws-bin", wsd.HttpWsMuxHandler{
            Config: config,
            Log:    logScope,
        })
        handler.ServeHTTP(rw, req)
    })
    if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
        fmt.Println("could not start server!", err)
        os.Exit(1)
    }
}
```
Make sure you run `go get` to fetch the dependency.
------------------------------------------------------------------------------------------------------------------------
Install Go
----------

Use apt, homebrew etc or install directly from source. Follow the instructions at http://golang.org/doc/install.

Building Websocketd
-------------------

Make sure you have set your `$GOPATH` to somewhere, e.g: `$HOME/gocode`:

```bash
$ mkdir -p $GOPATH/src/github.com/npolavarapu
$ cd $GOPATH/src/github.com/npolavarapu
$ git clone https://github.com/npolvarapu/websocketd
$ cd websocketd
```

Build:

```bash
$ make
```

Run:

```bash
$ ./websocketd
```
