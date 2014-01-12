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

> Don't like bash? The [examples](https://github.com/joewalnes/websocketd/tree/master/examples) directory contains the same program in [Python](https://raw.github.com/joewalnes/websocketd/master/examples/python/count.py), [Ruby](https://raw.github.com/joewalnes/websocketd/master/examples/ruby/count.rb), [Perl](https://raw.github.com/joewalnes/websocketd/master/examples/perl/count.pl) and [PHP](https://raw.github.com/joewalnes/websocketd/master/examples/php/count.php). If you're on Windows, there's also [JScript](https://raw.github.com/joewalnes/websocketd/master/examples/windows-jscript/count.js) and
[VBScript](https://raw.github.com/joewalnes/websocketd/master/examples/windows-vbscript/count.vbs).

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

[[console-count.png]]

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

[[example-count.png]]
