//go:build integration

package stocks

import (
	"log"
	"moneybits/core"
	"moneybits/tests/integration/setup"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tickerTestAPP *core.AppContainer

func TestMain(m *testing.M) {
	err := setup.LoadEnv(".env.integration")
	if err != nil {
		log.Fatalf("failed to load env file: %s", err.Error())
	}

	tickerTestAPP = core.NewAppContainer()

	code := m.Run()

	os.Exit(code)
}

func TestFetchTicketData(t *testing.T) {
	t.Skip("temp skipping while the mountabank setup it not done")

	restCli := setup.NewTestRestClient(tickerTestAPP.HTTPServer.Server)

	resp, err := restCli.Get("/ticker/BBAS3")
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusOK, resp.StatusCode)
}
