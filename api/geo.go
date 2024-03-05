package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"nimaajdari.com/nobadweather/errs"
)

type Query struct {
	Parsed struct {
		City         string `json:"city"`
		ExpectedType string `json:"expected_type"`
	} `json:"parsed"`
	Text string `json:"text"`
}

type Bbox struct {
	Lat1 float64 `json:"lat1"`
	Lat2 float64 `json:"lat2"`
	Lon1 float64 `json:"lon1"`
	Lon2 float64 `json:"lon2"`
}

type Datasource struct {
	Attribution string `json:"attribution"`
	License     string `json:"license"`
	Sourcename  string `json:"sourcename"`
	URL         string `json:"url"`
}

type Rank struct {
	Confidence          float64 `json:"confidence"`
	ConfidenceCityLevel float64 `json:"confidence_city_level"`
	Importance          float64 `json:"importance"`
	MatchType           string  `json:"match_type"`
	Popularity          float64 `json:"popularity"`
}

type Timezone struct {
	AbbreviationDST  string  `json:"abbreviation_DST"`
	AbbreviationSTD  string  `json:"abbreviation_STD"`
	Name             string  `json:"name"`
	NameAlt          string  `json:"name_alt"`
	OffsetDST        string  `json:"offset_DST"`
	OffsetDSTSeconds float64 `json:"offset_DST_seconds"`
	OffsetSTD        string  `json:"offset_STD"`
	OffsetSTDSeconds float64 `json:"offset_STD_seconds"`
}

type Result struct {
	AddressLine1  string     `json:"address_line1"`
	AddressLine2  string     `json:"address_line2"`
	Bbox          Bbox       `json:"bbox"`
	Category      string     `json:"category"`
	City          string     `json:"city"`
	Country       string     `json:"country"`
	CountryCode   string     `json:"country_code"`
	County        string     `json:"county,omitempty"`
	Datasource    Datasource `json:"datasource"`
	Formatted     string     `json:"formatted"`
	Hamlet        string     `json:"hamlet,omitempty"`
	Lat           float64    `json:"lat"`
	Lon           float64    `json:"lon"`
	Name          string     `json:"name,omitempty"`
	PlaceID       string     `json:"place_id"`
	PlusCode      string     `json:"plus_code"`
	PlusCodeShort string     `json:"plus_code_short,omitempty"`
	Postcode      string     `json:"postcode,omitempty"`
	Rank          Rank       `json:"rank"`
	Ref           string     `json:"ref,omitempty"`
	ResultType    string     `json:"result_type"`
	State         string     `json:"state,omitempty"`
	StateCode     string     `json:"state_code,omitempty"`
	Timezone      Timezone   `json:"timezone"`
}

type GeoApiResponse struct {
	Query   Query    `json:"query"`
	Results []Result `json:"results"`
}

type SearchResult struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	Country      string
	CountryCode  string
	County       string
	Formatted    string
	Lat          float64
	Lon          float64
	ResultType   string
}

func (sr SearchResult) String() string {
	return fmt.Sprintf("Type: %s, Formatted Address: %s, City: %s, County: %s, Country: %s, Country Code: %s, Latitude: %f, Longitude: %f",
		sr.ResultType, sr.Formatted, sr.City, sr.County, sr.Country, sr.CountryCode, sr.Lat, sr.Lon)
}

type GeoapifyClient struct {
	apiKey string
}

func NewGeoapifyClient(apiKey string) GeoapifyClient {
	return GeoapifyClient{apiKey: apiKey}
}

func (c GeoapifyClient) SearchLocation(query string) ([]SearchResult, error) {

	fmt.Printf("Searching for location query:  %s \n", query)

	queryUrlEncoded := url.QueryEscape(query)
	url := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&format=json&apiKey=%s", queryUrlEncoded, c.apiKey)

	response, err := http.Get(url)
	if err != nil {
		return nil, errs.StackInfo(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errs.StackInfo(err)
	}

	if response.StatusCode != 200 {
		return nil, errs.StackInfo(fmt.Errorf("got unexpected status code from geoapify API: %d and body: %s", response.StatusCode, string(body)))
	}

	var geoApiResponse GeoApiResponse
	err = json.Unmarshal(body, &geoApiResponse)
	if err != nil {
		return nil, errs.StackInfo(err)
	}

	searchResult := make([]SearchResult, 0, len(geoApiResponse.Results))
	for _, responseItem := range geoApiResponse.Results {
		result := SearchResult{
			AddressLine1: responseItem.AddressLine1,
			AddressLine2: responseItem.AddressLine2,
			City:         responseItem.City,
			Country:      responseItem.Country,
			CountryCode:  responseItem.CountryCode,
			County:       responseItem.County,
			Formatted:    responseItem.Formatted,
			Lat:          responseItem.Lat,
			Lon:          responseItem.Lon,
			ResultType:   responseItem.ResultType,
		}
		searchResult = append(searchResult, result)
	}

	return searchResult, nil
}
