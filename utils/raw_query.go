package utils

const (
	
	// Master Data Gudang
	CREATE_GUDANG = `INSERT INTO gudang(kode, nama) VALUES ($1, $2)`
	READ_GUDANG = `SELECT * FROM gudang`
	UPDATE_GUDANG = `UPDATE gudang SET nama = $1 WHERE kode = $2`
	DELETE_GUDANG = `DELETE FROM gudang WHERE kode = $1`

	// Master Data Barang
	CREATE_BARANG = `INSERT INTO barang(kode, nama, harga, jumlah, expired, kode_gudang) VALUES ($1, $2, $3, $4, $5, $6)`
	READ_BARANG = `SELECT b.kode, b.nama, b.harga, b.jumlah, b.expired, g.kode, g.nama FROM barang AS b INNER JOIN gudang AS g ON b.kode_gudang = g.kode;`
	UPDATE_BARANG = `UPDATE barang SET nama = $1, harga = $2, jumlah = $3, expired = $4, kode_gudang = $5 WHERE kode = $6`
	DELETE_BARANG = `DELETE FROM barang WHERE kode = $1`

)