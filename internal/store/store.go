package store

import "github.com/kevintamakuwala/retail-pulse/internal/models"

// StoreMaster is the in-memory store database
var StoreMaster = map[string]models.Store{
	"S00339218": {StoreID: "S00339218", StoreName: "Store 1", AreaCode: "A1"},
	"S01408764": {StoreID: "S01408764", StoreName: "Store 2", AreaCode: "A2"},
}

// Exists checks if a store exists in the master database
func Exists(storeID string) bool {
	_, exists := StoreMaster[storeID]
	return exists
}
