package schema

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/mszsgo/hgraph"
)

// Main启动
func ListenServe() {
	fmt.Println("******************************************************************************************")
	defer PrintStack()
	args()
	app.ListenAndServe()
}

func args() {
	var (
		// 服务名与端口号
		name string
		port int
	)
	flag.StringVar(&name, "name", app.Name, fmt.Sprintf("Set Application name. Default '%s'", app.Name))
	flag.IntVar(&port, "port", app.Port, fmt.Sprintf("Set Port. Default is %d", app.Port))

	flag.Parse()

	app.Name = name
	app.Port = port
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))

	if err := recover(); err != nil {
		log.Fatalf("** Main Fatalf-> %s", err)
	}
}

type Application struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

var app = &Application{Name: "member", Version: "0.0.1", Host: "0.0.0.0", Port: 80}

func (app *Application) ListenAndServe() {
	app.Handles()
	log.Printf("MicroService: %s  ListenAndServe %s:%d   Start server http://127.0.0.1:%d", app.Name, app.Host, app.Port, app.Port)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", app.Host, app.Port), nil))
}

func (app *Application) Handles() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		w.Write([]byte("MicroService:" + app.Name))
	})

	h := hgraph.GraphqlHttpHandler(Query{}, Mutation{})
	// Graphql服务
	http.Handle("/graphql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "token", r.Header.Get("MS-token"))
		// token、用户id、用户名等信息存储header，用于记录操作日志
		h.ContextHandler(ctx, w, r)
	}))

}
