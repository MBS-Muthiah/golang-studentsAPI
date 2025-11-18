package main
import "fmt"
import "api-project/internal/config"
import "net/http"
import "log"
import "os"
import "os/signal"
import "syscall"
import "time"
import "context"
import "log/slog"
import "api-project/internal/config/http/handlers/student"
import "api-project/storage/sqlite"


func main(){
	// fmt.Println("Hello, Students API!")
	//load config
	//database setup 
	//http server setup
	cfg :=config.MustLoad()
	storage,err:=sqlite.New(cfg)

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		log.Fatal("Failed to connect to database:", err)
	}
	slog.Info("Connected to database successfully")

	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())

	server := &http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("Starting server on", slog.String("address", cfg.HTTPServer.Address))
	fmt.Println("Starting server on", cfg.HTTPServer.Address)
	done :=make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt,syscall.SIGTERM,syscall.SIGINT)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
		fmt.Println("Failed to start server:", err)
		log.Fatal("Failed to start server:", err)
	}
	}()

	// err := server.ListenAndServe()
	<-done
	fmt.Println("Server stopped")
	slog.Info("Server stopped")

	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err:=server.Shutdown(ctx)
	if err != nil {
		fmt.Println("Failed to shutdown server:", err)
		log.Fatal("Failed to shutdown server:", err)
		slog.Error("Failed to shutdown server:", slog.String("error", err.Error()))
	}
	slog.Info("Server exited properly")
	// fmt.Println("Server started on", cfg.HTTPServer.Address)
}