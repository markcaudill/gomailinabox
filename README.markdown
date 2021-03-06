gomailinabox
============

Description
-----------

This is an unofficial client library for the [Mail-in-a-Box](https://mailinabox.email/) API.

I wrote this because I needed it for a custom Terraform provider. The API isn't really documented (as far as I can tell), but it's implemented as a Flask application and pretty simple to follow.

I've only implemented [the DNS API](https://github.com/mail-in-a-box/mailinabox/blob/v0.44/management/daemon.py#L269) since that's all I need, but I'm open to any pull requests to flesh this out.

Usage
-----

```go
package main

import (
    "github.com/markcaudill/gomailinabox"
    "log"
)

func main() {
    client := gomailinabox.NewClient(&gomailinabox.Config{URL: "https://mail.example.com", Username: "admin@example.com", Password: "abc123"})
    inRec := &gomailinabox.Record{Domain: "testdomain.example.com", Type: "A", Value: "1.1.1.1"}
    outRecs, err := client.CreateRecord(inRec)
    if err != nil {
        log.Fatalf("Error createdin Record %+v: %+v", inRec, err)
    }
    log.Printf("Created Records: %+v", outRecs)
}
```

Documentation
-------------

```
func NewClient(c *Config) *Client
    NewClient returns a new, configured Client

func (c *Client) CreateRecord(r *Record) ([]Record, error)
    CreateRecord creates a DNS record and returns the result of GetRecord(r).
    Also, if Record.Value isn't specified, the value is automatically populated
    by the API using what it perceives as the client IP.

func (c *Client) DeleteRecord(r *Record) ([]Record, error)
    DeleteRecord deletes an records that match r.

func (c *Client) GetRecord(r *Record) ([]Record, error)
    GetRecord returns a list of Records that match the criteria in r

func (c *Client) UpdateRecord(r *Record) ([]Record, error)
    UpdateRecord updates an existing Record. Due to the underlying API, it will
    also create the Record if it doesn't already exist. Also, if Record.Value
    isn't specified, the value is automatically populated by the API using what
    it perceives as the client IP.

type Config struct {
        URL      string
        Username string
        Password string
}

type Record struct {
        Domain string `json:"qname"`
        Type   string `json:"rtype"`
        Value  string `json:"value"`
}
    Record represents a DNS record (missing things like TTL but the mailinabox
    API only supports these values) The struct tags match the actual API
    response and are used by encoding/json.Unmarshal.
```
