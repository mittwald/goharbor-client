package goharborclient

import (
	"context"
	"flag"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/mittwald/goharbor-client/api/v1.10.0/client"
	"github.com/mittwald/goharbor-client/system"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	setupScript    = "scripts/setup-harbor.sh"
	teardownScript = "scripts/teardown-harbor.sh"
	host           = "localhost:30002"
	defaultUser    = "admin"
	password       = "Harbor12345"
)

var (
	swaggerClient = client.New(runtimeclient.New(host, "/api", []string{"http"}), strfmt.Default)
	authInfo      = runtimeclient.BasicAuth(defaultUser, password)

	integrationTest = flag.Bool("integration", false,
		"test against a real Harbor instance")
	harborVersion = flag.String("version", "1.10.2",
		"Harbor version, used in conjunction with -integration, "+
			"defaults to 1.10.2")
	skipSpinUp = flag.Bool("skip-spinup", false,
		"Skip kind cluster creation")
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
	c := system.NewClient(swaggerClient, authInfo)

	_, err := c.Health(ctx)
	require.NoError(t, err)
}
