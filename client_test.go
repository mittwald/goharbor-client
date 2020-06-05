package goharborclient

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	setupScript    = "scripts/setup-harbor.sh"
	teardownScript = "scripts/teardown-harbor.sh"
)

var (
	integrationTest = flag.Bool("integration", false,
		"test against a real Harbor instance")
	harborVersion = flag.String("version", "1.10.2",
		"Harbor version, used in conjunction with -integration, "+
			"defaults to 1.10.2")
	skipSpinUp = flag.Bool("skip-spinup", false,
		"Skip kind cluster creation")
	host            = "localhost:30002"
	defaultUser     = "admin"
	defaultPassword = "Harbor12345"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	flag.Parse()

	if *integrationTest && !*skipSpinUp {
		err := setupHarbor(*harborVersion)
		if err != nil {
			panic("error setting up harbor: " + err.Error())
		}
		defer teardownHarbor()
	}

	return m.Run()
}

func setupHarbor(version string) error {
	cmdPath, err := exec.LookPath(setupScript)
	if err != nil {
		return err
	}

	cmd := exec.Command(cmdPath, version)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func teardownHarbor() error {
	cmdPath, err := exec.LookPath(teardownScript)
	if err != nil {
		return err
	}

	cmd := exec.Command(cmdPath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func TestAPIHealth(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	_, err := c.Health(ctx)
	require.NoError(t, err)
}
