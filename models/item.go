package models

import "time"



// ItemStatus is an external GraphQL representation of ItemStatus.
type ItemStatus string

// ItemStatus graphql enum values.
const (
	ItemStatusUnknown   ItemStatus = "UNKNOWN"
	ItemStatusDraft     ItemStatus = "DRAFT"
	ItemStatusPublished ItemStatus = "PUBLISHED"
)

func (i ItemStatus) DatabaseStatus() Status {
	switch i {
	case ItemStatusDraft:
		return StatusDraft
	case ItemStatusPublished:
		return StatusPublished
	}

	return StatusUnknown
}

func (s Status) ItemStatus() ItemStatus {
	switch s {
	case StatusDraft:
		return ItemStatusDraft
	case StatusPublished:
		return ItemStatusPublished
	}

	return ItemStatusUnknown
}
