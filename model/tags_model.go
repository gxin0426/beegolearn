package model

import "strings"

func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)

	for _, tag := range tags {
		tagList := strings.Split(tag, ",")
		for _, v := range tagList {
			tagsMap[v]++
		}
	}
	return tagsMap
}
