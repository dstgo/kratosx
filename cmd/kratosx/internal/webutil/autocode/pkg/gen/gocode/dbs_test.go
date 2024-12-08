package gocode

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/dstgo/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen"
	"github.com/dstgo/kratosx/cmd/kratosx/internal/webutil/autocode/pkg/gen/types"
)

func TestGenDbs(t *testing.T) {
	initTable := func() *types.Table {
		var table types.Table
		content, err := os.ReadFile("internal/webutil/autocode/pkg/gen/auto.json")
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(content, &table); err != nil {
			panic(err)
		}
		return &table
	}

	builder := NewDbsBuilder(gen.NewBuilder(nil, initTable()))
	code, err := builder.GenDbs()
	if err != nil {
		panic(err)
	}
	fmt.Println(code)
}
