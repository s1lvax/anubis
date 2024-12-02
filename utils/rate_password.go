package utils

import (
	"math"
	"unicode"
)

func EvaluatePassword(password string) (int, float64, string, []string) {
	length := len(password)

	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false
	repeatedChars := 0
	characterCounts := make(map[rune]int)
	characterPoolSize := 0

	for _, char := range password {
		characterCounts[char]++
		if characterCounts[char] > 1 {
			repeatedChars++
		}
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if hasLower {
		characterPoolSize += 26
	}
	if hasUpper {
		characterPoolSize += 26
	}
	if hasDigit {
		characterPoolSize += 10
	}
	if hasSpecial {
		characterPoolSize += 32
	}

	var entropy float64
	if characterPoolSize > 0 {
		entropy = math.Log2(float64(characterPoolSize)) * float64(length)
	}

	score := 0
	if length >= 20 {
		score += 4
	} else if length >= 16 {
		score += 3
	} else if length >= 12 {
		score += 2
	} else if length >= 8 {
		score += 1
	}

	if hasLower {
		score += 1
	}
	if hasUpper {
		score += 1
	}
	if hasDigit {
		score += 1
	}
	if hasSpecial {
		score += 3
	}

	if repeatedChars > 0 {
		score -= repeatedChars / 5
	}

	if score < 0 {
		score = 0
	}
	if score > 10 {
		score = 10
	}

	var rating string
	switch {
	case score >= 9:
		rating = "Very Strong"
	case score >= 7:
		rating = "Strong"
	case score >= 5:
		rating = "Moderate"
	default:
		rating = "Weak"
	}

	var feedback []string
	if length < 12 {
		feedback = append(feedback, "Increase the password length to at least 12 characters.")
	}
	if !hasLower {
		feedback = append(feedback, "Include lowercase letters.")
	}
	if !hasUpper {
		feedback = append(feedback, "Include uppercase letters.")
	}
	if !hasDigit {
		feedback = append(feedback, "Include digits.")
	}
	if !hasSpecial {
		feedback = append(feedback, "Include special characters.")
	}
	if repeatedChars > 0 {
		feedback = append(feedback, "Avoid repeated characters.")
	}

	return score, entropy, rating, feedback
}
