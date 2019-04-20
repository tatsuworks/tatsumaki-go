package tatsumakigo

// Action is the action type when updating a guild.
type Action string

const (
	ActionAdd    Action = "add"
	ActionRemove Action = "remove"
	ActionSet    Action = "set"
)
