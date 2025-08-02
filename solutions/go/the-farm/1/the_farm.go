package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	numberOfCows int
	message      string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.message)
}

func DivideFood(f FodderCalculator, cows int) (float64, error) {
	totalAmount, err := f.FodderAmount(cows)

	if err != nil {
		return 0, err
	}

	factor, err := f.FatteningFactor()

	if err != nil {
		return 0, err
	}

	return totalAmount / float64(cows) * factor, nil
}

func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(f, cows)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			numberOfCows: cows,
			message:      "there are no negative cows",
		}
	}

	if cows == 0 {
		return &InvalidCowsError{
			numberOfCows: cows,
			message:      "no cows don't need food",
		}
	}

	return nil
}
