//Package models Estrutura do modelo de dados para receitas
package models

type Receita struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	CreatedBy   uint   `json:"created_by"`
}

