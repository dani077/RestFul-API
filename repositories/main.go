package repositories

import (
	"database/sql"
)

type Item struct {
	TblItemID    int     `json:"TblItemID"`
	ItemCode     string  `json:"ItemCode"`
	ItemName     string  `json:"ItemName"`
	BuyingPrice  float64 `json:"BuyingPrice"`
	SellingPrice float64 `json:"SellingPrice"`
	ItemAmount   int     `json:"ItemAmount"`
	Pieces       string  `json:"Pieces"`
}

type Officer struct {
	TblOfficerID    int    `json:"TblOfficeID"`
	OfficerCode     string `json:"OfficerCode"`
	OfficerName     string `json:"OfficerName"`
	OfficerPassword int    `json:"OfficerPassword"`
	OfficerStatus   int    `json:"OfficerStatus"`
}

type Selling struct {
	TblSellingID int     `json:"TblSellingID"`
	Invoice      string  `json:"Invoice "`
	InvoiceDate  string  `json:"InvoiceDate"`
	Item         int     `json:"Item"`
	Total        float64 `json:"Total"`
	Paid         float64 `json:"Paid"`
	Kembali      float64 `json:"Kembali"`
	OfficerCode  string  `json:"OfficerCode"`
}

type Detail struct {
	TblSellingDetailID int     `json:"tblSellingDetailID"`
	Invoice            string  `json:"Invoice"`
	ItemCode           string  `json:"ItemCode"`
	ItemName           string  `json:"ItemName"`
	ItemPrice          string  `json:"ItemPrice"`
	Subtotal           float64 `json:"Subtotal"`
}
type Transaksi struct {
	ItemName    string `json:"ItemName"`
	InvoiceDate string `json:"InvoiceDate"`
	OfficerName string `json:"OfficerName"`
}

/*
func (i *Item) GetItem(db *sql.DB) ([]Item, error) {
	return nil, db.QueryRow("select TblItemID, ItemCode, ItemName, BuyingPrice, SellingPrice, ItemAmount,Pieces from tblItem").Scan(&i.TblItemID, &i.ItemCode, &i.ItemName, &i.BuyingPrice, &i.SellingPrice, &i.ItemAmount, &i.Pieces)

}*/

func GetItem(db *sql.DB) ([]Item, error) {
	data, err := db.Query("Select tblItemID,ItemCode,ItemName,BuyingPrice,SellingPrice,ItemAmount,Pieces From tblItem")
	if err != nil {
		return nil, err
	}
	defer data.Close()

	item := []Item{}
	for data.Next() {
		var ditem Item

		if err := data.Scan(&ditem.TblItemID, &ditem.ItemCode, &ditem.ItemName, &ditem.BuyingPrice, &ditem.SellingPrice, &ditem.ItemAmount, &ditem.Pieces); err != nil {
			return nil, err
		}
		item = append(item, ditem)
	}

	return item, nil
}

func (o *Officer) GetOfficer(db *sql.DB) error {
	return db.QueryRow("select TblOfficerID, OfficerCode, OfficerName, OfficerPassword, OfficerStatus from tblOfficer").Scan(&o.TblOfficerID, &o.OfficerCode, &o.OfficerName, &o.OfficerPassword, &o.OfficerStatus)
}

func (s *Selling) GetSelling(db *sql.DB) error {
	return db.QueryRow("select TblSellingID, Invoice, InvoiceDate, Item, Total, Paid,Kembali,OfficerCode from tblSelling").Scan(&s.TblSellingID, &s.Invoice, &s.InvoiceDate, &s.Item, &s.Total, &s.Paid, &s.Kembali, &s.OfficerCode)
}

func (d *Detail) GetDetail(db *sql.DB) error {
	return db.QueryRow("select TblSellingDetailID,Invoice, ItemCode, ItemName, ItemPrice, Subtotal from tblSellingDetail").Scan(&d.TblSellingDetailID, &d.Invoice, &d.ItemCode, &d.ItemName, &d.ItemPrice, &d.Subtotal)
}

func (t *Transaksi) GetTrans(db *sql.DB) error {
	return db.QueryRow("select ItemName, InvoiceDate, OfficerName from Transaksi").Scan(&t.ItemName, &t.InvoiceDate, &t.OfficerName)
}

func (i *Item) UpdateItem(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE tblItem SET ItemCode=?, ItemName=? WHERE TblItemID=?",
			i.ItemCode, i.ItemName, i.TblItemID)

	return err
}

func (i *Item) DeleteItem(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM tblItem WHERE tblItemID=?", i.TblItemID)

	return err
}

func (i *Item) IItem(db *sql.DB) error {
	_, err := db.Query("INSERT INTO tblItem(ItemCode, ItemName) VALUES( ?,?) ", i.ItemCode, i.ItemName)

	return err
}
