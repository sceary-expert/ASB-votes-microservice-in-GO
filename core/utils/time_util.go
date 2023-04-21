package utils

import (
	"strconv"
	"time"

	"example.com/core/pkg/log"
)

// MomentToTime Covert moment unix to time.Time
func MomentToTime(moment string) (time.Time, error) {
	i, err := strconv.ParseInt(moment, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}

// UTCNow Get moment UTC NOW
func UTCNowUnix() int64 {
	return (time.Now().UTC().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)))
}

// UTCNow Get moment UTC NOW
func TimeUnix(inputTime time.Time) int64 {
	return (inputTime.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)))
}

// UTCUnixToTime Get UTC unix time in go time
func UTCUnixToTime(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(1000000))
}

// IsTimeExpired check time expiration in unix
func IsTimeExpired(timestamp int64, offsetInSeconds float64) bool {
	tt := UTCUnixToTime(timestamp)
	remainder := tt.Sub(time.Now())
	log.Info("remainder: %v calc : %v", remainder, (remainder.Seconds() + offsetInSeconds))

	return !((remainder.Seconds() + offsetInSeconds) > 0)
}

// IsTimeExpired check time expiration in golang time
func IsTimeExpiredInTime(tt time.Time, offsetInSeconds float64) bool {
	remainder := tt.Sub(time.Now())
	log.Info("remainder: %v calc : %v", remainder, (remainder.Seconds() + offsetInSeconds))

	return !((remainder.Seconds() + offsetInSeconds) > 0)
}
