// +build integration

package goharborclient

import (
	"flag"
	runtimeclient "github.com/go-openapi/runtime/client"
	"os"
	"os/exec"
	"testing"
)

const (
	setupScript    = "scripts/setup-harbor.sh"
	teardownScript = "scripts/teardown-harbor.sh"
	host           = "localhost:30002"
	defaultUser    = "admin"
	password       = "Harbor12345"
)

var (
	authInfo      = runtimeclient.BasicAuth(defaultUser, password)
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

	if !*skipSpinUp {
		err := setupHarbor(*harborVersion)
		if err != nil {
			panic("error setting up harbor: " + err.Error())
		}
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
