package main

import (
	"fmt"
	"log"
	"net/http"
    "encoding/json"
    "strconv"
    "strings"

    "alina.net/pkg/store/postgres"
    "alina.net/services/contact/internal"
    "alina.net/services/contact/internal/domain"

)

func main() {
	// Параметры подключения к БД
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "pa55word"
	dbname := "example"

	// Подключение к БД PostgreSQL
	db, err := postgres.ConnectDB(host, port, user, password, dbname)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	service := internal.NewService()

	// Инициализация и настройка обработчиков
	

    http.HandleFunc("/contacts", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            // Реализация создания контакта через use case
            contact := &domain.Contact{}
            err := json.NewDecoder(r.Body).Decode(contact)
            if err != nil {
                http.Error(w, "Invalid request body", http.StatusBadRequest)
                return
            }
            err = service.ContactUC.CreateContact(contact)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            fmt.Fprintf(w, "Contact Created")
            return
        }

        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    })

    http.HandleFunc("/contacts/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            // Извлечение contactID из URL
            contactID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/contacts/"))
            if err != nil {
                http.Error(w, "Invalid contact ID", http.StatusBadRequest)
                return
            }
            contact, err := service.ContactUC.ReadContact(contactID)
            if err != nil {
                http.Error(w, err.Error(), http.StatusNotFound)
                return
            }
            // Отправить контакт в формате JSON в ответе
            json.NewEncoder(w).Encode(contact)
            return
        }

        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    })

    http.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            // Реализация создания группы через use case
            group := &domain.Group{}
            err := json.NewDecoder(r.Body).Decode(group)
            if err != nil {
                http.Error(w, "Invalid request body", http.StatusBadRequest)
                return
            }
            err = service.GroupUC.CreateGroup(group)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            fmt.Fprintf(w, "Group Created")
            return
        }

        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    })

    http.HandleFunc("/groups/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            // Извлечение groupID из URL
            groupID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/groups/"))
            if err != nil {
                http.Error(w, "Invalid group ID", http.StatusBadRequest)
                return
            }
            group, err := service.GroupUC.ReadGroup(groupID)
            if err != nil {
                http.Error(w, err.Error(), http.StatusNotFound)
                return
            }
            // Отправить группу в формате JSON в ответе
            json.NewEncoder(w).Encode(group)
            return
        }

        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    })

    http.HandleFunc("/groups/add-contact", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            // Извлечение данных из запроса
            var requestData struct {
                ContactID int `json:"contactID"`
                GroupID   int `json:"groupID"`
            }
            err := json.NewDecoder(r.Body).Decode(&requestData)
            if err != nil {
                http.Error(w, "Invalid request body", http.StatusBadRequest)
                return
            }

            // Реализация добавления контакта в группу через use case
            err = service.ContactUC.AddContactToGroup(requestData.ContactID, requestData.GroupID)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            fmt.Fprintf(w, "Contact Added To Group")
            return
        }

        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    })

	// Запуск сервера
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
