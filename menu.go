package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	EaterID    int
	FoodMenuID int
}

type MenuItem struct {
	FoodMenuID int
	Count      int
}

// getTopMenuItems get the top n items are consumed
func getTopMenuItems(entries []Entry, count int) ([]MenuItem, error) {
	// eaterItems stores map [eaterID] -> foodMenuID
	eaterItems := make(map[int]map[int]bool)
	// menuItems num of times [menuID] -> count consumed
	menuItems := make(map[int]int)

	for _, entry := range entries {
		if _, ok := eaterItems[entry.EaterID]; !ok {
			eaterItems[entry.EaterID] = make(map[int]bool)
		}

		if _, ok := eaterItems[entry.EaterID][entry.FoodMenuID]; ok {
			return nil, fmt.Errorf("error: eater_id %d has foodmenu_id %d more than once", entry.EaterID, entry.FoodMenuID)
		}

		eaterItems[entry.EaterID][entry.FoodMenuID] = true
		menuItems[entry.FoodMenuID]++
	}

	menuItemsList := make([]MenuItem, 0, len(menuItems))
	for k, v := range menuItems {
		menuItemsList = append(menuItemsList, MenuItem{FoodMenuID: k, Count: v})
	}

	top3 := getTopNMenuItems(menuItemsList, count)
	return top3, nil
}

func getTopNMenuItems(items []MenuItem, n int) []MenuItem {
	// Sort the items slice in descending order based on the Count field
	// and use FoodMenuID as a tiebreaker
	sort.SliceStable(items, func(i, j int) bool {
		if items[i].Count == items[j].Count {
			return items[i].FoodMenuID < items[j].FoodMenuID
		}
		return items[i].Count > items[j].Count
	})

	if n > len(items) {
		n = len(items)
	}
	return items[:n]
}
