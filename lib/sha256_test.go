package lib

import (
	"testing"
)

func TestSha256EncodeBase64(t *testing.T) {
	input := "input"
	expected := "yWxtW+jQihLntc3Bsgf6ayQwl0yGgD2IkWdedv2ZLCA="

	output := Sha256EncodeBase64(input)

	if output != expected {
		t.Fatalf("input %s was expected to hash to %s but hashed to %s", input, expected, output)
	}
}
