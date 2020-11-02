// +build integration

package apiv2

import (
	"flag"
	"os"
	"os/exec"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	integrationtest "github.com/mittwald/goharbor-client/apiv2/testing"
)

var (
	authInfo      = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	harborVersion = flag.String("version", "2.1.1",
		"Harbor version, used in conjunction with -integration, "+
			"defaults to 2.1.1")
	skipSpinUp = flag.Bool("skip-spinup", false,
		"Skip kind cluster creation")
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	flag.Parse()

	if !*skipSpinUp {
		err := setupHarbor(*harborVersion)
		if err != nil {
			panic("error setting up harbor: " + err.Error())
		}
	}

	return m.Run()
}

func setupHarbor(version string) error {
	cmdPath, err := exec.LookPath(integrationtest.SetupScript)
	if err != nil {
		return err
	}

	cmd := exec.Command(cmdPath, version)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
