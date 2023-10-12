package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	startServer("localhost:8090")
}

func startServer(addr string) {
	r := setRouters()

	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 3,
		IdleTimeout:  time.Second * 60,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	wait := time.Second * 15
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	fmt.Println("shutting down")
	os.Exit(0)
}

func setRouters() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/test", handlerTest).Name("test-route")

	sub := r.PathPrefix("/").Methods("GET").Schemes("http").Subrouter()
	sub.HandleFunc("/{name:para1}/login", handlerLogin).Name("tag1")
	sub.HandleFunc("/{name}/logout", handlerLogout)
	sub.Use(MyMiddleware)
	sub.Use(MetricMiddleware)

	sub2 := r.PathPrefix("/").Methods("GET").Subrouter()
	sub2.HandleFunc("/temp", handlerTemp)
	sub2.HandleFunc("/users", handlerUsers)

	//r.Use(MyMiddleware)

	//r.HandleFunc("/{name}", handler).Methods("GET").Schemes("http")

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		fmt.Println(route.GetPathTemplate())
		return nil
	})
	return r
}

func handlerTemp(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 2)
	w.Write([]byte("temp"))
}

func handlerUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user1, user2"))
}

func handlerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route name:" + mux.CurrentRoute(r).GetName())

	w.Write([]byte("test"))
}

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 中间件可以根据 router 的 Name() 做一些处理
		fmt.Println("route name:", mux.CurrentRoute(r).GetName())

		fmt.Println("before next")
		next.ServeHTTP(w, r)
		fmt.Println("after next")
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	data       []byte
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK, nil}
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *loggingResponseWriter) Write(data []byte) (int, error) {
	w.data = data
	return w.ResponseWriter.Write(data)
}

func MetricMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		router := mux.CurrentRoute(request)

		fmt.Println(router.GetHostTemplate())
		fmt.Println(router.GetPathTemplate())
		fmt.Println(runtime.FuncForPC(reflect.ValueOf(router.GetHandler()).Pointer()).Name())
		//fmt.Println(router.GetMethods())
		fmt.Println(router.GetQueriesTemplates())
		lrw := NewLoggingResponseWriter(writer)
		next.ServeHTTP(lrw, request)

		fmt.Println("metric time cost:", time.Since(start).String(), request.RequestURI, lrw.statusCode, string(lrw.data))
	})
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route name:" + mux.CurrentRoute(r).GetName())

	fmt.Println("login...")
	vars := mux.Vars(r)
	w.WriteHeader(400)
	w.Write([]byte("hello " + vars["name"]))
}

func handlerLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout...")
	vars := mux.Vars(r)
	w.Write([]byte("bye " + vars["name"]))
}
