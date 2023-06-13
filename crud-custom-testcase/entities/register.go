package entities

type Register struct {
	Id             uint   `gorm:"primary_key"`
	Admin_id       uint   `gorm:"column:admin_id"`
	Super_admin_id uint   `gorm:"column:super_admin_id"`
	Status         string `gorm:"column:status"`
}
