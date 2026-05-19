// (Router) Multiplexer
// DefaultServeMux(внутри хранится map соответствий между URL <-> Handler)
//Способы регистрации путей:
//1) http.Handle(url, handle)
//2) http.HandleFunc(url,handFunc(w http.ResponseWriter, r *http.Request){})
//Правило строгого совпадения с URL: /api/data != /api/data/
//request -> URi /static/styles.css

// 1.URL "/"
// 2 .URL "/static/"

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux() // создание локального мультиплексора
	//подключение файлового сервера к роутеру на корневой путь
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fileServer)


	//используется префиксный путь
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Добро пожаловать в API. Вы запросили: %s", r.URL.Path)
	})

	//строгий путь для точного совпадения
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Страница о проекта"))
	})

	//универсальный обработчик(корень)
	//является префиксом для всего
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	// 	if r.URL.Path != "/"{
	// 		http.NotFound(w,r)
	// 		return
	// 	}

	// 	w.Write([]byte("Главная страница"))
	// })

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Scanln()
}
