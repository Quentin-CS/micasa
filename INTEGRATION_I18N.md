# Integration Guide: i18n Localization

This document guides the implementation of the i18n system throughout the micasa application.

## Quick Start

### 1. Initialize i18n in main.go

Add to `cmd/micasa/main.go`:

```go
import "github.com/micasa-dev/micasa/internal/i18n"

func init() {
    i18n.Init() // Initialize language detection at startup
}
```

### 2. Replace Hardcoded Strings

In `internal/app/model.go` and other UI files, replace hardcoded strings with i18n calls.

**Before:**
```go
m.setStatusInfo("Deleted shown.")
```

**After:**
```go
i18n.Get().DeletedShown()
// or for setStatusInfo:
m.setStatusInfo(i18n.Get().DeletedShown())
```

## Integration Areas

### Priority 1: Status Messages (Immediate)

These are high-visibility messages shown to users constantly. File: `internal/app/model.go`

Strings to replace:
- `"Deleted shown."` → `i18n.Get().DeletedShown()`
- `"Deleted hidden."` → `i18n.Get().DeletedHidden()`
- `"Settled shown."` → `i18n.Get().SettledShown()`
- `"Settled hidden."` → `i18n.Get().SettledHidden()`
- `"Pinned."` → `i18n.Get().Pinned()`
- `"Unpinned."` → `i18n.Get().Unpinned()`

Search command to find all:
```bash
grep -n "setStatusInfo\|setStatusError\|setStatusWarn" internal/app/model.go | head -20
```

### Priority 2: UI Labels (High)

File: `internal/app/forms.go`, `internal/app/tables.go`

Strings like:
- Button labels: "Save", "Cancel", "Delete", "Edit", "Add"
- Column headers
- Tab names: "Projects", "Maintenance", "Spending", "Dashboard"

### Priority 3: Help Text and Dialogs (Medium)

File: `internal/app/help.go` (if exists), form rendering

Strings like:
- Help text
- Confirmation dialogs
- Error messages

### Priority 4: Entity Names (Lower Priority)

File: `internal/data/entity_context.go`

Dynamic entity names that could be translated.

## Implementation Pattern

### Step 1: Add to translations.go

If a string isn't in `translations.go`, add it:

```go
// In translations struct:
type translations struct {
    // ... existing fields ...
    myNewString string
}

// In translationsEN:
var translationsEN = translations{
    // ... existing translations ...
    myNewString: "My English String",
}

// In translationsFR:
var translationsFR = translations{
    // ... existing translations ...
    myNewString: "Ma chaîne en français",
}
```

### Step 2: Add Accessor Method

In `i18n.go`:

```go
func (l *Language) MyNewString() string { return l.t.myNewString }
```

### Step 3: Use in Code

```go
import "github.com/micasa-dev/micasa/internal/i18n"

// Use it:
message := i18n.Get().MyNewString()
```

## Testing

### Unit Tests

The i18n package has comprehensive tests in `internal/i18n/i18n_test.go`.

Run tests:
```bash
go test -v ./internal/i18n/
```

### Integration Testing

To test the TUI with different languages:

```bash
# Test with French
LANG=fr_FR.UTF-8 ./micasa

# Test with English
LANG=en_US.UTF-8 ./micasa

# Test with unsupported language (should fall back to English)
LANG=de_DE.UTF-8 ./micasa
```

### Visual Verification

After each replacement:
1. Run the application with `LANG=en_US.UTF-8`
2. Run the application with `LANG=fr_FR.UTF-8`
3. Verify both languages display correctly
4. Check that UI layout hasn't broken (longer French strings, etc.)

## Implementation Checklist

- [ ] Initialize i18n in main.go
- [ ] Add necessary translation strings to `translations.go`
- [ ] Create accessor methods in `i18n.go`
- [ ] Replace status messages in `model.go`
- [ ] Replace UI labels in `forms.go`
- [ ] Replace table headers in `tables.go`
- [ ] Replace tab names in `tables.go`
- [ ] Update help text
- [ ] Test with both EN and FR locales
- [ ] Test with unsupported locales (should default to EN)
- [ ] Verify no UI layout issues from longer strings

## Common Pitfalls

### 1. Forgetting to Import i18n

```go
import "github.com/micasa-dev/micasa/internal/i18n"
```

### 2. String Formatting

For strings with parameters, don't hardcode the parameter in the translation:

**Bad:**
```go
// In translations
errorMsg: "Error loading %s",
// Usage
fmt.Sprintf(i18n.Get().ErrorMsg(), filename)
```

**Good:**
```go
// In translations
errorLoading: "Error loading",
// Usage (or create a specialized method)
fmt.Sprintf("%s %s", i18n.Get().ErrorLoading(), filename)
```

Or create specialized methods:
```go
func (l *Language) ErrorLoadingFile(filename string) string {
    return fmt.Sprintf("%s %s", l.t.errorLoading, filename)
}
```

### 3. Consistency Across Code

When replacing strings, ensure consistency:
- If "Delete" appears as a button label, use `i18n.Get().Delete()`
- If "delete" appears in sentence context, create a separate method or handle the case

### 4. Testing with Real Locales

Always test with actual locale environment variables:
```bash
LANG=fr_FR.UTF-8 go test ./...
```

## Measuring Progress

To see how many hardcoded UI strings remain:

```bash
# Count setStatusInfo calls that don't use i18n
grep "setStatusInfo" internal/app/*.go | grep -v "i18n" | wc -l

# List all string literals in key files
grep -n '"[A-Z][a-zA-Z ].*\."' internal/app/model.go | head -30
```

## Next Steps

1. Complete this integration phase
2. Add support for more languages (German, Spanish, Japanese)
3. Add date/time localization
4. Consider user language preference in config file
5. Add language selector in UI

## Questions?

Refer to:
- `internal/i18n/README.md` - Usage documentation
- `internal/i18n/i18n_test.go` - Examples of correct usage
- `golang.org/x/text/language` - Language tag documentation
