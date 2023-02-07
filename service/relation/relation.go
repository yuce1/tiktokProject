package relation

import (
	"tiktok-go/repository"
)

func AddFriend(relation *repository.Relation) error {
	return repository.NewRelationDaoInstance().CreateRelation(relation)
}

func GetRelationListById(userid int64) (*[]repository.Relation, error) {
	return repository.NewRelationDaoInstance().GetRelationList(userid)
}
