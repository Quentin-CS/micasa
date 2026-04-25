// Copyright 2026 Phillip Cloud
// Licensed under the Apache License, Version 2.0

package app

import (
	"charm.land/huh/v2"
	"github.com/micasa-dev/micasa/internal/data"
	"github.com/micasa-dev/micasa/internal/i18n"
	"github.com/micasa-dev/micasa/internal/locale"
)

type houseSection int

const (
	houseSectionIdentity houseSection = iota
	houseSectionStructure
	houseSectionUtilities
	houseSectionFinancial
)

// houseSectionOrder lists sections in the order they appear in the form.
var houseSectionOrder = []houseSection{
	houseSectionIdentity,
	houseSectionStructure,
	houseSectionUtilities,
	houseSectionFinancial,
}

func (s houseSection) title() string {
	lang := i18n.Get()
	switch s {
	case houseSectionIdentity:
		return lang.SecBasics()
	case houseSectionStructure:
		return lang.SecStructure()
	case houseSectionUtilities:
		return lang.SecUtilities()
	case houseSectionFinancial:
		return lang.SecFinancial()
	default:
		return ""
	}
}

// houseFieldDef describes a single editable field on HouseProfile.
// Used by both the full form and the overlay inline editor.
type houseFieldDef struct {
	key     string
	label   string
	labelFn func(us data.UnitSystem) string // unit-aware label override
	section houseSection
	// build creates a huh.Field bound to the given *string value.
	build func(m *Model, value *string) huh.Field
	// get reads the display value from a HouseProfile.
	get func(p data.HouseProfile, cur locale.Currency, us data.UnitSystem) string
	// ptr returns a pointer to this field's backing string in houseFormData.
	ptr func(fd *houseFormData) *string
	// validate checks a string value for this field. nil = no validation.
	validate func(string) error
	// toggle, if non-nil, means Enter toggles the value instead of opening
	// a textinput. The function flips the value and returns the new string.
	toggle func(current string) string
}

// displayLabel returns the label for the given unit system, using labelFn
// when available and falling back to the static label.
func (d houseFieldDef) displayLabel(us data.UnitSystem) string {
	if d.labelFn != nil {
		return d.labelFn(us)
	}
	return d.label
}

func houseFieldDefs() []houseFieldDef {
	return []houseFieldDef{
		// Identity — ordered to match form tab order (postal code after nickname
		// for autofill, then address lines, city, state).
		{
			key: "nickname", label: "Name", section: houseSectionIdentity,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().
					Title(requiredTitle(i18n.Get().FldNickname())).
					Description("Ex: Primary Residence").
					Value(v).
					Validate(requiredText("nickname"))
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.Nickname
			},
			ptr:      func(fd *houseFormData) *string { return &fd.Nickname },
			validate: requiredText("nickname"),
		},
		{
			key: "postal_code", label: "ZIP", section: houseSectionIdentity,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldPostalCode()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.PostalCode
			},
			ptr:      func(fd *houseFormData) *string { return &fd.PostalCode },
			validate: nil,
		},
		{
			key: "address_line1", label: "Addr 1", section: houseSectionIdentity,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldAddressLine1()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.AddressLine1
			},
			ptr:      func(fd *houseFormData) *string { return &fd.AddressLine1 },
			validate: nil,
		},
		{
			key: "address_line2", label: "Addr 2", section: houseSectionIdentity,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldAddressLine2()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.AddressLine2
			},
			ptr:      func(fd *houseFormData) *string { return &fd.AddressLine2 },
			validate: nil,
		},
		{
			key: "city", label: "City", section: houseSectionIdentity,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldCity()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.City
			},
			ptr:      func(fd *houseFormData) *string { return &fd.City },
			validate: nil,
		},
		{
			key: "state", label: "State", section: houseSectionIdentity,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldState()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.State
			},
			ptr:      func(fd *houseFormData) *string { return &fd.State },
			validate: nil,
		},
		// Structure
		{
			key: "year_built", label: "Year", section: houseSectionStructure,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().
					Title(i18n.Get().FldYearBuilt()).
					Placeholder(i18n.Get().Ph1998()).
					Value(v).
					Validate(optionalInt("year built"))
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return intToString(p.YearBuilt)
			},
			ptr:      func(fd *houseFormData) *string { return &fd.YearBuilt },
			validate: optionalInt("year built"),
		},
		{
			key: "square_feet", label: "Ft\u00B2", section: houseSectionStructure,
			labelFn: func(us data.UnitSystem) string {
				if us == data.UnitsMetric {
					return "m\u00B2"
				}
				return "Ft\u00B2"
			},
			build: func(m *Model, v *string) huh.Field {
				return huh.NewInput().
					Title(data.AreaFormTitle(m.unitSystem)).
					Placeholder(data.AreaPlaceholder(m.unitSystem)).
					Value(v).
					Validate(optionalInt(data.AreaFormTitle(m.unitSystem)))
			},
			get: func(p data.HouseProfile, _ locale.Currency, us data.UnitSystem) string {
				return intToString(data.SqFtToDisplayInt(p.SquareFeet, us))
			},
			ptr:      func(fd *houseFormData) *string { return &fd.SquareFeet },
			validate: optionalInt(data.AreaFormTitle(data.UnitsImperial)),
		},
		{
			key: "lot_square_feet", label: "Lot", section: houseSectionStructure,
			build: func(m *Model, v *string) huh.Field {
				return huh.NewInput().
					Title(data.LotAreaFormTitle(m.unitSystem)).
					Placeholder(data.LotAreaPlaceholder(m.unitSystem)).
					Value(v).
					Validate(optionalInt(data.LotAreaFormTitle(m.unitSystem)))
			},
			get: func(p data.HouseProfile, _ locale.Currency, us data.UnitSystem) string {
				return intToString(data.SqFtToDisplayInt(p.LotSquareFeet, us))
			},
			ptr:      func(fd *houseFormData) *string { return &fd.LotSquareFeet },
			validate: optionalInt(data.LotAreaFormTitle(data.UnitsImperial)),
		},
		{
			key: "bedrooms", label: "Bed", section: houseSectionStructure,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().
					Title(i18n.Get().FldBedrooms()).
					Placeholder(i18n.Get().Ph3()).
					Value(v).
					Validate(optionalInt("bedrooms"))
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return intToString(p.Bedrooms)
			},
			ptr:      func(fd *houseFormData) *string { return &fd.Bedrooms },
			validate: optionalInt("bedrooms"),
		},
		{
			key: "bathrooms", label: "Bath", section: houseSectionStructure,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().
					Title(i18n.Get().FldBathrooms()).
					Placeholder(i18n.Get().Ph25()).
					Value(v).
					Validate(optionalFloat("bathrooms"))
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return formatFloat(p.Bathrooms)
			},
			ptr:      func(fd *houseFormData) *string { return &fd.Bathrooms },
			validate: optionalFloat("bathrooms"),
		},
		{
			key: "foundation_type", label: "Fndtn", section: houseSectionStructure,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldFoundationType()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.FoundationType
			},
			ptr:      func(fd *houseFormData) *string { return &fd.FoundationType },
			validate: nil,
		},
		{
			key: "wiring_type", label: "Wire", section: houseSectionStructure,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldWiringType()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.WiringType
			},
			ptr:      func(fd *houseFormData) *string { return &fd.WiringType },
			validate: nil,
		},
		{
			key: "roof_type", label: "Roof", section: houseSectionStructure,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldRoofType()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.RoofType
			},
			ptr:      func(fd *houseFormData) *string { return &fd.RoofType },
			validate: nil,
		},
		{
			key: "exterior_type", label: "Ext", section: houseSectionStructure,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldExteriorType()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.ExteriorType
			},
			ptr:      func(fd *houseFormData) *string { return &fd.ExteriorType },
			validate: nil,
		},
		{
			key: "basement_type", label: "Bsmnt", section: houseSectionStructure,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldBasement()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				if p.BasementType != "" {
					return "Yes"
				}
				return "No"
			},
			ptr:      func(fd *houseFormData) *string { return &fd.BasementType },
			validate: nil,
			toggle: func(cur string) string {
				if cur != "" {
					return ""
				}
				return "Yes"
			},
		},
		// Utilities
		{
			key: "heating_type", label: "Heat", section: houseSectionUtilities,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldHeatingType()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.HeatingType
			},
			ptr:      func(fd *houseFormData) *string { return &fd.HeatingType },
			validate: nil,
		},
		{
			key: "cooling_type", label: "Cool", section: houseSectionUtilities,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldCoolingType()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.CoolingType
			},
			ptr:      func(fd *houseFormData) *string { return &fd.CoolingType },
			validate: nil,
		},
		{
			key: "water_source", label: "Water", section: houseSectionUtilities,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldWaterSource()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.WaterSource
			},
			ptr:      func(fd *houseFormData) *string { return &fd.WaterSource },
			validate: nil,
		},
		{
			key: "sewer_type", label: "Sewer", section: houseSectionUtilities,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldSewerType()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.SewerType
			},
			ptr:      func(fd *houseFormData) *string { return &fd.SewerType },
			validate: nil,
		},
		{
			key: "parking_type", label: "Parking", section: houseSectionUtilities,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldParkingType()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.ParkingType
			},
			ptr:      func(fd *houseFormData) *string { return &fd.ParkingType },
			validate: nil,
		},
		// Financial
		{
			key: "insurance_carrier", label: "Ins carrier", section: houseSectionFinancial,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldInsCarrier()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.InsuranceCarrier
			},
			ptr:      func(fd *houseFormData) *string { return &fd.InsuranceCarrier },
			validate: nil,
		},
		{
			key: "insurance_policy", label: "Ins policy", section: houseSectionFinancial,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldInsPolicy()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.InsurancePolicy
			},
			ptr:      func(fd *houseFormData) *string { return &fd.InsurancePolicy },
			validate: nil,
		},
		{
			key: "insurance_renewal", label: "Ins renewal", section: houseSectionFinancial,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().
					Title(i18n.Get().FldInsRenewal()).
					Value(v).
					Validate(optionalDate("insurance renewal"))
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return data.FormatDate(p.InsuranceRenewal)
			},
			ptr:      func(fd *houseFormData) *string { return &fd.InsuranceRenewal },
			validate: optionalDate("insurance renewal"),
		},
		{
			key: "property_tax", label: "Prop tax", section: houseSectionFinancial,
			build: func(m *Model, v *string) huh.Field {
				return huh.NewInput().
					Title(i18n.Get().FldPropertyTax()).
					Placeholder(i18n.Get().Ph4200()).
					Value(v).
					Validate(optionalMoney("property tax", m.cur))
			},
			get: func(p data.HouseProfile, cur locale.Currency, _ data.UnitSystem) string {
				return cur.FormatOptionalCents(p.PropertyTaxCents)
			},
			ptr:      func(fd *houseFormData) *string { return &fd.PropertyTax },
			validate: nil, // currency-dependent; validated by saveHouseFormData
		},
		{
			key: "hoa_name", label: "HOA", section: houseSectionFinancial,
			build: func(_ *Model, v *string) huh.Field {
				return huh.NewInput().Title(i18n.Get().FldHOAName()).Value(v)
			},
			get: func(p data.HouseProfile, _ locale.Currency, _ data.UnitSystem) string {
				return p.HOAName
			},
			ptr:      func(fd *houseFormData) *string { return &fd.HOAName },
			validate: nil,
		},
		{
			key: "hoa_fee", label: "HOA fee", section: houseSectionFinancial,
			build: func(m *Model, v *string) huh.Field {
				return huh.NewInput().
					Title(i18n.Get().FldHOAFee()).
					Placeholder(i18n.Get().Ph250()).
					Value(v).
					Validate(optionalMoney("HOA fee", m.cur))
			},
			get: func(p data.HouseProfile, cur locale.Currency, _ data.UnitSystem) string {
				return cur.FormatOptionalCents(p.HOAFeeCents)
			},
			ptr:      func(fd *houseFormData) *string { return &fd.HOAFee },
			validate: nil, // currency-dependent; validated by saveHouseFormData
		},
	}
}
