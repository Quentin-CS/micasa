// Copyright 2026 Phillip Cloud
// Licensed under the Apache License, Version 2.0

package i18n

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
)

func TestDetectSystemLanguage(t *testing.T) {
	tests := []struct {
		name     string
		locale   string
		expected language.Tag
	}{
		{
			name:     "French locale",
			locale:   "fr_FR.UTF-8",
			expected: language.French,
		},
		{
			name:     "English locale",
			locale:   "en_US.UTF-8",
			expected: language.English,
		},
		{
			name:     "French without encoding",
			locale:   "fr_FR",
			expected: language.French,
		},
		{
			name:     "German locale",
			locale:   "de_DE.UTF-8",
			expected: language.German,
		},
		{
			name:     "Portuguese locale",
			locale:   "pt_BR.UTF-8",
			expected: language.Portuguese,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag, err := parseLocale(tt.locale)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, tag)
		})
	}
}

func TestInit(t *testing.T) {
	oldLang := globalLang
	defer func() { globalLang = oldLang }()

	globalLang = nil
	err := Init()
	require.NoError(t, err)
	require.NotNil(t, globalLang)
}

func TestLanguageTranslations(t *testing.T) {
	// Test English
	lang := &Language{
		tag: language.English,
		t:   translationsEN,
	}

	assert.Equal(t, "Deleted shown.", lang.DeletedShown())
	assert.Equal(t, "Pinned.", lang.Pinned())
	assert.Equal(t, "Help", lang.Help())
	assert.Equal(t, "Projects", lang.Projects())

	// Test French
	lang.SetLanguage(language.French)
	assert.Equal(t, "Supprimés affichés.", lang.DeletedShown())
	assert.Equal(t, "Épinglé.", lang.Pinned())
	assert.Equal(t, "Aide", lang.Help())
	assert.Equal(t, "Projets", lang.Projects())
}

func TestLanguageTag(t *testing.T) {
	lang := &Language{
		tag: language.English,
		t:   translationsEN,
	}

	assert.Equal(t, language.English, lang.Tag())

	lang.SetLanguage(language.French)
	assert.Equal(t, language.French, lang.Tag())
}

func TestGetGlobalLanguage(t *testing.T) {
	oldLang := globalLang
	defer func() { globalLang = oldLang }()

	globalLang = nil
	lang := Get()
	require.NotNil(t, lang)
	// Should default to English if no system locale detected
	// (or whatever the system locale is)
}

func TestFormat(t *testing.T) {
	lang := &Language{
		tag: language.English,
		t:   translationsEN,
	}

	result := lang.Format("Hello %s", "World")
	assert.Equal(t, "Hello World", result)
}

func TestParseLocaleEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		locale  string
		wantErr bool
	}{
		{
			name:    "empty string",
			locale:  "",
			wantErr: true,
		},
		{
			name:    "invalid language code",
			locale:  "xyz",
			wantErr: true,
		},
		{
			name:    "locale with modifier",
			locale:  "en_US@latin",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseLocale(tt.locale)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetTranslationsForLanguage(t *testing.T) {
	tests := []struct {
		name     string
		tag      language.Tag
		expected string
	}{
		{
			name:     "English",
			tag:      language.English,
			expected: "Deleted shown.",
		},
		{
			name:     "American English",
			tag:      language.AmericanEnglish,
			expected: "Deleted shown.",
		},
		{
			name:     "French",
			tag:      language.French,
			expected: "Supprimés affichés.",
		},
		{
			name:     "Unknown language defaults to English",
			tag:      language.German,
			expected: "Deleted shown.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trans := getTranslations(tt.tag)
			assert.Equal(t, tt.expected, trans.deletedShown)
		})
	}
}

func TestSetLanguageUpdatesTranslations(t *testing.T) {
	lang := &Language{
		tag: language.English,
		t:   translationsEN,
	}

	// Start with English
	assert.Equal(t, "Deleted shown.", lang.DeletedShown())

	// Switch to French
	lang.SetLanguage(language.French)
	assert.Equal(t, "Supprimés affichés.", lang.DeletedShown())
	assert.Equal(t, language.French, lang.Tag())

	// Switch back to English
	lang.SetLanguage(language.English)
	assert.Equal(t, "Deleted shown.", lang.DeletedShown())
	assert.Equal(t, language.English, lang.Tag())
}

func TestDetectSystemLanguageFromEnv(t *testing.T) {
	// Save original env vars
	oldLang := os.Getenv("LANG")
	oldLC := os.Getenv("LC_ALL")
	oldMessages := os.Getenv("LC_MESSAGES")

	defer func() {
		os.Setenv("LANG", oldLang)
		os.Setenv("LC_ALL", oldLC)
		os.Setenv("LC_MESSAGES", oldMessages)
	}()

	// Clear all locale env vars
	os.Unsetenv("LANG")
	os.Unsetenv("LC_ALL")
	os.Unsetenv("LC_MESSAGES")

	// Test with LANG set to French
	os.Setenv("LANG", "fr_FR.UTF-8")
	tag, err := detectSystemLanguage()
	require.NoError(t, err)
	assert.Equal(t, language.French, tag)

	// Test with LC_ALL taking precedence
	os.Setenv("LC_ALL", "en_US.UTF-8")
	tag, err = detectSystemLanguage()
	require.NoError(t, err)
	assert.Equal(t, language.English, tag)
}
