package gen

import (
	"bytes"
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMain(t *testing.T) {
	tOption := ServerTplData{
		Address:       "0.0.0.0:9090",
		Protocol:      ThriftProtocolCompact,
		Buffered:      true,
		Framed:        false,
		Secure:        false,
		MockedService: "UserService",
	}
	_ = tOption
	tpl, err := template.New("tplServer").Parse(tplServer)
	assert.Nil(t, err)
	var b bytes.Buffer
	err = tpl.Execute(&b, tOption)
	assert.Nil(t, err)
	f, err := os.Create("main.txt")
	assert.Nil(t, err)
	f.Write(b.Bytes())
	defer f.Close()

}
