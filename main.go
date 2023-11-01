package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"scaffoldTest/model"
)

const (
	// Host name of the HTTP Server
	Host = "localhost"
	// Port of the HTTP Server
	Port = "8080"
)

type ResponseMessage struct {
	Message string
}
type Session struct {
	Year     string `json:"year"`
	Semester string `json:"semester"`
}

func ImaalumLogin() *http.Client {

	file, err := os.ReadFile("index.json")
	if err != nil {
		log.Println(err)
	}
	var UserStruct model.User
	json.Unmarshal(file, &UserStruct)
	decodeUser, _ := base64.StdEncoding.DecodeString(UserStruct.Username)
	decodePass, _ := base64.StdEncoding.DecodeString(UserStruct.Password)

	formVal := url.Values{
		"username":    {string(decodeUser)},
		"password":    {string(decodePass)},
		"execution":   {"e1s1"},
		"_eventId":    {"submit"},
		"geolocation": {""},
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		// error handling
	}

	client := &http.Client{
		Jar: jar,
	}

	urlObj, _ := url.Parse("https://imaluum.iium.edu.my/")
	resp_first, _ := client.Get("https://cas.iium.edu.my:8448/cas/login?service=https%3a%2f%2fimaluum.iium.edu.my%2fhome")
	defer resp_first.Body.Close()
	client.Jar.SetCookies(urlObj, resp_first.Cookies())
	cookies1 := resp_first.Cookies()
	resp, _ := client.PostForm("https://cas.iium.edu.my:8448/cas/login?service=https%3a%2f%2fimaluum.iium.edu.my%2fhome?service=https%3a%2f%2fimaluum.iium.edu.my%2fhome", formVal)
	defer resp.Body.Close()
	newCook := append(cookies1, resp.Cookies()...)
	client.Jar.SetCookies(urlObj, newCook)
	return client
}
func SetupCredentialsHandler(c *gin.Context) {

	var user model.User
	c.BindJSON(&user)
	encodeUser := base64.StdEncoding.EncodeToString([]byte(user.Username))
	encodePass := base64.StdEncoding.EncodeToString([]byte(user.Password))
	data := model.User{
		Username: encodeUser,
		Password: encodePass,
	}
	file, _ := json.MarshalIndent(data, "", " ")
	err2 := os.WriteFile("index.json", file, fs.ModePerm)
	if err2 != nil {
		log.Fatal(err2)
		c.JSON(200, ResponseMessage{Message: err2.Error()})

	}
	c.JSON(200, ResponseMessage{Message: "Credentials Saved Succesfully"})
}
func LoginWifiHandler(c *gin.Context) {
	var userJson model.User
	file, err := os.ReadFile("index.json")
	if err != nil {
		log.Fatal(err)

	}

	json.Unmarshal(file, &userJson)
	if err != nil {
		log.Println(userJson.Username)
		log.Println(err.Error())
	}
	decodeUser, _ := base64.StdEncoding.DecodeString(userJson.Username)
	decodePass, _ := base64.StdEncoding.DecodeString(userJson.Password)
	formVal := url.Values{
		"user":     {string(decodeUser)},
		"password": {string(decodePass)},
		"url":      {"http://www.iium.edu.my/"},
		"cmd":      {"authenticate"},
		"Login":    {"Log In"},
	}
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}
	resp, err := client.PostForm("https://captiveportalmahallahgombak.iium.edu.my/cgi-bin/login", formVal)
	if resp.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, ResponseMessage{Message: "You are now connected to IIUM-Student!"})
		return
	}
	c.JSON(http.StatusBadRequest, ResponseMessage{Message: "An Error Occured. It could be that you can logged into the network"})

}
func DownloadScheduleHandler(c *gin.Context) {
	var session Session

	err := c.BindJSON(&session)
	if err != nil {
		log.Println(err.Error())
		return
	}
	client := ImaalumLogin()
	sessionVal, semesterVal := session.Year, session.Semester
	response, _ := client.Get(fmt.Sprintf("https://imaluum.iium.edu.my/confirmationslip?ses=%s&sem=%s", sessionVal, semesterVal))
	if response.StatusCode == 200 {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(err.Error())
		}
		err = os.WriteFile("./files/cs.html", bodyBytes, 0644)
		if err != nil {
			log.Println(err.Error())
			os.Mkdir("files", 0644)
			os.WriteFile("./files/cs.html", bodyBytes, 0644)
		}
		c.JSON(200, struct {
			Message string
		}{
			Message: "Confirmation Slip was Downloaded Successfully",
		})
		return
	}
	c.JSON(200, struct {
		Message string
	}{
		Message: "Error Occured",
	})
	client.Get("https://cas.iium.edu.my:8448/cas/logout?service=http://imaluum.iium.edu.my/")
	return
}
func DownloadExamSlipHandler(context *gin.Context) {
	client := ImaalumLogin()
	response, _ := client.Get("https://imaluum.iium.edu.my/MyAcademic/course_timetable")
	defer response.Body.Close()
	defer client.Get("https://cas.iium.edu.my:8448/cas/logout?service=http://imaluum.iium.edu.my/")
	if response.StatusCode == 200 {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			os.Exit(1)
		}
		_ = os.WriteFile("./files/timetable.pdf", bodyBytes, 0644)
		context.JSON(200, ResponseMessage{Message: "Exam timetable downloaded successfully"})
		return
	}
	context.JSON(http.StatusBadRequest, ResponseMessage{Message: "Something went wrong"})

}
func DownloadFinanceHandler(context *gin.Context) {
	client := ImaalumLogin()
	response, _ := client.Get("https://imaluum.iium.edu.my/MyFinancial")
	defer response.Body.Close()
	defer client.Get("https://cas.iium.edu.my:8448/cas/logout?service=http://imaluum.iium.edu.my/")
	if response.StatusCode == 200 {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			os.Exit(1)
		}
		_ = os.WriteFile("./files/finance.pdf", bodyBytes, 0644)
		context.JSON(200, ResponseMessage{Message: "Financial Statement downloaded successfully"})
		return
	}
	context.JSON(http.StatusBadRequest, ResponseMessage{Message: "Something went wrong"})
}

func DownloadResultHandler(context *gin.Context) {
	var session Session
	context.BindJSON(&session)
	var UserStruct model.User
	file, err := os.ReadFile("index.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(file, &UserStruct)
	decodeUser, _ := base64.StdEncoding.DecodeString(UserStruct.Username)
	decodePass, _ := base64.StdEncoding.DecodeString(UserStruct.Password)
	form_val := url.Values{
		"mat_no":   {string(decodeUser)},
		"pin_no":   {string(decodePass)},
		"sessi":    {session.Year},
		"semester": {session.Semester},
		"login":    {"Login"},
	}
	resp, err := http.PostForm("https://myapps.iium.edu.my/anrapps/viewResult.php", form_val)
	if resp.StatusCode == 200 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
		}

		_ = os.WriteFile("./result.html", bodyBytes, 0644)

	}
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
func main() {
	err := os.Mkdir("files", 0644)
	if err != nil {
		log.Println(err.Error())
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/login", LoginWifiHandler)
	r.POST("/cs", DownloadScheduleHandler)
	r.GET("/es", DownloadExamSlipHandler)
	r.POST("/result", DownloadResultHandler)
	r.GET("/finance", DownloadFinanceHandler)
	r.POST("/setup", SetupCredentialsHandler)
	r.Use(cors.Default())
	r.Use(static.Serve("/", static.LocalFile("./web/dist", false)))
	open("http://localhost:8080")
	r.Run()
}
