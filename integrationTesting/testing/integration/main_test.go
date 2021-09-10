package integration

import (
	"flag"
	"httpServerBareMetal/internal/server"
	"os"
	"testing"
)

var s *server.Server

func TestMain(m *testing.M) {
	flag.Parse()

	s = server.New()

	os.Exit(m.Run())
}
