package domain

type Category struct {
	Base
	Name string `gorm:"size:50; not null; index;"`
}
