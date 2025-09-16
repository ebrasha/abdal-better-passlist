/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : normal.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Normal complexity password generator with basic combinations
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

// NormalGenerator handles normal complexity password generation
type NormalGenerator struct {
	commonNumbers []string
	specialChars  []string
}

// NewNormalGenerator creates a new normal complexity generator
func NewNormalGenerator() *NormalGenerator {
	return &NormalGenerator{
		commonNumbers: utils.GenerateCommonNumbers(),
		specialChars:  []string{"!", "@", "#", "$", "%", "&", "*", "123", "456", "789"},
	}
}

// GeneratePasswords generates passwords with normal complexity
func (ng *NormalGenerator) GeneratePasswords(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	// Basic name combinations
	ng.generateBasicNameCombinations(personalInfo, passwordChan)

	// Simple number combinations
	ng.generateSimpleNumberCombinations(personalInfo, passwordChan)

	// Basic special character combinations
	ng.generateBasicSpecialCombinations(personalInfo, passwordChan)

	// Family name combinations
	ng.generateFamilyNameCombinations(personalInfo, passwordChan)

	// Date combinations
	ng.generateDateCombinations(personalInfo, passwordChan)
}

// generateBasicNameCombinations creates basic name-based passwords
func (ng *NormalGenerator) generateBasicNameCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	firstName := strings.ToLower(personalInfo.FirstName)
	lastName := strings.ToLower(personalInfo.LastName)

	// Basic combinations
	combinations := []string{
		firstName,
		lastName,
		firstName + lastName,
		lastName + firstName,
		firstName + "123",
		lastName + "123",
		firstName + "1234",
		lastName + "1234",
		strings.Title(firstName),
		strings.Title(lastName),
		strings.Title(firstName) + strings.Title(lastName),
		strings.Title(lastName) + strings.Title(firstName),
	}

	for _, combo := range combinations {
		if combo != "" {
			passwordChan <- combo
		}
	}
}

// generateSimpleNumberCombinations creates simple number-based passwords
func (ng *NormalGenerator) generateSimpleNumberCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	// Extract numbers from mobile and home phone
	mobileNumbers := utils.ExtractNumbers(personalInfo.MobileNumber)
	homeNumbers := utils.ExtractNumbers(personalInfo.HomePhoneNumber)

	// Basic number combinations
	numbers := []string{"123", "1234", "12345", "111", "222", "333", "000", "0000"}
	numbers = append(numbers, mobileNumbers...)
	numbers = append(numbers, homeNumbers...)

	// Add common years
	year := utils.ExtractYearFromDate(personalInfo.BirthDate)
	if year != "" {
		numbers = append(numbers, year)
		numbers = append(numbers, year[2:]) // Last two digits
	}

	for _, num := range numbers {
		if num != "" {
			passwordChan <- num
		}
	}
}

// generateBasicSpecialCombinations creates basic special character combinations
func (ng *NormalGenerator) generateBasicSpecialCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	firstName := strings.ToLower(personalInfo.FirstName)
	lastName := strings.ToLower(personalInfo.LastName)

	// Basic special character combinations
	specials := []string{"!", "@", "#", "$", "123", "456"}

	for _, special := range specials {
		combinations := []string{
			firstName + special,
			lastName + special,
			special + firstName,
			special + lastName,
			firstName + special + lastName,
			lastName + special + firstName,
		}

		for _, combo := range combinations {
			if combo != "" {
				passwordChan <- combo
			}
		}
	}
}

// generateFamilyNameCombinations creates family name-based passwords
func (ng *NormalGenerator) generateFamilyNameCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	// Father and mother names
	if personalInfo.FatherName != "" {
		fatherName := strings.ToLower(personalInfo.FatherName)
		combinations := []string{
			fatherName,
			fatherName + "123",
			fatherName + "1234",
			strings.Title(fatherName),
		}

		for _, combo := range combinations {
			passwordChan <- combo
		}
	}

	if personalInfo.MotherName != "" {
		motherName := strings.ToLower(personalInfo.MotherName)
		combinations := []string{
			motherName,
			motherName + "123",
			motherName + "1234",
			strings.Title(motherName),
		}

		for _, combo := range combinations {
			passwordChan <- combo
		}
	}

	// Spouse names
	if personalInfo.SpouseName != "" {
		spouseName := strings.ToLower(personalInfo.SpouseName)
		combinations := []string{
			spouseName,
			spouseName + "123",
			strings.Title(spouseName),
		}

		for _, combo := range combinations {
			passwordChan <- combo
		}
	}

	// Pet name
	if personalInfo.PetName != "" {
		petName := strings.ToLower(personalInfo.PetName)
		combinations := []string{
			petName,
			petName + "123",
			strings.Title(petName),
		}

		for _, combo := range combinations {
			passwordChan <- combo
		}
	}
}

// generateDateCombinations creates date-based passwords
func (ng *NormalGenerator) generateDateCombinations(personalInfo *models.PersonalInfo, passwordChan chan<- string) {
	if personalInfo.BirthDate == "" {
		return
	}

	year := utils.ExtractYearFromDate(personalInfo.BirthDate)
	month := utils.ExtractMonthFromDate(personalInfo.BirthDate)
	day := utils.ExtractDayFromDate(personalInfo.BirthDate)

	// Basic date combinations
	dates := []string{year, month, day}
	if len(year) == 4 {
		dates = append(dates, year[2:]) // Last two digits of year
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
			}

			for _, combo := range combinations {
				passwordChan <- combo
			}
		}
	}

	// Date format combinations
	if year != "" && month != "" && day != "" {
		formats := []string{
			day + month + year,
			month + day + year,
			year + month + day,
			day + month + year[2:],
			month + day + year[2:],
		}

		for _, format := range formats {
			passwordChan <- format
		}
	}
}
