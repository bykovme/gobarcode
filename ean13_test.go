package gobarcode

import "testing"

const ean13WrongLengthBarcode = "389272"
const ean13DigitsOnly1 = "12345h1234567"
const ean13DigitsOnly2 = "(2345h1234567"
const ean13WrongChecksum = "4600605022419"
const ean13CorrectChecksum = "4600605022412"
const roundMin = 51
const roundEqual = 50
const roundMax = 59

func TestEan13Length(t *testing.T) {
	err := CheckEAN13(ean13WrongLengthBarcode)
	if err.Error() != ErrWrongLength {
		t.Error("Expected error: '" + ErrWrongLength + "', got instead '" + err.Error() + "'")
	}
}

func TestEan13Digits(t *testing.T) {
	err := CheckEAN13(ean13DigitsOnly1)
	if err.Error() != ErrNotNumber {
		t.Error("Expected error '" + ErrNotNumber + "', got instead '" + err.Error() + "'")
	}
	err = CheckEAN13(ean13DigitsOnly2)
	if err.Error() != ErrNotNumber {
		t.Error("Expected error '" + ErrNotNumber + "', got instead '" + err.Error() + "'")
	}
}

func TestEan13WrongChecksum(t *testing.T) {
	err := CheckEAN13(ean13WrongChecksum)
	if err.Error() != ErrWrongCheckSum {
		t.Error("Expected error '" + ErrWrongCheckSum + "', got instead '" + err.Error() + "'")
	}
}

func TestEan13CorrectChecksum(t *testing.T) {
	err := CheckEAN13(ean13CorrectChecksum)
	if err != nil {
		t.Error("Expected no error, got instead '" + err.Error() + "'")
	}
}

func TestRoundMin(t *testing.T) {
	res1 := round(roundMin)
	if res1 != 60 {
		t.Errorf("Rounded %d, expected %d, got instead %d", roundMin, 60, res1)
	}
}

func TestRoundMax(t *testing.T) {
	res1 := round(roundMax)
	if res1 != 60 {
		t.Errorf("Rounded %d, expected %d, got instead %d", roundMin, 60, res1)
	}
}

func TestRoundEqual(t *testing.T) {
	res1 := round(roundEqual)
	if res1 != 50 {
		t.Errorf("Rounded %d, expected %d, got instead %d", roundMin, 60, res1)
	}
}
