package utils

import (
	"github.com/fatih/color"
)

var (
	// === Status colors ===
	ShowingWarning = color.RGB(255, 128, 0)		// Amber — Warnings
    ShowingError = color.RGB(255, 255, 0)		// Red — Errors
	ShowingSuccess = color.RGB(0, 255, 0)   	// Green — Success
	ShowingInfo    = color.RGB(0, 170, 255)    // Blue — Info or Neutral

	// === UI accents ===
	ShowingTitle     = color.RGB(0, 255, 255)  // Cyan — Titles / Headers
	ShowingSubtitle  = color.RGB(150, 150, 255) // Light Blue — Subtitles
	ShowingHighlight = color.RGB(255, 255, 0)  // Yellow — Highlights or IDs
	ShowingMuted     = color.RGB(120, 120, 120) // Gray — Secondary text

	// === Task or Note specific ===
	ShowingNote = color.RGB(255, 255, 255) // White — Regular note text
	ShowingTask = color.RGB(180, 255, 180) // Light green — Tasks
	ShowingTag  = color.RGB(255, 128, 255) // Pink — Tags

	// === Prompt or interactive elements ===
	ShowingPrompt = color.RGB(0, 200, 255) // Light Blue — For questions
	ShowingInput  = color.RGB(255, 255, 255) // White — For user input
)