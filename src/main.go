package main

import (
 "com.rule.engine/controllers"
 "github.com/asim/go-micro/v3"
 "github.com/micro/micro/v3/service/logger"
 _ "github.com/jackc/pgx/v4/stdlib"
 "net/http"
 mhttp "github.com/go-micro/plugins/v3/server/http"
 "github.com/gorilla/mux"
 app "com.rule.engine/config/loader"
)


func main() {
	app.Setconfig()
	port :=app.GetVal("GO_MICRO_SERVICE_PORT")
	srv := micro.NewService(
		micro.Server(mhttp.NewServer()),
    )
	opts1 := []micro.Option{
		micro.Name("ruleengine"),
		micro.Version("latest"),
		micro.Address(":"+port),
	}
	srv.Init(opts1...)
	r := mux.NewRouter().StrictSlash(true)
	r.Use(corsMiddleware)
	registerRoutes(r)		
	var handlers http.Handler = r
	

    if err := micro.RegisterHandler(srv.Server(), handlers); err != nil {
		logger.Fatal(err)
	}
	
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}


func registerRoutes(router *mux.Router) {
		registerControllerRoutes(controllers.ManagementController{}, router)
}

func registerControllerRoutes(controller controllers.Controller, router *mux.Router) {
	controller.RegisterRoutes(router)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept,Authorization")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}