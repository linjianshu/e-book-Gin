package data

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUser_CreateSession(t *testing.T) {
	user := User{
		Id:        13,
		Uuid:      "d6440010-31a1-469e-5a3f-4c40f9d0cb22",
		Name:      "linjianshu",
		Email:     "1018814650@qq.com",
		Password:  "39dfa55283318d31afe5a3ff4a0e3253e2045e43",
		CreatedAt: time.Time{},
	}

	session, err := user.CreateSession()
	equal := assert.Equal(t, nil, err)
	fmt.Println(equal)
	fmt.Println(session.UserId)
}

func TestUser_Session(t *testing.T) {
	user := User{
		Id:        13,
		Uuid:      "d6440010-31a1-469e-5a3f-4c40f9d0cb22",
		Name:      "linjianshu",
		Email:     "1018814650@qq.com",
		Password:  "39dfa55283318d31afe5a3ff4a0e3253e2045e43",
		CreatedAt: time.Time{},
	}
	session, err := user.Session()
	assert.Equal(t, nil, err)
	assert.Equal(t, user.Id, session.UserId)
}

func TestSession_Check(t *testing.T) {
	user := User{
		Id:        13,
		Uuid:      "d6440010-31a1-469e-5a3f-4c40f9d0cb22",
		Name:      "linjianshu",
		Email:     "1018814650@qq.com",
		Password:  "39dfa55283318d31afe5a3ff4a0e3253e2045e43",
		CreatedAt: time.Time{},
	}
	session, _ := user.Session()
	check, err := session.Check()
	assert.Equal(t, nil, err)
	assert.Equal(t, true, check)
}

func TestSession_DeleteByUUID(t *testing.T) {
	user := User{
		Id:        13,
		Uuid:      "d6440010-31a1-469e-5a3f-4c40f9d0cb22",
		Name:      "linjianshu",
		Email:     "1018814650@qq.com",
		Password:  "39dfa55283318d31afe5a3ff4a0e3253e2045e43",
		CreatedAt: time.Time{},
	}
	session, _ := user.Session()
	session.DeleteByUUID()
}

func TestSession_User(t *testing.T) {
	user := User{
		Id:        13,
		Uuid:      "d6440010-31a1-469e-5a3f-4c40f9d0cb22",
		Name:      "linjianshu",
		Email:     "1018814650@qq.com",
		Password:  "39dfa55283318d31afe5a3ff4a0e3253e2045e43",
		CreatedAt: time.Time{},
	}
	session, _ := user.Session()
	u, err := session.User()
	assert.Equal(t, nil, err)
	assert.Equal(t, 13, u.Id)
}

func TestSessionDeleteAll(t *testing.T) {
	err := SessionDeleteAll()
	assert.Equal(t, nil, err)
}

func TestUser_Create(t *testing.T) {
	user := User{
		Name:     "jwt",
		Email:    "1337888217@qq.com",
		Password: "lalala123",
	}
	err := user.Create()
	assert.Equal(t, nil, err)
}

func TestUser_Delete(t *testing.T) {
	user := User{Id: 15}
	err := user.Delete()
	assert.Equal(t, nil, err)
}

func TestUser_Update(t *testing.T) {
	user := User{
		Id:       16,
		Name:     "jwt",
		Email:    "1337888217@qq.com",
		Password: "lalala123",
	}

	user.Name = "JWT"
	err := user.Update()
	assert.Equal(t, nil, err)
}

func TestUserDeleteAll(t *testing.T) {
	err := UserDeleteAll()
	assert.Equal(t, nil, err)
}

func TestUsers(t *testing.T) {
	users, err := Users()
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, len(users))
}

func TestUserByEmail(t *testing.T) {
	user, err := UserByEmail("1018814650@qq.com")
	assert.Equal(t, nil, err)
	assert.Equal(t, 13, user.Id)
}

func TestUserByUUID(t *testing.T) {
	user, err := UserByUUID("d6440010-31a1-469e-5a3f-4c40f9d0cb22")
	assert.Equal(t, nil, err)
	assert.Equal(t, 13, user.Id)
}
