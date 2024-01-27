package stores

import (
	"github.com/onlineeric/eric-gin-server/models"
)

var ItemDb = make(map[int]*models.Item)

func init() {
	ItemDb[1] = &models.Item{
		ID:          1,
		Name:        "item1",
		Description: "item1 description",
		Price:       1.11,
	}
	ItemDb[2] = &models.Item{
		ID:          2,
		Name:        "item2",
		Description: "item2 description",
		Price:       2.222,
	}
	ItemDb[3] = &models.Item{
		ID:          3,
		Name:        "item3",
		Description: "item3 description",
		Price:       3.3333,
	}
}
