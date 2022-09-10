package data

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// CreateSession Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	newSession := Session{
		Uuid:      CreateUUID(),
		Email:     user.Email,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	err = Db.Create(&newSession).Error
	if err != nil {
		return
	}

	err = Db.Model(&Session{}).Where("email = ?", user.Email).Find(&session).Error

	return

}

// Session Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}

	err = Db.Model(&Session{}).Where("user_id = ?", user.Id).Find(&session).Error

	return
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	Db.Model(&Session{}).Where("uuid = ?", session.Uuid).Find(&session)
	if session.Id != 0 {
		valid = true
	} else {
		valid = false
	}
	return
}

// DeleteByUUID Delete session from database
func (session *Session) DeleteByUUID() (err error) {
	err = Db.Delete(&Session{}, "uuid = ?", session.Uuid).Error
	return
}

// User Get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}

	err = Db.Model(&User{}).Where("id = ?", session.UserId).Find(&user).Error
	return
}

// SessionDeleteAll Delete all sessions from database
func SessionDeleteAll() (err error) {
	err = Db.Where("1=1").Delete(&Session{}).Error
	return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
	// information from postgres.

	user.Uuid = CreateUUID()
	user.Password = Encrypt(user.Password)
	user.CreatedAt = time.Now()

	err = Db.Create(&user).Error

	return
}

// Delete user from database
func (user *User) Delete() (err error) {

	err = Db.Delete(&User{}, &user).Error

	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	err = Db.Model(&User{}).Where("id = ?", user.Id).Updates(&user).Error
	return
}

// UserDeleteAll Delete all users from database
func UserDeleteAll() (err error) {
	err = Db.Where("1=1").Delete(&User{}).Error
	return
}

// Users Get all users in the database and returns it
func Users() (users []User, err error) {
	err = Db.Model(&User{}).Find(&users).Error
	return
}

// UserByEmail Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.Model(&User{}).Where("email = ?", email).Find(&user).Error
	return
}

// UserByUUID Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.Model(&User{}).Where("uuid = ?", uuid).Find(&user).Error
	return
}
