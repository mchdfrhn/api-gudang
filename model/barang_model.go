package model

type Barang struct {
	Kode    string    `json:"kode"`
	Nama    string `json:"nama"`
	Harga   int    `json:"harga"`
	Jumlah  int    `json:"jumlah"`
	Expired string `json:"expired"`
	Gudang  Gudang `json:"gudang"`
}
