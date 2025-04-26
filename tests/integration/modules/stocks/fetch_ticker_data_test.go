//go:build integration

package stocks

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"moneybits/core"
	"moneybits/core/modules/stocks/api/dtos"
	"moneybits/tests/integration/setup"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tickerTestAPP *core.AppContainer

func TestMain(m *testing.M) {
	testCtx := context.Background()

	err := setup.LoadEnv(".env.integration")
	if err != nil {
		log.Fatalf("failed to load env file: %s", err.Error())
	}

	setup.ComposeUp(testCtx)

	tickerTestAPP = core.NewAppContainer()

	code := m.Run()

	setup.ComposeDown(testCtx)

	os.Exit(code)
}

func TestFetchTicketData(t *testing.T) {
	restCli := setup.NewTestRestClient(tickerTestAPP.HTTPServer.Server)

	resp, err := restCli.Get("/ticker/BBAS3")
	assert.Nil(t, err)

	var tickerDataResp dtos.TickerDataResponse
	b, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	err = json.Unmarshal(b, &tickerDataResp)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusOK, resp.StatusCode)

	assert.EqualValues(t, "BBAS3", tickerDataResp.Symbol)
	assert.EqualValues(t, "BRASIL      ON      NM", tickerDataResp.ShorName)
	assert.EqualValues(t, "Banco do Brasil S.A.", tickerDataResp.LongName)
	assert.EqualValues(t, "BRL", tickerDataResp.Currency)
}
