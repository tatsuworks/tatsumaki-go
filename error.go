package tatsumakigo

import (
	"golang.org/x/xerrors"
)

var (
	errorRequestFailed = func(err error) error {
		return xerrors.Errorf("tatsumakigo: failed to create request: %w", err)
	}

	errorResponseFailed = func(err error) error {
		if err != nil {
			return xerrors.Errorf("tatsumakigo: failed to get response: %w", err)
		}

		return xerrors.New("tatsumakigo: failed to get response")
	}

	errorParseFailed = func(err error) error {
		return xerrors.Errorf("tatsumakigo: failed to parse response: %w", err)
	}

	errorAdjustInvalid = func() error {
		return xerrors.New("tatsumakigo: the amount to adjust for remove and set actions must be above 0")
	}

	errorAdjustBounds = func() error {
		return xerrors.New("tatsumakigo: the amount to adjust must be between 0 and 50,000 (inclusive)")
	}
)
