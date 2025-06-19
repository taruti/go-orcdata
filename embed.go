package orcdata

import _ "embed"
import (
	"github.com/klauspost/compress/zstd"
"io"
"bytes"
)

//go:embed orc.json.zst
var zstdJson []byte

type rc struct {
	*zstd.Decoder
}
func (r rc) Close() error {
	r.Decoder.Close()
	return nil
}

func EmbeddedJSON() io.ReadCloser {
	rd, _ := zstd.NewReader(bytes.NewReader(zstdJson), zstd.WithDecoderConcurrency(1))
	return rc{rd}
}
