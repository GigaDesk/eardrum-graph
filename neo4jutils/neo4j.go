package neo4jutils

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Neo4jInstance holds information regarding a Neo4j database connection.
type Neo4jInstance struct {
	Ctx context.Context
	Driver neo4j.DriverWithContext
	Db string
}
// The Init function sets up a connection to a Neo4j database using the provided credentials and stores the necessary information within the `Neo4jInstance` for subsequent interactions with the database
    //
    // This function expects the following environment variables to be set:
    //  - dbUri (string): The URI for the Neo4j database (e.g., "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io")
    //  - dbUser (string): The username for accessing the Neo4j database
    //  - dbPassword (string): The password for accessing the Neo4j database
func (n *Neo4jInstance) Init(dbUri, dbUser, dbPassword string) error {
    ctx := context.Background()
    driver, err := neo4j.NewDriverWithContext(
        dbUri,
        neo4j.BasicAuth(dbUser, dbPassword, ""))

    err = driver.VerifyConnectivity(ctx)
    if err != nil {
        return err
    }
    fmt.Println("Connection to neo4j established.")
	n.Ctx = ctx
	n.Driver = driver
	n.Db = "neo4j"
    return nil
}