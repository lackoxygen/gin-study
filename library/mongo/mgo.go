package mongo

import (
	"time"
	"github.com/lackoxygen/gin/library/config"
	"gopkg.in/mgo.v2"
	"fmt"
)

var (
	Session *mgo.Session
)


func InitConnection() {
	dailInfo := &mgo.DialInfo{
		Addrs:     []string{config.MongodbConfig.Host},
		Direct:    false,
		Timeout:   60 * time.Second,
		Database:  config.MongodbConfig.Database,
		Username:  config.MongodbConfig.User,
		Password:  config.MongodbConfig.Password,
		PoolLimit: 4096,
	}
	session, err := mgo.DialWithInfo(dailInfo)
	if err != nil {
		fmt.Printf("mgo dail error[%s]\n", err.Error())
	}
	Session = session
}

func Connection() *mgo.Session{
	s := Session.Copy()
	return s
}

func Exec(database string, collection string, handle func(c *mgo.Collection)) {
	s := Connection()
	defer s.Close()
	c := s.DB(database).C(collection)
	handle(c)
}