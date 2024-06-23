package request

type BarangRequest struct {
	Kode       string `json:"kode"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	Jumlah     int    `json:"jumlah"`
	Expired    string `json:"expired"`
	KodeGudang string `json:"kode_gudang"`
}
