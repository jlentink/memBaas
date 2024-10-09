package cache

import "memBaas/internal/proto"

type Status int

const (
	InvalidKey Status = iota
	NotUnique
	Ok
	NotFound
	OverwriteDenied
	Error
)

var statusText = map[Status]string{
	InvalidKey:      "Key is invalid.",
	NotUnique:       "Key is not unique",
	Ok:              "Okay",
	NotFound:        "Key not found",
	OverwriteDenied: "Overwrite denied",
	Error:           "Unknown error",
}

func (s Status) String() string {
	return statusText[s]
}

func (s Status) GetStatus() proto.GetResponseGetStatus {
	switch s {
	case NotFound:
		return proto.GetResponse_NOT_FOUND
	case Ok:
		return proto.GetResponse_OK
	default:
		return proto.GetResponse_ERROR
	}
}

func (s Status) SetStatus() proto.SetResponseSetStatus {
	switch s {
	case InvalidKey:
		return proto.SetResponse_INVALID_KEY
	case NotUnique:
		return proto.SetResponse_NOT_UNIQUE
	case Ok:
		return proto.SetResponse_OK
	default:
		return proto.SetResponse_ERROR
	}
}
