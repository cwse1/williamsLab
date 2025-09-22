package migrations

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"williamsLab/models"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
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

		var authors string
		var doi string
		var date time.Time
		var issue string

		if p.Article.JournalData.JournalIssue.PubDate.Month != "" {
			if p.Article.JournalData.JournalIssue.PubDate.Day != "" {
				date, _ = time.Parse("2006-Jan-02", strings.Join([]string{p.Article.JournalData.JournalIssue.PubDate.Year, p.Article.JournalData.JournalIssue.PubDate.Month, p.Article.JournalData.JournalIssue.PubDate.Day}, "-"))
			} else {
				date, _ = time.Parse("2006-Jan", strings.Join([]string{p.Article.JournalData.JournalIssue.PubDate.Year, p.Article.JournalData.JournalIssue.PubDate.Month}, "-"))
			}
		} else {
			date, _ = time.Parse("2006", p.Article.JournalData.JournalIssue.PubDate.Year)
		}

		inputDate, _ := types.ParseDateTime(date)
		for _, author := range p.Article.Authors {
			if author.Suffix != "" {
				authors += author.LastName + ", " + author.Initials + ". " + author.Suffix + "; "
			} else {
				authors += author.LastName + ", " + author.Initials + "; "
			}
		}

		for _, id := range p.ArticleIds {
			if id.Type == "doi" {
				doi = id.Value
			}
		}

		if p.Article.JournalData.JournalIssue.Issue != "" {
			issue = p.Article.JournalData.JournalIssue.Volume + ":" + p.Article.JournalData.JournalIssue.Issue
		} else {
			issue = p.Article.JournalData.JournalIssue.Volume
		}

		rec.Set("pmid", p.PMID)
		rec.Set("doi", doi)
		rec.Set("authors", authors)
		rec.Set("title", p.Article.Title)
		rec.Set("abstract", p.Article.Abstract)
		rec.Set("journal", p.Article.JournalData.JournalTitle)
		rec.Set("issue", issue)
		rec.Set("date", inputDate)
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
				Name: "authors",
			},
			&core.TextField{
				Name: "title",
			},
			&core.TextField{
				Name: "abstract",
			},
			&core.TextField{
				Name: "journal",
			},
			&core.TextField{
				Name: "issue",
			},
			&core.DateField{
				Name: "date",
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
