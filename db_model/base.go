package db_model

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"gopkg.in/mgo.v2"
	"strings"
)

type MongoClient struct {
	session *mgo.Session
}

var MongoDbClient *MongoClient

var SubjectName2DbName map[string]string

func (s *MongoClient) getSession() *mgo.Session {
	return s.session.Copy()
}

func (s *MongoClient) GetDb(dbname string) *mgo.Database {
	ses := s.getSession()
	return ses.DB(dbname)
}

func InitDbMondel() {
	appConfig := web.AppConfig
	dbhost, _ := appConfig.String("db_host")
	mdc, err := mgo.Dial(dbhost)
	if err == nil {
		MongoDbClient = &MongoClient{mdc}
		logs.Info("connect mongodb success: %s", dbhost)
	} else {
		logs.Error("connect mongodb fail: %s, err: %s", dbhost, err.Error())
	}
	SubjectName2DbName = map[string]string{}
	strSubjects, _ := appConfig.String("subjects")
	allSubjectNames := strings.FieldsFunc(strSubjects, func(x rune) bool { return x == ',' })
	for _, subjectName := range allSubjectNames {
		dbName, _ := appConfig.String(fmt.Sprintf("%s::db_name", subjectName))
		dbName = strings.TrimSpace(dbName)
		SubjectName2DbName[subjectName] = dbName
	}
}
