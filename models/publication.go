package models

type Publication struct {
	id       string
	PMID     string `db:"pmid" json:"pmid"`
	PMCID    string
	Authors  string `db:"authors" json:"authors"`
	Title    string `db:"title" json:"title"`
	Journal  string
	Date     string
	Issue    string
	DOI      string
	Abstract string
}
