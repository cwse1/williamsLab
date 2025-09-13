package models

type Publication struct {
	id       string
	PMID     string `db:"pmid" json:"pmid"`
	PMCID    string
	Authors  string `db:"authors" json:"authors"`
	Title    string `db:"title" json:"title"`
	Journal  string
	Year     int16
	Date     string
	Vol      string
	DOI      string
	Abstract string
}
