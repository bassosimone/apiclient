package apiclient

const defaultBaseURL = "https://ps1.ooni.io"

func (c *Client) baseURL() string {
	if c.BaseURL != "" {
		return c.BaseURL
	}
	return defaultBaseURL
}
