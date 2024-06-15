package models

import "time"

type Request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"customShort"`
	Expiry      time.Duration `json:"expiry"`
}
type Response struct {
	RateReset    time.Duration `json:"rateReset"`
	URL          string        `json:"url"`
	CustomShort  string        `json:"customShort"`
	Expiry       time.Duration `json:"expiry"`
	RateRemainig int           `json:"rateRemainig"`
}
