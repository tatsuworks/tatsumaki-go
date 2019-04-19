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
)
