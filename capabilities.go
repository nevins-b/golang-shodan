package shodan

import "encoding/json"

func (s Shodan) Ports() ([]int, error) {
	req, err := s.getURI("GET", "/shodan/ports", nil)
	if err != nil {
		return nil, err
	}
	body, err := s.performRequest(req)
	if err != nil {
		return nil, err
	}
	res := make([]int, 10, 100)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s Shodan) Protocols() (map[string]string, error) {
	req, err := s.getURI("GET", "/shodan/protocols", nil)
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

func (s Shodan) Services() (map[string]string, error) {
	req, err := s.getURI("GET", "/shodan/services", nil)
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
