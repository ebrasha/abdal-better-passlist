/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : generator.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Main password generator orchestrator with concurrency support
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
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/fatih/color"
)

// PasswordGenerator handles password generation with different complexity levels
type PasswordGenerator struct {
	complexity    string
	workers       int
	normalGen     *NormalGenerator
	sensitiveGen  *SensitiveGenerator
	aggressiveGen *AggressiveGenerator
}

// NewPasswordGenerator creates a new password generator instance
func NewPasswordGenerator(complexity string, workers int) *PasswordGenerator {
	return &PasswordGenerator{
		complexity:    complexity,
		workers:       workers,
		normalGen:     NewNormalGenerator(),
		sensitiveGen:  NewSensitiveGenerator(),
		aggressiveGen: NewAggressiveGenerator(),
	}
}

// GeneratePasswords generates passwords based on personal information and complexity level
func (pg *PasswordGenerator) GeneratePasswords(personalInfo *models.PersonalInfo) ([]string, error) {
	startTime := time.Now()

	color.Cyan("🚀 Starting password generation...")
	color.Yellow("📊 Complexity Level: %s", pg.complexity)
	color.Yellow("⚡ Workers: %d", pg.workers)
	color.Yellow("👤 Target: %s %s", personalInfo.FirstName, personalInfo.LastName)

	// Create channels for concurrent processing
	passwordChan := make(chan string, 1000)
	done := make(chan bool)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < pg.workers; i++ {
		wg.Add(1)
		go pg.worker(i, personalInfo, passwordChan, &wg)
	}

	// Collect passwords
	var passwords []string
	go func() {
		defer func() {
			if r := recover(); r != nil {
				utils.LogError("Password collection panic", fmt.Errorf("panic: %v", r))
				done <- true
			}
		}()

		for password := range passwordChan {
			passwords = append(passwords, password)
			if len(passwords)%1000 == 0 {
				// Use a safe total for progress display
				safeTotal := 10000
				if len(passwords) > safeTotal {
					safeTotal = len(passwords) + 1000
				}
				utils.ShowProgress(len(passwords), safeTotal, "Generating passwords")
			}
		}
		done <- true
	}()

	// Wait for all workers to complete
	wg.Wait()
	close(passwordChan)
	<-done

	// Remove duplicates
	passwords = utils.RemoveDuplicates(passwords)

	duration := time.Since(startTime)

	color.Green("✅ Password generation completed!")
	color.Cyan("📈 Generated: %d unique passwords", len(passwords))
	color.Cyan("⏱️  Duration: %v", duration)
	color.Cyan("🔧 Workers used: %d", pg.workers)

	return passwords, nil
}

// worker processes password generation in a separate goroutine
func (pg *PasswordGenerator) worker(id int, personalInfo *models.PersonalInfo, passwordChan chan<- string, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			utils.LogError(fmt.Sprintf("Worker %d panic", id), fmt.Errorf("panic: %v", r))
		}
		wg.Done()
	}()

	// Generate passwords based on complexity level
	switch pg.complexity {
	case "normal":
		pg.normalGen.GeneratePasswords(personalInfo, passwordChan)
	case "sensitive":
		pg.sensitiveGen.GeneratePasswords(personalInfo, passwordChan)
	case "aggressive":
		pg.aggressiveGen.GeneratePasswords(personalInfo, passwordChan)
	default:
		utils.LogError("Unknown complexity level", fmt.Errorf("complexity: %s", pg.complexity))
		color.Red("❌ Unknown complexity level: %s", pg.complexity)
	}
}

// GetComplexityInfo returns information about the current complexity level
func (pg *PasswordGenerator) GetComplexityInfo() string {
	switch pg.complexity {
	case "normal":
		return "Normal complexity - Basic combinations for standard security testing"
	case "sensitive":
		return "Sensitive complexity - Enhanced combinations for thorough security testing"
	case "aggressive":
		return "Aggressive complexity - Maximum combinations for comprehensive security testing"
	default:
		return "Unknown complexity level"
	}
}

// GetWorkerInfo returns information about the worker configuration
func (pg *PasswordGenerator) GetWorkerInfo() string {
	return fmt.Sprintf("Using %d concurrent workers (max recommended: %d)", pg.workers, runtime.NumCPU()*2)
}

// ValidateConfiguration validates the generator configuration
func (pg *PasswordGenerator) ValidateConfiguration() error {
	if pg.workers < 1 || pg.workers > 16 {
		return fmt.Errorf("invalid number of workers: %d (must be between 1 and 16)", pg.workers)
	}

	if pg.complexity != "normal" && pg.complexity != "sensitive" && pg.complexity != "aggressive" {
		return fmt.Errorf("invalid complexity level: %s (must be normal, sensitive, or aggressive)", pg.complexity)
	}

	return nil
}

// GetEstimatedTime returns estimated generation time based on complexity and workers
func (pg *PasswordGenerator) GetEstimatedTime() time.Duration {
	baseTime := time.Minute * 2 // Base time for normal complexity with 4 workers

	switch pg.complexity {
	case "normal":
		baseTime = time.Minute * 1
	case "sensitive":
		baseTime = time.Minute * 5
	case "aggressive":
		baseTime = time.Minute * 15
	}

	// Adjust for number of workers (more workers = faster)
	workerFactor := float64(4) / float64(pg.workers)
	if workerFactor < 0.5 {
		workerFactor = 0.5
	}

	return time.Duration(float64(baseTime) * workerFactor)
}

// GetMemoryUsage returns estimated memory usage
func (pg *PasswordGenerator) GetMemoryUsage() string {
	var memUsage string

	switch pg.complexity {
	case "normal":
		memUsage = "~50MB"
	case "sensitive":
		memUsage = "~200MB"
	case "aggressive":
		memUsage = "~500MB"
	}

	return memUsage
}
