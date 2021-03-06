package tabulate

import (
	"strings"

	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/utils/json"
)

type writer interface {
	Write([]string) error
	Flush()
	Error() error
}

type mapWriter struct {
	m      map[string]string
	err    error
	out    stdio.Io
	joiner string
}

func newMapWriter(writer stdio.Io, joiner string) *mapWriter {
	return &mapWriter{
		m:      make(map[string]string),
		out:    writer,
		joiner: joiner,
	}
}

func (mw *mapWriter) Write(array []string) error {
	if len(array) < 2 || array[0] == "" {
		// silently ignore
		return nil
	}

	mw.m[array[0]] = strings.Join(array[1:], mw.joiner)
	return nil
}

func (mw *mapWriter) Flush() {
	var b []byte

	b, mw.err = json.Marshal(mw.m, mw.out.IsTTY())
	if mw.err != nil {
		return
	}

	_, mw.err = mw.out.Write(b)
	return
}

func (mw *mapWriter) Error() error {
	return mw.err
}
