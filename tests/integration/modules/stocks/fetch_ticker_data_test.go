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

func TestMain(m *testing.M) {
	err := setup.LoadEnv(".env.integration")
	if err != nil {
		log.Fatalf("failed to load env file: %s", err.Error())
	}

	code := m.Run()

	os.Exit(code)
}

func TestFetchTicketData(t *testing.T) {
	app := core.NewAppContainer()

	restCli := setup.NewTestRestClient(app.HTTPServer.Server)

	resp, err := restCli.Get("/ticker/BBAS3")
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusOK, resp.StatusCode)
}
