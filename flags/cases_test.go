package flags

type flagsTest struct {
	input    []string
	expected Flags
}

var flagsTestCases = []flagsTest{
	{
		input:    []string{"-p", "3001", "-a", "127.0.0.1", "-path", "/var/html"},
		expected: Flags{Dir: "/var/html", Port: "3001", Addr: "127.0.0.1"},
	},
}
