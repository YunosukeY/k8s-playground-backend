package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func getClient() *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	return &http.Client{Jar: jar}
}

var client = getClient()

type getAPI[R any] struct {
	url url.URL
}

func NewGetAPI[R any](host string, path string) getAPI[R] {
	u := url.URL{}
	u.Scheme = "http"
	u.Host = host
	u.Path = path
	return getAPI[R]{url: u}
}

func (api getAPI[R]) Request(withAuth bool) (*R, error) {
	var resp *http.Response
	var err error
	if withAuth {
		req, _ := http.NewRequest(http.MethodGet, api.url.String(), nil)
		req.Header.Add("X-Auth", "test")
		resp, err = client.Do(req)
	} else {
		resp, err = client.Get(api.url.String())
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("wrong status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil
	}
	var r R
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type postAPI[B any, R any] struct {
	url url.URL
}

func NewPostAPI[B any, R any](host string, path string) postAPI[B, R] {
	u := url.URL{}
	u.Scheme = "http"
	u.Host = host
	u.Path = path
	return postAPI[B, R]{url: u}
}

func (api postAPI[B, R]) Request(b B, withAuth bool) (*R, error) {
	bs, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	var resp *http.Response
	if withAuth {
		req, _ := http.NewRequest(http.MethodPost, api.url.String(), bytes.NewBuffer(bs))
		req.Header.Add("X-Auth", "test")
		resp, err = client.Do(req)
	} else {
		resp, err = client.Post(api.url.String(), "application/json", bytes.NewBuffer(bs))
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("wrong status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil
	}
	var r R
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
