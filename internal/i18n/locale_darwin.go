// Copyright 2026 Phillip Cloud
// Licensed under the Apache License, Version 2.0

//go:build darwin

package i18n

import (
	"os/exec"
	"strings"
)

// platformLocale returns the macOS user locale from System Preferences via
// `defaults read -g AppleLocale`, which reflects the language selected in
// System Settings regardless of the POSIX LANG variable.
func platformLocale() string {
	out, err := exec.Command("defaults", "read", "-g", "AppleLocale").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
