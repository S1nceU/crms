package model

type Customer struct {
	CustomerId  int     `json:"customerId" gorm:"primary_key; auto_increase;not null"`
	Name        string  `json:"name"        gorm:"column:name; not null"`
	Gender      string  `json:"gender"      gorm:"column:gender; not null"`
	Birthday    string  `json:"birthday"    gorm:"column:birthday; not null"`
	ID          string  `json:"ID"          gorm:"column:ID; type:varchar(100); uniqueIndex"`
	Address     string  `json:"address"     gorm:"column:address"`
	Phonenumber string  `json:"phonenumber" gorm:"column:phonenumber"`
	Carnumber   string  `json:"carnumber"   gorm:"column:carnumber"`
	Citizenship string  `json:"citizenship" gorm:"column:Citizenship; not null"`
	Note        string  `json:"note"        gorm:"column:ex"`
	History     History `                   gorm:"foreignKey:CustomerId"`
}

type History struct {
	HistoryId  int    `json:"historyId"  gorm:"primary_key;auto_increase; not null"`
	CustomerId int    `json:"customerId" gorm:"column:customer_id; not null"`
	Date       string `json:"date"        gorm:"column:date; not null"`
	Nofpeople  int    `json:"Nofpeople"   gorm:"column:Nofpeople; not null"`
	Price      int    `json:"price"       gorm:"column:price; not null"`
	Note       string `json:"note"        gorm:"column:ex"`
}
