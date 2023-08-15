package main

import (
	"bytes"
	"e-book-Gin/data"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

// TODO:6. 现在主流的依赖管理方式是mod，可以参考go的开源代码结构
//ljs :不知道项目目录中出现了go.mod算不算修改好了?
import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
)

type FileInfo struct {
	Id   int
	Type string
	Size int64
	Name string
}

type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

// TODO:5. 这里在内存里记录文件信息，有没有考虑中间文件更新、修改或者删除等，内存里的内容没同步更新呢？
// ljs:好的 删除/更新文件的时候我仔细设计一下
var fileInfos []FileInfo

// TODO:7. 可以参考下 go 的 web 工程目录结构，都在一个文件里会随着代码复杂度上升，可读性会下降
// ljs:把handlerFunc从匿名方法中抽出来可以吗?
func main() {

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	//index主页
	r.GET("/index", index)
	//下载文件
	r.GET("/download", download)
	//打开上传页面
	r.GET("/upload", upload)
	//上传文件
	r.POST("/saveItem", saveItem)

	r.GET("/blog", blog)

	//打开发布话题界面
	r.GET("/thread/new", newThread)
	//创建帖子操作
	r.POST("/thread/create", createThread)

	r.GET("/thread/read", readThread)

	r.POST("/thread/post", postThread)

	r.GET("/login", login)

	r.GET("/signup", signup)

	r.POST("/authenticate", authenticate)

	r.POST("/signup_account", signupAccount)

	r.GET("/translatePage", translagePage)
	r.POST("/translateResult", translateResult)
	r.GET("/littleMouse", littleMouse)
	r.GET("/datashare", datashare)
	r.GET("/Vote", vote)
	r.GET("/uploadPage", uploadPage)
	err := r.Run(":6666")
	if err != nil {
		panic(err)
	}
}

func uploadPage(context *gin.Context) {
	context.Redirect(http.StatusTemporaryRedirect, "http://124.220.207.169:8766/uploadPage")
}

func datashare(context *gin.Context) {
	context.Redirect(http.StatusTemporaryRedirect, "http://124.220.207.169:8081/")
}
func vote(context *gin.Context) {
	context.Redirect(http.StatusTemporaryRedirect, "http://124.220.207.169:9999/Vote")
}

func littleMouse(context *gin.Context) {
	context.Redirect(http.StatusTemporaryRedirect, "http://124.220.207.169:8000/index")
}

func translateResult(context *gin.Context) {
	word := context.PostForm("word")
	//context.
	request := DictRequest{
		TransType: "en2zh",
		Source:    word,
		UserID:    "",
	}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	data := bytes.NewBuffer(buf)
	client := &http.Client{}
	//var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("os-type", "web")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Bad status Code:", resp.StatusCode, "body", string(bodyText))
	}

	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatalln(err)
	}

	context.HTML(http.StatusOK, "caiyunxiaoyi.html", dictResponse)
}

func translagePage(context *gin.Context) {
	context.HTML(http.StatusOK, "caiyunxiaoyi.html", nil)
}

func newThread(context *gin.Context) {
	_, err := session(context)
	if err != nil {
		fmt.Println(err.Error())
		context.Redirect(http.StatusFound, "/login")
		return
	}
	context.HTML(http.StatusFound, "newThread.html", nil)
}

func createThread(context *gin.Context) {
	sess, err := session(context)
	if err != nil {
		context.Redirect(http.StatusFound, "/login")
		return
	}
	user, err := sess.User()
	if err != nil {
		fmt.Println("cannot get user from session", err.Error())
		return
	}
	topic := context.PostForm("topic")
	_, err = user.CreateThread(topic)
	if err != nil {
		fmt.Println("cannot create thread , err: ", err.Error())
		return
	}

	context.Redirect(http.StatusFound, "/blog")

}

func signup(context *gin.Context) {
	context.HTML(http.StatusOK, "signup.html", nil)
}

func signupAccount(context *gin.Context) {
	user := data.User{
		Uuid:     "",
		Name:     context.PostForm("name"),
		Email:    context.PostForm("email"),
		Password: context.PostForm("password"),
	}

	err := user.Create()
	if err != nil {
		fmt.Printf("create user error ,err : %v\n", err.Error())
		return
	}

	context.Redirect(http.StatusFound, "/login")
}

func authenticate(context *gin.Context) {
	email := context.PostForm("email")
	user, err := data.UserByEmail(email)
	if err != nil {
		fmt.Printf("get user by email error, email: %s , err: %v", email, err.Error())
	}

	encrypt := data.Encrypt(context.PostForm("password"))
	if user.Password == encrypt {
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println("cannot create session , err :\n", err.Error())
			return
		}

		context.SetCookie("_cookie", session.Uuid, 0, "", "", false, false)
		context.Redirect(http.StatusFound, "/index")
	} else {
		context.Redirect(http.StatusFound, "/login")
	}
}

func login(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

var mu sync.Mutex

func index(context *gin.Context) {
	var wg sync.WaitGroup
	dir, err := ioutil.ReadDir("public")
	if err != nil {
		wg.Add(1)
		go Log(fmt.Sprintf("readDir public err , detail : %s\n", err.Error()), &wg)
		context.Error(err)
		wg.Wait()
		return
	}

	fileInfos = []FileInfo{}
	var fileInfo FileInfo
	for i, info := range dir {
		fileInfo.Id = i
		fileInfo.Type = info.Mode().Type().String()
		fileInfo.Size = info.Size() / 1024
		fileInfo.Name = info.Name()
		fileInfos = append(fileInfos, fileInfo)
	}

	context.HTML(http.StatusOK, "index.html", fileInfos)
}

func download(context *gin.Context) {
	var wg sync.WaitGroup
	id := context.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		wg.Add(1)
		go Log(fmt.Sprintf("strconv atoi %s error , detail :%s\n", id, err.Error()), &wg)
		context.Error(err)
		wg.Wait()
		return
	}

	// TODO:1. 这里数组如果出界，程序会异常退出
	// ljs:这里我考虑的是atoi是从界面上获取的id应该不会导致数组出界 老师是担心postman这类api测试工具随便给个id这类场景吗?
	if atoi >= len(fileInfos) {
		context.JSON(http.StatusNotImplemented, gin.H{"error": "id越界,找不到指定文件"})
		return
	}

	info := fileInfos[atoi]

	//TODO: 2. 可以考虑换一个日志输出库
	//ljs:老师这里我先使用原生api打到log文件里可以不 后续改进成使用go-log相关的插件或者库

	wg.Add(1)
	go Log(fmt.Sprintf("download fileName : %s .\n", info.Name), &wg)

	join := filepath.Join("public", info.Name)

	wg.Wait()

	file, err := ioutil.ReadFile(join)
	if err != nil {
		wg.Add(1)
		go Log(fmt.Sprintf("readFile %s error , detail : %s\n", join, err.Error()), &wg)
		context.Error(err)
		wg.Wait()
		return
	}

	//TODO: 3. 由于文件大小因素，考虑是否压缩？
	//压缩文件是用go来压缩 发送压缩包还是使用http的一些header指定传输的时候gzip
	context.Header("Content-Encoding", "gzip")
	context.Header("Content-Type", "application/octet-stream")
	context.Header("Content-Disposition", "attachment;filename="+info.Name)
	context.Data(http.StatusOK, "application/octet-stream", file)
}

func upload(context *gin.Context) {
	context.HTML(http.StatusOK, "upload.html", nil)
}

func saveItem(context *gin.Context) {
	var wg sync.WaitGroup

	file, err := context.FormFile("file")
	if err != nil {
		wg.Add(1)
		go Log(fmt.Sprintf("FormFile failed , error : %v\n", err.Error()), &wg)
		context.Error(err)
		wg.Wait()
		return
	}

	dst := path.Join("public", file.Filename)
	err = context.SaveUploadedFile(file, dst)
	if err != nil {
		wg.Add(1)
		go Log(fmt.Sprintf("FormFile failed , error : %v\n", err.Error()), &wg)
		context.Error(err)
		wg.Wait()
		return
	}

	//简言之，302和307状态码都可用于临时重定向。但是对于有些客户端实现，在收到302状态码时，客户端可能会将请求方法强制修改为"GET"；对于307状态码，禁止客户端修改请求方法为"GET"。
	context.Redirect(http.StatusFound, "/index")
}

func Log(info string, wg *sync.WaitGroup) {
	//上锁
	mu.Lock()
	defer mu.Unlock()

	//做完了
	defer wg.Done()
	logFileName := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))
	logFileName = path.Join("log", logFileName)
	openFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND, 0666)
	defer openFile.Close()
	if err != nil {
		fmt.Printf("%s : openFile %s failed , err : %v\n", time.Now().Format("15:04:05"), logFileName, err)
	} else {
		fmt.Fprintf(openFile, "%s : %s", time.Now().Format("15:04:05"), info)
	}
}

func blog(context *gin.Context) {
	threads, err := data.Threads()
	if err != nil {
		fmt.Println(err.Error())
	}
	context.HTML(http.StatusOK, "blog.html", threads)
}

func readThread(context *gin.Context) {
	uuid := context.Query("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		fmt.Println("readThread error , err", err.Error())
	} else {
		context.HTML(http.StatusOK, "thread.html", []data.Thread{thread})
	}
}

func postThread(context *gin.Context) {
	sess, err := session(context)
	if err != nil {
		context.Redirect(http.StatusFound, "/login")
	} else {
		user, err := sess.User()
		if err != nil {
			fmt.Println("cannot get user from session")
			return
		}
		body := context.PostForm("body")
		uuid := context.PostForm("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			fmt.Println("get thread by uuid error ", err.Error())
			return
		}
		_, err = user.CreatePost(thread, body)
		if err != nil {
			fmt.Println("cannot create post")
			return
		}

		url := fmt.Sprintf("/thread/read?id=%s", uuid)
		context.Redirect(http.StatusFound, url)
	}
}
