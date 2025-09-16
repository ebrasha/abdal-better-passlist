/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : models.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Data models and structures for the password generation system
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package models

// Country represents a country with its specific document requirements
type Country struct {
	Code              string         `json:"code"`
	Name              string         `json:"name"`
	RequiredDocuments []DocumentType `json:"required_documents"`
}

// DocumentType represents a specific document type required for a country
type DocumentType struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

// PersonalInfo contains all personal information collected from the user
type PersonalInfo struct {
	// Basic Information
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`

	// Family Information
	FatherName         string `json:"father_name"`
	MotherName         string `json:"mother_name"`
	SpouseName         string `json:"spouse_name"`
	SecondSpouseName   string `json:"second_spouse_name"`
	ThirdSpouseName    string `json:"third_spouse_name"`
	GirlfriendName     string `json:"girlfriend_name"`
	FavoritePersonName string `json:"favorite_person_name"`

	// Contact Information
	MobileNumber    string `json:"mobile_number"`
	HomePhoneNumber string `json:"home_phone_number"`

	// Pet Information
	PetName string `json:"pet_name"`

	// Country and Documents
	Country   *Country          `json:"country"`
	Documents map[string]string `json:"documents"`
}

// PasswordComplexity represents the different complexity levels
type PasswordComplexity struct {
	Level        string `json:"level"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	MaxPasswords int    `json:"max_passwords"`
}

// GenerationConfig contains configuration for password generation
type GenerationConfig struct {
	Complexity       string `json:"complexity"`
	Workers          int    `json:"workers"`
	OutputDir        string `json:"output_dir"`
	MaxPasswords     int    `json:"max_passwords"`
	IncludeSpecial   bool   `json:"include_special"`
	IncludeNumbers   bool   `json:"include_numbers"`
	IncludeUppercase bool   `json:"include_uppercase"`
	IncludeLowercase bool   `json:"include_lowercase"`
}

// PasswordResult contains the result of password generation
type PasswordResult struct {
	Passwords  []string `json:"passwords"`
	Count      int      `json:"count"`
	Complexity string   `json:"complexity"`
	Duration   string   `json:"duration"`
	FilePath   string   `json:"file_path"`
}

// GetComplexityLevels returns all available complexity levels
func GetComplexityLevels() []PasswordComplexity {
	return []PasswordComplexity{
		{
			Level:        "normal",
			Name:         "Normal",
			Description:  "Minimal combinations for basic security",
			MaxPasswords: 10000,
		},
		{
			Level:        "sensitive",
			Name:         "Sensitive",
			Description:  "Complete combinations for enhanced security",
			MaxPasswords: 100000,
		},
		{
			Level:        "aggressive",
			Name:         "Aggressive",
			Description:  "All possible combinations for maximum coverage",
			MaxPasswords: 1000000,
		},
	}
}

// GetSupportedCountries returns all supported countries
func GetSupportedCountries() []Country {
	return []Country{
		{
			Code: "ir",
			Name: "Iran",
			RequiredDocuments: []DocumentType{
				{Type: "national_id", Name: "National ID Card", Description: "Iranian National ID Card", Required: true},
				{Type: "birth_certificate", Name: "Birth Certificate", Description: "Birth certificate or family registry", Required: true},
				{Type: "passport", Name: "Passport", Description: "Iranian passport", Required: false},
				{Type: "driving_license", Name: "Driving License", Description: "Iranian driving license", Required: false},
				{Type: "birth_certificate_number", Name: "Birth Certificate Number", Description: "Birth certificate serial number", Required: false},
				{Type: "citizenship_card", Name: "Citizenship Card", Description: "Citizenship card number", Required: false},
			},
		},
		{
			Code: "us",
			Name: "United States",
			RequiredDocuments: []DocumentType{
				{Type: "drivers_license", Name: "Driver's License", Description: "State driver's license or state ID", Required: true},
				{Type: "passport", Name: "Passport", Description: "US passport", Required: false},
				{Type: "ssn", Name: "Social Security Number", Description: "Social Security Number (SSN)", Required: true},
				{Type: "birth_certificate", Name: "Birth Certificate", Description: "Birth certificate", Required: false},
				{Type: "voter_registration", Name: "Voter Registration", Description: "Voter registration card", Required: false},
				{Type: "tax_id", Name: "Tax ID", Description: "Tax identification number", Required: false},
			},
		},
		{
			Code: "uk",
			Name: "United Kingdom",
			RequiredDocuments: []DocumentType{
				{Type: "passport", Name: "Passport", Description: "UK passport", Required: true},
				{Type: "driving_license", Name: "Driving Licence", Description: "UK driving licence", Required: false},
				{Type: "nino", Name: "National Insurance Number", Description: "National Insurance Number (NINO)", Required: true},
				{Type: "birth_certificate", Name: "Birth Certificate", Description: "Birth certificate", Required: false},
				{Type: "voter_registration", Name: "Voter Registration", Description: "Voter registration", Required: false},
			},
		},
		{
			Code: "de",
			Name: "Germany",
			RequiredDocuments: []DocumentType{
				{Type: "personalausweis", Name: "Personalausweis", Description: "German national ID card", Required: true},
				{Type: "passport", Name: "Passport", Description: "German passport", Required: false},
				{Type: "tax_id", Name: "Tax ID", Description: "Steueridentifikationsnummer (Tax ID)", Required: true},
				{Type: "registration", Name: "Registration", Description: "Meldebescheinigung (registration)", Required: false},
			},
		},
		{
			Code: "fr",
			Name: "France",
			RequiredDocuments: []DocumentType{
				{Type: "national_id", Name: "Carte Nationale d'Identité", Description: "French national ID card", Required: true},
				{Type: "passport", Name: "Passport", Description: "French passport", Required: false},
				{Type: "tax_number", Name: "Tax Number", Description: "Numéro fiscal (tax number)", Required: true},
				{Type: "family_book", Name: "Family Book", Description: "Livret de famille / Birth certificate", Required: false},
			},
		},
		{
			Code: "se",
			Name: "Sweden",
			RequiredDocuments: []DocumentType{
				{Type: "personnummer", Name: "Personnummer", Description: "Personal identity number", Required: true},
				{Type: "id_card", Name: "ID Card", Description: "Swedish ID card", Required: false},
				{Type: "passport", Name: "Passport", Description: "Swedish passport", Required: false},
				{Type: "bankid", Name: "BankID", Description: "Digital identity (BankID)", Required: false},
			},
		},
		{
			Code: "in",
			Name: "India",
			RequiredDocuments: []DocumentType{
				{Type: "aadhaar", Name: "Aadhaar", Description: "Biometric ID (Aadhaar)", Required: true},
				{Type: "pan", Name: "PAN", Description: "Permanent Account Number (tax ID)", Required: true},
				{Type: "passport", Name: "Passport", Description: "Indian passport", Required: false},
				{Type: "voter_id", Name: "Voter ID", Description: "Voter ID (EPIC)", Required: false},
				{Type: "driving_license", Name: "Driving Licence", Description: "Indian driving licence", Required: false},
				{Type: "birth_certificate", Name: "Birth Certificate", Description: "Birth certificate", Required: false},
			},
		},
		{
			Code: "jp",
			Name: "Japan",
			RequiredDocuments: []DocumentType{
				{Type: "my_number", Name: "My Number", Description: "Social ID (My Number)", Required: true},
				{Type: "passport", Name: "Passport", Description: "Japanese passport", Required: false},
				{Type: "driving_license", Name: "Driver's Licence", Description: "Japanese driver's licence", Required: false},
				{Type: "pension_id", Name: "Pension/Tax IDs", Description: "Pension and tax identification", Required: false},
				{Type: "family_registry", Name: "Family Registry", Description: "Koseki/Juminhyo documents", Required: false},
			},
		},
		{
			Code: "ae",
			Name: "United Arab Emirates",
			RequiredDocuments: []DocumentType{
				{Type: "emirates_id", Name: "Emirates ID", Description: "Biometric Emirates ID", Required: true},
				{Type: "passport", Name: "Passport", Description: "UAE passport", Required: false},
				{Type: "driving_license", Name: "Driving Licence", Description: "UAE driving licence", Required: false},
				{Type: "residence_visa", Name: "Residence Visa", Description: "Residence visa / Unified ID", Required: false},
			},
		},
		{
			Code: "ca",
			Name: "Canada",
			RequiredDocuments: []DocumentType{
				{Type: "drivers_license", Name: "Driver's Licence", Description: "Canadian driver's licence", Required: true},
				{Type: "passport", Name: "Passport", Description: "Canadian passport", Required: false},
				{Type: "sin", Name: "Social Insurance Number", Description: "Social Insurance Number (SIN)", Required: true},
				{Type: "provincial_id", Name: "Provincial ID", Description: "Provincial ID card", Required: false},
				{Type: "birth_certificate", Name: "Birth Certificate", Description: "Birth certificate", Required: false},
			},
		},
		{
			Code: "au",
			Name: "Australia",
			RequiredDocuments: []DocumentType{
				{Type: "drivers_license", Name: "Driver's Licence", Description: "Australian driver's licence", Required: true},
				{Type: "passport", Name: "Passport", Description: "Australian passport", Required: false},
				{Type: "tfn", Name: "Tax File Number", Description: "Tax File Number (TFN)", Required: true},
				{Type: "medicare_card", Name: "Medicare Card", Description: "Medicare card", Required: false},
				{Type: "birth_certificate", Name: "Birth Certificate", Description: "Birth certificate", Required: false},
			},
		},
		{
			Code: "general",
			Name: "General (Other Countries)",
			RequiredDocuments: []DocumentType{
				{Type: "national_id", Name: "National ID", Description: "National ID card", Required: true},
				{Type: "passport", Name: "Passport", Description: "International passport", Required: false},
				{Type: "tax_id", Name: "Tax ID", Description: "Tax identification number", Required: false},
				{Type: "social_security", Name: "Social Security", Description: "Social security number", Required: false},
				{Type: "residence_registration", Name: "Residence Registration", Description: "Local residence registration", Required: false},
			},
		},
	}
}
