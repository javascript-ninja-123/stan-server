package subject

// Subject  will be exported
type Subject string

// EntityUploaded will be exported
var (
	EntityUploaded Subject = "EntityUploaded"
	EntityMoved    Subject = "EntityMoved"
	EntityDeleted  Subject = "EntityDeleted"
	EntityReNamed  Subject = "EntityReNamed"
	EntityUpdated  Subject = "EntityUpdated"
)
