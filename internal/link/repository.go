package link

import "go/adv-demo/pkg/db"



type LinkRepository struct {
	DataBase *db.Db
}


func NewLinkRepository(database *db.Db) *LinkRepository{
	return &LinkRepository {
		DataBase: database,
	}
}


func (repo *LinkRepository) CreateLink(link *Link) {
	
}
