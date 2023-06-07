package jwt

import (
	"crypto"
	"gdl"
	"gdl/assert"
	"gdl/er"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

const UserInfoPkg = "user:info:id:"

type CustomClaims struct {
	Id uint
	*jwt.RegisteredClaims
}

const tokenFinish = "tokenFinish"

type Auth struct {
	Code        uint
	Name        string
	Description string
}
type UserInfo struct {
	Id       uint
	Name     string
	Email    string
	Auths    []Auth
	Ban      bool
	ExpireAt time.Time
}

type Token struct {
	PrivateKey crypto.PrivateKey
	PublicKey  crypto.PublicKey
	Cache      gdl.Cache
}

func (t Token) CreateToken(userInfo UserInfo) (string, error) {
	claims := CustomClaims{
		Id: userInfo.Id,
	}
	claims.RegisteredClaims = &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(userInfo.ExpireAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "openai_web",
		Subject:   "somebody",
		Audience:  []string{"somebody_else"},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	signedString, err := token.SignedString(t.PrivateKey)
	if err != nil {
		return "", err
	}
	return signedString, nil

}

func (t Token) Verify(tokenStr string) (ok bool, err error) {
	var claims CustomClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, t.dealToken)
	return token.Valid, err
}
func (t Token) GetInfo(tokenStr string) (UserInfo, error) {
	var claims CustomClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claims, t.dealToken)
	if err != nil {
		return UserInfo{}, err
	}
	if !token.Valid {
		return UserInfo{}, errors.New("token inValid")
	}
	k, bo := gdl.GetK[UserInfo](t.Cache, UserInfoPkg+strconv.Itoa(int(claims.Id)))
	//开发专用
	//user, err := dao.User.GetById(claims.Id)
	//auths, err := dao.Auth.GetAuthByUserId(user.ID)
	//return *UserInfoForm(user, auths), nil
	//开发专用end
	assert.True(bo, er.ServiceErrType, "身份过期，请重新登录")
	return k, err
}
func (t Token) dealToken(token *jwt.Token) (interface{}, error) {
	//一些token处理
	return t.PublicKey, nil
}

type DataClaims struct {
	Data any
	jwt.RegisteredClaims
}

func (t Token) CreateEmailToken(ti time.Duration, data any) (string, error) {
	claims := DataClaims{
		Data: data,
	}
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ti)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "openai_web",
		Subject:   "somebody",
		Audience:  []string{"somebody_else"},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	signedString, err := token.SignedString(t.PrivateKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}
func (t Token) GetEmailInfo(tokenStr string) (DataClaims, error) {
	var claims DataClaims
	_, err := jwt.ParseWithClaims(tokenStr, &claims, t.dealToken)
	return claims, err
}
