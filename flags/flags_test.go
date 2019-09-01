package flags

import "testing"

func TestGetFlags(t *testing.T) {
	var expected = Flags{Dir: "./", Port: "8080", Addr: "0.0.0.0"}
	if observed := GetFlags(); observed != expected {
		t.Fatalf("GetFlags() = %v, want %v", observed, expected)
	}
}
