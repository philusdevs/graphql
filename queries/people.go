package queries

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"https://github.com/philusdevs/graphql/tree/master/models"
)

// UserField is a field that returns a single user
var UserField = &graphql.Field{
	Type: userObject,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Description: "The name of the user you want to retrieve",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		return fetchPeople(p.Args["name"].(string)), nil
	},
}

// Find a single user for the given userID
func fetchPeople(name string) models.People {

	response, err := http.Get("https://(https://swapi.dev/api/people")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var people models.People

	err = json.Unmarshal(responseData, &people)

	if err != nil {
		log.Fatal(err)
	}

	return people

}
