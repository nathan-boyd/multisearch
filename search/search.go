package search

import (
	"bytes"
	"text/template"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/nathan-boyd/multisearch/config"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

type SearchTemplateData struct {
	Query string
}

func GetTemplatesFromCollection(allCollections config.SearchTemplateCollection, selected_collection string) (urls []config.SearchTemplate) {
	From(allCollections.Items).
		WhereT(func(c config.SearchTemplateContainer) bool {
			return c.Name == selected_collection
		}).
		SelectManyT(func(c config.SearchTemplateContainer) Query {
			return From(c.SearchTemplates)
		}).
		ToSlice(&urls)
	return
}

func SearchCollection(cmd *cobra.Command, search_query string, selected_collection string, cfgFile string) (err error) {
	var allCollections config.SearchTemplateCollection
	if err, allCollections = config.GetTemplateCollection(cfgFile); nil != err {
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
