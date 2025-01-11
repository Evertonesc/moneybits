//go:build integration

package transactions

import (
	"context"
	"moneybits/core"
	"moneybits/tests/integration/setup"
	"moneybits/tests/integration/setup/containers"
	"moneybits/tests/integration/setup/dirs"
	"testing"
)

type userFixture struct {
	Users []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"users"`
}

const (
	fixturesPath = "fixtures/transactions.json"
)

func TestCreateTransactionsAPI(t *testing.T) {
	ctx := context.Background()

	setup.ComposeUp(ctx, t)
	defer setup.ComposeDown(ctx)

	db, err := containers.PostgresDB(ctx)
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}

	var userFixture userFixture
	err = db.Load(
		ctx,
		dirs.BuildFixturePath("fixtures/users.json"),
		&userFixture,
	)
	if err != nil {
		t.Error(err)
	}

	for _, user := range userFixture.Users {
		err := db.Execute(
			ctx,
			"INSERT INTO users (name, email) VALUES ($1, $2)",
			user.Name, user.Email,
		)
		if err != nil {
			t.Errorf("failed to insert user fixture: %v", err)
		}
	}

	app := core.NewAppContainer()
	client := setup.NewTestRestClient(app.HTTPServer.Server)
	defer client.Server.Close()

	t.Run("get user", func(t *testing.T) {
		// var response map[string]interface{}
		// resp, err := client.Get("/users/1", &response)
		// if err != nil {
		// 	t.Errorf("failed to get user: %v", err)
		// }

		// if resp.StatusCode != 200 {
		// 	t.Errorf("expected status code 200, got %d", resp.StatusCode)
		// }

		// if response["name"] != userFixture.Users[0].Name {
		// 	t.Errorf("expected name %s, got %s", userFixture.Users[0].Name, response["name"])
		// }
	})

	t.Run("create user", func(t *testing.T) {
		// newUser := map[string]string{
		// 	"name":  "New User",
		// 	"email": "new@example.com",
		// }

		// var response map[string]interface{}
		// resp, err := client.Post("/users", newUser, &response)
		// if err != nil {
		// 	t.Fatalf("failed to create user: %v", err)
		// }

		// if resp.StatusCode != 201 {
		// 	t.Errorf("expected status code 201, got %d", resp.StatusCode)
		// }

		// if response["name"] != newUser["name"] {
		// 	t.Errorf("expected name %s, got %s", newUser["name"], response["name"])
		// }
	})
}
