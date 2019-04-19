package tatsumaki_go

import "errors"

var (
	errorRequestFailed = func(err error) error {
		return errors.New("tatsumakigo: Failed to create request:\n" + err.Error())
	}

	errorResponseFailed = func(err error) error {
		if err != nil {
			return errors.New("tatsumakigo: Failed to get response:\n" + err.Error())
		} else {
			return errors.New("tatsumakigo: Failed to get response")
		}
	}
)
