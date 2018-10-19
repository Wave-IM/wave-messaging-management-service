package models

// Mapping : Mapping between external and minternal user ID
type Mapping struct {
	OriginalUserID       string `json:"originalUserID"`
	InternalHermesUserID string `json:"internalHermesUserID"`
}