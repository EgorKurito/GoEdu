package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	_ "github.com/mattn/go-sqlite3"

)

var tutorials []Tutorial

type Tutorial struct {
	ID			int
	Title		string
	Author		Author
	Comments	[]Comment
}

type Author struct {
	Name		string
	Tutorials	[]int
}

type Comment struct {
	Body string
}

func populate() []Tutorial  {
	author := &Author{Name: "Egor Kurito", Tutorials: []int{1}}
	tutorial := Tutorial{
		ID:		1,
		Title: "Go GraphQL Tutorial",
		Author: *author,
		Comments: []Comment{
			Comment{Body: "First Comment"},
		},
	}
	
	var tutorials []Tutorial
	tutorials = append(tutorials, tutorial)
	return tutorials
}

var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Tutorials": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	})

var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: authorType,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
			},
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{

	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type: 			tutorialType,
			Description: 	"Create a new Tutorial",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				tutorial := Tutorial{
					Title: p.Args["title"].(string),
				}
				tutorials = append(tutorials, tutorial)
				return tutorials, nil
			},
		},
	},
})

func main()  {
	//tutorials = populate()
	// Schema
	fields := graphql.Fields{
		"tutorial": &graphql.Field{
			Type: tutorialType,
			Description: "Get Tutorial By ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["id"].(int)
				if ok {
					db, err := sql.Open("sqlite3", "tutorials.db")
					if err != nil {
						log.Fatal(err)
					}
					defer db.Close()
					var tutorial Tutorial
					err = db.QueryRow("SELECT ID, Title FROM tutorials where ID = ?", id).Scan(&tutorial.ID, &tutorial.Title)
					if err != nil {
						log.Fatal(err)
					}
					return tutorial, nil
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type: graphql.NewList(tutorialType),
			Description: "Get Tutorial List",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db, err := sql.Open("sqlite3", "tutorials.db")
				if err != nil {
					log.Fatal(err)
				}
				defer db.Close()
				var tutorials []Tutorial
				results, err := db.Query("SELECT * FROM tutorials")
				if err != nil {
					log.Fatal(err)
				}
				for results.Next() {
					var tutorial Tutorial
					err = results.Scan(&tutorial.ID, &tutorial.Title)
					if err != nil {
						log.Fatal(err)
					}
					log.Println(tutorial)
					tutorials = append(tutorials, tutorial)
				}
				return tutorials, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
		Mutation: mutationType,
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	// Query
	query := `
    {
        tutorial(id: 1) {
            id
            title
        }
    }
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}

