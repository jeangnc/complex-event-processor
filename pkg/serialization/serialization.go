package serialization

import (
	"encoding/json"
	"jeangnc/pattern-matcher/pkg/tree"
	"jeangnc/pattern-matcher/pkg/types"
	"log"
	"os"
)

func LoadJsonFile(filename string) *tree.ConditionTree {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", filename, err.Error())
	}

	t := tree.NewConditionTree()
	d := json.NewDecoder(f)

	for d.More() {
		c := types.Condition{}
		d.Decode(&c)
		t.Append(c)
	}

	return t
}
