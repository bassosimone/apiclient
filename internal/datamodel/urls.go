package datamodel

// URLSRequest is the URLS request.
type URLSRequest struct {
	CategoryCodes string `query:"category_codes"`
	CountryCode   string `query:"country_code"`
	Limit         int64  `query:"limit"`
}

// URLSResponse is the URLS response.
type URLSResponse struct {
	Results []URLSResponseURL `json:"results"`
}

// URLSResponseURL is a single URL in the URLS response.
type URLSResponseURL struct {
	CategoryCode string `json:"category_code"`
	CountryCode  string `json:"country_code"`
	URL          string `json:"url"`
}
