package orcdata

import "strings"

type Data struct {
	SailNo       string // Needs to be normalized
	YachtName    string
	Class        string
	LOA          float64
	OSN          float64
	GPH          float64
	FinRatingTOT float64 `json:"FIN_FinRating_TOT"`
	Allowances   OrcAllowance
}

func NormalizeSailNo(raw string) string {
	var b strings.Builder
	for i := 0; i < len(raw); i++ {
		ch := raw[i]
		switch {
		case ch >= 'a' && ch <= 'z':
			ch -= 'a' - 'A'
			fallthrough
		case ch >= '0' && ch <= '9':
			fallthrough
		case ch >= 'A' && ch <= 'Z':
			b.WriteByte(ch)
		case ch == ' ' || ch == '-' || ch == '\t':
			continue
		default:
			b.WriteByte('X')
		}
	}
	return b.String()
}
