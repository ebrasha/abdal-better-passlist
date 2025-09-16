/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : aggressive.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Aggressive complexity password generator with maximum combinations
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

// AggressiveGenerator handles aggressive complexity password generation
type AggressiveGenerator struct {
	commonNumbers []string
	specialChars  []string
}

// NewAggressiveGenerator creates a new aggressive complexity generator
func NewAggressiveGenerator() *AggressiveGenerator {
	return &AggressiveGenerator{
		commonNumbers: utils.GenerateCommonNumbers(),
		specialChars:  utils.GenerateSpecialCharacters(),
	}
}

// GeneratePasswords generates passwords with aggressive complexity
func (ag *AggressiveGenerator) GeneratePasswords(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	// Maximum name combinations
	ag.generateMaximumNameCombinations(personalInfo, passwordChan)

	// Extensive number combinations
	ag.generateExtensiveNumberCombinations(personalInfo, passwordChan)

	// Maximum special character combinations
	ag.generateMaximumSpecialCombinations(personalInfo, passwordChan)

	// Comprehensive family combinations
	ag.generateComprehensiveFamilyCombinations(personalInfo, passwordChan)

	// Advanced date combinations
	ag.generateAdvancedDateCombinations(personalInfo, passwordChan)

	// Document-based combinations
	ag.generateDocumentCombinations(personalInfo, passwordChan)

	// Mixed case combinations
	ag.generateMixedCaseCombinations(personalInfo, passwordChan)

	// Leet speak combinations
	ag.generateLeetSpeakCombinations(personalInfo, passwordChan)

	// Keyboard pattern combinations
	ag.generateKeyboardPatternCombinations(personalInfo, passwordChan)

	// Advanced document processing
	ag.generateAdvancedDocumentCombinations(personalInfo, passwordChan)
}

// generateMaximumNameCombinations creates maximum name-based passwords
func (ag *AggressiveGenerator) generateMaximumNameCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	_ = strings.ToLower(personalInfo.FirstName)
	_ = strings.ToLower(personalInfo.LastName)

	// Generate all name variations
	firstNameVariations := utils.GenerateVariations(personalInfo.FirstName)
	lastNameVariations := utils.GenerateVariations(personalInfo.LastName)

	// Maximum combinations with all variations
	for _, firstVar := range firstNameVariations {
		for _, lastVar := range lastNameVariations {
			// Basic combinations
			combinations := []string{
				firstVar + lastVar,
				lastVar + firstVar,
				firstVar + "_" + lastVar,
				lastVar + "_" + firstVar,
				firstVar + "-" + lastVar,
				lastVar + "-" + firstVar,
				firstVar + "." + lastVar,
				lastVar + "." + firstVar,
				firstVar + " " + lastVar,
				lastVar + " " + firstVar,
			}

			for _, combo := range combinations {
				if combo != "" {
					passwordChan <- combo
				}
			}
		}
	}

	// Add extensive numbers to name variations
	extensiveNumbers := []string{
		"123", "1234", "12345", "123456", "1234567", "12345678",
		"111", "222", "333", "444", "555", "666", "777", "888", "999",
		"000", "0000", "00000", "000000", "0000000", "00000000",
		"01", "02", "03", "04", "05", "06", "07", "08", "09", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30",
		"31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
		"50", "60", "70", "80", "90", "100", "200", "300", "500", "1000",
	}

	for _, firstVar := range firstNameVariations {
		for _, num := range extensiveNumbers {
			combinations := []string{
				firstVar + num,
				num + firstVar,
				firstVar + "_" + num,
				num + "_" + firstVar,
				firstVar + "-" + num,
				num + "-" + firstVar,
				firstVar + "." + num,
				num + "." + firstVar,
			}

			for _, combo := range combinations {
				if combo != "" {
					passwordChan <- combo
				}
			}
		}
	}
}

// generateExtensiveNumberCombinations creates extensive number-based passwords
func (ag *AggressiveGenerator) generateExtensiveNumberCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	// Extract and process all numbers
	mobileNumbers := utils.ExtractNumbers(personalInfo.MobileNumber)
	homeNumbers := utils.ExtractNumbers(personalInfo.HomePhoneNumber)

	allNumbers := append(mobileNumbers, homeNumbers...)

	// Process each number extensively
	for _, num := range allNumbers {
		if num == "" {
			continue
		}

		// Generate extensive number variations
		numberParts := utils.SplitNumberIntoParts(num)

		for _, part := range numberParts {
			if part != "" {
				passwordChan <- part

				// Add special characters to numbers
				for _, special := range ag.specialChars {
					combinations := []string{
						part + special,
						special + part,
						part + "_" + special,
						special + "_" + part,
					}

					for _, combo := range combinations {
						passwordChan <- combo
					}
				}
			}
		}

		// Combine with names extensively
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
				firstName + "-" + part,
				part + "-" + firstName,
				lastName + "-" + part,
				part + "-" + lastName,
				firstName + "." + part,
				part + "." + firstName,
				lastName + "." + part,
				part + "." + lastName,
			}

			for _, combo := range combinations {
				if combo != "" {
					passwordChan <- combo
				}
			}
		}
	}
}

// generateMaximumSpecialCombinations creates maximum special character combinations
func (ag *AggressiveGenerator) generateMaximumSpecialCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	firstName := strings.ToLower(personalInfo.FirstName)
	lastName := strings.ToLower(personalInfo.LastName)

	// Use all special characters
	for _, special := range ag.specialChars {
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
			firstName + "-" + special,
			special + "-" + firstName,
			lastName + "-" + special,
			special + "-" + lastName,
			firstName + "." + special,
			special + "." + firstName,
			lastName + "." + special,
			special + "." + lastName,
		}

		for _, combo := range combinations {
			if combo != "" {
				passwordChan <- combo
			}
		}
	}

	// Multiple special characters
	multipleSpecials := []string{"!@#", "$%^", "&*(", ")_+", "123!", "456@", "789#"}
	for _, multiSpecial := range multipleSpecials {
		combinations := []string{
			firstName + multiSpecial,
			lastName + multiSpecial,
			multiSpecial + firstName,
			multiSpecial + lastName,
			firstName + multiSpecial + lastName,
			lastName + multiSpecial + firstName,
		}

		for _, combo := range combinations {
			if combo != "" {
				passwordChan <- combo
			}
		}
	}
}

// generateComprehensiveFamilyCombinations creates comprehensive family name combinations
func (ag *AggressiveGenerator) generateComprehensiveFamilyCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	// Father name combinations
	if personalInfo.FatherName != "" {
		fatherVariations := utils.GenerateVariations(personalInfo.FatherName)
		for _, variation := range fatherVariations {
			// Basic combinations
			combinations := []string{
				variation,
				variation + "123",
				variation + "1234",
				variation + "12345",
				"123" + variation,
				"1234" + variation,
				"12345" + variation,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}

			// Add special characters
			for _, special := range ag.specialChars {
				combinations := []string{
					variation + special,
					special + variation,
					variation + "_" + special,
					special + "_" + variation,
				}

				for _, combo := range combinations {
					passwordChan <- combo
				}
			}
		}
	}

	// Mother name combinations
	if personalInfo.MotherName != "" {
		motherVariations := utils.GenerateVariations(personalInfo.MotherName)
		for _, variation := range motherVariations {
			// Basic combinations
			combinations := []string{
				variation,
				variation + "123",
				variation + "1234",
				variation + "12345",
				"123" + variation,
				"1234" + variation,
				"12345" + variation,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}

			// Add special characters
			for _, special := range ag.specialChars {
				combinations := []string{
					variation + special,
					special + variation,
					variation + "_" + special,
					special + "_" + variation,
				}

				for _, combo := range combinations {
					passwordChan <- combo
				}
			}
		}
	}

	// Spouse combinations
	spouses := []string{personalInfo.SpouseName, personalInfo.SecondSpouseName, personalInfo.ThirdSpouseName}
	for _, spouse := range spouses {
		if spouse != "" {
			spouseVariations := utils.GenerateVariations(spouse)
			for _, variation := range spouseVariations {
				// Basic combinations
				combinations := []string{
					variation,
					variation + "123",
					variation + "1234",
					variation + "12345",
					"123" + variation,
					"1234" + variation,
					"12345" + variation,
				}

				for _, combo := range combinations {
					passwordChan <- combo
				}

				// Add special characters
				for _, special := range ag.specialChars {
					combinations := []string{
						variation + special,
						special + variation,
						variation + "_" + special,
						special + "_" + variation,
					}

					for _, combo := range combinations {
						passwordChan <- combo
					}
				}
			}
		}
	}

	// Girlfriend/Boyfriend combinations
	if personalInfo.GirlfriendName != "" {
		gfVariations := utils.GenerateVariations(personalInfo.GirlfriendName)
		for _, variation := range gfVariations {
			// Basic combinations
			combinations := []string{
				variation,
				variation + "123",
				variation + "1234",
				variation + "12345",
				"123" + variation,
				"1234" + variation,
				"12345" + variation,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}

			// Add special characters
			for _, special := range ag.specialChars {
				combinations := []string{
					variation + special,
					special + variation,
					variation + "_" + special,
					special + "_" + variation,
				}

				for _, combo := range combinations {
					passwordChan <- combo
				}
			}
		}
	}

	// Pet name combinations
	if personalInfo.PetName != "" {
		petVariations := utils.GenerateVariations(personalInfo.PetName)
		for _, variation := range petVariations {
			// Basic combinations
			combinations := []string{
				variation,
				variation + "123",
				variation + "1234",
				variation + "12345",
				"123" + variation,
				"1234" + variation,
				"12345" + variation,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}

			// Add special characters
			for _, special := range ag.specialChars {
				combinations := []string{
					variation + special,
					special + variation,
					variation + "_" + special,
					special + "_" + variation,
				}

				for _, combo := range combinations {
					passwordChan <- combo
				}
			}
		}
	}

	// Favorite person combinations
	if personalInfo.FavoritePersonName != "" {
		favVariations := utils.GenerateVariations(personalInfo.FavoritePersonName)
		for _, variation := range favVariations {
			// Basic combinations
			combinations := []string{
				variation,
				variation + "123",
				variation + "1234",
				variation + "12345",
				"123" + variation,
				"1234" + variation,
				"12345" + variation,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}

			// Add special characters
			for _, special := range ag.specialChars {
				combinations := []string{
					variation + special,
					special + variation,
					variation + "_" + special,
					special + "_" + variation,
				}

				for _, combo := range combinations {
					passwordChan <- combo
				}
			}
		}
	}
}

// generateAdvancedDateCombinations creates advanced date-based passwords
func (ag *AggressiveGenerator) generateAdvancedDateCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
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

	// Combine with names extensively
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
				firstName + "-" + date,
				date + "-" + firstName,
				lastName + "-" + date,
				date + "-" + lastName,
				firstName + "." + date,
				date + "." + firstName,
				lastName + "." + date,
				date + "." + lastName,
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}

			// Add special characters to date combinations
			for _, special := range ag.specialChars {
				combinations := []string{
					firstName + date + special,
					lastName + date + special,
					special + firstName + date,
					special + lastName + date,
					firstName + special + date,
					lastName + special + date,
				}

				for _, combo := range combinations {
					passwordChan <- combo
				}
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
			day + "-" + month + "-" + year,
			month + "-" + day + "-" + year,
			year + "-" + month + "-" + day,
			day + "." + month + "." + year,
			month + "." + day + "." + year,
			year + "." + month + "." + day,
		}

		for _, format := range formats {
			passwordChan <- format
		}
	}
}

// generateDocumentCombinations creates document-based passwords
func (ag *AggressiveGenerator) generateDocumentCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
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
					firstName + "_" + num,
					num + "_" + firstName,
					lastName + "_" + num,
					num + "_" + lastName,
					docType + "_" + num,
					num + "_" + docType,
				}

				for _, combo := range combinations {
					passwordChan <- combo
				}

				// Add special characters to document combinations
				for _, special := range ag.specialChars {
					combinations := []string{
						firstName + num + special,
						lastName + num + special,
						special + firstName + num,
						special + lastName + num,
						docType + num + special,
						special + docType + num,
					}

					for _, combo := range combinations {
						passwordChan <- combo
					}
				}
			}
		}
	}
}

// generateMixedCaseCombinations creates mixed case combinations
func (ag *AggressiveGenerator) generateMixedCaseCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
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

	// Add extensive numbers to mixed case
	extensiveNumbers := []string{
		"123", "1234", "12345", "123456", "1234567", "12345678",
		"111", "222", "333", "444", "555", "666", "777", "888", "999",
		"000", "0000", "00000", "000000", "0000000", "00000000",
		"01", "02", "03", "04", "05", "06", "07", "08", "09", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30",
		"31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
		"50", "60", "70", "80", "90", "100", "200", "300", "500", "1000",
	}

	for _, combo := range combinations {
		for _, num := range extensiveNumbers {
			passwordChan <- combo + num
			passwordChan <- num + combo
			passwordChan <- combo + "_" + num
			passwordChan <- num + "_" + combo
			passwordChan <- combo + "-" + num
			passwordChan <- num + "-" + combo
			passwordChan <- combo + "." + num
			passwordChan <- num + "." + combo
		}
	}
}

// generateLeetSpeakCombinations creates leet speak combinations
func (ag *AggressiveGenerator) generateLeetSpeakCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	firstName := strings.ToLower(personalInfo.FirstName)
	lastName := strings.ToLower(personalInfo.LastName)

	// Leet speak mappings
	leetMap := map[string]string{
		"a": "4", "e": "3", "i": "1", "o": "0", "s": "5", "t": "7",
		"A": "4", "E": "3", "I": "1", "O": "0", "S": "5", "T": "7",
	}

	// Convert names to leet speak
	firstNameLeet := firstName
	lastNameLeet := lastName

	for original, replacement := range leetMap {
		firstNameLeet = strings.ReplaceAll(firstNameLeet, original, replacement)
		lastNameLeet = strings.ReplaceAll(lastNameLeet, original, replacement)
	}

	// Leet speak combinations
	combinations := []string{
		firstNameLeet,
		lastNameLeet,
		firstNameLeet + lastNameLeet,
		lastNameLeet + firstNameLeet,
		firstNameLeet + "123",
		lastNameLeet + "123",
		"123" + firstNameLeet,
		"123" + lastNameLeet,
	}

	for _, combo := range combinations {
		if combo != "" {
			passwordChan <- combo
		}
	}
}

// generateKeyboardPatternCombinations creates keyboard pattern combinations
func (ag *AggressiveGenerator) generateKeyboardPatternCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	// Common keyboard patterns
	keyboardPatterns := []string{
		"qwerty", "asdfgh", "zxcvbn", "qwertyui", "asdfghjk", "zxcvbnm",
		"123456", "654321", "13579", "24680", "qwe123", "asd123", "zxc123",
		"qwerty123", "asdfgh123", "zxcvbn123", "123qwe", "123asd", "123zxc",
	}

	firstName := strings.ToLower(personalInfo.FirstName)
	lastName := strings.ToLower(personalInfo.LastName)

	for _, pattern := range keyboardPatterns {
		combinations := []string{
			firstName + pattern,
			lastName + pattern,
			pattern + firstName,
			pattern + lastName,
			firstName + "_" + pattern,
			pattern + "_" + firstName,
			lastName + "_" + pattern,
			pattern + "_" + lastName,
		}

		for _, combo := range combinations {
			passwordChan <- combo
		}
	}
}

// generateAdvancedDocumentCombinations creates advanced document-based combinations
func (ag *AggressiveGenerator) generateAdvancedDocumentCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	for docType, docValue := range personalInfo.Documents {
		if docValue == "" {
			continue
		}

		// Extract numbers from documents
		docNumbers := utils.ExtractNumbers(docValue)

		for _, num := range docNumbers {
			if num != "" {
				// Advanced document number processing
				numberParts := utils.SplitNumberIntoParts(num)

				for _, part := range numberParts {
					if part != "" {
						// Basic document number
						passwordChan <- part

						// Document number with names
						firstName := strings.ToLower(personalInfo.FirstName)
						lastName := strings.ToLower(personalInfo.LastName)

						combinations := []string{
							firstName + part,
							lastName + part,
							part + firstName,
							part + lastName,
							docType + part,
							part + docType,
							firstName + "_" + part,
							part + "_" + firstName,
							lastName + "_" + part,
							part + "_" + lastName,
							docType + "_" + part,
							part + "_" + docType,
						}

						for _, combo := range combinations {
							passwordChan <- combo
						}

						// Add special characters to document combinations
						for _, special := range ag.specialChars {
							combinations := []string{
								firstName + part + special,
								lastName + part + special,
								special + firstName + part,
								special + lastName + part,
								docType + part + special,
								special + docType + part,
							}

							for _, combo := range combinations {
								passwordChan <- combo
							}
						}
					}
				}
			}
		}
	}
}
