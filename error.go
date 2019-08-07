package tatsumakigo

import (
	"golang.org/x/xerrors"
)

var (
	errorRequestFailed = func(err error) error {
		return xerrors.Errorf("tatsumakigo: Failed to create request: %w", err)
	}

	errorResponseFailed = func(err error) error {
		if err != nil {
			return xerrors.Errorf("tatsumakigo: Failed to get response: %w", err)
		}

		return xerrors.New("tatsumakigo: Failed to get response")
	}

	errorParseFailed = func(err error) error {
		return xerrors.Errorf("tatsumakigo: Failed to parse response: %w", err)
	}

	errorAdjustInvalid = func() error {
		return xerrors.New("tatsumakigo: The amount to adjust for remove and set actions must be above 0")
	}

	errorAdjustBounds = func() error {
		return xerrors.New("tatsumakigo: The amount to adjust must be between 0 and 50,000 (inclusive)")
	}
)
