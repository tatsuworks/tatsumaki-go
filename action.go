package tatsumakigo

// Action is the action type for updating a user's guild points or score.
type Action string

const (
	// ActionAdd is the action for adding a user's guild points or score.
	ActionAdd Action = "add"
	// ActionRemove is the action for removing a user's guild points or score.
	ActionRemove Action = "remove"
	// ActionSet is the action for setting a user's guild points or score.
	ActionSet Action = "set"
)
