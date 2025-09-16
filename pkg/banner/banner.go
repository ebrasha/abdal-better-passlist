/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : banner.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : ASCII banner display with cyberpunk styling for the application
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package banner

import (
	"fmt"
	"github.com/fatih/color"
)

// ShowBanner displays the application banner with cyberpunk styling
func ShowBanner() {
	// Clear screen (only if not in test mode)
	// fmt.Print("\033[2J\033[H")

	// Create cyberpunk colors
	cyan := color.New(color.FgCyan, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	yellow := color.New(color.FgYellow, color.Bold)
	red := color.New(color.FgRed, color.Bold)
	magenta := color.New(color.FgMagenta, color.Bold)

	// ASCII Art Banner

	cyan.Println("░█▀▀█ █▀▀▄ █▀▀▄ █▀▀█ █░░ 　 ▒█▀▀█ █▀▀ ▀▀█▀▀ ▀▀█▀▀ █▀▀ █▀▀█ 　 ▒█▀▀█ █▀▀█ █▀▀ █▀▀ ▒█░░░ ░▀░ █▀▀ ▀▀█▀▀ ")
	cyan.Println("▒█▄▄█ █▀▀▄ █░░█ █▄▄█ █░░ 　 ▒█▀▀▄ █▀▀ ░░█░░ ░░█░░ █▀▀ █▄▄▀ 　 ▒█▄▄█ █▄▄█ ▀▀█ ▀▀█ ▒█░░░ ▀█▀ ▀▀█ ░░█░░ ")
	cyan.Println("▒█░▒█ ▀▀▀░ ▀▀▀░ ▀░░▀ ▀▀▀ 　 ▒█▄▄█ ▀▀▀ ░░▀░░ ░░▀░░ ▀▀▀ ▀░▀▀ 　 ▒█░░░ ▀░░▀ ▀▀▀ ▀▀▀ ▒█▄▄█ ▀▀▀ ▀▀▀ ░░▀░░")
	cyan.Println("▒█░▒█ ▀▀▀░ ▀▀▀░ ▀░░▀ ▀▀▀ 　 ▒█▄▄█ ▀▀▀ ░░▀░░ ░░▀░░ ▀▀▀ ▀░▀▀ 　 ▒█░░░ ▀░░▀ ▀▀▀ ▀▀▀ ▒█▄▄█ ▀▀▀ ▀▀▀ ░░▀░░")

	// Application info
	fmt.Println()

	green.Println("╭──────────────────────────────────────────────────────────────────────────────╮")
	green.Println("│                                                                              │")
	green.Println("│                           🚀 ABDAL BETTER PASSLIST 🚀                        │")
	green.Println("│                                                                              │")
	yellow.Println("│   A comprehensive password list generator based on personal information      │")
	yellow.Println("│   Designed for authorized security testing and educational purposes          │")
	green.Println("│                                                                              │")
	green.Println("╰──────────────────────────────────────────────────────────────────────────────╯")

	// Developer information
	fmt.Println()
	cyan.Println("╭──────────────────────────────────────────────────────────────────────────────╮")
	cyan.Println("│                              👨‍💻 DEVELOPER INFO 👨‍💻                            │")
	cyan.Println("│                                                                              │")
	magenta.Println("│  Programmer: Ebrahim Shafiei (EbraSha)                                       │")
	magenta.Println("│  Email:      Prof.Shafiei@Gmail.com                                          │")
	green.Println("│  GitHub:     https://github.com/ebrasha                                      │")
	green.Println("│  Twitter:    https://x.com/ProfShafiei                                       │")
	yellow.Println("│  LinkedIn:   https://www.linkedin.com/in/profshafiei/                        │")
	yellow.Println("│  Telegram:   https://t.me/ProfShafiei                                        │")
	cyan.Println("│                                                                              │")
	red.Println("│  \"Coding is an engaging and beloved hobby for me. I passionately and         │")
	red.Println("│   insatiably pursue knowledge in cybersecurity and programming.\"             │")
	cyan.Println("│                                                                              │")
	cyan.Println("╰──────────────────────────────────────────────────────────────────────────────╯")

	fmt.Println()
	color.New(color.FgCyan, color.Bold).Println("🔐 Initializing Abdal Better PassList...")
	color.New(color.FgYellow).Println("⚡ Ready to generate comprehensive password lists!")
	fmt.Println()
}
