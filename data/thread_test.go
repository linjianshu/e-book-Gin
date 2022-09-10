package data

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNumReplies(t *testing.T) {
	thread := Thread{
		Id:        3,
		Uuid:      "52fdfc07-2182-454f-563f-5f0f9a621d72",
		Topic:     "今亡亦死,举大义亦死,等死,死国可乎?",
		UserId:    5,
		CreatedAt: time.Time{},
	}
	replies := thread.NumReplies()
	equal := assert.Equal(t, 2, replies)
	fmt.Println(equal)
}

func TestPosts(t *testing.T) {
	thread := Thread{
		Id:        3,
		Uuid:      "52fdfc07-2182-454f-563f-5f0f9a621d72",
		Topic:     "今亡亦死,举大义亦死,等死,死国可乎?",
		UserId:    5,
		CreatedAt: time.Time{},
	}
	posts, err := thread.Posts()
	equal := assert.Equal(t, nil, err)
	fmt.Println(equal)

	for _, post := range posts {
		fmt.Println(post.Id)
	}
	b := assert.Equal(t, len(posts), 2)
	fmt.Println(b)
}

func TestThreads(t *testing.T) {
	threads, err := Threads()
	equal := assert.Equal(t, nil, err)
	fmt.Println(equal)

	for _, thread := range threads {
		fmt.Println(thread.Topic)
	}

	b := assert.Equal(t, 4, len(threads))
	fmt.Println(b)
}

func TestThreadByUUID(t *testing.T) {
	thread := Thread{
		Id:        3,
		Uuid:      "52fdfc07-2182-454f-563f-5f0f9a621d72",
		Topic:     "今亡亦死,举大义亦死,等死,死国可乎?",
		UserId:    5,
		CreatedAt: time.Time{},
	}

	th, err := ThreadByUUID(thread.Uuid)
	equal := assert.Equal(t, nil, err)
	fmt.Println(equal)

	b := assert.Equal(t, thread.Id, th.Id)
	fmt.Println(b)
}

func TestUser_CreateThread(t *testing.T) {
	user := User{
		Id:        13,
		Uuid:      "d6440010-31a1-469e-5a3f-4c40f9d0cb22",
		Name:      "linjianshu",
		Email:     "1018814650@qq.com",
		Password:  "39dfa55283318d31afe5a3ff4a0e3253e2045e43",
		CreatedAt: time.Time{},
	}
	thread, err := user.CreateThread("人活着是为了什么")
	equal := assert.Equal(t, nil, err)
	fmt.Println(equal)

	b := assert.Equal(t, "人活着是为了什么", thread.Topic)
	fmt.Println(b)

}

func TestUser_CreatePost(t *testing.T) {
	user := User{
		Id:        13,
		Uuid:      "d6440010-31a1-469e-5a3f-4c40f9d0cb22",
		Name:      "linjianshu",
		Email:     "1018814650@qq.com",
		Password:  "39dfa55283318d31afe5a3ff4a0e3253e2045e43",
		CreatedAt: time.Time{},
	}

	thread := Thread{
		Id: 3,
	}
	post, err := user.CreatePost(thread, "嗯嗯好")
	equal := assert.Equal(t, nil, err)
	fmt.Println(equal)

	b := assert.Equal(t, "嗯嗯好", post.Body)
	fmt.Println(b)
}

func TestThread_User(t *testing.T) {
	expected := User{
		Id:   13,
		Uuid: "d6440010-31a1-469e-5a3f-4c40f9d0cb22",
	}
	thread := Thread{UserId: 13}
	user := thread.User()
	equal := assert.Equal(t, expected.Uuid, user.Uuid)
	fmt.Println(equal)
}

func TestPost_User(t *testing.T) {
	expected := User{
		Id:   13,
		Uuid: "d6440010-31a1-469e-5a3f-4c40f9d0cb22",
	}

	post := Post{
		UserId: 13,
	}
	user := post.User()
	equal := assert.Equal(t, expected.Uuid, user.Uuid)
	fmt.Println(equal)
}
