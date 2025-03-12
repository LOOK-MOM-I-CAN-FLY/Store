package main

import (
	"html/template"
	"log"
	"net/http"
)

// Product представляет товар в магазине.
type Product struct {
	Name     string
	ImageURL string
}

func main() {
	// Пример данных товаров. В будущем можно заменить на данные из базы.
	products := []Product{
		{Name: "Товар 1", ImageURL: "/static/images/product-placeholder.png"},
		{Name: "Товар 2", ImageURL: "/static/images/product-placeholder.png"},
		{Name: "Товар 3", ImageURL: "/static/images/product-placeholder.png"},
	}

	// Парсинг HTML шаблона.
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Обработка статических файлов из папки static.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Обработчик главной страницы.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.Execute(w, products); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
