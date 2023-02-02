package search

import (
	"bytes"
	"text/template"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/nathan-boyd/multisearch/config"
	"github.com/pkg/browser"
)

type SearchTemplateData struct {
	Query string
}

func GetTemplatesFromCollection(allCollections config.SearchTemplateCollection, selected_collection string) (urls []config.SearchTemplate) {
	linq.From(allCollections.Items).
		WhereT(func(c config.SearchTemplateContainer) bool {
			return c.Name == selected_collection
		}).
		SelectManyT(func(c config.SearchTemplateContainer) linq.Query {
			return linq.From(c.SearchTemplates)
		}).
		ToSlice(&urls)
	return
}

func SearchCollection(search_query string, selected_collection string, cfgFile string) (err error) {
	var allCollections config.SearchTemplateCollection
	if allCollections, err = config.GetTemplateCollection(cfgFile); nil != err {
		return
	}
	searchTemplates := GetTemplatesFromCollection(allCollections, selected_collection)
	for _, template := range searchTemplates {
		if err = OpenBrowser(template, search_query); nil != err {
			return
		}
	}
	return
}

func OpenBrowser(st config.SearchTemplate, searchQuery string) (err error) {
	var t *template.Template
	if t, err = template.New("query").Parse(st.Template); nil != err {
		return
	}
	var templateOutput bytes.Buffer
	if err = t.Execute(&templateOutput, SearchTemplateData{Query: searchQuery}); nil != err {
		return
	}
	return browser.OpenURL(templateOutput.String())
}
