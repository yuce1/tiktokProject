package video

import (
	"tiktok-go/repository"
)

func PublishVideo(video *repository.Video, userid int64) error {
	return repository.TransacCreateVideoUpdateWorkCount(video, userid, repository.ACTION_CREATE)
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

func GetVideoBySet(videoIdSet []int64) (*[]repository.Video, error) {
	return repository.NewVideoDaoInstance().ListByVideoIdSet(videoIdSet)
}
