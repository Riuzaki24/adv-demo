package main

import (
	"context"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/link"
	"go/adv-demo/internal/user"
	"go/adv-demo/pkg/db"
	"go/adv-demo/pkg/middleware"
	"net/http"
	"time"
)


func main() {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2 * time.Second)
	defer cancel()

	done := make (chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <- done:
		fmt.Println("done task")
	case <- ctxWithTimeout.Done():
		fmt.Println("Timeout")
	}
}

func main2() {
	conf := configs.LoadConfig()   
	db := db.NewDb(conf)
	router := http.NewServeMux()


	// repositories

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)


	// Services

	authService := auth.NewAuthService(userRepository)

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf, 
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	//middlewares
	stack := middleware.Chain (
		middleware.CORS, 
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
