package repository

import "gorm.io/gorm"

// 抽离了多表操作

var (
	ACTION_CREATE int = 1
	ACTION_DEL    int = -1
)

func TransacCreateVideoUpdateWorkCount(video *Video, userid int64, action int) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Model(video).Create(video).Error; e != nil {
			return e
		}
		if e := tx.Model(&User{Id: userid}).UpdateColumn("work_count", gorm.Expr("work_count + ?", action)).Error; e != nil {
			return e
		}
		return nil
	})
	return err
}

func TransacCreateFavorUpdateFavorCount(f *Favorite, userid int64, action int) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Model(f).Save(f).Error; e != nil {
			return e
		}
		if e := tx.Model(&User{Id: userid}).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", action)).Error; e != nil {
			return e
		}
		return nil
	})
	return err
}

func TransacDelFavorUpdateFavorCount(f *Favorite, userid int64, action int) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if e := tx.Model(f).Delete(f).Error; e != nil {
			return e
		}
		if e := tx.Model(&User{Id: userid}).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", action)).Error; e != nil {
			return e
		}
		return nil
	})
	return err
}
