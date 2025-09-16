/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : disclaimer.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Legal disclaimer and user agreement system for the application
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package disclaimer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// ShowDisclaimer displays the legal disclaimer and gets user agreement
func ShowDisclaimer() bool {
	// Create colors for styling
	red := color.New(color.FgRed, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)

	green := color.New(color.FgGreen, color.Bold)

	// Clear screen and show disclaimer header
	// fmt.Print("\033[2J\033[H")

	red.Println("                           🚨 WARNING 🚨                                   ")
	red.Println("                                                                              	")
	red.Println("  ANY MISUSE OF THIS SOFTWARE IS THE SOLE RESPONSIBILITY OF THE USER           ")
	red.Println("  AND MUST BE USED IN ACCORDANCE WITH LOCAL LAWS AND REGULATIONS               ")
	red.Println("                                                                               ")

	fmt.Println()
	green.Println("Do you agree to these terms and conditions?")
	yellow.Println("Type 'y' to agree or 'n' to decline and exit:")
	fmt.Println()

	// Get user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your response: ")
	response, err := reader.ReadString('\n')
	if err != nil {
		red.Println("Error reading input:", err)
		return false
	}

	// Clean and check response
	response = strings.TrimSpace(strings.ToLower(response))

	if response == "y" || response == "yes" {
		green.Println("\n✅ Agreement accepted. Proceeding with the application...")
		fmt.Println()
		return true
	} else if response == "n" || response == "no" {
		red.Println("\n❌ Agreement declined. Exiting application...")
		return false
	} else {
		red.Println("\n❌ Invalid response. Please type 'y' for yes or 'n' for no")
		fmt.Println()
		return ShowDisclaimer() // Recursive call to show disclaimer again
	}
}
