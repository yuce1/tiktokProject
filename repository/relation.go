package repository

import (
	"sync"

	"gorm.io/gorm"
)

type Relation struct {
	Id        int64 `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	FromId    int64 `json:"uid"`
	ToId      int64 `json:"vid"`
	CreatedAt int64 `json:"created_at,omitempty"`
	DeleteAt  int64 `json:"delete_at,omitempty"`
}

type RelationDao struct {
	db *gorm.DB
}

var (
	relationDao  *RelationDao
	relationOnce sync.Once
)

func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(func() {
		relationDao = &RelationDao{
			db: db,
		}
	})
	return relationDao
}

func (t *RelationDao) CreateRelation(relation *Relation) error {

	result := t.db.Create(relation)
	return result.Error
}

func (t *RelationDao) GetFollowList(userid int64) (*[]Relation, error) {
	var relations []Relation
	res := t.db.Where("from_id = ?", userid).Order("created_at DESC").Find(&relations)
	return &relations, res.Error
}

func (t *RelationDao) GetFollowerList(userid int64) (*[]Relation, error) {
	var relations []Relation
	res := t.db.Where("to_id = ? ", userid).Order("created_at DESC").Find(&relations)
	return &relations, res.Error
}

func (t *RelationDao) CheckRelation(from_id int64, to_id int64) bool {
	var relation Relation
	res := t.db.Where("from_id = ? and to_id = ? ", from_id, to_id).Take(&relation)
	return res.RowsAffected == 1
}

func (t *RelationDao) DeleteRelation(from_id int64, to_id int64) error {
	var relation Relation
	result := t.db.Where("from_id = ? and to_id = ? ", from_id, to_id).Take(&relation)
	t.db.Delete(&relation)
	return result.Error
}

func (t *RelationDao) GetFriendList(userid int64) (*[]Relation, error) {
	var relations []Relation
	res := t.db.Where("to_id = ? ", userid).Order("created_at DESC").Find(&relations)
	return &relations, res.Error
}
