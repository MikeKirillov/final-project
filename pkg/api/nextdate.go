package api

import (
	"errors"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"
)

const LAYOUT = "20060102"
const DAYS_LIMIT = 400

func nextDayHandler(w http.ResponseWriter, req *http.Request) {
	now := req.FormValue("now")
	date := req.FormValue("date")
	repeat := req.FormValue("repeat")

	dateNow, err := time.Parse(LAYOUT, now)
	if err != nil {
		http.Error(w, "the value 'now' parsing error "+req.RequestURI, http.StatusBadRequest)
		return
	}

	result, err := NextDate(dateNow, date, repeat)
	if err != nil {
		http.Error(w, err.Error()+req.RequestURI, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// Split the 'repeat' value on slices.
	// The 1st element is always a repeating type (y, d, m, w).
	// The rest of the elements are values that are used depending on the repeating type
	splitedRep := strings.Split(repeat, " ")

	var increasedDate time.Time
	// using as result even with error because
	// of https://go.dev/ref/spec#The_zero_value
	var resultDate string

	// the 'repeat' value check rules
	err := repChecks(repeat, splitedRep[0])
	if err != nil {
		return resultDate, err
	}

	// For the year-rule, we check it in this way
	// or for an array's length that is always equal to 1 in this specific case
	if splitedRep[0] == "y" {
		date, err := time.Parse(LAYOUT, dstart)
		if err != nil {
			return resultDate, err
		}

		increasedDate = date.AddDate(1, 0, 0)
		// The returned date must be greater than the date specified in the 'now' variable.
		// In simpler terms, when searching for the next date, you should start counting from
		// 'dstart' and repeat intervals until the date becomes greater than 'now'
		for increasedDate.Before(now) {
			increasedDate = increasedDate.AddDate(1, 0, 0)
		}

		resultDate = increasedDate.Format(LAYOUT)
	} else {
		daysCount, err := strconv.Atoi(splitedRep[1])
		if err != nil {
			return resultDate, err
		}

		if daysCount > DAYS_LIMIT {
			return resultDate, errors.New("the 'repeat' value contains an interval that exceeds the maximum allowed limit")
		}

		date, err := time.Parse(LAYOUT, dstart)
		if err != nil {
			return resultDate, err
		}

		increasedDate = date.AddDate(0, 0, daysCount)
		// same check for 'increasedDate < now'
		for increasedDate.Before(now) {
			increasedDate = increasedDate.AddDate(0, 0, daysCount)
		}

		resultDate = increasedDate.Format(LAYOUT)
	}
	return resultDate, nil
}

func repChecks(repeat string, repType string) error {
	if len(strings.TrimSpace(repeat)) == 0 {
		return errors.New("the 'repeat' value can not be an empty string")
	}

	repTypes := []string{"y", "d"}

	if !slices.Contains(repTypes, repType) {
		return errors.New("the 'repeat' value contains an invalid character")
	}

	if repType != repTypes[0] && len(repeat) == 1 {
		return errors.New("the 'repeat' value is missing a day range specification")
	}

	return nil
}
