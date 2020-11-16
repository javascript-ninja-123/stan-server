package event

// UploadEntityEventData will be used to parse
type UploadEntityEventData struct {
	entityId   string
	entityURL  string
	createdAt  string
	version    int
	uploaderId string
	path       string
}

// UploadEntityEvent will be used to parse
type UploadEntityEvent struct {
	subject string
	data    UploadEntityEventData
}
