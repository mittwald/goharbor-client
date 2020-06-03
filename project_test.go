package goharborclient

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestAPIProjectNew(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, name, 3, 3)
	defer c.Projects().Delete(ctx, p)

	require.NoError(t, err)
	assert.Equal(t, name, p.Name)
	assert.False(t, p.Deleted)
}

func TestAPIProjectGet(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, name, 3, 3)
	require.NoError(t, err)
	defer c.Projects().Delete(ctx, p)

	p2, err := c.Projects().Get(ctx, name)
	require.NoError(t, err)
	assert.Equal(t, p, p2)
}

func TestAPIProjectDelete(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"
	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, name, 3, 3)
	require.NoError(t, err)

	err = c.Projects().Delete(ctx, p)
	require.NoError(t, err)

	p, err = c.Projects().Get(ctx, name)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "project not found")
		_, ok := err.(*ProjectError)
		assert.True(t, ok)
	}
}

func TestAPIProjectList(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	namePrefix := "test-project"
	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%s-%d", namePrefix, i)
		p, err := c.Projects().NewProject(ctx, name, 3, 3)
		require.NoError(t, err)
		defer func() {
			err := c.Projects().Delete(ctx, p)
			if err != nil {
				panic("error in cleanup routine: " + err.Error())
			}
		}()
	}

	projects, err := c.Projects().List(ctx, namePrefix)
	require.NoError(t, err)
	assert.Len(t, projects, 10)
	for _, v := range projects {
		assert.Contains(t, v.Name, namePrefix)
	}
}

func TestAPIProjectUpdate(t *testing.T) {
	name := "test-project"
	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, name, 3, 3)
	require.NoError(t, err)
	defer c.Projects().Delete(ctx, p)
	require.Equal(t, "", p.Metadata.AutoScan)

	p.Metadata.AutoScan = "true"
	err = c.Projects().Update(ctx, p, 2, 2)
	require.NoError(t, err)
	p2, err := c.Projects().Get(ctx, name)
	require.NoError(t, err)
	assert.Equal(t, p, p2)
}
