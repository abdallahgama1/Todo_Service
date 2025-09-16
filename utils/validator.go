package utils

import (
	"errors"
	"strings"
	"time"
)

func ValidatePriority(priority string) error {
	validPriorities := []string{"Low", "Medium", "High"}
	for _, p := range validPriorities {
		if strings.EqualFold(p, priority) {
			return nil
		}
	}
	return errors.New("priority must be Low, Medium, or High")
}

func ValidateDueDate(dueDateStr string) error {
	if dueDateStr == "" {
		return nil
	}

	dueDate, err := time.Parse(time.RFC3339, dueDateStr)
	if err != nil {
		return errors.New("invalid dueDate format, must be ISO 8601 (e.g., 2025-07-20T12:00:00Z)")
	}

	referenceDate := time.Date(2025, time.July, 15, 0, 0, 0, 0, time.UTC)

	if dueDate.Before(referenceDate) {
		return errors.New("due date cannot be in the past")
	}

	return nil
}
