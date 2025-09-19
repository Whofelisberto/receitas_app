//Package config Estrutura do banco de dados e configuração
package config

import (
	"fmt"
	"log"

	"receitasfitness/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB * gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=admin password=admin dbname=receitasfitness port=3360 sslmode=disable"
	DB , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("falha para se conectar ao banco de dados")
	}
	fmt.Println("Conexão com o banco de dados realizada com sucesso!")

err = DB.AutoMigrate(&models.User{}, &models.Receita{})
if err != nil {
	log.Fatal("falha ao migrar tabelas", err)
}
fmt.Println("Conectado ao PostgreSQL com sucesso!")
}
