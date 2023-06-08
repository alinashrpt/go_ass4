package postgres

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" // Подключаем драйвер PostgreSQL
)

func ConnectDB(host, port, user, password, dbname string) (*sql.DB, error) {
    connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        return nil, err
    }

    // Проверяем подключение к БД
    err = db.Ping()
    if err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}
