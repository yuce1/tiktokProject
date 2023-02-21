package user

import (
	"strings"
	"tiktok-go/middleware/jwt"
	"tiktok-go/repository"

	"crypto/sha256"
	"encoding/hex"
)

// make sha256hash to protect name and passwd
func GetSHA256HashCode(message []byte) string {
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

func CheckUserExist(username string) bool {
	return repository.NewUserDaoInstance().CheckUserExist(username)
}

func GetUserByName(username string) (*repository.User, bool) {
	return repository.NewUserDaoInstance().GetUserByName(username)
}

func GetUserbyId(userid int64) (*repository.User, bool) {
	return repository.NewUserDaoInstance().GetUserById(userid)
}

func GetUserWithVerify(username string, password string) (*repository.User, bool) {
	return repository.NewUserDaoInstance().VerifyUser(username, password)
}

func GenerateToken(userid int64) string {
	tk, _ := jwt.NewInstance().GenToken(userid, 60*60*24)
	return tk
}

func GetUserByToken(token string) (*repository.User, bool) {
	tokens := strings.Split(token, "_")
	if len(tokens) != 2 {
		return nil, false
	}
	return GetUserWithVerify(tokens[0], tokens[1])
}

func RegisterUser(username string, password string) (int64, string, error) {
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

func (f *RegisterFlow) Do() (int64, string, error) {
	if err := f.CreateUser(); err != nil {
		return 0, "", err
	}
	if err := f.CreateToken(); err != nil {
		return 0, "", err
	}
	return f.user.Id, f.token, nil
}

func (f *RegisterFlow) CreateUser() error {
	return repository.NewUserDaoInstance().CreateUser(f.user)
}

func (f *RegisterFlow) CreateToken() error {
	f.token, _ = jwt.NewInstance().GenToken(f.user.Id, 60*60*24)
	return nil
}
