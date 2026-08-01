package config

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigInit(t *testing.T) {
	tmpDir := t.TempDir()

	envDir := filepath.Join(tmpDir, "env")
	err := os.MkdirAll(envDir, 0755)
	require.NoError(t, err)

	defaultYaml := `
server:
  address: "localhost:8080"
http:
  auth:
    secret: "test-secret"
security:
  whitelisted_hosts:
    - "localhost:8080"
  ssl_redirects: false
`
	devYaml := `
server:
  address: "localhost:9090"
`

	err = os.WriteFile(filepath.Join(envDir, "default.yaml"), []byte(defaultYaml), 0644)
	require.NoError(t, err)

	err = os.WriteFile(filepath.Join(envDir, "development.yaml"), []byte(devYaml), 0644)
	require.NoError(t, err)

	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	os.Chdir(tmpDir)

	Init("development")

	cfg := GetConfig()
	assert.Equal(t, "localhost:9090", cfg.GetString("server.address"))
	assert.Equal(t, "test-secret", cfg.GetString("http.auth.secret"))
	assert.Equal(t, []string{"localhost:8080"}, cfg.GetStringSlice("security.whitelisted_hosts"))
	assert.False(t, cfg.GetBool("security.ssl_redirects"))
}

func TestConfigGetConfigBeforeInit(t *testing.T) {
	mu.Lock()
	cfg = nil
	initOnce = sync.Once{}
	initErr = nil
	mu.Unlock()

	// log.Fatal calls os.Exit, which terminates the test process
	// We can't easily test this without refactoring config to return errors
	// This test is a placeholder to document the behavior
}