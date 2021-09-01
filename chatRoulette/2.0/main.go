package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/net/websocket"
)

const listenAddr = "localhost:8080"

var partner = make(chan io.ReadWriteCloser)

type socket struct {
	*websocket.Conn
	done chan bool
}

func (s socket) Read(b []byte) (int, error)  { return s.Conn.Read(b) }
func (s socket) Write(b []byte) (int, error) { return s.Conn.Write(b) }
func (s socket) Close() error                { s.done <- true; return nil }

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/sock", websocket.Handler(socketHandler))
	http.ListenAndServe(listenAddr, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, listenAddr)
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<script>
websocket = new WebSocket("ws://{{.}}/sock");
websocket.onmessage = function(m) { console.log("Received:", m.data); }
</script>
</html>
`))

func socketHandler(ws *websocket.Conn) {
	s := socket{ws, make(chan bool)}
	go match(s)
	<-s.done
}

func handler(c *websocket.Conn) {
	var s string
	fmt.Fscan(c, &s)
	fmt.Println("Received:", s)
	fmt.Fprint(c, "How do you do?")
}

func match(c io.ReadWriteCloser) {
	fmt.Fprintln(c, "Waiting for partner")
	select {
	case partner <- c:
	case p := <-partner:
		chat(p, c)
	}
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Found one! Say hi.")
	fmt.Fprintln(b, "Found one! Say hi.")
	errc := make(chan error, 1)
	go cp(a, b, errc)
	go cp(b, a, errc)
	if err := <-errc; err != nil {
		log.Println(err)
	}
	a.Close()
	b.Close()
}

func cp(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err
}
