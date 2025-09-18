package migrations

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

type SearchQuery struct {
	XMLName xml.Name `xml:"eSearchResult"`
	IdList  struct {
		Ids []string `xml:"Id"`
	} `xml:"IdList"`
}

type ArticleSet struct {
	XMLName       xml.Name `xml:"PubmedArticleSet"`
	PubmedArticle []struct {
		PMID    string `xml:"PMID"`
		Article struct {
			Journal struct {
				JournalIssue struct {
					PubDate struct {
						Year  string `xml:"Year"`
						Month string `xml:"Month"`
						Day   string `xml:"Day"`
					} `xml:"PubDate"`
				} `xml:"JournalIssue"`
			} `xml:"Journal"`
			Title string `xml:"ArticleTitle"`
		} `xml:"MedlineCitation>Article"`
	} `xml:"PubmedArticle"`
}

func getIds() []string {
	r, err := http.Get("https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi?db=pubmed&term=Williams+DC+Jr[author]&retmax=50")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var search SearchQuery
	if err := xml.Unmarshal([]byte(body), &search); err != nil {
		log.Fatal(err)
	}
	return search.IdList.Ids
}

func getArticles() ArticleSet {
	ids := getIds()
	r, err := http.Get(fmt.Sprintf("https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=pubmed&id=%v&retmode=xml", strings.Join(ids, ",")))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	var list ArticleSet
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := xml.Unmarshal([]byte(body), &list); err != nil {
		log.Fatal(err)
	}
	return list
}

func populatePublications(app core.App, pubs *core.Collection) {
	set := getArticles()
	for _, p := range set.PubmedArticle {
		rec := core.NewRecord(pubs)
		rec.Set("pmid", p.PMID)
		app.Save(rec)
	}
}

func init() {
	m.Register(func(app core.App) error {
		pubs := core.NewBaseCollection("publications")

		pubs.Fields.Add(&core.TextField{Name: "pmid"})

		app.Save(pubs)

		populatePublications(app, pubs)

		return nil
	}, func(app core.App) error {
		pubs, _ := app.FindCollectionByNameOrId("publications")

		return app.Delete(pubs)
	})
}
