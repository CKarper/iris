package destination

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type defaultDestination struct {
	baseDestination `yaml:",inline"`
	URL             string `yaml:"url"`
	Secret          string `yaml:"secret"`
}

func getHmac(secret string, payload []byte) string {
	if secret != "" {
		key := []byte(secret)
		mac := hmac.New(sha256.New, key)
		mac.Write(payload)
		hmac := base64.URLEncoding.EncodeToString(mac.Sum(nil))
		return hmac
	}
	return ""
}

func (d *defaultDestination) Exec(payload interface{}) {
	fmt.Printf("Executing default destination to %s\n", d.URL)
	mJSON, _ := json.Marshal(payload)
	contentReader := bytes.NewReader(mJSON)
	req, _ := http.NewRequest("POST", d.URL, contentReader)
	req.Header.Set("X-IRIS-HMAC", getHmac(d.Secret, mJSON))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	client.Do(req)
}
