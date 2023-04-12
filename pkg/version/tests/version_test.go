package tests

import (
	"github.com/blackmarllboro/create-project-struct/pkg/version"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestGetVersion(t *testing.T) {
	v, err := version.GoVersion()
	if err != nil {
		t.Fatalf("GoVersion() returned an error: %v", err)
	}

	var versionParts []string

	t.Log("verify that the version string is in the correct format")
	{
		re := regexp.MustCompile(`go \d+\.\d+(\.\d+)?`)
		if !re.MatchString(v) {
			t.Fatalf("GoVersion() returned an invalid version string: %s", v)
		}
	}

	t.Log("verify that the version string is properly formatted")
	{
		versionDigits := strings.TrimPrefix(v, "go")
		versionParts = strings.Split(versionDigits, ".")
		if len(versionParts) < 2 || len(versionParts) > 3 {
			t.Fatalf("GoVersion() returned an invalid version string: %s", v)
		}
	}

	t.Log("verify that the version string starts with \"go\"")
	{
		if !strings.HasPrefix(v, "go") {
			t.Fatalf("GoVersion() returned an invalid version string: %s", v)
		}
	}

	t.Log("verify that the patch version number, if present, is an integer")
	{
		if len(versionParts) == 3 {
			patch := versionParts[2]
			if _, err := strconv.Atoi(patch); err != nil {
				t.Fatalf("GoVersion() returned an invalid patch version number: %s", patch)
			}
		}
	}

	t.Log("SUCCESS")
}
