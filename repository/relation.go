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

func (t *RelationDao) GetRelationList(userid int64) (*[]Relation, error) {
	var relations []Relation
	res := t.db.Where("from_id = ?", userid).Order("created_at DESC").Find(&relations)
	return &relations, res.Error
}
