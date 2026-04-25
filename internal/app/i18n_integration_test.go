// Copyright 2026 Phillip Cloud
// Licensed under the Apache License, Version 2.0

package app

import (
	"testing"

	"github.com/micasa-dev/micasa/internal/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
)

// TestI18nIntegration verifies that i18n is properly integrated into the app package.
func TestI18nIntegration(t *testing.T) {
	// Initialize i18n
	err := i18n.Init()
	require.NoError(t, err)

	// Get the language instance
	lang := i18n.Get()
	require.NotNil(t, lang)

	// Verify that the translation methods are accessible
	assert.NotEmpty(t, lang.DeletedShown())
	assert.NotEmpty(t, lang.DeletedHidden())
	assert.NotEmpty(t, lang.Pinned())
	assert.NotEmpty(t, lang.Unpinned())
	assert.NotEmpty(t, lang.Cleared())
	assert.NotEmpty(t, lang.SettledShown())
	assert.NotEmpty(t, lang.SettledHidden())

	// Verify tab names
	assert.NotEmpty(t, lang.Projects())
	assert.NotEmpty(t, lang.Maintenance())
	assert.NotEmpty(t, lang.Documents())

	// Verify additional messages
	assert.NotEmpty(t, lang.AllColumnsVisible())
	assert.NotEmpty(t, lang.NothingToCopy())
	assert.NotEmpty(t, lang.PressIToEdit())
	assert.NotEmpty(t, lang.PressOToOpen())
	assert.NotEmpty(t, lang.NothingSelected())
	assert.NotEmpty(t, lang.LinkedItemNotFound())
	assert.NotEmpty(t, lang.CannotHideLastColumn())
	assert.NotEmpty(t, lang.ResolvedWithReopen())
	assert.NotEmpty(t, lang.DeletedWithRestore())
	assert.NotEmpty(t, lang.ResolveIncidentFirstThenDel())
	assert.NotEmpty(t, lang.DeleteItemFirstThenDelete())
	assert.NotEmpty(t, lang.HouseProfileRequired())
}

// TestI18nLanguageSwitch verifies that language switching works correctly.
func TestI18nLanguageSwitch(t *testing.T) {
	// Initialize i18n
	err := i18n.Init()
	require.NoError(t, err)

	lang := i18n.Get()

	// Test English
	lang.SetLanguage(language.English)
	assert.Equal(t, "Deleted shown.", lang.DeletedShown())
	assert.Equal(t, "Pinned.", lang.Pinned())
	assert.Equal(t, "Projects", lang.Projects())
	assert.Equal(t, "All columns visible.", lang.AllColumnsVisible())
	assert.Equal(t, "Nothing to copy.", lang.NothingToCopy())

	// Test French
	lang.SetLanguage(language.French)
	assert.Equal(t, "Supprimés affichés.", lang.DeletedShown())
	assert.Equal(t, "Épinglé.", lang.Pinned())
	assert.Equal(t, "Projets", lang.Projects())
	assert.Equal(t, "Toutes les colonnes visibles.", lang.AllColumnsVisible())
	assert.Equal(t, "Rien à copier.", lang.NothingToCopy())
	assert.Equal(t, "Rien de sélectionné.", lang.NothingSelected())
	assert.Equal(t, "Résolu. Appuyez sur d pour rouvrir.", lang.ResolvedWithReopen())
	assert.Equal(t, "Supprimé. Appuyez sur d pour restaurer.", lang.DeletedWithRestore())

	// Switch back to English
	lang.SetLanguage(language.English)
	assert.Equal(t, "Deleted shown.", lang.DeletedShown())
}
