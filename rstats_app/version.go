package rstats_app

import (
	"runtime/debug"
	"strings"
)

func normalizeVersion() string {
	version := "unknown-version"
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, s := range info.Settings {
			if s.Key == "-tags" {
				tags := strings.Split(s.Value, ",")
				for _, tag := range tags {
					if strings.HasPrefix(tag, "vcs.describe=") {
						version = tag[len("vcs.describe="):]
						break
					}
				}
				break
			}
		}
	}
	if !strings.HasPrefix(version, "v") {
		return version + ".pre-release"
	}
	return version
}

var normalizedVersion = normalizeVersion()

func GetVersion() string {
	return normalizedVersion
}
