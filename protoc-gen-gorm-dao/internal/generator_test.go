package internal

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/pluginpb"
)

//go:embed testdata/CodeGeneratorRequest.json
var generatorRequest []byte

func Test_GenerateFile(t *testing.T) {
	request := &pluginpb.CodeGeneratorRequest{}
	if err := protojson.Unmarshal(generatorRequest, request); err != nil {
		t.Error(err)
	}
	plugin, err := protogen.Options{}.New(request)
	if err != nil {
		panic(err)
	}
	for _, f := range plugin.Files {
		g := GenerateFile(plugin, f)
		content, err := g.Content()
		assert.Nil(t, err)
		t.Log(string(content))
	}
}
