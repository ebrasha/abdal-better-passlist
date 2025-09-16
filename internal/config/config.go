/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : config.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Configuration management for the password generation system
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package config

import (
	"abdal-better-passlist/internal/models"
	"fmt"
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	DefaultComplexity      string
	DefaultWorkers         int
	DefaultOutputDir       string
	MaxWorkers             int
	MinWorkers             int
	MaxPasswordsNormal     int
	MaxPasswordsSensitive  int
	MaxPasswordsAggressive int
	EnableLogging          bool
	LogLevel               string
	EnableProgressBar      bool
	EnableColorOutput      bool
}

// NewConfig creates a new configuration instance with default values
func NewConfig() *Config {
	return &Config{
		DefaultComplexity:      "normal",
		DefaultWorkers:         4,
		DefaultOutputDir:       "./output",
		MaxWorkers:             16,
		MinWorkers:             1,
		MaxPasswordsNormal:     10000,
		MaxPasswordsSensitive:  100000,
		MaxPasswordsAggressive: 1000000,
		EnableLogging:          true,
		LogLevel:               "info",
		EnableProgressBar:      true,
		EnableColorOutput:      true,
	}
}

// LoadFromEnvironment loads configuration from environment variables
func (c *Config) LoadFromEnvironment() {
	// Load default complexity
	if complexity := os.Getenv("ABDAL_COMPLEXITY"); complexity != "" {
		if isValidComplexity(complexity) {
			c.DefaultComplexity = complexity
		}
	}

	// Load default workers
	if workers := os.Getenv("ABDAL_WORKERS"); workers != "" {
		if w, err := strconv.Atoi(workers); err == nil && w >= c.MinWorkers && w <= c.MaxWorkers {
			c.DefaultWorkers = w
		}
	}

	// Load output directory
	if outputDir := os.Getenv("ABDAL_OUTPUT_DIR"); outputDir != "" {
		c.DefaultOutputDir = outputDir
	}

	// Load max workers
	if maxWorkers := os.Getenv("ABDAL_MAX_WORKERS"); maxWorkers != "" {
		if w, err := strconv.Atoi(maxWorkers); err == nil && w > 0 {
			c.MaxWorkers = w
		}
	}

	// Load logging settings
	if enableLogging := os.Getenv("ABDAL_ENABLE_LOGGING"); enableLogging != "" {
		if b, err := strconv.ParseBool(enableLogging); err == nil {
			c.EnableLogging = b
		}
	}

	// Load log level
	if logLevel := os.Getenv("ABDAL_LOG_LEVEL"); logLevel != "" {
		if isValidLogLevel(logLevel) {
			c.LogLevel = logLevel
		}
	}

	// Load progress bar setting
	if enableProgress := os.Getenv("ABDAL_ENABLE_PROGRESS"); enableProgress != "" {
		if b, err := strconv.ParseBool(enableProgress); err == nil {
			c.EnableProgressBar = b
		}
	}

	// Load color output setting
	if enableColor := os.Getenv("ABDAL_ENABLE_COLOR"); enableColor != "" {
		if b, err := strconv.ParseBool(enableColor); err == nil {
			c.EnableColorOutput = b
		}
	}
}

// GetMaxPasswordsForComplexity returns the maximum number of passwords for a given complexity level
func (c *Config) GetMaxPasswordsForComplexity(complexity string) int {
	switch complexity {
	case "normal":
		return c.MaxPasswordsNormal
	case "sensitive":
		return c.MaxPasswordsSensitive
	case "aggressive":
		return c.MaxPasswordsAggressive
	default:
		return c.MaxPasswordsNormal
	}
}

// ValidateWorkers validates the number of workers
func (c *Config) ValidateWorkers(workers int) error {
	if workers < c.MinWorkers {
		return fmt.Errorf("number of workers must be at least %d", c.MinWorkers)
	}
	if workers > c.MaxWorkers {
		return fmt.Errorf("number of workers cannot exceed %d", c.MaxWorkers)
	}
	return nil
}

// GetGenerationConfig creates a GenerationConfig from the current config
func (c *Config) GetGenerationConfig(complexity string, workers int, outputDir string) *models.GenerationConfig {
	return &models.GenerationConfig{
		Complexity:       complexity,
		Workers:          workers,
		OutputDir:        outputDir,
		MaxPasswords:     c.GetMaxPasswordsForComplexity(complexity),
		IncludeSpecial:   true,
		IncludeNumbers:   true,
		IncludeUppercase: true,
		IncludeLowercase: true,
	}
}

// isValidComplexity checks if a complexity level is valid
func isValidComplexity(complexity string) bool {
	return complexity == "normal" || complexity == "sensitive" || complexity == "aggressive"
}

// isValidLogLevel checks if a log level is valid
func isValidLogLevel(level string) bool {
	validLevels := []string{"debug", "info", "warn", "error"}
	for _, validLevel := range validLevels {
		if level == validLevel {
			return true
		}
	}
	return false
}

// GetComplexityDescription returns a description for a complexity level
func (c *Config) GetComplexityDescription(complexity string) string {
	switch complexity {
	case "normal":
		return "Minimal combinations for basic security - Fast generation, smaller wordlist"
	case "sensitive":
		return "Complete combinations for enhanced security - Balanced generation, medium wordlist"
	case "aggressive":
		return "All possible combinations for maximum coverage - Slow generation, large wordlist"
	default:
		return "Unknown complexity level"
	}
}

// GetComplexityLevels returns all available complexity levels with their descriptions
func (c *Config) GetComplexityLevels() []models.PasswordComplexity {
	return []models.PasswordComplexity{
		{
			Level:        "normal",
			Name:         "Normal",
			Description:  c.GetComplexityDescription("normal"),
			MaxPasswords: c.MaxPasswordsNormal,
		},
		{
			Level:        "sensitive",
			Name:         "Sensitive",
			Description:  c.GetComplexityDescription("sensitive"),
			MaxPasswords: c.MaxPasswordsSensitive,
		},
		{
			Level:        "aggressive",
			Name:         "Aggressive",
			Description:  c.GetComplexityDescription("aggressive"),
			MaxPasswords: c.MaxPasswordsAggressive,
		},
	}
}
