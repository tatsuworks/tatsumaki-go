package tatsumakigo

import (
	"golang.org/x/xerrors"
)

// ErrAdjustBounds is returned by the adjust guild points and score methods if the amount to adjust is invalid.
var ErrAdjustBounds = xerrors.New("tatsumakigo: the amount to adjust must be between 0 and 50,000 (inclusive)")

// ErrAdjustInvalid is returned by the adjust guild points and score methods if the amount to adjust
// for an action is invalid.
var ErrAdjustInvalid = xerrors.New("tatsumakigo: the amount to adjust for add and remove actions must be above 0")

var ErrLeaderboardLimit = xerrors.New("tatsumakigo: the limit cannot be less than -1")

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
