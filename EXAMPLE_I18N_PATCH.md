# Example: Integrating i18n into main.go

This shows the minimal changes needed to initialize i18n at application startup.

## Changes to cmd/micasa/main.go

### Step 1: Add Import

At the top of `main.go`, in the imports section (after line 24), add:

```go
"github.com/micasa-dev/micasa/internal/i18n"
```

The imports should look like:

```go
import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime/debug"

	tea "charm.land/bubbletea/v2"
	"charm.land/fang/v2"
	"charm.land/lipgloss/v2"
	"github.com/micasa-dev/micasa/internal/app"
	"github.com/micasa-dev/micasa/internal/config"
	"github.com/micasa-dev/micasa/internal/data"
	"github.com/micasa-dev/micasa/internal/extract"
	"github.com/micasa-dev/micasa/internal/i18n"  // ADD THIS LINE
	"github.com/spf13/cobra"
)
```

### Step 2: Initialize i18n in main()

In the `main()` function (starting at line 106), add the i18n initialization at the very beginning:

```go
func main() {
	// Initialize i18n (language detection) first
	if err := i18n.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "warning: failed to initialize i18n: %v\n", err)
		// Non-fatal error - continues with default (English)
	}

	root := newRootCmd()
	if err := fang.Execute(
		// ... rest of the function unchanged
	); err != nil {
		// ... rest of the function unchanged
	}
}
```

## Example: Using i18n in model.go

Here's how to replace a hardcoded string with i18n translation:

### Before

In `internal/app/model.go` around line 551:

```go
m.setStatusInfo("Deleted shown.")
```

### After

```go
import "github.com/micasa-dev/micasa/internal/i18n"

// In the handler or function:
m.setStatusInfo(i18n.Get().DeletedShown())
```

## Complete Example with Multiple Strings

Here's a more complete example of refactoring status message handlers:

```go
// BEFORE
case keyD:
	if m.showDeleted {
		m.showDeleted = false
		m.setStatusInfo("Deleted hidden.")
	} else {
		m.showDeleted = true
		m.setStatusInfo("Deleted shown.")
	}

// AFTER
case keyD:
	lang := i18n.Get()
	if m.showDeleted {
		m.showDeleted = false
		m.setStatusInfo(lang.DeletedHidden())
	} else {
		m.showDeleted = true
		m.setStatusInfo(lang.DeletedShown())
	}
```

## Testing the Integration

### Test 1: English (Default)

```bash
# Build
go build -o ./micasa ./cmd/micasa

# Run with English
./micasa

# Or explicitly
LANG=en_US.UTF-8 ./micasa
```

### Test 2: French

```bash
LANG=fr_FR.UTF-8 ./micasa
```

The status messages should now appear in French. For example, instead of "Deleted shown." you should see "Supprimés affichés."

### Test 3: Unsupported Locale (German)

```bash
LANG=de_DE.UTF-8 ./micasa
```

The application should fall back to English since German is not yet supported in the translations.

## Performance Note

The `i18n.Init()` function is called once at startup and:
- Parses the system locale (milliseconds)
- Selects the appropriate translation set (microseconds)
- Has negligible performance impact

After initialization, `i18n.Get()` returns a pre-initialized singleton, so accessing strings is O(1).

## Error Handling

If i18n initialization fails (e.g., malformed environment variable):
- A warning is printed to stderr
- The application continues with English as the default language
- This is intentional: missing translations should never crash the app

## Next Steps

After applying this patch:

1. **Test the integration** with both EN and FR locales
2. **Replace hardcoded strings** progressively in model.go (highest priority)
3. **Add more translations** for other UI strings (forms, help text, etc.)
4. **Update CLAUDE.md** with the new i18n pattern once validated
5. **Consider adding more languages** (German, Spanish, Portuguese, etc.)

## Rollback

If needed, this change is completely non-breaking:
- Remove the import statement
- Remove the `i18n.Init()` call from main()
- The application continues working exactly as before

All code using i18n would need to be reverted to hardcoded strings, but the core application is unaffected.

## Common Issues

### Issue: i18n.Get() returns nil

**Solution:** Make sure `i18n.Init()` was called before the first access to `i18n.Get()`. 

Put this check in the test:
```go
func TestI18nInitialized(t *testing.T) {
	i18n.Init()
	lang := i18n.Get()
	require.NotNil(t, lang)
	assert.NotEmpty(t, lang.DeletedShown())
}
```

### Issue: Translations are empty

**Solution:** Verify that `internal/i18n/translations.go` contains the translation strings.

Run:
```bash
grep "deletedShown:" internal/i18n/translations.go
```

### Issue: Wrong language detected

**Solution:** Check the environment variable:

```bash
echo $LANG
echo $LC_ALL
echo $LC_MESSAGES
```

Or test manually:
```bash
LANG=fr_FR.UTF-8 go run ./cmd/micasa
```

## Questions?

See `internal/i18n/README.md` for comprehensive usage documentation.
