//go:build integration

package setup

import (
	"moneybits/tests/integration/setup/dirs"
	"path/filepath"

	"github.com/compose-spec/compose-go/v2/dotenv"
)

func LoadEnv(envFilepath string) error {
	rootDir := dirs.RootProjectDir()
	envFile := filepath.Join(rootDir, envFilepath)

	return dotenv.Load(envFile)
}
