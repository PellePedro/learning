package gen

import (
	"bytes"
	"html/template"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	UserService = "UserService"
)

func TestMockService(t *testing.T) {
	sp, err := GenerateMockService(UserService, "./data/social_network.thrift")
	assert.Nil(t, err)

	tpl, err := template.New("tplServiceMock").Parse(tplServiceMock)
	assert.Nil(t, err)
	var b bytes.Buffer
	err = tpl.Execute(&b, sp)
	assert.Nil(t, err)
	os.WriteFile("mock.txt", b.Bytes(), 0666)
}
