package setup

import (
	"context"
	"fmt"
	"log"
	"moneybits/tests/integration/setup/dirs"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/compose-spec/compose-go/v2/dotenv"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	// ComposeTestStack provides all dependencie that is needed in the project.
	ComposeTestStack tc.ComposeStack

	// strategiesMap receive the service name defined in the docker-compose file with
	// its strategy of initialization
	strategiesMap = map[string]wait.Strategy{
		"postgres": wait.ForLog("database system is ready to accept connections").WithStartupTimeout(1 * time.Minute),
	}
)

func ComposeUp(ctx context.Context, t *testing.T) error {
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")

	rootDir := dirs.RootProjectDir()

	err := dotenv.Load(filepath.Join(rootDir, ".env.local"))
	if err != nil {
		t.Fatal(err)
	}

	composePath, err := dockerComposeFilePath()
	if err != nil {
		t.Fatal(err)
	}

	compose, err := tc.NewDockerComposeWith(
		tc.WithStackFiles(composePath),
		tc.StackIdentifier("moneybits"),
	)
	if err != nil {
		return err
	}

	for serviceName, strategy := range strategiesMap {
		compose.WaitForService(serviceName, strategy)
	}

	err = compose.Up(ctx, tc.Wait(true))
	if err != nil {
		t.Fatalf("failed to start compose stack: %v", err)
	}

	ComposeTestStack = compose

	return nil
}

func ComposeDown(ctx context.Context) {
	for _, service := range ComposeTestStack.Services() {
		container, err := ComposeTestStack.ServiceContainer(ctx, service)
		if err != nil {
			log.Fatalf("error getting container in terminate process: %s", err.Error())
		}

		err = container.Terminate(ctx)
		if err != nil {
			log.Fatalf("fail to terminate %s container: %s", service, err.Error())
		}
	}

	log.Println("compose down finished")
}

func dockerComposeFilePath() (string, error) {
	projectRoot := dirs.RootProjectDir()
	composePath := filepath.Join(projectRoot, "docker-compose.yml")

	if _, err := os.Stat(composePath); os.IsNotExist(err) {
		return "", fmt.Errorf("docker-compose.yml not found at: %s", composePath)
	}

	return composePath, nil
}
