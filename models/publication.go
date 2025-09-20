package models

import (
	"encoding/xml"
)

type Publication struct {
	id       string
	PMID     string `db:"pmid"`
	Authors  string `db:"authors"`
	Title    string `db:"title"`
	Journal  string `db:"journal"`
	Date     string `db:"date"`
	Issue    string `db:"issue"`
	DOI      string `db:"doi"`
	Abstract string `db:"abstract"`
}

type Date struct {
	Year  string
	Month string
	Day   string
}

type SearchQuery struct {
	XMLName xml.Name `xml:"eSearchResult"`
	IdList  struct {
		Ids []string `xml:"Id"`
	} `xml:"IdList"`
}

type ArticleSet struct {
	XMLName       xml.Name `xml:"PubmedArticleSet"`
	PubmedArticle []struct {
		PMID    string `xml:"MedlineCitation>PMID"`
		Article struct {
			JournalData struct {
				JournalIssue struct {
					Volume  int32 `xml:"Volume"`
					Issue   int32 `xml:"Issue"`
					PubDate struct {
						Year  string `xml:"Year"`
						Month string `xml:"Month"`
						Day   string `xml:"Day,omitempty"`
					} `xml:"PubDate"`
				} `xml:"JournalIssue"`
				JournalTitle string `xml:"ISOAbbreviation"`
			} `xml:"Journal"`
			Title    string `xml:"ArticleTitle"`
			Abstract string `xml:"Abstract>AbstractText"`
			Authors  []struct {
				LastName  string `xml:"LastName"`
				FirstName string `xml:"FirstName"`
				Initials  string `xml:"Initials"`
				Suffix string `xml:"Suffix"`
			} `xml:"AuthorList>Author"`
		} `xml:"MedlineCitation>Article"`
		ArticleIds []struct {
			Type  string `xml:"IdType,attr"`
			Value string `xml:",chardata"`
		} `xml:"PubmedData>ArticleIdList>ArticleId"`
	} `xml:"PubmedArticle"`
}
