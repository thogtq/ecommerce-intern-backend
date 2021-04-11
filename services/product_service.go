package services

import "strings"

func GetParentCategories(categories []string) []string {
	parentCategories := []string{}
	for _, category := range categories {
		parent := strings.Split(category, "/")[0]
		if !contains(parentCategories,parent){
			parentCategories = append(parentCategories, parent)
		}
		
	}
	return parentCategories
}
