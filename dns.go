package shodan

import (
	"encoding/json"
	"strings"
)

func (s Shodan) DNSResolve(hostnames []string) (map[string]string, error) {
	h := strings.Join(hostnames, ",")
	params := make(map[string]interface{})
	params["hostnames"] = h
	req, err := s.getURI("GET", "/dns/resolve", params)
	if err != nil {
		return nil, err
	}
	body, err := s.performRequest(req)
	if err != nil {
		return nil, err
	}
	res := make(map[string]string)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s Shodan) DNSReverse(ips []string) (map[string][]string, error) {
	h := strings.Join(ips, ",")
	params := make(map[string]interface{})
	params["ips"] = h
	req, err := s.getURI("GET", "/dns/reverse", params)
	if err != nil {
		return nil, err
	}
	body, err := s.performRequest(req)
	if err != nil {
		return nil, err
	}
	res := make(map[string][]string)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
