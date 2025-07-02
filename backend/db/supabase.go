package db

import (
	"database/sql"
	"log"
	"time"

	"backend/config"

	_ "github.com/lib/pq" // Driver PostgreSQL
)

var DB *sql.DB // Variável global para a instância do banco de dados

func InitDB(cfg *config.Config) {
	var err error
	DB, err = sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}

	// Configurações opcionais para pool de conexões (melhora performance)
	DB.SetMaxOpenConns(25)                      // Número máximo de conexões abertas
	DB.SetMaxIdleConns(25)                      // Número máximo de conexões ociosas
	DB.SetConnMaxLifetime(5 * 60 * time.Second) // Tempo máximo de vida de uma conexão

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados Supabase: %v", err)
	}

	log.Println("Conectado com sucesso ao banco de dados Supabase PostgreSQL!")
}

func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Erro ao fechar a conexão com o banco de dados: %v", err)
		}
		log.Println("Conexão com o banco de dados fechada.")
	}
}
