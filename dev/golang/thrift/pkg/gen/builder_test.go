package gen

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	goMod = "go.mod"
)

var (
	stagingDir string
)

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err.Error)
	}
	stagingDir = filepath.Join(usr.HomeDir, "staging")
}

func TestGenerateThrift(t *testing.T) {
	path, err := os.Getwd()
	assert.Nil(t, err)
	thriftIdl := filepath.Join(path, "data", "social_network.thrift")
	env := []string{"STAGING=/root"}
	CmdExec(stagingDir, "thrift", env, "-r", "-gen", "go", "-out", stagingDir, thriftIdl)
}

func TestBuildThriftServer(t *testing.T) {

	//1. Create staging directory
	os.RemoveAll(stagingDir)
	err := os.Mkdir(stagingDir, 0755)
	assert.Nil(t, err)
	os.WriteFile(filepath.Join(stagingDir, goMod), []byte(tplGoMod), 0600)
	execModTidy(stagingDir)

	//2.
}

var tplGoMod = `module skyramp.com/thrift
require (
  github.com/apache/thrift v0.16.0
)
go 1.18
`
