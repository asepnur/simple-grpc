package webserver

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	pb "github.com/asepnur/simple-grpc/user-services/grpc"
	usr "github.com/asepnur/simple-grpc/user-services/grpc-function"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Config struct
type Config struct {
	Port     string
	GrpcPort string
}

type requestLogger struct {
	Handle http.Handler
	Logger *log.Logger
}

// Start function
func Start(cfg Config) {
	log.Println("Initializing web server")
	l := log.New(os.Stdout, "[iskandar] ", 0)

	port := ":" + cfg.Port
	if len(cfg.Port) < 1 {
		port = ":" + os.Getenv("PORT")
	}

	grpcPort := ":" + cfg.GrpcPort
	if len(cfg.GrpcPort) < 1 {
		grpcPort = ":" + os.Getenv("PORT")
	}
	r := httprouter.New()
	loadRouter(r)
	go func() {
		InitGRPC(grpcPort)
	}()

	http.ListenAndServe(port, requestLogger{Handle: r, Logger: l})

}
func (rl requestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// === for development only ===
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	// === end for development only ===
	rl.Handle.ServeHTTP(w, r)
	log.Printf("[simple-grpc] %s %s in %v", r.Method, r.URL.Path, time.Since(start))
}

// InitGRPC func
func InitGRPC(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &usr.UserServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
