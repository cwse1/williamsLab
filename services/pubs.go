package services

import (
	"encoding/xml"
	"io"
	"net/http"
)

type PubService struct {
}

func Migrate() string {
	r, _ := http.Get("https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi?db=pubmed&term=Williams+DC+Jr[author]&retmax=50")
	var list string
	r.Body.Close()
	body, _ := io.ReadAll(r.Body)
	xml.Unmarshal([]byte(body), &list)
	return list
}
