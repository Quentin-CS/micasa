// Copyright 2026 Phillip Cloud
// Licensed under the Apache License, Version 2.0

//go:build !darwin

package i18n

// platformLocale returns an empty string on non-macOS platforms; locale
// detection relies solely on POSIX environment variables (LANG, LC_ALL, etc.).
func platformLocale() string {
	return ""
}
