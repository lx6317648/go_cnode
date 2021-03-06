package database

import (
	//"database/sql"
	"fmt"
	"log"
	"github.com/garyburd/redigo/redis"
	//_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
	//"log"
)

type MongoLog struct {
}

func (MongoLog)Output(calldepth int, s string) error {
	log.SetFlags(log.Lshortfile)
	return log.Output(calldepth,s)
}

//var SqlDB *sql.DB
var MogSession *mgo.Session
var Redis redis.Conn
var Mgodb *mgo.Database
func init() {
}
//var mgodb *mgo.Database
func Config(mogo_url string,redis_url string) {
	var err error
	var mgoerr error
	// SqlDB, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test?parseTime=true")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// err = SqlDB.Ping()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	MogSession, mgoerr = mgo.Dial(mogo_url)
	if mgoerr != nil {
		panic(mgoerr)
	}
	MogSession.SetMode(mgo.Monotonic, true)
	Mgodb=MogSession.DB("egg_cnode")
	Redis, err = redis.Dial("tcp", redis_url)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	
	//  defer mogSession.Close()
	//  session.SetMode(mgo.Monotonic, true)
	//  mgodb = session.DB("egg_cnode")
	//  countNum, _ :=mgodb.C("users").Count()
	//  log.Println(countNum)
	mgo.SetDebug(false)  // 设置DEBUG模式
	//mgo.SetLogger(new(MongoLog)) // 设置日志. 	
}
