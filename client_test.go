package gomailinabox

import (
	"os"
	"testing"
)

func errorIfNotEqualStrings(a string, b string, t *testing.T) {
	if a != b {
		t.Errorf("%s != %s", a, b)
	}
}

func errorIfEqualStrings(a string, b string, t *testing.T) {
	if a == b {
		t.Errorf("%s == %s", a, b)
	}
}

func configFromEnv() *Config {
	url := os.Getenv("GOMIAB_URL")
	username := os.Getenv("GOMIAB_USERNAME")
	password := os.Getenv("GOMIAB_PASSWORD")
	if url == "" {
		return nil
	}
	return &Config{URL: url, Username: username, Password: password}
}

func errorIfErrorNotNil(e error, t *testing.T) {
	if e != nil {
		t.Errorf("+%v", e)
	}
}

func TestCreateRecord(t *testing.T) {
	config := configFromEnv()
	if config == nil {
		t.Skipf("skipping since GOMIAB_URL is not configured")
	}
	client := NewClient(config)

	inRec := &Record{Domain: "testgomailinabox.mrkc.me", Type: "A", Value: "1.1.1.1"}

	t.Logf("Using Record %+v", inRec)
	outRecs, err := client.CreateRecord(inRec)
	errorIfErrorNotNil(err, t)
	errorIfNotEqualStrings(outRecs[0].Domain, inRec.Domain, t)
	errorIfNotEqualStrings(outRecs[0].Type, inRec.Type, t)
	errorIfNotEqualStrings(outRecs[0].Value, inRec.Value, t)

	client.DeleteRecord(inRec)
}

func TestUpdateRecord(t *testing.T) {
	config := configFromEnv()
	if config == nil {
		t.Skipf("skipping since GOMIAB_URL is not configured")
	}
	client := NewClient(config)

	oldIP, newIP := "1.1.1.1", "1.2.3.4"
	inRec := &Record{Domain: "testgomailinabox.mrkc.me", Type: "A", Value: oldIP}

	t.Logf("Using Record %+v", inRec)
	client.CreateRecord(inRec)

	outRecs, err := client.UpdateRecord(&Record{Domain: inRec.Domain, Type: inRec.Type, Value: newIP})
	errorIfErrorNotNil(err, t)
	errorIfNotEqualStrings(outRecs[0].Domain, inRec.Domain, t)
	errorIfNotEqualStrings(outRecs[0].Type, inRec.Type, t)
	errorIfNotEqualStrings(outRecs[0].Value, newIP, t)

	client.DeleteRecord(&outRecs[0])
}
