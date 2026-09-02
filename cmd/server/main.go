package main

import (
	"URL_shortner/internal/repository"
	"URL_shortner/internal/service"
	"URL_shortner/internal/transport"
	"URL_shortner/pkg/auth"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//
	// infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// переменное окружение на пароля от бд
	dbPassword := os.Getenv("DB_PASSWORD")

	if dbPassword == "" {
		errorLog.Fatal("DB_PASSWORD не установлен")
	}

	// строка подключения к бд
	dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/url_shortner", dbPassword)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		errorLog.Fatal("ошибка подлкючения к базе данных", err)
	}

	err = db.Ping()
	if err != nil {
		errorLog.Fatal("Ошибка подключения", err)
	}

	defer db.Close()

	auth.InitStore()

	LinkRepo := repository.NewLinksRepo(db)
	LinksService := service.NewLinksService(LinkRepo)
	LinkHandler := transport.NewLinksTransport(LinksService)

	http.HandleFunc("/api/link", LinkHandler.Link)
	http.HandleFunc("/{shortCode}", LinkHandler.GetLinkByShortlink)

	http.HandleFunc("/home", homeHandler)

	fileServer := http.FileServer(http.Dir("./web/static"))

	http.Handle("/static/",
		http.StripPrefix("/static/", fileServer))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web/js"))))

	fmt.Println("сервер запущен на http://127.0.0.1:8020/home")
	http.ListenAndServe(":8020", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/html/index.html")
}
