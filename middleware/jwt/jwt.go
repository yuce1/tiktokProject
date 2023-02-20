package jwt

import (
	"log"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	synID    int64 = 0
	key      []byte
	keyOnce  sync.Once
	instance JWT
)

type Claim struct {
	Session int64 `json:"session"`
	User    int64 `json:"user"`
	jwt.RegisteredClaims
}

type JWT struct{}

func NewInstance() *JWT {
	keyOnce.Do(func() {
		key = make([]byte, 4)
		rand.Read(key)
		instance = JWT{}
	})
	return &instance
}

// flash the key will make token invaild which created before
func ReflashKey() {
	rand.Read(key)
}

// the session is a vaild session mark, it will be in the redis. (but now it just a number)
func (j *JWT) GenToken(userid int64, exp int64) (string, error) {
	atomic.AddInt64(&synID, 1)
	t := time.Now()
	claim := Claim{
		Session: synID,
		User:    userid,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(t),
			ExpiresAt: jwt.NewNumericDate(t.Add(time.Duration(exp) * time.Second)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(key)
}

func (j *JWT) ParseToken(tk string) (*jwt.Token, error) {
	result, err := jwt.ParseWithClaims(tk, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	return result, err
}

// Verify will veify the JWT token, if the token is vaild, we think this userid id correct.
// It will pass the  userid to the next func
func Verify(c *gin.Context) {

	// should we create two func, one for Query and one for PostForm?
	var tk string
	if tk = c.PostForm("token"); tk == "" {
		if tk = c.Query("token"); tk == "" {
			c.Set("Visitor", true)
			return // no token provide, this mean this is a visitor
		}
	}
	result, _ := NewInstance().ParseToken(tk)

	if !result.Valid {
		c.AbortWithStatus(http.StatusForbidden) // invaild token

		log.Printf("[WARN] Invaild JWT : %s", tk)

		return
	}

	c.Set("Visitor", false)
	c.Set("UserID", result.Claims.(*Claim).User)

}
