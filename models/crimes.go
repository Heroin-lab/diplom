package models

type Crimes struct {
	Lat           string `json:"lat"`
	Lon           string `json:"lon"`
	TitleCriminal string `json:"titleCriminal"`
	CrimeNumber   string `json:"crimeNumber"`
}
