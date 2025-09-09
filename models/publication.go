package models

type Publication struct {
	PMID    int32
	year    int16
	authors string
	title   string
	journal string
	date    string
}
