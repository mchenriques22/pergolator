package codegen

import (
	"bytes"
	"go/parser"
	"go/token"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testStructFileName = "struct_test_gen.go"
const testStruct = `package codegen
type MyStruct struct {
	BasicString string
	BasicInt  int
}

type MyStruct2 struct {
	BasicStruct MyStruct
}
`

func TestRunWithSimpleStruct(t *testing.T) {
	// Set up the test
	file, err := os.OpenFile(testStructFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	require.NoError(t, err)
	defer func() {
		require.NoError(t, file.Close())
	}()
	defer func() {
		require.NoError(t, os.Remove(testStructFileName))
	}()
	_, err = file.WriteString(testStruct)

	// Run the test
	var buffer bytes.Buffer
	err = Run(&buffer, "codegen", "", []string{"github.com/mchenriques22/pergolator/codegen.MyStruct", "github.com/mchenriques22/pergolator/codegen.MyStruct2"}, 2, true)
	require.NoError(t, err)
	assert.Equal(t, 7183, buffer.Len())

	_, err = parser.ParseFile(token.NewFileSet(), "", buffer.String(), parser.AllErrors)
	assert.NoError(t, err)

}

func TestRunWithHTTPRequest(t *testing.T) {
	var _ http.Request
	// Run the test
	var buffer bytes.Buffer
	err := Run(&buffer, "codegen", "", []string{"net/http.Request"}, 2, true)
	require.NoError(t, err)
	assert.Equal(t, 76283, buffer.Len())

	_, err = parser.ParseFile(token.NewFileSet(), "", buffer.String(), parser.AllErrors)
	assert.NoError(t, err)
}
