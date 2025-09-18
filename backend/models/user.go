//Package models Estrutura do banco de dados para usuários
package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Role		 string `json:"role"` //  "admin" ou "user"
}
