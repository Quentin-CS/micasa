# i18n - Internationalization

This package provides internationalization (i18n) support for the micasa application, automatically detecting the system language and providing translated UI strings.

## Features

- **Automatic system language detection** from environment variables (`LANG`, `LC_ALL`, `LC_MESSAGES`)
- **Graceful fallback** to English if system language is not supported
- **Easy string access** via methods like `i18n.Get().Help()`, `i18n.Get().Projects()`
- **Manual language switching** via `language.SetLanguage(tag)`
- **Support for English and French** (easily extensible to other languages)

## Usage

### Initialization

At application startup (in `main.go` or early initialization):

```go
import "github.com/micasa-dev/micasa/internal/i18n"

func init() {
    i18n.Init()
}
```

### Accessing Translations

Use the global singleton to access translated strings:

```go
lang := i18n.Get()

// Status messages
msg := lang.DeletedShown()  // "Deleted shown." (EN) or "Supprimés affichés." (FR)
msg := lang.Pinned()        // "Pinned." (EN) or "Épinglé." (FR)

// UI labels
msg := lang.Help()          // "Help" (EN) or "Aide" (FR)
msg := lang.Save()          // "Save" (EN) or "Enregistrer" (FR)

// Tab names
msg := lang.Projects()      // "Projects" (EN) or "Projets" (FR)

// Dialogs
msg := lang.ConfirmDelete() // "Are you sure?" (EN) or "Êtes-vous sûr ?" (FR)
```

### Adding New Translations

1. Add the new string to the `translations` struct in `translations.go`
2. Add the English version to `translationsEN`
3. Add the French version to `translationsFR`
4. Create an accessor method in `i18n.go`, e.g.:
   ```go
   func (l *Language) MyNewString() string { return l.t.myNewString }
   ```

### System Language Detection

The system language is automatically detected from:

1. `LC_ALL` environment variable (highest priority)
2. `LC_MESSAGES` environment variable
3. `LANG` environment variable (lowest priority)

Example locale strings that are recognized:
- `en_US.UTF-8` → English
- `fr_FR.UTF-8` → French
- `de_DE.UTF-8` → German (defaults to English, easy to add)
- `pt_BR.UTF-8` → Portuguese (defaults to English, easy to add)

### Manual Language Switching

For testing or user preferences:

```go
lang := i18n.Get()
lang.SetLanguage(language.French)  // Switch to French
lang.SetLanguage(language.English) // Switch back to English
```

## Supported Languages

Currently supported:
- **English** (en) - Default
- **French** (fr)

To add support for another language:
1. Add a new translation variable in `translations.go` (e.g., `translationsDE`)
2. Update the `getTranslations()` function to handle the new language
3. Add tests for the new language

## Environment Variables

The i18n package respects standard POSIX locale environment variables:

```bash
# Force French
export LANG=fr_FR.UTF-8
./micasa

# Force English
export LANG=en_US.UTF-8
./micasa
```

## Testing

Run the i18n tests:

```bash
go test -v ./internal/i18n/
```

Tests cover:
- Locale detection and parsing
- Translation lookups
- Language switching
- Environment variable handling
- Edge cases

## Architecture

The package uses `golang.org/x/text/language` for robust language tag handling and matching. This is a standard Go library for internationalization.

### Structure

```
i18n/
├── i18n.go          # Main Language type and initialization
├── translations.go  # Translation strings (EN + FR)
├── i18n_test.go     # Tests
└── README.md        # This file
```

## Future Enhancements

Possible improvements:
- Add more languages (German, Spanish, Japanese, etc.)
- Support for user language preferences in config file
- Plural form handling
- Date/time localization beyond current locale.go
- Number formatting by locale
