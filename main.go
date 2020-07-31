package main

import (
	"github.com/gotripe/db"
	"github.com/gotripe/handler"
	"github.com/gotripe/repo"
	"github.com/gotripe/router"
)

func main() {
	r := router.New()

	v1 := r.Group("/api")

	d, err := db.DialDefaultMongoDB()
	if err != nil {
		r.Logger.Fatal(err)
	}

	u := repo.NewUserRepo(d)
	m := repo.NewMovieRepo(d)
	s := repo.NewSubRepo(d)
	p := repo.NewPlanRepo(d)

	h := handler.NewHandler(u, m, s, p)

	h.Register(v1)

	r.Logger.Fatal(r.Start(":8080"))
}
