package util

import (
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	base62Regex      = regexp.MustCompile(`/go/([a-zA-Z0-9]{1,10})`)
	maxInt64InBase62 = "aZl8N0y58M7"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ConvertToBase62 converts a given number to its base 62 representation.
func ConvertToBase62(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	result := make([]byte, 0, 11)
	for num > 0 {
		result = append([]byte{base62Chars[num%62]}, result...)
		num = num / 62
	}

	return string(result)
}

func IsValidBase62String(str string) bool {
	if len(str) > len(maxInt64InBase62) || len(str) == 0 {
		return false
	}

	for _, c := range str {
		if strings.IndexRune(base62Chars, c) == -1 {
			return false
		}
	}

	for i, c := range str {
		strVal := int64(strings.IndexRune(base62Chars, c))
		maxVal := int64(strings.IndexRune(base62Chars, rune(maxInt64InBase62[i])))

		if strVal > maxVal {
			return false
		} else if strVal < maxVal {
			return true
		}
	}

	return true
}

func ConvertFromBase62(str string) int64 {
	var num int64 = 0
	for _, c := range str {
		num = num*62 + int64(strings.IndexRune(base62Chars, c))
	}
	return num
}

func isBase62Encoded(input string) bool {
	return base62Regex.MatchString(input)
}

// CalculateNextStep calculates the next step based on the lastTime, currentTime, step, and minStep.
// It returns the calculated next step value.
func CalculateNextStep(lastTime, currentTime time.Time, step, minStep int64) int64 {
	duration := currentTime.Sub(lastTime)
	var nextStep int64

	// Determine the next step based on the duration.
	switch {
	case duration < 15*time.Minute:
		nextStep = step * 2
	case duration > 30*time.Minute:
		nextStep = step / 2
	default:
		nextStep = step
	}

	// Ensure the next step is not less than the minimum step.
	if nextStep < minStep {
		nextStep = minStep
	}

	return nextStep
}

// GenerateIDs generates a list of IDs between start and end, excluding those in the skipList.
func GenerateIDs(start, end int64, skipList []int64) []int64 {
	ids := make([]int64, 0)
	skipMap := make(map[int64]struct{}, len(skipList))
	for _, id := range skipList {
		skipMap[id] = struct{}{}
	}
	for i := start; i <= end; i++ {
		if _, ok := skipMap[i]; !ok {
			ids = append(ids, i)
		}
	}
	return ids
}

// ExtractGoPath extracts the Go path from the given string.
// It takes a string parameter and returns a string.
func ExtractGoPath(str string) string {
	matches := base62Regex.FindStringSubmatch(str)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// IsValidURL checks if a string is a valid http(s) address.
// Returns true if the string is a valid URL, otherwise false.
func IsValidURL(str string) bool {
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if u.Scheme == "" || u.Host == "" {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	return true
}
