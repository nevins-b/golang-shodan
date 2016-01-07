package shodan

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Shodan struct {
	client *http.Client
	apiKey string
	Debug  bool
}

func NewShodan(client *http.Client, apiKey string) *Shodan {
	if client == nil {
		client = &http.Client{}
	}
	return &Shodan{
		client: client,
		apiKey: apiKey,
		Debug:  false,
	}
}

func (s Shodan) getURI(method string, uri string, params map[string]interface{}) (*http.Request, error) {
	urlStr := fmt.Sprintf("https://api.shodan.io%s", uri)
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, err
	}
	values := req.URL.Query()
	values.Add("key", s.apiKey)
	if method == "GET" && params != nil {
		for key := range params {
			values.Add(key, params[key].(string))
		}
	}
	req.URL.RawQuery = values.Encode()
	if method == "POST" && params != nil {
		for key := range params {
			req.Form.Add(key, params[key].(string))
		}
	}
	if s.Debug {
		fmt.Printf("Shodan request uri: %s", req.URL.String())
	}
	return req, nil
}

func (s Shodan) performRequest(request *http.Request) ([]byte, error) {
	resp, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if s.Debug {
		fmt.Printf("Status Code: %d\n", resp.StatusCode)
		fmt.Printf("Response Body: %s\n", string(body))
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Got response code %d with body %s", resp.StatusCode, string(body))
	}
	return body, nil
}
