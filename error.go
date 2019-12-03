package tatsumakigo

import (
	"golang.org/x/xerrors"
)

func errorRequestFailed(err error) error {
	return xerrors.Errorf("tatsumakigo: failed to create request: %w", err)
}

func errorResponseFailed(err error) error {
	if err != nil {
		return xerrors.Errorf("tatsumakigo: failed to get response: %w", err)
	}

	return xerrors.New("tatsumakigo: failed to get response")
}

func errorParseFailed(err error) error {
	return xerrors.Errorf("tatsumakigo: failed to parse response: %w", err)
}

func errorAdjustInvalid() error {
	return xerrors.New("tatsumakigo: the amount to adjust for add and remove actions must be above 0")
}

func errorAdjustBounds() error {
	return xerrors.New("tatsumakigo: the amount to adjust must be between 0 and 50,000 (inclusive)")
}
