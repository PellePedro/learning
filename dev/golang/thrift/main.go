package main

import (
	"os"
	"os/user"
	"path/filepath"
	"thrift/pkg/gen"
)

const (
	staging            = "staging"
	mockFileName       = "api.go"
	idlFile            = "social_network.thrift"
	UserService        = "UserService"
	ComposePostService = "ComposePostService"
)

func main() {

	usr, err := user.Current()
	cwd, err1 := os.Getwd()
	if err != nil && err1 != nil {
		panic(err.Error)
	}

	// 1. Generate GO code from thrift IDL
	stagingDir := filepath.Join(usr.HomeDir, staging)
	idl := filepath.Join(cwd, idlFile)
	os.RemoveAll(stagingDir)
	os.MkdirAll(stagingDir, 0750)
	gen.CompileThriftToGo(stagingDir, idl, []string{})

	// 2.
	mockedApiFile := filepath.Join(stagingDir, mockFileName)
	mockedApiFile = mockFileName
	gen.GenerateMockAPI("PostStorageService", mockedApiFile, idl)

	// 3.

}
