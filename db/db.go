package db

import (
	"github.com/NavExplorer/navexplorer-api-go/config"
	"github.com/globalsign/mgo"
)

type DBConnection struct {
	session *mgo.Session
}

var server = config.Get().Database.Host
var dbName = config.Get().Database.Name


func NewConnection() (conn *DBConnection) {
	session, err := mgo.Dial(server)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return &DBConnection{session}
}

func (conn *DBConnection) DB() (database *mgo.Database){
	return conn.session.DB(dbName)
}

func (conn *DBConnection) Use(tableName string) (collection *mgo.Collection) {
	return conn.session.DB(dbName).C(tableName)

}

func (conn *DBConnection) Close() {
	conn.session.Close()
	return
}
