package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/RLutsuk/ozon-project/graph"
	commentrepinmem "github.com/RLutsuk/ozon-project/internal/comment/infrastructure/inmemoryrep"
	commentrepdb "github.com/RLutsuk/ozon-project/internal/comment/infrastructure/postgresrep"
	commentrep "github.com/RLutsuk/ozon-project/internal/comment/repository"
	commentresolver "github.com/RLutsuk/ozon-project/internal/comment/resolver"
	commentusecase "github.com/RLutsuk/ozon-project/internal/comment/usecase"
	postrepinmem "github.com/RLutsuk/ozon-project/internal/post/infrastructure/inmemoryrep"
	postrepdb "github.com/RLutsuk/ozon-project/internal/post/infrastructure/postgresrep"
	postrep "github.com/RLutsuk/ozon-project/internal/post/repository"
	postresolver "github.com/RLutsuk/ozon-project/internal/post/resolver"
	postusecase "github.com/RLutsuk/ozon-project/internal/post/usecase"
	userloaders "github.com/RLutsuk/ozon-project/internal/user/dataloader"
	userrepinmem "github.com/RLutsuk/ozon-project/internal/user/infrastructure/inmemoryrep"
	userrepdb "github.com/RLutsuk/ozon-project/internal/user/infrastructure/postgresrep"
	userrep "github.com/RLutsuk/ozon-project/internal/user/repository"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {

	storageType := os.Getenv("STORAGE_TYPE")
	// serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	postgresUser := os.Getenv("POSTGRES_USERNAME")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresDB := os.Getenv("POSTGRES_DATABASE")

	if serverPort == "" || postgresUser == "" || postgresPassword == "" || postgresHost == "" || postgresPort == "" || postgresDB == "" {
		postgresHost = "localhost"
		postgresUser = "db_pg"
		postgresPassword = "db_postgres"
		postgresDB = "db_post"
		postgresPort = "5432"
		serverPort = "8080"
		storageType = "postgres"
	}

	var postRep postrep.RepositoryI
	var commentRep commentrep.RepositoryI
	var userRep userrep.RepositoryI
	switch storageType {
	case "postgres":
		config := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s", postgresHost, postgresUser, postgresPassword, postgresDB, postgresPort)
		db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		}
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		postRep = postrepdb.New(db)
		commentRep = commentrepdb.New(db)
		userRep = userrepdb.New(db)
		logrus.Info("Using PostgreSQL DB")
	default:
		postRep = postrepinmem.New()
		commentRep = commentrepinmem.New()
		userRep = userrepinmem.New()
		logrus.Info("Using In-Memory storage")
	}

	postUC := postusecase.New(postRep, userRep, commentRep)
	commentUC := commentusecase.New(commentRep, userRep, postRep)
	postResolver := postresolver.New(postUC)
	commentResolver := commentresolver.New(commentUC)

	resolver := &graph.Resolver{
		PostResolver:    postResolver,
		CommentResolver: commentResolver,
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Websocket{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// srv = userloaders.Middleware(userRep, srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", userloaders.Middleware(userRep, srv))

	if serverPort == "" {
		serverPort = defaultPort
	}
	logrus.Info("Server has started on host: http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}
