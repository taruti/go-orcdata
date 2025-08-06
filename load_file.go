package orcdata

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

type orcFile struct {
	Rms []Data `json:"rms"`
}

func ReadOrcJSON(r io.Reader) ([]Data, error) {
	var f orcFile
	rd := bufio.NewReader(r)
	// Ignore UTF-8 BOMs (orc.org has these).
	x, _, err := rd.ReadRune()
	if err != nil {
		return nil, err
	}
	if x != 0xFEFF {
		rd.UnreadRune()
	}
	err = json.NewDecoder(rd).Decode(&f)
	for i := range f.Rms {
		f.Rms[i].SailNo = NormalizeSailNo(f.Rms[i].SailNo)
	}
	return f.Rms, err
}

func LoadOrcJSONFile(filename string) ([]Data, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadOrcJSON(f)
}
