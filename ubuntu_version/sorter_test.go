// SPDX-License-Identifier: Apache-2.0

package ubuntu_version

import (
	"fmt"
	"testing"
)

func TestParseVersion1(t *testing.T) {
	version := "14-20240412-0ubuntu1"
	parsedVersion := ParseVersion(version)
	fmt.Println(parsedVersion)
}

func TestParseVersion2(t *testing.T) {
	version := "13ubuntu10"
	parsedVersion := ParseVersion(version)
	fmt.Println(parsedVersion)
}

func TestCompareVersion1(t *testing.T) {
	small := "openssl-3.0.2-0ubuntu1.15"
	large := "openssl-3.0.13-0ubuntu3.1"

	res := CompareVersions(ParseVersion(small), ParseVersion(large))
	if res != true {
		t.Errorf("expected true, got false")
	}
}
