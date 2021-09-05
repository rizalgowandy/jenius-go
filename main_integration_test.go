package jenius

import (
	"flag"
	"os"
	"testing"

	"github.com/kokizzu/gotro/L"
	"github.com/rizalgowandy/jenius-go/pkg/api"
)

// How to run all integration test:
// $ KEY=REAL_API_KEY go test -v . -run . -integration

var (
	integration bool
	client      *Client
)

func TestMain(m *testing.M) {
	flag.BoolVar(&integration, "integration", false, "enable integration test")
	flag.Parse()

	if !integration {
		os.Exit(m.Run())
	}

	var err error
	client, err = NewClient(api.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Debug:        true,
	})
	if L.IsError(err, "client: create failure") {
		os.Exit(1)
	}

	os.Exit(m.Run())
}
