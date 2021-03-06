// +build integration

package test

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/magefile/mage/mage"
)

func TestDockerTagsMageInvocation(t *testing.T) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	invocation := mage.Invocation{Stderr: stderr, Stdout: stdout}
	exitCode := mage.Invoke(invocation)

	if exitCode != 0 {
		t.Fatalf("got exit code %v, err: %s", exitCode, stderr)
	}

	checkDockerTags(t, stdout.Bytes())
}

func checkDockerTags(t *testing.T, actual []byte) {
	expectedFormat := `1\.0\.1,1\.0\.1-circleci-(\d+),gauge-1\.0\.1,chrome-81\.0\.4044\.92,go-1\.0\.2,node-1\.0\.3,taiko-1\.0\.4`
	regex := regexp.MustCompile(expectedFormat)

	if !regex.Match(actual) {
		t.Fatalf("%q is not in the expected format: %q", actual, expectedFormat)
	}
}
