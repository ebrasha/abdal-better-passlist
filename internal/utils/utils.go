/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : utils.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Utility functions for input handling, file operations, and common tasks
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/color"
)

// GetUserInput prompts user for input and returns the trimmed result
func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// SavePasswordsToFile saves a list of passwords to a file
func SavePasswordsToFile(passwords []string, filePath string) error {
	defer func() {
		if r := recover(); r != nil {
			LogError("SavePasswordsToFile panic", fmt.Errorf("panic: %v", r))
		}
	}()

	// Validate input
	if len(passwords) == 0 {
		LogError("Empty password list", fmt.Errorf("no passwords to save"))
		return fmt.Errorf("no passwords to save")
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		LogError("Failed to create directory", err)
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Open file for writing
	file, err := os.Create(filePath)
	if err != nil {
		LogError("Failed to create file", err)
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Write passwords to file
	for i, password := range passwords {
		if _, err := file.WriteString(password + "\n"); err != nil {
			LogError(fmt.Sprintf("Failed to write password at index %d", i), err)
			return fmt.Errorf("failed to write password: %v", err)
		}
	}

	LogInfo("Passwords saved successfully")
	return nil
}

// ValidateEmail validates an email address format
func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// ValidatePhoneNumber validates a phone number format
func ValidatePhoneNumber(phone string) bool {
	// Remove all non-digit characters
	digits := regexp.MustCompile(`\D`).ReplaceAllString(phone, "")

	// Check if it has reasonable length (7-15 digits)
	return len(digits) >= 7 && len(digits) <= 15
}

// ValidateDate validates a date in YYYY-MM-DD format
func ValidateDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}

// ExtractNumbers extracts all numbers from a string
func ExtractNumbers(input string) []string {
	re := regexp.MustCompile(`\d+`)
	return re.FindAllString(input, -1)
}

// ReverseString reverses a string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// SplitNumberIntoParts splits a number into different parts
func SplitNumberIntoParts(number string) []string {
	var parts []string

	// Add the full number
	parts = append(parts, number)

	// Add reversed number
	parts = append(parts, ReverseString(number))

	// Add individual digits
	for _, digit := range number {
		parts = append(parts, string(digit))
	}

	// Add number split by length
	if len(number) > 2 {
		// Split into two parts
		mid := len(number) / 2
		parts = append(parts, number[:mid])
		parts = append(parts, number[mid:])

		// Split into three parts if long enough
		if len(number) > 4 {
			third := len(number) / 3
			parts = append(parts, number[:third])
			parts = append(parts, number[third:third*2])
			parts = append(parts, number[third*2:])
		}
	}

	// Add first and last digits
	if len(number) > 0 {
		parts = append(parts, string(number[0]))
		parts = append(parts, string(number[len(number)-1]))
	}

	// Add first two and last two digits
	if len(number) >= 2 {
		parts = append(parts, number[:2])
		parts = append(parts, number[len(number)-2:])
	}

	return parts
}

// GenerateVariations generates different variations of a string
func GenerateVariations(input string) []string {
	var variations []string

	if input == "" {
		return variations
	}

	// Original
	variations = append(variations, input)

	// Lowercase
	variations = append(variations, strings.ToLower(input))

	// Uppercase
	variations = append(variations, strings.ToUpper(input))

	// Title case
	variations = append(variations, strings.Title(strings.ToLower(input)))

	// First letter uppercase
	if len(input) > 0 {
		firstUpper := strings.ToUpper(string(input[0])) + strings.ToLower(input[1:])
		variations = append(variations, firstUpper)
	}

	// Reverse
	variations = append(variations, ReverseString(input))

	// Remove spaces
	noSpaces := strings.ReplaceAll(input, " ", "")
	variations = append(variations, noSpaces)

	// Replace spaces with underscores
	underscore := strings.ReplaceAll(input, " ", "_")
	variations = append(variations, underscore)

	// Replace spaces with hyphens
	hyphen := strings.ReplaceAll(input, " ", "-")
	variations = append(variations, hyphen)

	// Remove vowels
	noVowels := removeVowels(input)
	variations = append(variations, noVowels)

	return variations
}

// removeVowels removes vowels from a string
func removeVowels(input string) string {
	vowels := "aeiouAEIOU"
	result := ""
	for _, char := range input {
		if !strings.ContainsRune(vowels, char) {
			result += string(char)
		}
	}
	return result
}

// GenerateSpecialCharacters returns common special characters
func GenerateSpecialCharacters() []string {
	return []string{
		"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=",
		"[", "]", "{", "}", "|", "\\", ":", ";", "\"", "'", "<", ">", ",", ".",
		"?", "/", "~", "`",
	}
}

// GenerateCommonNumbers returns common number patterns
func GenerateCommonNumbers() []string {
	return []string{
		"123", "1234", "12345", "123456",
		"111", "222", "333", "444", "555", "666", "777", "888", "999",
		"000", "0000", "00000", "000000",
		"01", "02", "03", "04", "05", "06", "07", "08", "09", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30",
		"31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
		"50", "60", "70", "80", "90", "100",
		"2000", "2001", "2002", "2003", "2004", "2005",
		"2006", "2007", "2008", "2009", "2010", "2011",
		"2012", "2013", "2014", "2015", "2016", "2017",
		"2018", "2019", "2020", "2021", "2022", "2023", "2024",
	}
}

// RemoveDuplicates removes duplicate strings from a slice
func RemoveDuplicates(input []string) []string {
	keys := make(map[string]bool)
	var result []string

	for _, item := range input {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
}

// SortStrings sorts a slice of strings
func SortStrings(input []string) []string {
	sorted := make([]string, len(input))
	copy(sorted, input)
	sort.Strings(sorted)
	return sorted
}

// IsValidInput checks if input is valid (not empty and not just whitespace)
func IsValidInput(input string) bool {
	return strings.TrimSpace(input) != ""
}

// FormatPhoneNumber formats a phone number by removing non-digit characters
func FormatPhoneNumber(phone string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(phone, "")
}

// ExtractYearFromDate extracts year from a date string
func ExtractYearFromDate(dateStr string) string {
	if ValidateDate(dateStr) {
		parts := strings.Split(dateStr, "-")
		if len(parts) >= 1 {
			return parts[0]
		}
	}
	return ""
}

// ExtractMonthFromDate extracts month from a date string
func ExtractMonthFromDate(dateStr string) string {
	if ValidateDate(dateStr) {
		parts := strings.Split(dateStr, "-")
		if len(parts) >= 2 {
			return parts[1]
		}
	}
	return ""
}

// ExtractDayFromDate extracts day from a date string
func ExtractDayFromDate(dateStr string) string {
	if ValidateDate(dateStr) {
		parts := strings.Split(dateStr, "-")
		if len(parts) >= 3 {
			return parts[2]
		}
	}
	return ""
}

// ContainsOnlyDigits checks if a string contains only digits
func ContainsOnlyDigits(input string) bool {
	for _, char := range input {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

// ContainsOnlyLetters checks if a string contains only letters
func ContainsOnlyLetters(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

// GetFileSize returns the size of a file in bytes
func GetFileSize(filepath string) (int64, error) {
	file, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return file.Size(), nil
}

// FormatFileSize formats file size in human readable format
func FormatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

// ShowProgress displays a progress bar
func ShowProgress(current, total int, message string) {
	if total == 0 {
		return
	}

	// Ensure current is not negative
	if current < 0 {
		current = 0
	}

	// Ensure current doesn't exceed total
	if current > total {
		current = total
	}

	percentage := float64(current) / float64(total) * 100
	barLength := 50
	filledLength := int(percentage / 100 * float64(barLength))

	// Ensure filledLength is within bounds
	if filledLength < 0 {
		filledLength = 0
	}
	if filledLength > barLength {
		filledLength = barLength
	}

	// Calculate empty length safely
	emptyLength := barLength - filledLength
	if emptyLength < 0 {
		emptyLength = 0
	}

	// Create progress bar safely
	var bar string
	if filledLength > 0 {
		bar += strings.Repeat("█", filledLength)
	}
	if emptyLength > 0 {
		bar += strings.Repeat("░", emptyLength)
	}

	color.Cyan("\r%s [%s] %.1f%% (%d/%d)", message, bar, percentage, current, total)

	if current == total {
		fmt.Println()
	}
}
