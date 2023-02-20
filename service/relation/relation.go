package relation

import (
	"tiktok-go/repository"
)

func Follow(relation *repository.Relation) error {
	repository.NewUserDaoInstance().UpdateFollowCount(relation.FromId, 1)
	repository.NewUserDaoInstance().UpdateFollowerCount(relation.ToId, 1)
	return repository.NewRelationDaoInstance().CreateRelation(relation)
}

func GetFollowListById(userid int64) (*[]repository.Relation, error) {
	return repository.NewRelationDaoInstance().GetFollowList(userid)
}

func GetFollowerListById(userid int64) (*[]repository.Relation, error) {
	return repository.NewRelationDaoInstance().GetFollowerList(userid)
}

func UnFollow(from_id int64, to_id int64) error {
	repository.NewUserDaoInstance().UpdateFollowCount(from_id, -1)
	repository.NewUserDaoInstance().UpdateFollowerCount(to_id, -1)
	return repository.NewRelationDaoInstance().DeleteRelation(from_id, to_id)
}

func GetFriendListById(userid int64) (*[]repository.Relation, error) {
	return repository.NewRelationDaoInstance().GetFriendList(userid)
}
