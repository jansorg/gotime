package test_setup

import (
	"io/ioutil"
	"os"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/jansorg/tom/go-tom/context"
	"github.com/jansorg/tom/go-tom/i18n"
	"github.com/jansorg/tom/go-tom/query"
	store2 "github.com/jansorg/tom/go-tom/store"
)

func CreateTestContext(lang language.Tag) (*context.GoTimeContext, error) {
	dir, err := ioutil.TempDir("", "gotime")
	if err != nil {
		return nil, err
	}

	store, err := store2.NewStore(dir)
	if err != nil {
		return nil, err
	}

	return &context.GoTimeContext{
		Store:           store,
		StoreHelper:     store2.NewStoreHelper(store),
		Query:           query.NewStoreQuery(store),
		Language:        lang,
		LocalePrinter:   message.NewPrinter(lang),
		Locale:          i18n.FindLocale(lang),
		DurationPrinter: i18n.NewDurationPrinter(lang),
	}, nil
}

func CleanupTestContext(ctx *context.GoTimeContext) {
	dir := ctx.Store.DirPath()
	os.RemoveAll(dir)
}
