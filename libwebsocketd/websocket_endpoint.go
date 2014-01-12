
package libwebsocketd

import (
	"io"

	"code.google.com/p/go.net/websocket"
)

type WebSocketEndpoint struct {
	ws     *websocket.Conn
	output chan string
	log    *LogScope
}

func NewWebSocketEndpoint(ws *websocket.Conn, log *LogScope) *WebSocketEndpoint {
	return &WebSocketEndpoint{
		ws:     ws,
		output: make(chan string),
		log:    log}
}

func (we *WebSocketEndpoint) Terminate() {
}

func (we *WebSocketEndpoint) Output() chan string {
	return we.output
}

func (we *WebSocketEndpoint) Send(msg string) bool {
	err := websocket.Message.Send(we.ws, msg)
	if err != nil {
		we.log.Trace("websocket", "Cannot send: %s", err)
		return false
	}
	return true
}

func (we *WebSocketEndpoint) ReadOutput(config *Config) {
	for {
		var msg string
		err := websocket.Message.Receive(we.ws, &msg)
		if err != nil {
			if err != io.EOF {
				we.log.Debug("websocket", "Cannot receive: %s", err)
			}
			break
		}
		we.output <- msg
	}
	close(we.output)
}
