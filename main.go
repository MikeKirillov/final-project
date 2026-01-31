package main

import (
	// "final-project/pkg/db"
	// "final-project/pkg/server"
	// "log"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

const LAYOUT = "20060102"
const DAYS_LIMIT = 400

func main() {
	// err := db.Init("scheduler.db")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// server.Start()
	date1, _ := NextDate(time.Now(), "20230301", "y") // 20250301
	fmt.Println("expected: 20260301, returned: " + date1)

	date2, _ := NextDate(time.Now(), "20260113", "d 7") // 20240120
	fmt.Println("expected: 20260203, returned: " + date2)

	_, err := NextDate(time.Now(), "20260113", "")
	fmt.Println(err)

	_, err = NextDate(time.Now(), "20260113", "    ")
	fmt.Println(err)

	_, err = NextDate(time.Now(), "20260113", "d")
	fmt.Println(err)

	_, err = NextDate(time.Now(), "20260113", "d 401")
	fmt.Println(err)

	_, err = NextDate(time.Now(), "20260113", "w ...")
	fmt.Println(err)

}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// Split the 'repeat' value on slices.
	// The 1st element is always a repeating type (y, d, m, w).
	// The rest of the elements are values that are used depending on the repeating type
	splitedRep := strings.Split(repeat, " ")

	// the 'repeat' value check rules
	err := repChecks(repeat, splitedRep[0])
	if err != nil {
		return "", err
	}

	var increasedDate time.Time

	// For the year-rule, we check it in this way
	// or for an array's length that is always equal to 1 in this specific case
	if splitedRep[0] == "y" {
		date, err := time.Parse(LAYOUT, dstart)
		if err != nil {
			return "", err
		}

		increasedDate = date
		// The returned date must be greater than the date specified in the 'now' variable.
		// In simpler terms, when searching for the next date, you should start counting from
		// 'dstart' and repeat intervals until the date becomes greater than 'now'
		for increasedDate.Before(now) {
			increasedDate = increasedDate.AddDate(1, 0, 0)
		}
	} else {
		daysCount, err := strconv.Atoi(splitedRep[1])
		if err != nil {
			return "", err
		}

		if daysCount > DAYS_LIMIT {
			return "", errors.New("the 'repeat' value contains an interval that exceeds the maximum allowed limit")
		}

		date, err := time.Parse(LAYOUT, dstart)
		if err != nil {
			return "", err
		}

		increasedDate = date
		// same check for 'increasedDate < now'
		for increasedDate.Before(now) {
			increasedDate = increasedDate.AddDate(0, 0, daysCount)
		}
	}
	return increasedDate.Format(LAYOUT), nil
}

func repChecks(repeat string, repType string) error {
	if strings.TrimSpace(repeat) == "" {
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
