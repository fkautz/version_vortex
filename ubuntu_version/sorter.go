// SPDX-License-Identifier: Apache-2.0

package ubuntu_version

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// UbuntuVersion represents a parsed version string
type UbuntuVersion struct {
	Epoch    int
	Upstream string
	Debian   string
	Ubuntu   string
}

// ParseVersion parses a version string into its components
func ParseVersion(version string) UbuntuVersion {
	var epoch int
	upstream := version
	debian := ""
	ubuntu := ""

	epochParts := strings.SplitN(version, ":", 2)
	if len(epochParts) == 2 {
		// assume epoch is zero if it is excluded
		epoch, _ = strconv.Atoi(epochParts[0])
		upstream = epochParts[1]
	}

	parts := strings.SplitN(upstream, "-", 2)
	upstream = parts[0]

	if len(parts) > 1 {
		if strings.Contains(parts[1], "ubuntu") {
			debianParts := strings.Split(parts[1], "ubuntu")
			debian = debianParts[0]
			ubuntu = debianParts[1]
		} else {
			debian = parts[1]
		}
	}

	return UbuntuVersion{
		Epoch:    epoch,
		Upstream: upstream,
		Debian:   debian,
		Ubuntu:   ubuntu,
	}
}

func CompareVersionStrings(v1, v2 string) bool {
	return CompareVersions(ParseVersion(v1), ParseVersion(v2))
}

// CompareVersions compares two version strings based on the Ubuntu version format
func CompareVersions(v1, v2 UbuntuVersion) bool {
	if v1.Epoch != v2.Epoch {
		return v1.Epoch < v2.Epoch
	}
	if v1.Upstream != v2.Upstream {
		return compareSubversions(v1.Upstream, v2.Upstream)
	}
	if v1.Debian != v2.Debian {
		return compareSubversions(v1.Debian, v2.Debian)
	}
	return compareSubversions(v1.Ubuntu, v2.Ubuntu)
}

// compareSubversions compares subversion strings correctly, handling numeric and alphanumeric parts
func compareSubversions(v1, v2 string) bool {
	re := regexp.MustCompile(`(\d+|\D+)`)
	v1Parts := re.FindAllString(v1, -1)
	v2Parts := re.FindAllString(v2, -1)

	for i := 0; i < len(v1Parts) && i < len(v2Parts); i++ {
		v1Part := v1Parts[i]
		v2Part := v2Parts[i]

		v1Num, v1Err := strconv.Atoi(v1Part)
		v2Num, v2Err := strconv.Atoi(v2Part)

		if v1Err == nil && v2Err == nil {
			if v1Num != v2Num {
				return v1Num < v2Num
			}
		} else {
			if v1Part != v2Part {
				return v1Part < v2Part
			}
		}
	}

	return len(v1Parts) < len(v2Parts)
}

func Sort(versions []string) {
	sort.Slice(versions, func(i, j int) bool {
		vi := ParseVersion(versions[i])
		vj := ParseVersion(versions[j])
		return CompareVersions(vi, vj)
	})
}
