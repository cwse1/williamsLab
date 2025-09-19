package migrations

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"williamsLab/models"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

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
	var search models.SearchQuery
	if err := xml.Unmarshal([]byte(body), &search); err != nil {
		log.Fatal(err)
	}
	return search.IdList.Ids
}

func getArticles() models.ArticleSet {
	ids := getIds()
	r, err := http.Get(fmt.Sprintf("https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=pubmed&id=%v&retmode=xml", strings.Join(ids, ",")))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	var list models.ArticleSet
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
		var doi string
		for _, id := range p.ArticleIds {
			if id.Type == "doi" {
				doi = id.Value
			}
		}
		rec.Set("doi", doi)
		rec.Set("title", p.Article.Title)
		app.Save(rec)
	}
}

func init() {
	m.Register(func(app core.App) error {
		pubs := core.NewBaseCollection("publications")

		pubs.Fields.Add(
			&core.TextField{
				Name: "pmid",
			},
			&core.TextField{
				Name: "doi",
			},
			&core.TextField{
				Name: "title",
			},
		)

		app.Save(pubs)

		populatePublications(app, pubs)

		return nil
	}, func(app core.App) error {
		pubs, _ := app.FindCollectionByNameOrId("publications")

		return app.Delete(pubs)
	})
}
