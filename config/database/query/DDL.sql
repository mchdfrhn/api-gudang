CREATE TABLE gudang (
    kode varchar(10) PRIMARY KEY,
    nama varchar(100) NOT NULL
);

CREATE TABLE barang (
    kode varchar(10) PRIMARY KEY,
    nama varchar(100) NOT NULL,
    harga int NOT NULL,
    jumlah int NOT NULL,
    expired date NOT NULL,
    kode_gudang varchar(10),
    FOREIGN KEY (kode_gudang) REFERENCES gudang(kode)
);