package main

// This value can be set for releases at build time using:
//   go {build|run} -ldflags "-X main.version 1.2.3.4".
// If unset, Version() shall return "DEVBUILD".
var version string

func Version() string {
	if version == "" {
		return "DEVBUILD"
	}
	return version
}
