package product

import "time"

type Product struct {
	ID           int
	Name         string
	Category     string
	Brand        string
	Spec         string
	Rp           string
	Disc         string
	Rank         int
	DateRegister time.Time
	DateUpdate   time.Time
	Photos       []Photos
}

type Photos struct {
	Id           int
	ProductID    int
	FileName     string
	IsCover      int
	DateRegister time.Time
	DateUpdate   time.Time
}
