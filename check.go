package gobarcode

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func round(f int) int {
	if f%10 == 0 {
		return int(f)
	}
	f1 := float64(f) / 10
	flooredF := math.Floor(f1 + 1.0)
	return int(flooredF * 10)
}

func CheckEAN13(barcode string) (err error) {

	barcode = strings.Trim(barcode, " ")

	if len(barcode) != 13 {
		return errors.New(ErrWrongLength)
	}

	if _, err := strconv.Atoi(barcode); err != nil {
		return errors.New(ErrNotNumber)
	}

	even := false
	sum := 0
	for i := 0; i < 12; i++ {
		digit := barcode[i : i+1]
		intDigit, err := strconv.Atoi(digit)
		if err != nil {
			return errors.New(ErrNotNumber)
		}
		if even {
			sum += intDigit * 3
		} else {
			sum += intDigit
		}

		even = !even
	}

	roundedSum := round(sum)

	checkSum := roundedSum - sum

	checkDigit, err := strconv.Atoi(barcode[12:13])
	if err != nil {
		return errors.New(ErrNotNumber)
	}

	if checkDigit != checkSum {
		return errors.New(ErrWrongCheckSum)
	}

	return nil
}
