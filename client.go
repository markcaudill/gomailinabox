package gomailinabox

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"strings"
)

// Client is a wrapper around a configured resty.Client
type Client struct {
	restClient *resty.Client
}

// NewClient returns a new, configured Client
func NewClient(c *Config) *Client {
	r := resty.New().
		SetHostURL(c.URL).
		SetHeader("User-Agent", "gomailinabox v0.0.1 (https://github.com/markcaudill/gomailinabox)").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBasicAuth(c.Username, c.Password)
	return &Client{restClient: r}
}

// CreateRecord creates a DNS record and returns the result of GetRecord(r).
// Also, if Record.Value isn't specified, the value is automatically
// populated by the API using what it perceives as the client IP.
func (c *Client) CreateRecord(r *Record) ([]Record, error) {
	urlParts := []string{"/admin/dns/custom", r.Domain, r.Type}
	if _, err := c.restClient.NewRequest().SetBody(r.Value).
		Post(strings.Join(urlParts, "/")); err != nil {
		return []Record{}, err
	}
	return c.GetRecord(r)
}

// GetRecord returns a list of Records that match the criteria in r
func (c *Client) GetRecord(r *Record) ([]Record, error) {
	urlParts := []string{"/admin/dns/custom", r.Domain, r.Type}
	resp, err := c.restClient.NewRequest().Get(strings.Join(urlParts, "/"))
	if err != nil {
		return []Record{}, err
	}
	var v []Record
	err = json.Unmarshal(resp.Body(), &v)
	if err != nil {
		return []Record{}, err
	}
	return v, err

}

// UpdateRecord updates an existing Record. Due to the underlying API, it will
// also create the Record if it doesn't already exist. Also, if Record.Value
// isn't specified, the value is automatically populated by the API using what
// it perceives as the client IP.
func (c *Client) UpdateRecord(r *Record) ([]Record, error) {
	urlParts := []string{"/admin/dns/custom", r.Domain, r.Type}
	if _, err := c.restClient.NewRequest().SetBody(r.Value).
		Put(strings.Join(urlParts, "/")); err != nil {
		return []Record{}, err
	}
	return c.GetRecord(r)
}

// DeleteRecord deletes an records that match r.
func (c *Client) DeleteRecord(r *Record) ([]Record, error) {
	urlParts := []string{"/admin/dns/custom", r.Domain, r.Type}
	// Get array of Records that will be deleted
	recs, err := c.GetRecord(r)
	if _, err = c.restClient.NewRequest().SetBody(r.Value).
		Delete(strings.Join(urlParts, "/")); err != nil {
		return []Record{}, err
	}
	return recs, err
}
