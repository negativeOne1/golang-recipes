package integration

import (
	"flag"
	"httpServerBareMetal/internal/server"
	"os"
	"testing"
)

var s *server.Server

var serverPort = flag.String("port", ":8080", "server port")

func TestMain(m *testing.M) {
	flag.Parse()

	s = server.New()

	os.Exit(m.Run())
}
