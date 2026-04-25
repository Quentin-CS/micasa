// Copyright 2026 Phillip Cloud
// Licensed under the Apache License, Version 2.0

package i18n

import "golang.org/x/text/language"

// translations holds all UI strings in a specific language.
type translations struct {
	// Status messages
	deletedShown      string
	deletedHidden     string
	settledShown      string
	settledHidden     string
	pinned            string
	unpinned          string
	filtered          string
	cleared           string
	copiedToClipboard string
	columnHidden      string
	columnShown       string

	// UI labels
	help   string
	quit   string
	save   string
	cancel string
	delete string
	edit   string
	add    string
	close  string
	search string
	filter string
	sort   string
	undo   string
	redo   string

	// Tab names
	projects    string
	maintenance string
	spending    string
	dashboard   string
	documents   string
	quotes      string
	incidents   string
	appliances  string
	vendors     string

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

	// Form field labels
	fldTitle           string
	fldProjectType     string
	fldStatus          string
	fldBudget          string
	fldActualCost      string
	fldStartDate       string
	fldEndDate         string
	fldDescription     string
	fldProject         string
	fldVendorName      string
	fldContactName     string
	fldEmail           string
	fldPhone           string
	fldWebsite         string
	fldEntity          string
	fldFileToAttach    string
	fldReplacementFile string
	fldTotal           string
	fldLabor           string
	fldMaterials       string
	fldOther           string
	fldReceivedDate    string
	fldNotes           string
	fldItem            string
	fldCategory        string
	fldSeason          string
	fldAppliance       string
	fldVendor          string
	fldSchedule        string
	fldInterval        string
	fldDueDate         string
	fldLastServiced    string
	fldManualURL       string
	fldManualNotes     string
	fldCost            string
	fldName            string
	fldBrand           string
	fldModelNumber     string
	fldSerialNumber    string
	fldLocation        string
	fldPurchaseDate    string
	fldWarrantyExpiry  string
	fldDateServiced    string
	fldPerformedBy     string
	fldDateNoticed     string
	fldDateResolved    string
	fldSeverity        string

	// House form fields
	fldNickname       string
	fldPostalCode     string
	fldAddressLine1   string
	fldAddressLine2   string
	fldCity           string
	fldState          string
	fldYearBuilt      string
	fldBedrooms       string
	fldBathrooms      string
	fldFoundationType string
	fldWiringType     string
	fldRoofType       string
	fldExteriorType   string
	fldBasement       string
	fldHeatingType    string
	fldCoolingType    string
	fldWaterSource    string
	fldSewerType      string
	fldParkingType    string
	fldInsCarrier     string
	fldInsPolicy      string
	fldInsRenewal     string
	fldPropertyTax    string
	fldHOAName        string
	fldHOAFee         string

	// Form sections
	secTimeline  string
	secVendor    string
	secQuote     string
	secDetails   string
	secSchedule  string
	secIdentity  string
	secContext   string
	secLinks     string
	secBasics    string
	secStructure string
	secUtilities string
	secFinancial string

	// Schedule types
	schedNone      string
	schedRecurring string
	schedFixedDue  string

	// Options and placeholders
	optNone          string
	optSelfHomeowner string
	phYYYYMM         string
	ph6m             string
	phKitchen        string
	phRefrigerator   string
	phAcmePlumbing   string
	ph1250           string
	ph1400           string
	ph3250           string
	ph2000           string
	ph1000           string
	ph250            string
	ph899            string
	ph125            string
	ph4200           string
	ph1998           string
	ph3              string
	ph25             string

	// Validation messages
	valRequired          string
	valEndDateAfterStart string

	// Project statuses
	statusIdeating  string
	statusPlanned   string
	statusQuoted    string
	statusUnderway  string
	statusDelayed   string
	statusCompleted string
	statusAbandoned string

	// Incident statuses
	statusOpen       string
	statusInProgress string
	statusResolved   string

	// Incident severities
	severityUrgent   string
	severitySoon     string
	severityWhenever string

	// Seasons
	seasonSpring string
	seasonSummer string
	seasonFall   string
	seasonWinter string

	// Help labels
	helpConfirm  string
	helpCancel   string
	helpNavigate string
	helpSection  string
	helpEdit     string
	helpClose    string
	helpNav      string
	helpToggle   string
	helpCollapse string
	helpTabs     string

	// Error messages
	errNoActiveTab           string
	errNothingSelected       string
	errCannotEditDeletedItem string
	errAddFromMaintenanceTab string
	errNoTargetTab           string

	// Form descriptions
	descOnlyNicknameRequired string

	// Status messages (continued)
	deleted                     string
	pressToRestore              string
	editItemFromMaintenanceTab  string
	hidingHiddenFiles           string
	showingHiddenFiles          string
	houseSetup                  string
	lastServicedDateSynced      string
	nothingToFollow             string
	permanentlyDeleted          string
	reopened                    string
	resolved                    string
	pressToReopen               string
	restored                    string
	saved                       string
	installTesseract            string
	layoutOn                    string
	layoutOff                   string
	extracted                   string
	checkingExtractionModel     string
	unitsStatus                 string
	modelPullError              string
	loadDocumentForExtraction   string
	extractionLLMError          string
	extractionIncomplete        string
	createTempFileError         string
	writeTempFileError          string
	editorError                 string
	readTempFileError           string
	listModelsError             string
	openError                   string
	syncBlobErrors              string
	syncError                   string
	postalCodeLookupError       string
	modelPullAlreadyInProgress  string
	llmSkippedSuffix            string
	nothingSelected             string
	linkedItemNotFound          string
	cannotHideLastColumn        string
	resolvedWithReopen          string
	deletedWithRestore          string
	resolveIncidentFirstThenDel string
	deleteItemFirstThenDelete   string
	houseProfileRequired        string
	extractionFailed            string
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
