package myip

import "testing"

func TestParseEmptyRemoteAddr(t *testing.T) {
	if result := ParseIPFromRemoteAddr(""); result != "" {
		t.Fatal(result)
	}
}

func TestParseInvalidRemoteAddr(t *testing.T) {
	if result := ParseIPFromRemoteAddr("asd"); result != "" {
		t.Fatal(result)
	}
}

func TestParseValidAddr(t *testing.T) {
	if result := ParseIPFromRemoteAddr("12.23.34.54:1234"); result != "12.23.34.54" {
		t.Fatal(result)
	}
}
