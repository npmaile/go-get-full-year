package gogetfullyear

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type YearProvider interface {
	GetYearEnterprise() (YearAPIReturnEnterprise, error)
	GetYearBasic() (YearAPIReturnFull, error)
}

type yearProvider struct {
	enterprise      bool
	yearProviderURL string
	c               http.Client
}

func NewYearProvider(url string, enterprise bool, timeout time.Duration) YearProvider {
	return &yearProvider{
		enterprise:      enterprise,
		c:               http.Client{Timeout: timeout},
		yearProviderURL: url,
	}
}

func (y *yearProvider) GetYearEnterprise() (YearAPIReturnEnterprise, error) {
	if !y.enterprise {
		panic("enterprise get year called, but enterprise was not passed to the constructor. Exiting now")
	}
	resp, err := y.c.Get(y.yearProviderURL)
	if err != nil {
		panic("unable to retrieve Year from server. This is almost certainly an issue on your end. Exiting")
	}
	if resp.StatusCode != 200 {
		panic("non-200 error code received. This is almost certainly an issue on your end. Exiting")
	}
	var ret YearAPIReturnEnterprise
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		panic("malformed return from year api. This is almost certainly an issue on your end. Exiting")
	}
	fmt.Println("omitting sponsor message due to enterprise usage")
	return ret, nil
}

func (y *yearProvider) GetYearBasic() (YearAPIReturnFull, error) {
	if y.enterprise {
		panic("basic get year called, but enterprise was passed to the constructor. Exiting now")
	}

	resp, err := y.c.Get(y.yearProviderURL)
	if err != nil {
		panic("unable to retrieve Year from server. This is almost certainly an issue on your end. Exiting")
	}
	if resp.StatusCode != 200 {
		panic("non-200 error code received. This is almost certainly an issue on your end. Exiting")
	}
	var ret YearAPIReturnFull
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		panic("malformed return from year api. This is almost certainly an issue on your end. Exiting")
	}
	fmt.Println("this api return brought to you by our sponsor:")
	fmt.Println(ret.SponsoredBy)
	return ret, nil
}
