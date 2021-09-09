package integration

import (
	"bytes"
	"encoding/json"
	"httpServerBareMetal/internal/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRoot(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Errorf("error creating request %v", err)
	}

	w := httptest.NewRecorder()

	s.GetRoot(w, r)

	if e, a := http.StatusOK, w.Code; e != a {
		t.Errorf("expected status code: %v, got status code: %v", e, a)
	}

	assert.Equal(t, "Hello Stranger", w.Body.String())
}

var (
	r   *http.Request
	b   bytes.Buffer
	err error
)

func TestSendJSON(t *testing.T) {
	td := []struct {
		Name         string
		Body         server.RequestInput
		ExpectedCode int
	}{
		{
			Name: "Foo",
			Body: server.RequestInput{
				Name: "Foo",
			},
			ExpectedCode: http.StatusOK,
		},
		{
			Name:         "400",
			Body:         server.RequestInput{},
			ExpectedCode: http.StatusOK,
		},
	}

	for _, test := range td {
		fn := func(t *testing.T) {
			if err := json.NewEncoder(&b).Encode(test.Body); err != nil {
				t.Errorf("error encoding request body: %v", err)
			}

			if r, err = http.NewRequest(http.MethodPost, "/json", &b); err != nil {
				t.Errorf("error creating request %v", err)
			}

			w := httptest.NewRecorder()
			s.SendJSON(w, r)

			assert.Equal(t, test.ExpectedCode, w.Code)
		}
		t.Run(test.Name, fn)
	}
}
