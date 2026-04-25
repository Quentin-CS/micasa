// Copyright 2026 Phillip Cloud
// Licensed under the Apache License, Version 2.0

// Package i18n provides internationalization support for the micasa application.
package i18n

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/language"
)

// Global singleton for the current language instance.
var globalLang *Language

// Language holds the current language and translation strings.
type Language struct {
	tag language.Tag
	t   translations
}

// Init initializes the global language instance based on system locale.
func Init() error {
	lang, err := detectSystemLanguage()
	if err != nil {
		lang = language.English
	}

	globalLang = &Language{
		tag: lang,
		t:   getTranslations(lang),
	}
	return nil
}

// Get returns the global language instance.
func Get() *Language {
	if globalLang == nil {
		Init()
	}
	return globalLang
}

// Tag returns the language tag.
func (l *Language) Tag() language.Tag {
	return l.tag
}

// Language returns the human-readable language name.
func (l *Language) LanguageName() string {
	return l.tag.String()
}

// SetLanguage sets the language manually (useful for testing or user preferences).
func (l *Language) SetLanguage(tag language.Tag) {
	l.tag = tag
	l.t = getTranslations(tag)
}

// detectSystemLanguage detects the system language from environment variables.
func detectSystemLanguage() (language.Tag, error) {
	// Check LANG, LC_ALL, LC_MESSAGES in order
	for _, envVar := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if locale := os.Getenv(envVar); locale != "" {
			return parseLocale(locale)
		}
	}

	// Fallback to English
	return language.English, nil
}

// parseLocale parses a locale string (e.g., "fr_FR.UTF-8") into a language tag.
func parseLocale(locale string) (language.Tag, error) {
	// Extract the language part before any underscore or dot
	// e.g., "fr_FR.UTF-8" -> "fr", "en_US" -> "en"
	parts := strings.FieldsFunc(locale, func(r rune) bool {
		return r == '_' || r == '-' || r == '.' || r == '@'
	})

	if len(parts) == 0 {
		return language.English, fmt.Errorf("empty locale")
	}

	tag, err := language.Parse(parts[0])
	if err != nil {
		return language.English, fmt.Errorf("failed to parse locale: %w", err)
	}

	return tag, nil
}

// ============================================================================
// String accessors (these will reference the translations map)
// ============================================================================

// Status messages
func (l *Language) DeletedShown() string             { return l.t.deletedShown }
func (l *Language) DeletedHidden() string            { return l.t.deletedHidden }
func (l *Language) SettledShown() string             { return l.t.settledShown }
func (l *Language) SettledHidden() string            { return l.t.settledHidden }
func (l *Language) Pinned() string                   { return l.t.pinned }
func (l *Language) Unpinned() string                 { return l.t.unpinned }
func (l *Language) Filtered() string                 { return l.t.filtered }
func (l *Language) Cleared() string                  { return l.t.cleared }
func (l *Language) CopiedToClipboard() string        { return l.t.copiedToClipboard }
func (l *Language) ColumnHidden() string             { return l.t.columnHidden }
func (l *Language) ColumnShown() string              { return l.t.columnShown }

// UI labels
func (l *Language) Help() string                     { return l.t.help }
func (l *Language) Quit() string                     { return l.t.quit }
func (l *Language) Save() string                     { return l.t.save }
func (l *Language) Cancel() string                   { return l.t.cancel }
func (l *Language) Delete() string                   { return l.t.delete }
func (l *Language) Edit() string                     { return l.t.edit }
func (l *Language) Add() string                      { return l.t.add }
func (l *Language) Close() string                    { return l.t.close }
func (l *Language) Search() string                   { return l.t.search }
func (l *Language) Filter() string                   { return l.t.filter }
func (l *Language) Sort() string                     { return l.t.sort }
func (l *Language) Undo() string                     { return l.t.undo }
func (l *Language) Redo() string                     { return l.t.redo }

// Tab names
func (l *Language) Projects() string                 { return l.t.projects }
func (l *Language) Maintenance() string              { return l.t.maintenance }
func (l *Language) Spending() string                 { return l.t.spending }
func (l *Language) Dashboard() string                { return l.t.dashboard }
func (l *Language) Documents() string                { return l.t.documents }
func (l *Language) Quotes() string                   { return l.t.quotes }
func (l *Language) Incidents() string                { return l.t.incidents }
func (l *Language) Appliances() string               { return l.t.appliances }
func (l *Language) Vendors() string                  { return l.t.vendors }

// Dialogs and prompts
func (l *Language) ConfirmDelete() string            { return l.t.confirmDelete }
func (l *Language) NoResults() string                { return l.t.noResults }
func (l *Language) Loading() string                  { return l.t.loading }
func (l *Language) Error() string                    { return l.t.errStr }

// Help text
func (l *Language) HelpText() string                 { return l.t.helpText }

// Column titles
func (l *Language) ColType() string                  { return l.t.colType }
func (l *Language) ColTitle() string                 { return l.t.colTitle }
func (l *Language) ColStatus() string                { return l.t.colStatus }
func (l *Language) ColBudget() string                { return l.t.colBudget }
func (l *Language) ColActual() string                { return l.t.colActual }
func (l *Language) ColStart() string                 { return l.t.colStart }
func (l *Language) ColEnd() string                   { return l.t.colEnd }
func (l *Language) ColTotal() string                 { return l.t.colTotal }
func (l *Language) ColLabor() string                 { return l.t.colLabor }
func (l *Language) ColMat() string                   { return l.t.colMat }
func (l *Language) ColOther() string                 { return l.t.colOther }
func (l *Language) ColRecv() string                  { return l.t.colRecv }
func (l *Language) ColItem() string                  { return l.t.colItem }
func (l *Language) ColCategory() string              { return l.t.colCategory }
func (l *Language) ColName() string                  { return l.t.colName }
func (l *Language) ColNotes() string                 { return l.t.colNotes }
func (l *Language) ColUpdated() string               { return l.t.colUpdated }
func (l *Language) ColProject() string               { return l.t.colProject }
func (l *Language) ColVendor() string                { return l.t.colVendor }
func (l *Language) ColAppliance() string             { return l.t.colAppliance }
func (l *Language) ColBrand() string                 { return l.t.colBrand }
func (l *Language) ColModel() string                 { return l.t.colModel }
func (l *Language) ColSerial() string                { return l.t.colSerial }
func (l *Language) ColLocation() string              { return l.t.colLocation }
func (l *Language) ColPurchased() string             { return l.t.colPurchased }
func (l *Language) ColWarranty() string              { return l.t.colWarranty }
func (l *Language) ColAge() string                   { return l.t.colAge }
func (l *Language) ColMaint() string                 { return l.t.colMaint }
func (l *Language) ColNext() string                  { return l.t.colNext }
func (l *Language) ColLast() string                  { return l.t.colLast }
func (l *Language) ColEvery() string                 { return l.t.colEvery }
func (l *Language) ColSeason() string                { return l.t.colSeason }
func (l *Language) ColSeverity() string              { return l.t.colSeverity }
func (l *Language) ColResolved() string              { return l.t.colResolved }
func (l *Language) ColPerformed() string             { return l.t.colPerformed }
func (l *Language) ColNoticed() string               { return l.t.colNoticed }
func (l *Language) ColCost() string                  { return l.t.colCost }
func (l *Language) ColDate() string                  { return l.t.colDate }
func (l *Language) ColContact() string               { return l.t.colContact }
func (l *Language) ColEmail() string                 { return l.t.colEmail }
func (l *Language) ColPhone() string                 { return l.t.colPhone }
func (l *Language) ColWebsite() string               { return l.t.colWebsite }
func (l *Language) ColJobs() string                  { return l.t.colJobs }
func (l *Language) ColOps() string                   { return l.t.colOps }
func (l *Language) ColEntity() string                { return l.t.colEntity }
func (l *Language) ColLog() string                   { return l.t.colLog }
func (l *Language) ColSize() string                  { return l.t.colSize }

// Additional messages
func (l *Language) AllColumnsVisible() string        { return l.t.allColumnsVisible }
func (l *Language) NothingToCopy() string            { return l.t.nothingToCopy }
func (l *Language) PressIToEdit() string             { return l.t.pressIToEdit }
func (l *Language) PressOToOpen() string             { return l.t.pressOToOpen }

// Status messages (continued)
func (l *Language) Deleted() string                  { return l.t.deleted }
func (l *Language) PressToRestore() string           { return l.t.pressToRestore }
func (l *Language) EditItemFromMaintenanceTab() string { return l.t.editItemFromMaintenanceTab }
func (l *Language) HidingHiddenFiles() string        { return l.t.hidingHiddenFiles }
func (l *Language) ShowingHiddenFiles() string       { return l.t.showingHiddenFiles }
func (l *Language) HouseSetup() string               { return l.t.houseSetup }
func (l *Language) LastServicedDateSynced() string   { return l.t.lastServicedDateSynced }
func (l *Language) NothingToFollow() string          { return l.t.nothingToFollow }
func (l *Language) PermanentlyDeleted() string       { return l.t.permanentlyDeleted }
func (l *Language) Reopened() string                 { return l.t.reopened }
func (l *Language) Resolved() string                 { return l.t.resolved }
func (l *Language) PressToReopen() string            { return l.t.pressToReopen }
func (l *Language) Restored() string                 { return l.t.restored }
func (l *Language) Saved() string                    { return l.t.saved }
func (l *Language) InstallTesseract() string         { return l.t.installTesseract }
func (l *Language) LayoutOn() string                 { return l.t.layoutOn }
func (l *Language) LayoutOff() string                { return l.t.layoutOff }
func (l *Language) Extracted() string                { return l.t.extracted }
func (l *Language) CheckingExtractionModel() string  { return l.t.checkingExtractionModel }
func (l *Language) NothingSelected() string          { return l.t.nothingSelected }
func (l *Language) LinkedItemNotFound() string       { return l.t.linkedItemNotFound }
func (l *Language) CannotHideLastColumn() string     { return l.t.cannotHideLastColumn }
func (l *Language) ResolvedWithReopen() string       { return l.t.resolvedWithReopen }
func (l *Language) DeletedWithRestore() string       { return l.t.deletedWithRestore }
func (l *Language) ResolveIncidentFirstThenDel() string { return l.t.resolveIncidentFirstThenDel }
func (l *Language) DeleteItemFirstThenDelete() string { return l.t.deleteItemFirstThenDelete }
func (l *Language) HouseProfileRequired() string     { return l.t.houseProfileRequired }
func (l *Language) ExtractionFailed() string        { return l.t.extractionFailed }

// TranslateColumnTitle translates English column titles to the current language.
func (l *Language) TranslateColumnTitle(title string) string {
	switch title {
	case "Type":
		return l.ColType()
	case "Title":
		return l.ColTitle()
	case "Status":
		return l.ColStatus()
	case "Budget":
		return l.ColBudget()
	case "Actual":
		return l.ColActual()
	case "Start":
		return l.ColStart()
	case "End":
		return l.ColEnd()
	case "Total":
		return l.ColTotal()
	case "Labor":
		return l.ColLabor()
	case "Mat":
		return l.ColMat()
	case "Other":
		return l.ColOther()
	case "Recv":
		return l.ColRecv()
	case "Item":
		return l.ColItem()
	case "Category":
		return l.ColCategory()
	case "Name":
		return l.ColName()
	case "Notes":
		return l.ColNotes()
	case "Updated":
		return l.ColUpdated()
	case "Project":
		return l.ColProject()
	case "Vendor":
		return l.ColVendor()
	case "Appliance":
		return l.ColAppliance()
	case "Brand":
		return l.ColBrand()
	case "Model":
		return l.ColModel()
	case "Serial":
		return l.ColSerial()
	case "Location":
		return l.ColLocation()
	case "Purchased":
		return l.ColPurchased()
	case "Warranty":
		return l.ColWarranty()
	case "Age":
		return l.ColAge()
	case "Maint":
		return l.ColMaint()
	case "Next":
		return l.ColNext()
	case "Last":
		return l.ColLast()
	case "Every":
		return l.ColEvery()
	case "Season":
		return l.ColSeason()
	case "Severity":
		return l.ColSeverity()
	case "Resolved":
		return l.ColResolved()
	case "Performed By":
		return l.ColPerformed()
	case "Noticed":
		return l.ColNoticed()
	case "Cost":
		return l.ColCost()
	case "Date":
		return l.ColDate()
	case "Contact":
		return l.ColContact()
	case "Email":
		return l.ColEmail()
	case "Phone":
		return l.ColPhone()
	case "Website":
		return l.ColWebsite()
	case "Jobs":
		return l.ColJobs()
	case "Ops":
		return l.ColOps()
	case "Entity":
		return l.ColEntity()
	case "Log":
		return l.ColLog()
	case "Size":
		return l.ColSize()
	default:
		return title // Return original if not found
	}
}

// Utility function for formatting strings with parameters
func (l *Language) Format(key string, args ...interface{}) string {
	return fmt.Sprintf(key, args...)
}
