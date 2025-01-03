package main

type YearAPIReturnFull struct {
	YearAPIReturnEnterprise
	SponsoredBy string `json:"sponsored_by"`
}

type YearAPIReturnEnterprise struct {
	Year       int    `json:"year"`
	YearString string `json:"year_string"`
}
