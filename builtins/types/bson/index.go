package bson

import (
	"github.com/lmorg/murex/lang"
	"labix.org/v2/mgo/bson"
)

func readIndex(p *lang.Process, params []string) error {
	var jInterface interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = bson.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	return lang.IndexTemplateObject(p, params, &jInterface, bson.Marshal)
}
