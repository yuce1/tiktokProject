package favorite

import "tiktok-go/repository"

func Do(f *repository.Favorite, userid int64) error {
	return repository.TransacCreateFavorUpdateFavorCount(f, userid, repository.ACTION_CREATE)
	// return repository.NewFavouriteInstance().Create(f)
}

func Undo(f *repository.Favorite, userid int64) error {
	return repository.TransacDelFavorUpdateFavorCount(f, userid, repository.ACTION_DEL)
	// return repository.NewFavouriteInstance().Delete(f)
}

func ListByUserId(userId int64) (*[]repository.Favorite, error) {
	return repository.NewFavouriteInstance().ListFavouriteByUserId(userId)
}

func ListByVideoId(videoId int64) (*[]repository.Favorite, error) {
	return repository.NewFavouriteInstance().ListFavouriteByVideoId(videoId)
}
