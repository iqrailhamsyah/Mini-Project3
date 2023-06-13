package entities

type Roles struct {
	ID        uint   `gorm:"primary_key"`
	Role_name string `gorm:"column:role_name"`
}
