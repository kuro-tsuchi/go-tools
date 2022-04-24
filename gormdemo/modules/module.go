package modules

type Model struct {
	ID        string `gorm:"primarykey"`
	CreatedAt string
	UpdatedAt string
}
