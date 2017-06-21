package v1

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"net/http"
	"strconv"
	"time"
)

// generateNonce generates a nonce to be included in a given request for the requestEndpoint.
// Returns the generated nonce, as well as the nonce seed.
func (s *Session) generateNonce(requestEndpoint string) (string, int, error) {
	nonce := int(time.Now().Unix()) + 1000
	mac := hmac.New(sha512.New, []byte(s.APISecret))
	_, err := mac.Write([]byte(strconv.Itoa(nonce) + requestEndpoint))
	if err != nil {
		return "", 0, err
	}
	expectedMAC := mac.Sum(nil)

	return hex.EncodeToString(expectedMAC), nonce, nil
}

/*
	doTRTRequest generates and executes a GET request to a given endpoint,
	while adding the proper headers to successfully authenticate with
	TRT APIs.

	To be able to successfully make an authenticated request to TRT APIs,
	one must add these headers:
    - X-TRT-KEY: the api key
    - X-TRT-SIGN: a signature of the request
    - X-TRT-NONCE: an integer
*/
func (s *Session) doTRTRequest(requestEndpoint string) (io.ReadCloser, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestEndpoint, nil)

	if err != nil {
		return nil, err
	}

	nonce, nonceSeed, err := s.generateNonce(requestEndpoint)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-TRT-KEY", s.APIKey)
	req.Header.Add("X-TRT-SIGN", nonce)
	req.Header.Add("X-TRT-NONCE", string(nonceSeed))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
