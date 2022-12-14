package elements

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/oleiade/reflections"
	"reflect"
)

const tableTagName = "table"

func retrieveColumnNames[V any](m []V) ([]string, []string) {
	uniqueKeys := make(map[string]bool)
	uniqueKeysArray := []string{}
	uniqueTagsArray := []string{}
	for _, k := range m {
		t := reflect.TypeOf(k)
		names := make([]string, t.NumField())
		for i := range names {
			names[i] = t.Field(i).Name
			if ok := uniqueKeys[t.Field(i).Name]; !ok {
				uniqueKeys[t.Field(i).Name] = true
				uniqueKeysArray = append(uniqueKeysArray, t.Field(i).Name)
				if t.Field(i).Tag.Get(tableTagName) != "" {
					uniqueTagsArray = append(uniqueTagsArray, t.Field(i).Tag.Get(tableTagName))
				} else {
					uniqueTagsArray = append(uniqueTagsArray, t.Field(i).Name)
				}
			}
		}
	}
	return uniqueKeysArray, uniqueTagsArray
}
func getField(v any, field string) any {
	value, err := reflections.GetField(v, field)
	if err != nil {
		fmt.Println(err)
	}
	return value
}
//Table converts a map of structs to a table
func Table[V any](m []V, class string) app.HTMLTable {
	// get all the keys
	data, tags := retrieveColumnNames(m)
	//keys := maps.Keys(data)
	//fmt.Printf("m %+v", m)
	//values := maps.Values(data)
	//fmt.Println("values ", values)
	//we are mapping out all the unique column names (keys[i])
	//we then need to make each row from data based on the current column name and the data field
	var rows []app.UI
	for _, v := range m {
		entries := []app.UI{}
		for _, j := range data {
			val := getField(v, j)
			switch val.(type) {
				case app.UI:
					entries = append(entries, app.Td().Style("border", "1px solid black").Body(val.(app.UI)))
				default:
					entries = append(entries, app.Td().Style("border", "1px solid black").Text(val))
			}

		}
		rows = append(rows, app.Tr().Body(entries...))
	}
	tableHeadings := app.THead().Body(
		app.Tr().Body(
			//for all of the values in the map
			app.Range(data).Slice(func(j int) app.UI {
				return app.Th().Class("text-center").Text(tags[j])
			}),
		),
	)
	tableBody := app.TBody().Style("border", "1px solid black").Body(
		rows...,
	)
	return app.Table().Class("table").Class(class).Body(
		tableHeadings,
		tableBody,
	)
}

