
package libwebsocketd

type Endpoint interface {
	Terminate()
	Output() chan string
	Send(msg string) bool
}
