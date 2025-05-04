package utils

import (
	"math"
	"strings"
)

func ToSlug(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder

	spaceFound := false
	for _, r := range s {
		switch {
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'):
			if spaceFound {
				b.WriteRune('-')
				spaceFound = false
			}
			b.WriteRune(r)
		case r == ' ':
			spaceFound = true
		}
	}

	return b.String()
}

func roundToTwoDecimal(val float64) float32 {
	return float32(math.Round(val*100) / 100)
}
