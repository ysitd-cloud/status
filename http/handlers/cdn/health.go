package cdn

type cfResponse struct {
	Result result `json:"result"`
}

type result struct {
	Totals resultData `json:"totals"`
}

type resultData struct {
	Since    string   `json:"since"`
	Until    string   `json:"until"`
	Requests requests `json:"requests"`
}

type requests struct {
	All         int            `json:"all"`
	Cached      int            `json:"cached"`
	Uncached    int            `json:"uncached"`
	ContentType map[string]int `json:"content_type"`
	Country     map[string]int `json:"country"`
	SSL         requestsSSL    `json:"ssl"`
	HttpStatus  map[string]int `json:"http_status"`
}

type requestsSSL struct {
	Encrypted   int `json:"encrypted"`
	Unencrypted int `json:"unencrypted"`
}
