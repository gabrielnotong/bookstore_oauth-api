package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
)

func NewSession() *gocql.Session {
	// Connect to cassandra cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	fmt.Println("Cassandra connection successfully created.")

	return session
}
