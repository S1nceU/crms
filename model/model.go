package model

type Customer struct {
	CustomerId  int     `json:"CustomerId"  gorm:"primary_key; auto_increase;not null"`
	Name        string  `json:"Name"        gorm:"column:name; not null"`
	Gender      string  `json:"Gender"      gorm:"column:gender; not null"`
	Birthday    string  `json:"Birthday"    gorm:"column:birthday; not null"`
	ID          string  `json:"ID"          gorm:"column:ID; type:varchar(100); uniqueIndex; not null"`
	Address     string  `json:"Address"     gorm:"column:address"`
	PhoneNumber string  `json:"PhoneNumber" gorm:"column:phonenumber"`
	CarNumber   string  `json:"CarNumber"   gorm:"column:carnumber"`
	Citizenship string  `json:"Citizenship" gorm:"column:Citizenship; not null"`
	Note        string  `json:"Note"        gorm:"column:ex"`
	History     History `                   gorm:"foreignKey:CustomerId"`
}

type History struct {
	HistoryId      int    `json:"HistoryId"      gorm:"primary_key;auto_increase; not null"`
	CustomerId     int    `json:"CustomerId"     gorm:"column:customer_id; not null"`
	Date           string `json:"Date"           gorm:"column:date; not null"`
	NumberOfPeople int    `json:"NumberOfPeople" gorm:"column:Nofpeople; not null"`
	Price          int    `json:"Price"          gorm:"column:price; not null"`
	Room           string `json:"Room"           gorm:"column:room; not null"`
	Note           string `json:"Note"           gorm:"column:ex"`
}
