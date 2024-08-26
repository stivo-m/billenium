package domain

type Transaction struct {
	Base
	CategoryId string          `gorm:"type:uuid;not null;index;"`
	UserId     string          `gorm:"type:uuid;not null;index;"`
	Name       string          `gorm:"size:100;not null"`
	Type       string          `gorm:"size:20;not null;"`
	Amount     string          `gorm:"size:100"`
	TxnType    TxnTypeOption   `gorm:"type:txn_type null; default:incoming"`
	TxnStatus  TxnStatusOption `gorm:"type:txn_status null; default:draft"`
}
