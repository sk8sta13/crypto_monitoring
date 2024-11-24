package main

type Rule struct {
    Operator string  `json:"operator"`
    Value    float64 `json:"value"`
}

type Request struct {
	Url string `json:"url"`
	Rules []Rule `json:"rules"`
}

type Config struct {
	Interval int `json:"interval"`
	Requests []Request `json:"requests"`
}