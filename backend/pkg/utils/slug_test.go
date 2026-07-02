package utils

import "testing"

func TestSlugify(t *testing.T) {
	cases := map[string]string{
		"Kelistrikan Industri":   "kelistrikan-industri",
		"  K3 & 5S  ":            "k3-5s",
		"Smart Grid / AMI!":      "smart-grid-ami",
		"already-slug":           "already-slug",
		"UPPER_snake Case":       "upper-snake-case",
	}
	for in, want := range cases {
		if got := Slugify(in); got != want {
			t.Errorf("Slugify(%q) = %q, want %q", in, got, want)
		}
	}
}
