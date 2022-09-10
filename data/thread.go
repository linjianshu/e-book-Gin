package data

import (
	"fmt"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// CreatedAtDate 页面加载的时候 页面里如果模版 会自动加载 结构体的方法
// format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// NumReplies 页面加载的时候 页面里如果模版 会自动加载 结构体的方法
// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
	var count64 int64
	Db.Model(&Post{}).Where("thread_id = ?", thread.Id).Count(&count64)
	count = int(count64)
	return
}

// Posts get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	Db.Model(&Post{}).Select("id", "uuid", "body", "user_id", "thread_id", "created_at").Where("thread_id = ?", thread.Id).Find(&posts)
	return
}

// CreateThread Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	thread := Thread{
		Uuid:      CreateUUID(),
		Topic:     topic,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}

	Db.Create(&thread)

	Db.Model(&Thread{}).Where("user_id = ?", user.Id).First(&conv)

	if conv.Uuid != thread.Uuid {
		err = fmt.Errorf("create topic %s\n", topic)
		return
	}
	return

	//TODO:和原版有点不同
}

// CreatePost Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	newPost := Post{
		Uuid:      CreateUUID(),
		Body:      body,
		UserId:    user.Id,
		ThreadId:  conv.Id,
		CreatedAt: time.Now(),
	}
	Db.Create(&newPost)

	Db.Model(&Post{}).Where("user_id = ? and body = ? and thread_id = ?", user.Id, body, conv.Id).Find(&post)
	if post.Uuid != newPost.Uuid {
		err = fmt.Errorf("create post %s to thread %s error.\n", body, conv.Topic)
	}
	return
}

// Threads Get all threads in the database and returns it
func Threads() (threads []Thread, err error) {
	Db.Model(&Thread{}).Select("id", "uuid", "topic", "user_id", "created_at").Order("created_at desc").Find(&threads)
	return
}

// ThreadByUUID Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	Db.Model(&Thread{}).Select("id", "uuid", "topic", "user_id", "created_at").Where("uuid = ?", uuid).Scan(&conv)
	return
}

// User 页面加载的时候 页面里如果模版 会自动加载 结构体的方法
// Get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	Db.Model(&User{}).Where("id = ?", thread.UserId).First(&user)
	return
}

// User Get the user who wrote the post
func (post *Post) User() (user User) {
	user = User{}
	Db.Model(&User{}).Where("id = ?", post.UserId).First(&user)
	return
}
