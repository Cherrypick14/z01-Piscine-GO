package piscine

type food struct {
	preptime int
}

var menu = map[string]food{
	"burger":  {15},
	"chips":   {10},
	"nuggets": {12},
}

func FoodDeliveryTime(order string) int {
	keys := make([]string, 0, len(menu))
	for key := range menu {
		keys = append(keys, key)
	}
	for i := 0; i < len(keys); i++ {
		if keys[i] == order {
			return menu[order].preptime
		}
	}
	return 404
}
