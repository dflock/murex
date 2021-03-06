package datatools

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/utils/alter"
)

func init() {
	lang.GoFunctions["alter"] = cmdAlter
}

func cmdAlter(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	v, err := lang.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	s, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	new, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	path, err := alter.SplitPath(s)
	if err != nil {
		return err
	}

	v, err = alter.Alter(p.Context, v, path, new)
	if err != nil {
		return err
	}

	b, err := lang.MarshalData(p, dt, v)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
