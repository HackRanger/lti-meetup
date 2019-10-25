package models

var (
	ItemList map[int]*Item
)

func init() {
	ItemList = make(map[int]*Item)
	i := Item{1,"BasicItem"}
	ItemList[i.Id] = &i
}

type Item struct {
	Id       int
	ItemName string
}