//Package models Estrutura do banco de dados para usuários
package models


type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role		 string `json:"role"` //  "admin" ou "user"
}
