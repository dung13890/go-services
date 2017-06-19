package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
)

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type config struct {
	Server      string
	MongoDBHost string
	MongoDBUser string
	MongoDBPwd  string
	Database    string
}

var AppConfig config
var session *mgo.Session

func Init() {
	AppConfig = loadConfig()
	session = createSession()
}

func ResponseError(w http.ResponseWriter, handle error, message string, status int) {
	errRs := errorResponse{
		Error:   handle.Error(),
		Message: message,
		Status:  status,
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if j, err := json.Marshal(errRs); err == nil {
		w.Write(j)
	}
}

func loadConfig() config {
	conf := config{}
	file, err := ioutil.ReadFile("common/config.json")
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	return conf
}

func createSession() *mgo.Session {
	value, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.MongoDBUser,
		Password: AppConfig.MongoDBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createSession]: %s\n", err)
	}
	return value
}

func GetSession() *mgo.Session {
	if session == nil {
		session = createSession()
	}
	return session
}
