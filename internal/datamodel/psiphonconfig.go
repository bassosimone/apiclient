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
type PsiphonConfigResponse struct {
	ClientPlatform                     string
	ClientVersion                      int64
	PropagationChannelId               string
	SponsorId                          string
	RemoteServerListURLs               []PsiphonConfigRemoteServerListURL
	RemoteServerListDownloadFilename   string
	RemoteServerListSignaturePublicKey string
	TargetApiProtocol                  string
	EstablishTunnelTimeoutSeconds      int64
	LocalHttpProxyPort                 int64
	LocalSocksProxyPort                int64
	UseIndistinguishableTLS            bool
}
