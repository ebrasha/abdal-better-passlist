/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : sensitive.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Sensitive complexity password generator with enhanced combinations
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package generators

import (
	"abdal-better-passlist/internal/models"
	"abdal-better-passlist/internal/utils"
	"strings"
)

// SensitiveGenerator handles sensitive complexity password generation
type SensitiveGenerator struct {
	commonNumbers []string
	specialChars  []string
}

// NewSensitiveGenerator creates a new sensitive complexity generator
func NewSensitiveGenerator() *SensitiveGenerator {
	return &SensitiveGenerator{
		commonNumbers: utils.GenerateCommonNumbers(),
		specialChars:  utils.GenerateSpecialCharacters(),
	}
}

// GeneratePasswords generates passwords with sensitive complexity
func (sg *SensitiveGenerator) GeneratePasswords(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	// Enhanced name combinations
	sg.generateEnhancedNameCombinations(personalInfo, passwordChan)

	// Advanced number combinations
	sg.generateAdvancedNumberCombinations(personalInfo, passwordChan)

	// Enhanced special character combinations
	sg.generateEnhancedSpecialCombinations(personalInfo, passwordChan)

	// Comprehensive family combinations
	sg.generateComprehensiveFamilyCombinations(personalInfo, passwordChan)

	// Advanced date combinations
	sg.generateAdvancedDateCombinations(personalInfo, passwordChan)

	// Document-based combinations
	sg.generateDocumentCombinations(personalInfo, passwordChan)

	// Mixed case combinations
	sg.generateMixedCaseCombinations(personalInfo, passwordChan)
}

// generateEnhancedNameCombinations creates enhanced name-based passwords
func (sg *SensitiveGenerator) generateEnhancedNameCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	_ = strings.ToLower(personalInfo.FirstName)
	_ = strings.ToLower(personalInfo.LastName)

	// Generate name variations
	firstNameVariations := utils.GenerateVariations(personalInfo.FirstName)
	lastNameVariations := utils.GenerateVariations(personalInfo.LastName)

	// Enhanced combinations
	for _, firstVar := range firstNameVariations {
		for _, lastVar := range lastNameVariations {
			combinations := []string{
				firstVar + lastVar,
				lastVar + firstVar,
				firstVar + "_" + lastVar,
				lastVar + "_" + firstVar,
				firstVar + "-" + lastVar,
				lastVar + "-" + firstVar,
				firstVar + "." + lastVar,
				lastVar + "." + firstVar,
			}

			for _, combo := range combinations {
				if combo != "" {
					passwordChan <- combo
				}
			}
		}
	}

	// Add numbers to name variations
	numbers := []string{"123", "1234", "12345", "111", "222", "333", "000", "0000", "01", "02", "03"}
	for _, firstVar := range firstNameVariations {
		for _, num := range numbers {
			combinations := []string{
				firstVar + num,
				num + firstVar,
				firstVar + "_" + num,
				num + "_" + firstVar,
			}

			for _, combo := range combinations {
				if combo != "" {
					passwordChan <- combo
				}
			}
		}
	}
}

// generateAdvancedNumberCombinations creates advanced number-based passwords
func (sg *SensitiveGenerator) generateAdvancedNumberCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	// Extract and process all numbers
	mobileNumbers := utils.ExtractNumbers(personalInfo.MobileNumber)
	homeNumbers := utils.ExtractNumbers(personalInfo.HomePhoneNumber)

	allNumbers := append(mobileNumbers, homeNumbers...)

	// Process each number
	for _, num := range allNumbers {
		if num == "" {
			continue
		}

		// Generate number variations
		numberParts := utils.SplitNumberIntoParts(num)

		for _, part := range numberParts {
			if part != "" {
				passwordChan <- part
			}
		}

		// Combine with names
		firstName := strings.ToLower(personalInfo.FirstName)
		lastName := strings.ToLower(personalInfo.LastName)

		for _, part := range numberParts {
			combinations := []string{
				firstName + part,
				lastName + part,
				part + firstName,
				part + lastName,
				firstName + "_" + part,
				part + "_" + firstName,
				lastName + "_" + part,
				part + "_" + lastName,
			}

			for _, combo := range combinations {
				if combo != "" {
					passwordChan <- combo
				}
			}
		}
	}
}

// generateEnhancedSpecialCombinations creates enhanced special character combinations
func (sg *SensitiveGenerator) generateEnhancedSpecialCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	firstName := strings.ToLower(personalInfo.FirstName)
	lastName := strings.ToLower(personalInfo.LastName)

	// Use more special characters
	specials := []string{"!", "@", "#", "$", "%", "&", "*", "123", "456", "789", "000", "111"}

	for _, special := range specials {
		combinations := []string{
			firstName + special,
			lastName + special,
			special + firstName,
			special + lastName,
			firstName + special + lastName,
			lastName + special + firstName,
			firstName + "_" + special,
			special + "_" + firstName,
			lastName + "_" + special,
			special + "_" + lastName,
		}

		for _, combo := range combinations {
			if combo != "" {
				passwordChan <- combo
			}
		}
	}
}

// generateComprehensiveFamilyCombinations creates comprehensive family name combinations
func (sg *SensitiveGenerator) generateComprehensiveFamilyCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	// Father name combinations
	if personalInfo.FatherName != "" {
		fatherVariations := utils.GenerateVariations(personalInfo.FatherName)
		for _, variation := range fatherVariations {
			combinations := []string{
				variation,
				variation + "123",
				variation + "1234",
				variation + "12345",
				"123" + variation,
				"1234" + variation,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}
		}
	}

	// Mother name combinations
	if personalInfo.MotherName != "" {
		motherVariations := utils.GenerateVariations(personalInfo.MotherName)
		for _, variation := range motherVariations {
			combinations := []string{
				variation,
				variation + "123",
				variation + "1234",
				variation + "12345",
				"123" + variation,
				"1234" + variation,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}
		}
	}

	// Spouse combinations
	spouses := []string{personalInfo.SpouseName, personalInfo.SecondSpouseName, personalInfo.ThirdSpouseName}
	for _, spouse := range spouses {
		if spouse != "" {
			spouseVariations := utils.GenerateVariations(spouse)
			for _, variation := range spouseVariations {
				combinations := []string{
					variation,
					variation + "123",
					variation + "1234",
					"123" + variation,
					"1234" + variation,
				}

				for _, combo := range combinations {
					passwordChan <- combo
				}
			}
		}
	}

	// Girlfriend/Boyfriend combinations
	if personalInfo.GirlfriendName != "" {
		gfVariations := utils.GenerateVariations(personalInfo.GirlfriendName)
		for _, variation := range gfVariations {
			combinations := []string{
				variation,
				variation + "123",
				variation + "1234",
				"123" + variation,
				"1234" + variation,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}
		}
	}

	// Pet name combinations
	if personalInfo.PetName != "" {
		petVariations := utils.GenerateVariations(personalInfo.PetName)
		for _, variation := range petVariations {
			combinations := []string{
				variation,
				variation + "123",
				variation + "1234",
				"123" + variation,
				"1234" + variation,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}
		}
	}

	// Favorite person combinations
	if personalInfo.FavoritePersonName != "" {
		favVariations := utils.GenerateVariations(personalInfo.FavoritePersonName)
		for _, variation := range favVariations {
			combinations := []string{
				variation,
				variation + "123",
				variation + "1234",
				"123" + variation,
				"1234" + variation,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}
		}
	}
}

// generateAdvancedDateCombinations creates advanced date-based passwords
func (sg *SensitiveGenerator) generateAdvancedDateCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	if personalInfo.BirthDate == "" {
		return
	}

	year := utils.ExtractYearFromDate(personalInfo.BirthDate)
	month := utils.ExtractMonthFromDate(personalInfo.BirthDate)
	day := utils.ExtractDayFromDate(personalInfo.BirthDate)

	// Generate all date variations
	dates := []string{year, month, day}
	if len(year) == 4 {
		dates = append(dates, year[2:]) // Last two digits
	}

	// Add zero-padded versions
	if len(month) == 1 {
		dates = append(dates, "0"+month)
	}
	if len(day) == 1 {
		dates = append(dates, "0"+day)
	}

	// Combine with names
	firstName := strings.ToLower(personalInfo.FirstName)
	lastName := strings.ToLower(personalInfo.LastName)

	for _, date := range dates {
		if date != "" {
			combinations := []string{
				firstName + date,
				lastName + date,
				date + firstName,
				date + lastName,
				firstName + "_" + date,
				date + "_" + firstName,
				lastName + "_" + date,
				date + "_" + lastName,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}
		}
	}

	// Advanced date format combinations
	if year != "" && month != "" && day != "" {
		formats := []string{
			day + month + year,
			month + day + year,
			year + month + day,
			day + month + year[2:],
			month + day + year[2:],
			year[2:] + month + day,
			day + "/" + month + "/" + year,
			month + "/" + day + "/" + year,
			year + "/" + month + "/" + day,
		}

		for _, format := range formats {
			passwordChan <- format
		}
	}
}

// generateDocumentCombinations creates document-based passwords
func (sg *SensitiveGenerator) generateDocumentCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	for docType, docValue := range personalInfo.Documents {
		if docValue == "" {
			continue
		}

		// Extract numbers from documents
		docNumbers := utils.ExtractNumbers(docValue)

		for _, num := range docNumbers {
			if num != "" {
				// Basic document number
				passwordChan <- num

				// Document number with names
				firstName := strings.ToLower(personalInfo.FirstName)
				lastName := strings.ToLower(personalInfo.LastName)

				combinations := []string{
					firstName + num,
					lastName + num,
					num + firstName,
					num + lastName,
					docType + num,
					num + docType,
				}

				for _, combo := range combinations {
					passwordChan <- combo
				}
			}
		}
	}
}

// generateMixedCaseCombinations creates mixed case combinations
func (sg *SensitiveGenerator) generateMixedCaseCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	firstName := personalInfo.FirstName
	lastName := personalInfo.LastName

	// Mixed case variations
	combinations := []string{
		strings.ToUpper(firstName) + strings.ToLower(lastName),
		strings.ToLower(firstName) + strings.ToUpper(lastName),
		strings.ToUpper(firstName) + strings.ToUpper(lastName),
		strings.ToLower(firstName) + strings.ToLower(lastName),
		strings.Title(firstName) + strings.ToLower(lastName),
		strings.ToLower(firstName) + strings.Title(lastName),
	}

	for _, combo := range combinations {
		if combo != "" {
			passwordChan <- combo
		}
	}

	// Add numbers to mixed case
	numbers := []string{"123", "1234", "12345", "111", "222", "333"}
	for _, combo := range combinations {
		for _, num := range numbers {
			passwordChan <- combo + num
			passwordChan <- num + combo
		}
	}
}
