/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : countries.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Country selection and document management system
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package countries

import (
	"abdal-better-passlist/internal/models"
	"abdal-better-passlist/internal/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// SelectCountry allows user to select a country interactively
func SelectCountry() (*models.Country, error) {
	countries := models.GetSupportedCountries()

	color.Cyan("\n=== Country Selection ===")
	color.Yellow("Please select your country from the list below:\n")

	// Display countries with Iran first
	color.Green("1. 🇮🇷 Iran")

	for i, country := range countries {
		if country.Code == "ir" {
			continue // Skip Iran as it's already displayed first
		}

		flag := getCountryFlag(country.Code)
		fmt.Printf("%d. %s %s\n", i+1, flag, country.Name)
	}

	color.Cyan("\nEnter the number of your country (1-%d): ", len(countries))

	var choice int
	var err error

	for {
		input := utils.GetUserInput("Country number: ")

		// Check if input is empty
		if strings.TrimSpace(input) == "" {
			color.Red("❌ Please enter a number. You cannot leave this field empty.")
			continue
		}

		// Try to convert to integer
		choice, err = strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			color.Red("❌ Invalid input! Please enter a valid number (not letters or symbols).")
			color.Yellow("💡 Tip: Enter a number between 1 and %d", len(countries))
			continue
		}

		// Check if choice is within valid range
		if choice < 1 || choice > len(countries) {
			color.Red("❌ Invalid choice! Please select a number between 1 and %d", len(countries))
			color.Yellow("💡 You entered: %d, but valid range is 1-%d", choice, len(countries))
			continue
		}

		// If we reach here, the input is valid
		break
	}

	// Adjust for Iran being displayed first
	var selectedCountry *models.Country
	if choice == 1 {
		// Iran selected
		for _, country := range countries {
			if country.Code == "ir" {
				selectedCountry = &country
				break
			}
		}
	} else {
		// Other countries
		index := choice - 1
		if countries[index].Code == "ir" {
			index++ // Skip Iran in the list
		}
		if index < len(countries) {
			selectedCountry = &countries[index]
		}
	}

	if selectedCountry == nil {
		return nil, fmt.Errorf("country not found")
	}

	color.Green("\n✅ Selected country: %s", selectedCountry.Name)
	color.Cyan("📋 Required documents for %s:", selectedCountry.Name)

	for _, doc := range selectedCountry.RequiredDocuments {
		status := "Optional"
		if doc.Required {
			status = "Required"
			color.Red("  • %s (%s)", doc.Name, status)
		} else {
			color.Yellow("  • %s (%s)", doc.Name, status)
		}
	}

	return selectedCountry, nil
}

// GetCountryByCode returns a country by its code
func GetCountryByCode(code string) (*models.Country, error) {
	countries := models.GetSupportedCountries()

	code = strings.ToLower(strings.TrimSpace(code))

	for _, country := range countries {
		if country.Code == code {
			return &country, nil
		}
	}

	return nil, fmt.Errorf("country with code '%s' not found", code)
}

// getCountryFlag returns the flag emoji for a country code
func getCountryFlag(code string) string {
	flags := map[string]string{
		"ir":      "🇮🇷",
		"us":      "🇺🇸",
		"uk":      "🇬🇧",
		"de":      "🇩🇪",
		"fr":      "🇫🇷",
		"se":      "🇸🇪",
		"in":      "🇮🇳",
		"jp":      "🇯🇵",
		"ae":      "🇦🇪",
		"ca":      "🇨🇦",
		"au":      "🇦🇺",
		"general": "🌍",
	}

	if flag, exists := flags[code]; exists {
		return flag
	}

	return "🌍" // Default flag for unknown countries
}

// ValidateCountryCode checks if a country code is valid
func ValidateCountryCode(code string) bool {
	_, err := GetCountryByCode(code)
	return err == nil
}

// GetCountryList returns a formatted list of all supported countries
func GetCountryList() string {
	countries := models.GetSupportedCountries()
	var result strings.Builder

	result.WriteString("Supported Countries:\n")
	result.WriteString("==================\n\n")

	for _, country := range countries {
		flag := getCountryFlag(country.Code)
		result.WriteString(fmt.Sprintf("%s %s (%s)\n", flag, country.Name, country.Code))
	}

	return result.String()
}

// GetRequiredDocumentsForCountry returns the required documents for a specific country
func GetRequiredDocumentsForCountry(countryCode string) ([]models.DocumentType, error) {
	country, err := GetCountryByCode(countryCode)
	if err != nil {
		return nil, err
	}

	return country.RequiredDocuments, nil
}

// IsDocumentRequired checks if a specific document is required for a country
func IsDocumentRequired(countryCode, documentType string) (bool, error) {
	documents, err := GetRequiredDocumentsForCountry(countryCode)
	if err != nil {
		return false, err
	}

	for _, doc := range documents {
		if doc.Type == documentType {
			return doc.Required, nil
		}
	}

	return false, fmt.Errorf("document type '%s' not found for country '%s'", documentType, countryCode)
}
