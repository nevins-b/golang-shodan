package shodan

import (
	"encoding/json"
	"fmt"
)

import "time"

type ShodanHost struct {
	City         string
	RegionCode   string `json:"region_code"`
	AreaCode     string `json:"area_code"`
	Longitude    int
	Latitude     int
	CountryCode3 string `json:"country_code3"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	PostalCode   int    `json:"postal_code"`
	DMACode      string `json:"dma_code"`
	IP           string
	Hostnames    []string
	OS           string
	Ports        []int
}

type ShodanHostSearchResponse struct {
	Matches []*ShodanHostSearchMatch
	Facets  map[string][]*ShodanHostFacet
	Total   int
}

type ShodanHostSearchMatch struct {
	OS        string
	Timestamp time.Time
	ISP       string
	ASN       string
	Hostnames []string
	IP        int
	Domains   []string
	Org       string
	Data      string
	Port      int
	IPStr     string `json:"ip_str"`
	Location  *ShodanLocation
}

type ShodanLocation struct {
	City         string
	RegionCode   string `json:"region_code"`
	AreaCode     string `json:"area_code"`
	Longitude    int
	Latitude     int
	CountryCode3 string `json:"country_code3"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	PostalCode   int    `json:"postal_code"`
	DMACode      string `json:"dma_code"`
}

type ShodanHostFacet struct {
	Count int
	Value string
}

func (s Shodan) search(uri string, query string, facets string) (*ShodanHostSearchResponse, error) {
	params := make(map[string]interface{})
	params["query"] = query
	if len(facets) > 0 {
		params["facets"] = facets
	}

	req, err := s.getURI("GET", uri, params)
	if err != nil {
		return nil, err
	}
	body, err := s.performRequest(req)
	if err != nil {
		return nil, err
	}
	res := &ShodanHostSearchResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s Shodan) Host(ip string, history bool, minify bool) (*ShodanHost, error) {
	params := make(map[string]interface{})
	params["history"] = history
	params["minify"] = minify

	uri := fmt.Sprintf("/shodan/host/%s", ip)
	req, err := s.getURI("GET", uri, params)
	if err != nil {
		return nil, err
	}
	body, err := s.performRequest(req)
	if err != nil {
		return nil, err
	}
	res := &ShodanHost{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s Shodan) Search(query string, facets string) (*ShodanHostSearchResponse, error) {
	uri := "/shodan/host/search"
	return s.search(uri, query, facets)
}

func (s Shodan) Count(query string, facets string) (*ShodanHostSearchResponse, error) {
	uri := "/shodan/host/count"
	return s.search(uri, query, facets)
}
