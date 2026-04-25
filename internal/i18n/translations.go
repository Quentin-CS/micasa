// Copyright 2026 Phillip Cloud
// Licensed under the Apache License, Version 2.0

package i18n

import "golang.org/x/text/language"

// translations holds all UI strings in a specific language.
type translations struct {
	// Status messages
	deletedShown       string
	deletedHidden      string
	settledShown       string
	settledHidden      string
	pinned             string
	unpinned           string
	filtered           string
	cleared            string
	copiedToClipboard  string
	columnHidden       string
	columnShown        string

	// UI labels
	help       string
	quit       string
	save       string
	cancel     string
	delete     string
	edit       string
	add        string
	close      string
	search     string
	filter     string
	sort       string
	undo       string
	redo       string

	// Tab names
	projects     string
	maintenance  string
	spending     string
	dashboard    string
	documents    string
	quotes       string
	incidents    string
	appliances   string
	vendors      string

	// Dialogs and prompts
	confirmDelete string
	noResults     string
	loading       string
	errStr        string

	// Help text
	helpText string

	// Column titles
	colType      string
	colTitle     string
	colStatus    string
	colBudget    string
	colActual    string
	colStart     string
	colEnd       string
	colTotal     string
	colLabor     string
	colMat       string
	colOther     string
	colRecv      string
	colItem      string
	colCategory  string
	colName      string
	colNotes     string
	colUpdated   string
	colProject   string
	colVendor    string
	colAppliance string
	colBrand     string
	colModel     string
	colSerial    string
	colLocation  string
	colPurchased string
	colWarranty  string
	colAge       string
	colMaint     string
	colNext      string
	colLast      string
	colEvery     string
	colSeason    string
	colSeverity  string
	colResolved  string
	colPerformed string
	colNoticed   string
	colCost      string
	colDate      string
	colContact   string
	colEmail     string
	colPhone     string
	colWebsite   string
	colJobs      string
	colOps       string
	colEntity    string
	colLog       string
	colSize      string

	// Additional messages
	allColumnsVisible string
	nothingToCopy     string
	pressIToEdit      string
	pressOToOpen      string

	// Status messages (continued)
	deleted                       string
	pressToRestore                string
	editItemFromMaintenanceTab    string
	hidingHiddenFiles             string
	showingHiddenFiles            string
	houseSetup                    string
	lastServicedDateSynced        string
	nothingToFollow               string
	permanentlyDeleted            string
	reopened                      string
	resolved                      string
	pressToReopen                 string
	restored                      string
	saved                         string
	installTesseract              string
	layoutOn                      string
	layoutOff                     string
	extracted                     string
	checkingExtractionModel       string
	nothingSelected               string
	linkedItemNotFound            string
	cannotHideLastColumn          string
	resolvedWithReopen            string
	deletedWithRestore            string
	resolveIncidentFirstThenDel   string
	deleteItemFirstThenDelete     string
	houseProfileRequired          string
	extractionFailed              string
}

// getTranslations returns the appropriate translation set for the given language.
func getTranslations(tag language.Tag) translations {
	// Match against base language, ignoring region
	// Check if tag starts with French code
	tagStr := tag.String()
	if len(tagStr) >= 2 {
		prefix := tagStr[:2]
		if prefix == "fr" {
			return translationsFR
		}
	}
	// Default to English for everything else
	return translationsEN
}
