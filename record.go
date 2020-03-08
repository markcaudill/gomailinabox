package gomailinabox

// Record represents a DNS record (missing things like TTL but the mailinabox API only supports these values)
// The struct tags match the actual API response and are used by encoding/json.Unmarshal.
type Record struct {
	Domain string `json:"qname"`
	Type   string `json:"rtype"`
	Value  string `json:"value"`
}
