package orcdata

import (
	"encoding/json"
	"testing"
)

func TestSailNo(t *testing.T) {
	n := NormalizeSailNo(" fin   - 3 1 8  ")
	if n != "FIN318" {
		t.Fatal("Sail no normalization failure", n)
	}
}

func TestLoad(t *testing.T) {
	var d Data
	err := json.Unmarshal([]byte(sampleDataString), &d)
	if err != nil {
		t.Fatal(err)
	}
}

const sampleDataString = `{
}`
