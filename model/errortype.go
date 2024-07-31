package model

const (
	BindFailed StatusCode = "bind_failed"
	// Failure sends the custom error message and API message from the logic
	Failure        StatusCode = "failure"
	Ok             StatusCode = "ok"
	RecordCreated  StatusCode = "record_created"
	RecordUpdated  StatusCode = "record_updated"
	RecordDeleted  StatusCode = "record_deleted"
	RecordNotFound StatusCode = "record_not_found"
	// UnexpectedError is a server error
	UnexpectedError StatusCode = "unexpected_error"
	// AuthError is any of authorization errors
	AuthError StatusCode = "authorization_error"
)
