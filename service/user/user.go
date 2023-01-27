package user

import (
	"strings"
	"tiktok-go/repository"
)

func CheckUserExist(username string) bool {
	return repository.NewUserDaoInstance().CheckUserExist(username)
}

func GetUserByName(username string) (*repository.User, bool) {
	return repository.NewUserDaoInstance().GetUserByName(username)
}

func GetUserWithVerify(username string, password string) (*repository.User, bool) {
	return repository.NewUserDaoInstance().VerifyUser(username, password)
}

func GenerateToken(username string, password string) string {
	return strings.Join([]string{username, password}, "_")
}

func GetUserByToken(token string) (*repository.User, bool) {
	tokens := strings.Split(token, "_")
	if len(tokens) != 2 {
		return nil, false
	}
	return GetUserWithVerify(tokens[0], tokens[1])
}

func RegisterUser(username string, password string) (string, error) {
	user := repository.NewUser()
	user.Name = username
	user.Passwd = password
	return NewRegisterFlow(user).Do()
}

func NewRegisterFlow(user *repository.User) *RegisterFlow {
	return &RegisterFlow{user: user}
}

type RegisterFlow struct {
	user *repository.User

	token string
}

func (f *RegisterFlow) Do() (string, error) {
	if err := f.CreateUser(); err != nil {
		return "", err
	}
	if err := f.CreateToken(); err != nil {
		return "", err
	}
	return f.token, nil
}

func (f *RegisterFlow) CreateUser() error {
	return repository.NewUserDaoInstance().CreateUser(f.user)
}

func (f *RegisterFlow) CreateToken() error {
	f.token = f.user.Name + f.user.Passwd
	return nil
}
