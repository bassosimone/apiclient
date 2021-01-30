package datamodel

// PsiphonConfigRequest is the request for the PsiphonConfig API
type PsiphonConfigRequest struct{}

// PsiphonConfigRemoteServerListURL is the URL of a remote server
type PsiphonConfigRemoteServerListURL struct {
	URL               string
	OnlyAfterAttempts int64
	SkipVerify        bool
}

// PsiphonConfigResponse is the response from the PsiphonConfig API
type PsiphonConfigResponse map[string]interface{}
