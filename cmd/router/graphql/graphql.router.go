package graphqlrouter

import (
	"context"
	"strings"

	userservice "github.com/MohitVachhani/go-learn/cmd/service/user"
	userInterface "github.com/MohitVachhani/go-learn/pkg/structs/user"
	"github.com/friendsofgo/graphiql"
	"github.com/gorilla/mux"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

type UserResolver struct {
	u *userInterface.User
}

func (r *UserResolver) EmailId() string { return r.u.EmailID }

func (q *query) User(ctx context.Context, args struct{ EmailId string }) *UserResolver {

	user := userservice.GetUser(userInterface.UserFilters{
		EmailID: strings.ToLower(args.EmailId),
	})

	return &UserResolver{&user}
}

func InitializeGraphqlRouter(router *mux.Router) *mux.Router {

	s := `
	type User {
		EmailId: String!
	}

	type Query {
		user(EmailId: String!): User
	}

	schema {
		query: Query
	}
	`
	schema := graphql.MustParseSchema(s, &query{})
	router.Handle("/graphql", &relay.Handler{Schema: schema})

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	router.Handle("/", graphiqlHandler)

	return router
}
