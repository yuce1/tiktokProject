package relation

import (
	"tiktok-go/repository"
)

func Follow(relation *repository.Relation) error {
	return repository.NewRelationDaoInstance().CreateRelation(relation)
}

func GetFollowListById(userid int64) (*[]repository.Relation, error) {
	return repository.NewRelationDaoInstance().GetFollowList(userid)
}

func GetFollowerListById(userid int64) (*[]repository.Relation, error) {
	return repository.NewRelationDaoInstance().GetFollowerList(userid)
}

func UnFollow(from_id int64, to_id int64) error {
	return repository.NewRelationDaoInstance().DeleteRelation(from_id, to_id)
}
