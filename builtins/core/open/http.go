package open

import (
	"io"

	"github.com/lmorg/murex/builtins/core/httpclient"
	"github.com/lmorg/murex/lang"
)

func http(p *lang.Process, url string) (io.ReadCloser, string, error) {
	resp, err := httpclient.Request(p.Context, "GET", url, nil, p.Config, true)

	if err != nil {
		return nil, "", err
	}

	dt := lang.MimeToMurex(resp.Header.Get("Content-Type"))

	// TODO: insert something about content-length detection

	return resp.Body, dt, nil
}
