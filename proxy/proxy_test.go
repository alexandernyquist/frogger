package frogger

import(
	"fmt"
	"regexp"
	"testing"
)

func TestProxy_ShouldDump(t *testing.T) {
	// Match beginning
	testMatch("googleusercontent.com", "*.googleusercontent.com")
	testMatch("ps1.googleusercontent.com", "*.googleusercontent.com")
	testMatch("yoflo.test.googleusercontent.com", "*.googleusercontent.com")

	// Match end
	testMatch("www.googleusercontent.fi", "www.googleusercontent.*")
	testMatch("www.googleusercontent.fi/foo", "www.googleusercontent.*")
	testMatch("www.googleusercontent.co.uk", "www.googleusercontent.*")
}

func testMatch(host, dumps string) {
	pattern := getRegexPattern(dumps)
	match, _ := regexp.MatchString(pattern, host)
	fmt.Printf("%s -> %s: %v\n", dumps, host, match)
}

func getRegexPattern(pattern string) string {
	// TODO: Strip dot after star?
	return "^" + pattern + "$"
}