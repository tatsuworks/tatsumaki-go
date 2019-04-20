package tatsumakigo

import "errors"

var (
	errorRequestFailed = func(err error) error {
		return errors.New("tatsumakigo: Failed to create request: " + err.Error())
	}

	errorResponseFailed = func(err error) error {
		if err != nil {
			return errors.New("tatsumakigo: Failed to get response: " + err.Error())
		} else {
			return errors.New("tatsumakigo: Failed to get response")
		}
	}

	errorParseFailed = func(err error) error {
		return errors.New("tatsumakigo: Failed to parse response: " + err.Error())
	}

	errorAdjustInvalid = func() error {
		return errors.New("tatsumakigo: The amount to adjust for remove and set actions must be above 0")
	}

	errorAdjustBounds = func() error {
		return errors.New("tatsumakigo: The amount to adjust must be between 0 and 50,000 (inclusive)")
	}
)
