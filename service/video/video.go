package video

import "tiktok-go/repository"

func PublishVideo(video *repository.Video) error {
	return repository.NewVideoDaoInstance().CreateVideo(video)
}

func CheckVideo(hash_code string) (*repository.Video, error) {
	return repository.NewVideoDaoInstance().CheckVideoHash(hash_code)
}

func GetVideosByUsername(username string) (*[]repository.Video, error) {
	return repository.NewVideoDaoInstance().GetVideosByAuthor(username)
}

func GetStreamFeed(timeStamp int64) (*[]repository.Video, error) {
	return repository.NewVideoDaoInstance().GetVideoList(timeStamp)
}
