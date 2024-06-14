package timezone

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	customerrors "github.com/seemsod1/api-project/internal/errors"
)

// GetTimezoneDiff returns the difference between the local time and the specified hour.
func GetTimezoneDiff(needHour int) int {
	localHour := time.Now().Hour()
	timeZoneDiff := calculateTimeZoneDiff(localHour, needHour)
	return adjustTimeZoneDiff(localHour, needHour, timeZoneDiff)
}

func calculateTimeZoneDiff(localHour, needHour int) int {
	return needHour - localHour
}

// adjustTimeZoneDiff adjusts the time difference if it exceeds 12 hours in either direction.
func adjustTimeZoneDiff(localHour, needHour, timeZoneDiff int) int {
	if timeZoneDiff > 12 || timeZoneDiff < -12 {
		if timeZoneDiff < 0 {
			timeZoneDiff *= -1
		}
		switch {
		case localHour > needHour:
			return 24 - timeZoneDiff
		case localHour < needHour:
			return timeZoneDiff - 24
		default:
			return 0
		}
	}
	return timeZoneDiff
}

// ProcessTimezoneHeader processes the timezone header and returns the offset.
func ProcessTimezoneHeader(r *http.Request) (int, error) {
	userTimezone := r.Header.Get("Accept-Timezone")
	if userTimezone == "" {
		return 0, nil
	}

	offsetStr, err := extractOffsetString(userTimezone)
	if err != nil {
		return 0, err
	}

	offset, err := parseOffset(offsetStr)
	if err != nil {
		return 0, err
	}

	return offset, nil
}

func extractOffsetString(userTimezone string) (string, error) {
	switch {
	case userTimezone == "UTC":
		return "0", nil
	case len(userTimezone) == 5 && (strings.HasPrefix(userTimezone, "UTC+") || strings.HasPrefix(userTimezone, "UTC-")):
		return userTimezone[3:], nil
	default:
		return "", customerrors.ErrInvalidTimezone
	}
}

func parseOffset(offsetStr string) (int, error) {
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset > 12 || offset < -12 {
		return 0, customerrors.ErrInvalidTimezone
	}
	return offset, nil
}