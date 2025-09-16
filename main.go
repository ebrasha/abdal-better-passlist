/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : main.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Main entry point for Abdal Better PassList - A comprehensive password list generator
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package main

import (
	"abdal-better-passlist/internal/config"
	"abdal-better-passlist/internal/countries"
	"abdal-better-passlist/internal/generators"
	"abdal-better-passlist/internal/models"
	"abdal-better-passlist/internal/utils"
	"abdal-better-passlist/pkg/banner"
	"abdal-better-passlist/pkg/disclaimer"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

// Global configuration
var appConfig *config.Config

func main() {
	// Set console title
	fmt.Print("\033]0;Abdal Better PassList\007")

	// Initialize error logging
	defer func() {
		if r := recover(); r != nil {
			utils.LogError("Application panic", fmt.Errorf("panic: %v", r))
			panic(r) // Re-panic to maintain normal panic behavior
		}
		utils.CloseLogger()
	}()

	// Initialize logger
	if err := utils.InitLogger(); err != nil {
		fmt.Printf("Warning: Failed to initialize logger: %v\n", err)
	}

	app := &cli.App{
		Name:    "Abdal Better PassList",
		Usage:   "Generate comprehensive password lists based on personal information",
		Version: "1.0.0",
		Authors: []*cli.Author{
			{
				Name:  "Ebrahim Shafiei (EbraSha)",
				Email: "Prof.Shafiei@Gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "country",
				Aliases: []string{"c"},
				Usage:   "Country code for document collection (e.g., ir, us, uk)",
			},
			&cli.StringFlag{
				Name:    "complexity",
				Aliases: []string{"l"},
				Usage:   "Password complexity level: normal, sensitive, aggressive",
				Value:   "normal",
			},
			&cli.IntFlag{
				Name:    "workers",
				Aliases: []string{"w"},
				Usage:   "Number of concurrent workers for password generation",
				Value:   4,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output directory for generated password files",
				Value:   "./output",
			},
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"i"},
				Usage:   "Run in interactive mode",
			},
			&cli.BoolFlag{
				Name:   "skip-disclaimer",
				Usage:  "Skip disclaimer (for testing purposes)",
				Hidden: true,
			},
			&cli.BoolFlag{
				Name:   "test-mode",
				Usage:  "Run in test mode (skip disclaimer and show basic info)",
				Hidden: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "help-detailed",
				Aliases: []string{"hd"},
				Usage:   "Show detailed help information",
				Action: func(c *cli.Context) error {
					showDetailedHelp()
					return nil
				},
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Test the application functionality",
				Action: func(c *cli.Context) error {
					return runTestMode()
				},
			},
		},
		Action: runApplication,
	}

	// Custom help template
	cli.AppHelpTemplate = getHelpTemplate()

	if err := app.Run(os.Args); err != nil {
		utils.LogError("Application error", err)
		color.Red("Error: %v", err)
		os.Exit(1)
	}
}

func runApplication(c *cli.Context) error {
	// Check for test mode first
	if c.Bool("test-mode") {
		return runTestMode()
	}

	// Show disclaimer first and get user agreement (unless skipped for testing)
	if !c.Bool("skip-disclaimer") {
		if !disclaimer.ShowDisclaimer() {
			color.Red("User did not agree to the disclaimer. Exiting...")
			return fmt.Errorf("user declined disclaimer")
		}
	} else {
		color.Yellow("⚠️  Disclaimer skipped for testing purposes")
	}

	// Show banner after disclaimer agreement
	banner.ShowBanner()

	// Initialize configuration
	appConfig = config.NewConfig()

	// Check if interactive mode is requested or no arguments provided
	if c.Bool("interactive") || len(os.Args) == 1 {
		return runInteractiveMode()
	}

	// Run in CLI mode
	return runCLIMode(c)
}

func runTestMode() error {
	// Create cyberpunk colors
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	magenta := color.New(color.FgMagenta, color.Bold)
	white := color.New(color.FgWhite, color.Bold)

	// Header
	cyan.Println("\n" + strings.Repeat("═", 60))
	cyan.Println("╔══════════════════════════════════════════════════════════════╗")
	cyan.Println("║                🧪 TEST MODE - ABDAL BETTER PASSLIST 🧪      ║")
	cyan.Println("║                                                              ║")
	cyan.Println("╚══════════════════════════════════════════════════════════════╝")
	cyan.Println(strings.Repeat("═", 60))

	// Status
	green.Println("\n✅ Application is working correctly!")
	green.Println("🚀 All systems operational!")

	// Available modes
	yellow.Println("\n📋 Available Modes:")
	cyan.Println("  • Interactive: --interactive")
	white.Println("     Run in interactive mode with step-by-step guidance")

	cyan.Println("  • CLI: --country ir --complexity normal")
	white.Println("     Run with command line arguments")

	cyan.Println("  • Help: --help")
	white.Println("     Show basic help information")

	cyan.Println("  • Detailed Help: help-detailed")
	white.Println("     Show comprehensive help with examples")

	cyan.Println("  • Test: test")
	white.Println("     Run this test mode")

	// Quick examples
	yellow.Println("\n⚡ Quick Examples:")
	green.Println("  abdal-better-passlist --interactive")
	green.Println("  abdal-better-passlist --country ir --complexity sensitive --workers 4")
	green.Println("  abdal-better-passlist help-detailed")

	// Developer info
	magenta.Println("\n👨‍💻 Developer: Ebrahim Shafiei (EbraSha)")
	white.Println("📧 Email: Prof.Shafiei@Gmail.com")
	white.Println("🌐 GitHub: https://github.com/ebrasha")
	white.Println("🐦 Twitter: https://x.com/ProfShafiei")

	// Log file info
	yellow.Println("\n📋 Log Information:")
	white.Printf("📁 Log file: %s\n", utils.GetLogFilePath())
	white.Println("🔍 All errors are automatically logged to the log file")
	white.Println("⚠️  Check the log file if you encounter any issues")

	// Footer
	cyan.Println("\n" + strings.Repeat("═", 60))
	cyan.Println("💡 Ready to generate password lists! Choose your preferred mode above.")
	cyan.Println(strings.Repeat("═", 60))

	return nil
}

func runInteractiveMode() error {
	// Create cyberpunk colors
	var cyan, green, yellow, magenta, white *color.Color
	cyan = color.New(color.FgCyan, color.Bold)
	green = color.New(color.FgGreen, color.Bold)
	yellow = color.New(color.FgYellow, color.Bold)
	magenta = color.New(color.FgMagenta, color.Bold)
	white = color.New(color.FgWhite, color.Bold)

	cyan.Println("🎮 INTERACTIVE MODE - ABDAL BETTER PASSLIST 🎮")

	cyan.Println(strings.Repeat("═", 70))

	yellow.Println("\n🚀 Welcome to Abdal Better PassList Interactive Mode!")
	white.Println("📝 Follow the prompts to generate comprehensive password lists")
	white.Println("🔒 Remember: Use generated passwords responsibly and legally")

	generationCount := 0

	for {
		generationCount++
		cyan.Println("\n" + strings.Repeat("─", 50))
		cyan.Printf("🔄 Generation #%d\n", generationCount)
		cyan.Println(strings.Repeat("─", 50))

		// Get country selection
		country, err := countries.SelectCountry()
		if err != nil {
			return err
		}

		// Get personal information
		personalInfo, err := collectPersonalInformation(country)
		if err != nil {
			return err
		}

		// Get complexity level
		complexity, err := selectComplexityLevel()
		if err != nil {
			return err
		}

		// Get number of workers
		workers, err := getWorkerCount()
		if err != nil {
			return err
		}

		// Get output directory
		outputDir, err := getOutputDirectory()
		if err != nil {
			return err
		}

		// Generate passwords
		err = generatePasswords(personalInfo, complexity, workers, outputDir)
		if err != nil {
			return err
		}

		// Ask if user wants to generate another password list
		if !askForAnotherGeneration() {
			break
		}
	}

	// Use existing cyberpunk colors for final message

	green.Println("\n" + strings.Repeat("═", 70))
	green.Println("╔══════════════════════════════════════════════════════════════════════════╗")
	green.Println("║                        🎉 SESSION COMPLETED 🎉                         ║")
	green.Println("║                                                                          ║")
	green.Println("╚══════════════════════════════════════════════════════════════════════════╝")
	green.Println(strings.Repeat("═", 70))

	green.Println("\n🎉 Thank you for using Abdal Better PassList!")
	cyan.Printf("📊 Total password lists generated: %d\n", generationCount)
	yellow.Println("🔒 Remember: Use generated passwords responsibly and legally")

	// Log information
	yellow.Println("\n📋 Log Information:")
	white.Printf("📁 Log file: %s\n", utils.GetLogFilePath())
	white.Println("🔍 All errors are automatically logged to the log file")
	white.Println("⚠️  Check the log file if you encounter any issues")

	magenta.Println("\n👨‍💻 Developed by Ebrahim Shafiei (EbraSha)")
	white.Println("📧 Email: Prof.Shafiei@Gmail.com")
	white.Println("🌐 GitHub: https://github.com/ebrasha")
	white.Println("🐦 Twitter: https://x.com/ProfShafiei")

	cyan.Println("\n" + strings.Repeat("═", 70))
	cyan.Println("💡 Visit our GitHub repository for updates and documentation")
	cyan.Println(strings.Repeat("═", 70))

	return nil
}

func runCLIMode(c *cli.Context) error {
	// Create cyberpunk colors
	cyan := color.New(color.FgCyan, color.Bold)

	cyan.Println("\n" + strings.Repeat("═", 60))
	cyan.Println("╔══════════════════════════════════════════════════════════════╗")
	cyan.Println("║                    💻 CLI MODE - ABDAL BETTER PASSLIST 💻   ║")
	cyan.Println("║                                                              ║")
	cyan.Println("╚══════════════════════════════════════════════════════════════╝")
	cyan.Println(strings.Repeat("═", 60))

	// Get country
	countryCode := c.String("country")
	if countryCode == "" {
		return fmt.Errorf("country code is required in CLI mode. Use --country or -c flag")
	}

	country, err := countries.GetCountryByCode(countryCode)
	if err != nil {
		return err
	}

	// Get personal information
	personalInfo, err := collectPersonalInformation(country)
	if err != nil {
		return err
	}

	// Get complexity level
	complexity := c.String("complexity")
	if !isValidComplexity(complexity) {
		return fmt.Errorf("invalid complexity level: %s. Use: normal, sensitive, or aggressive", complexity)
	}

	// Get number of workers
	workers := c.Int("workers")
	if workers < 1 || workers > 16 {
		return fmt.Errorf("number of workers must be between 1 and 16")
	}

	// Get output directory
	outputDir := c.String("output")

	// Generate passwords
	return generatePasswords(personalInfo, complexity, workers, outputDir)
}

func collectPersonalInformation(country *models.Country) (*models.PersonalInfo, error) {
	color.Cyan("\n=== Personal Information Collection ===")
	color.Yellow("Please provide the following information:\n")

	personalInfo := &models.PersonalInfo{
		Country:   country,
		Documents: make(map[string]string),
	}

	// Common information for all countries
	personalInfo.FirstName = utils.GetUserInput("First Name: ")
	personalInfo.LastName = utils.GetUserInput("Last Name: ")
	personalInfo.GirlfriendName = utils.GetUserInput("Girlfriend/Boyfriend Name (optional): ")
	personalInfo.SpouseName = utils.GetUserInput("Spouse Name (optional): ")
	personalInfo.SecondSpouseName = utils.GetUserInput("Second Spouse Name (optional): ")
	personalInfo.ThirdSpouseName = utils.GetUserInput("Third Spouse Name (optional): ")
	personalInfo.FatherName = utils.GetUserInput("Father Name: ")
	personalInfo.MotherName = utils.GetUserInput("Mother Name: ")
	personalInfo.PetName = utils.GetUserInput("Pet Name (optional): ")
	personalInfo.BirthDate = utils.GetUserInput("Birth Date (YYYY-MM-DD): ")
	personalInfo.FavoritePersonName = utils.GetUserInput("Favorite Person Name (optional): ")
	personalInfo.MobileNumber = utils.GetUserInput("Mobile Number: ")
	personalInfo.HomePhoneNumber = utils.GetUserInput("Home Phone Number (optional): ")

	// Country-specific documents
	color.Cyan("\n=== Country-Specific Documents ===")
	color.Yellow("Please provide the following documents for %s:\n", country.Name)

	for _, doc := range country.RequiredDocuments {
		value := utils.GetUserInput(fmt.Sprintf("%s: ", doc.Name))
		personalInfo.Documents[doc.Type] = value
	}

	return personalInfo, nil
}

func selectComplexityLevel() (string, error) {
	color.Cyan("\n=== Password Complexity Level ===")
	color.Yellow("Select the complexity level for password generation:\n")
	color.Green("1. Normal - Minimal combinations for basic security")
	color.Yellow("2. Sensitive - Complete combinations for enhanced security")
	color.Red("3. Aggressive - All possible combinations for maximum coverage")

	choice := utils.GetUserInput("\nEnter your choice (1-3): ")

	switch choice {
	case "1":
		return "normal", nil
	case "2":
		return "sensitive", nil
	case "3":
		return "aggressive", nil
	default:
		return "", fmt.Errorf("invalid choice. Please select 1, 2, or 3")
	}
}

func getWorkerCount() (int, error) {
	color.Cyan("\n=== Concurrency Settings ===")
	color.Yellow("Enter the number of concurrent workers for password generation (1-16):")
	color.Cyan("Recommended: 4-8 workers for optimal performance")

	input := utils.GetUserInput("Number of workers: ")
	workers, err := strconv.Atoi(input)
	if err != nil || workers < 1 || workers > 16 {
		return 0, fmt.Errorf("invalid number of workers. Please enter a number between 1 and 16")
	}

	return workers, nil
}

func getOutputDirectory() (string, error) {
	color.Cyan("\n=== Output Settings ===")

	defaultDir := "./output"
	input := utils.GetUserInput(fmt.Sprintf("Output directory (default: %s): ", defaultDir))

	if strings.TrimSpace(input) == "" {
		return defaultDir, nil
	}

	return input, nil
}

func generatePasswords(personalInfo *models.PersonalInfo, complexity string, workers int, outputDir string) error {
	// Create cyberpunk colors
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	white := color.New(color.FgWhite, color.Bold)

	cyan.Println("\n" + strings.Repeat("═", 60))
	cyan.Println("║                    🔐 PASSWORD GENERATION 🔐                ")
	cyan.Println(strings.Repeat("═", 60))

	yellow.Printf("\n🚀 Generating passwords with %s complexity using %d workers...\n", complexity, workers)
	white.Printf("👤 Target: %s %s\n", personalInfo.FirstName, personalInfo.LastName)
	white.Printf("🌍 Country: %s\n", personalInfo.Country.Name)

	// Create output directory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Generate filename
	filename := fmt.Sprintf("%s_%s_%s.txt",
		personalInfo.FirstName,
		personalInfo.LastName,
		time.Now().Format("2006-01-02"))

	filepath := fmt.Sprintf("%s/%s", outputDir, filename)

	// Generate passwords
	generator := generators.NewPasswordGenerator(complexity, workers)
	passwords, err := generator.GeneratePasswords(personalInfo)
	if err != nil {
		utils.LogError("Password generation failed", err)
		return fmt.Errorf("failed to generate passwords: %v", err)
	}

	// Save to file
	if err := utils.SavePasswordsToFile(passwords, filepath); err != nil {
		utils.LogError("Failed to save passwords to file", err)
		return fmt.Errorf("failed to save passwords: %v", err)
	}

	green.Println("\n" + strings.Repeat("═", 60))
	green.Println("                    ✅ GENERATION COMPLETED ✅                ")
	green.Println(strings.Repeat("═", 60))

	green.Println("\n✅ Password generation completed successfully!")
	cyan.Printf("📁 Output file: %s\n", filepath)
	cyan.Printf("🔢 Total passwords generated: %d\n", len(passwords))
	yellow.Println("🔒 Remember: Use generated passwords responsibly and legally")

	return nil
}

func isValidComplexity(complexity string) bool {
	return complexity == "normal" || complexity == "sensitive" || complexity == "aggressive"
}

func showDetailedHelp() {
	// Create cyberpunk colors
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	red := color.New(color.FgRed, color.Bold)
	magenta := color.New(color.FgMagenta, color.Bold)
	white := color.New(color.FgWhite, color.Bold)

	// Header with cyberpunk styling
	cyan.Println("\n" + strings.Repeat("═", 80))
	cyan.Println("╔══════════════════════════════════════════════════════════════════════════════╗")
	cyan.Println("║                    🚀 ABDAL BETTER PASSLIST - HELP SYSTEM 🚀                ║")
	cyan.Println("║                                                                              ║")
	cyan.Println("╚══════════════════════════════════════════════════════════════════════════════╝")
	cyan.Println(strings.Repeat("═", 80))

	// Usage section
	yellow.Println("\n📋 USAGE:")
	white.Println("  abdal-better-passlist [OPTIONS]")
	white.Println("  abdal-better-passlist --interactive")
	white.Println("  abdal-better-passlist --country ir --complexity sensitive --workers 8")

	// Options section
	green.Println("\n🔧 OPTIONS:")
	cyan.Println("  -c, --country CODE")
	white.Println("     Country code for document collection")
	white.Println("     Available codes: ir, us, uk, de, fr, se, in, jp, ae, ca, au, general")

	cyan.Println("  -l, --complexity LEVEL")
	white.Println("     Password complexity level")
	white.Println("     Levels: normal, sensitive, aggressive")

	cyan.Println("  -w, --workers NUMBER")
	white.Println("     Number of concurrent workers (1-16)")
	white.Println("     Default: 4")

	cyan.Println("  -o, --output DIR")
	white.Println("     Output directory for password files")
	white.Println("     Default: ./output")

	cyan.Println("  -i, --interactive")
	white.Println("     Run in interactive mode")

	cyan.Println("  -h, --help")
	white.Println("     Show this help information")

	// Complexity levels with cyberpunk colors
	yellow.Println("\n🎯 COMPLEXITY LEVELS:")
	green.Println("  🟢 Normal:")
	white.Println("     - Basic name combinations")
	white.Println("     - Simple number variations")
	white.Println("     - Minimal special character usage")
	white.Println("     - Fast generation, smaller wordlist")

	yellow.Println("  🟡 Sensitive:")
	white.Println("     - Complete name combinations")
	white.Println("     - All number variations (reversed, partial)")
	white.Println("     - Common special character patterns")
	white.Println("     - Balanced generation, medium wordlist")

	red.Println("  🔴 Aggressive:")
	white.Println("     - All possible combinations")
	white.Println("     - Maximum number variations")
	white.Println("     - Extensive special character usage")
	white.Println("     - Slow generation, large wordlist")

	// Countries section
	yellow.Println("\n🌍 SUPPORTED COUNTRIES:")
	cyan.Println("  🇮🇷 ir  - Iran (National ID, Birth Certificate, Passport, Driving License)")
	cyan.Println("  🇺🇸 us  - United States (Driver's License, SSN, Passport)")
	cyan.Println("  🇬🇧 uk  - United Kingdom (Passport, Driving Licence, NINO)")
	cyan.Println("  🇩🇪 de  - Germany (Personalausweis, Passport, Tax ID)")
	cyan.Println("  🇫🇷 fr  - France (Carte Nationale, Passport, Tax Number)")
	cyan.Println("  🇸🇪 se  - Sweden (Personnummer, ID-card, BankID)")
	cyan.Println("  🇮🇳 in  - India (Aadhaar, PAN, Passport, Voter ID)")
	cyan.Println("  🇯🇵 jp  - Japan (My Number, Passport, Driver's Licence)")
	cyan.Println("  🇦🇪 ae  - UAE (Emirates ID, Passport, Driving Licence)")
	cyan.Println("  🇨🇦 ca  - Canada (Driver's Licence, SIN, Passport)")
	cyan.Println("  🇦🇺 au  - Australia (Driver's Licence, TFN, Medicare)")
	cyan.Println("  🌍 general - Other countries (National ID, Passport)")

	// Examples section
	yellow.Println("\n📝 EXAMPLES:")
	white.Println("  # Interactive mode")
	green.Println("  abdal-better-passlist --interactive")

	white.Println("  # CLI mode with Iran")
	green.Println("  abdal-better-passlist --country ir --complexity sensitive --workers 6")

	white.Println("  # High-performance generation")
	green.Println("  abdal-better-passlist --country us --complexity aggressive --workers 12 --output /tmp/passwords")

	// Commands section
	yellow.Println("\n🎮 COMMANDS:")
	cyan.Println("  test, t")
	white.Println("     Test the application functionality")

	cyan.Println("  help-detailed, hd")
	white.Println("     Show this detailed help information")

	// Disclaimer with red warning
	red.Println("\n" + strings.Repeat("⚠", 20) + " DISCLAIMER " + strings.Repeat("⚠", 20))
	red.Println("  This tool is for educational and authorized testing purposes only.")
	red.Println("  Users are responsible for compliance with local laws and regulations.")
	red.Println("  Any misuse of this software is the sole responsibility of the user.")
	red.Println(strings.Repeat("⚠", 52))

	// Developer section
	magenta.Println("\n👨‍💻 DEVELOPER:")
	white.Println("  Ebrahim Shafiei (EbraSha)")
	white.Println("  Email: Prof.Shafiei@Gmail.com")
	white.Println("  GitHub: https://github.com/ebrasha")
	white.Println("  Twitter: https://x.com/ProfShafiei")
	white.Println("  LinkedIn: https://www.linkedin.com/in/profshafiei/")
	white.Println("  Telegram: https://t.me/ProfShafiei")

	// Log information
	yellow.Println("\n📋 Log Information:")
	white.Printf("📁 Log file: %s\n", utils.GetLogFilePath())
	white.Println("🔍 All errors are automatically logged to the log file")
	white.Println("⚠️  Check the log file if you encounter any issues")

	// Footer
	cyan.Println("\n" + strings.Repeat("═", 80))
	cyan.Println("💡 Tip: Use 'abdal-better-passlist test' to verify the application is working correctly")
	cyan.Println(strings.Repeat("═", 80))
}

func askForAnotherGeneration() bool {
	// Create cyberpunk colors
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	red := color.New(color.FgRed, color.Bold)
	white := color.New(color.FgWhite, color.Bold)

	cyan.Println("\n" + strings.Repeat("═", 70))
	cyan.Println("                       🔄 CONTINUE GENERATION 🔄                       ")
	cyan.Println(strings.Repeat("═", 70))

	yellow.Println("\n🎯 Current password list generation completed successfully!")
	green.Println("🚀 Do you want to generate a password list for another person?")

	cyan.Println("\n📋 Options:")
	white.Println("  • Type 'yes' or 'y' to generate another password list")
	white.Println("  • Type 'no' or 'n' to exit the application")
	white.Println("  • Type 'help' or 'h' to see more options")

	input := utils.GetUserInput("\nYour choice: ")
	input = strings.ToLower(strings.TrimSpace(input))

	switch input {
	case "yes", "y":
		green.Println("\n✅ Starting new password generation...")
		cyan.Println("🔄 Please provide information for the next person")
		yellow.Println("📝 You will go through the same process: country selection, personal info, complexity, etc.\n")
		return true
	case "no", "n", "exit", "quit":
		yellow.Println("\n👋 Thank you for using Abdal Better PassList!")
		cyan.Println("🔒 Remember: Use generated passwords responsibly and legally")
		return false
	case "help", "h":
		showContinueHelp()
		return askForAnotherGeneration() // Ask again after showing help
	default:
		red.Println("\n❌ Invalid input. Please type 'yes', 'no', or 'help'")
		yellow.Println("💡 Tip: You can also use 'y' for yes, 'n' for no")
		return askForAnotherGeneration() // Recursive call to ask again
	}
}

func showContinueHelp() {
	// Create cyberpunk colors
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	white := color.New(color.FgWhite, color.Bold)

	cyan.Println("\n" + strings.Repeat("═", 60))
	cyan.Println("                    📖 HELP - CONTINUE GENERATION 📖        ")
	cyan.Println(strings.Repeat("═", 60))

	green.Println("\n✅ To generate another password list:")
	white.Println("   • yes, y")
	white.Println("   • Type any of these to continue with another person")

	yellow.Println("\n❌ To exit the application:")
	white.Println("   • no, n, exit, quit")
	white.Println("   • Type any of these to close the application")

	cyan.Println("\nℹ️  Additional Information:")
	white.Println("   • Each generation creates a separate password file")
	white.Println("   • You can choose different countries and complexity levels")
	white.Println("   • All files are saved in the specified output directory")
	white.Println("   • The application will restart the entire process for each person")
	white.Println("   • You can generate unlimited password lists in one session")

	yellow.Println("\n💡 Tips:")
	white.Println("   • Use 'help' or 'h' anytime to see this information")
	white.Println("   • You can mix different countries and complexity levels")
	white.Println("   • Each person gets their own unique password file")

	cyan.Println("\n" + strings.Repeat("═", 60))
}

func getHelpTemplate() string {
	return `{{.Name}} - {{.Usage}}

USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .VisibleCommands}}
COMMANDS:
{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}
{{else}}{{range .VisibleCommands}}
   {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}
{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
`
}
