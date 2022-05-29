package cassandra

import (
	"github.com/gocql/gocql"
)

var session *gocql.Session // make a global variable session

func init() {
	// connect to the cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	// cluster := gocql.NewCluster("127.0.0.1")
	// cluster.Keyspace = "oauth"
	// cluster.Consistency = gocql.Quorum
	return session
}
