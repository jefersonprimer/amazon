package main

import (
	"backend/config"
	"backend/db"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Carrega as configurações do ambiente
	cfg := config.LoadConfig()

	// Inicializa a conexão com o banco de dados Supabase
	db.InitDB(cfg)
	defer db.CloseDB() // Garante que a conexão será fechada ao encerrar o app

	// Inicializa o Gin router
	router := gin.Default()

	// Configura todas as rotas da sua API
	routes.SetupRoutes(router)

	// Inicia o servidor
	log.Printf("Servidor iniciado na porta %s...", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}
