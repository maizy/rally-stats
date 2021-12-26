package rstats_app

import (
	"runtime/debug"
	"strings"
)

func normalizeVersion() string {
	var version = "unknown-version"
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		if buildInfoVersion := buildInfo.Main.Version; buildInfoVersion != "" {
			version = buildInfoVersion
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
