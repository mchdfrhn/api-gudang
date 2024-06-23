-- Buat Database
CREATE DATABASE db_gudang;





-- Buat Tabel

CREATE TABLE gudang (
    kode serial PRIMARY KEY,
    nama varchar(100) NOT NULL
);

CREATE TABLE barang (
    kode serial PRIMARY KEY,
    nama varchar(100) NOT NULL,
    harga int NOT NULL,
    jumlah int NOT NULL,
    expired date NOT NULL,
    kode_gudang int,
    FOREIGN KEY (kode_gudang) REFERENCES gudang(kode)
);

CREATE INDEX idx_barang_kodegudang ON barang(kode_gudang);





-- Buat Prosedur utk Menampilkan Data dengan Dynamic Query + Paging

CREATE OR REPLACE FUNCTION get_barang_list(
    limit int,
    offset int
)
RETURNS TABLE (
    kode_gudang int,
    nama_gudang varchar,
    kode_barang int,
    nama_barang varchar,
    harga_barang int,
    jumlah_barang int,
    expired_barang date
) AS $$
BEGIN
    RETURN QUERY EXECUTE format(
        'SELECT g.kode AS kode_gudang, 
                g.nama AS nama_gudang, 
                b.kode AS kode_barang, 
                b.nama AS nama_barang, 
                b.harga AS harga_barang, 
                b.jumlah AS jumlah_barang, 
                b.expired AS expired_barang
         FROM gudang AS g
         JOIN barang AS b ON g.kode = b.kode_gudang
         ORDER BY g.kode, b.kode
         LIMIT %L OFFSET %L',
         limit, offset
    );
END;
$$ LANGUAGE plpgsql;





-- Buat Trigger Menampilkan Barang Kadaluarsa saat Input Barang

CREATE OR REPLACE FUNCTION check_expired_barang()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.expired < CURRENT_DATE THEN
        RAISE NOTICE 'Barang % sudah kadaluarsa pada tanggal %', NEW.nama, NEW.expired;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_check_expired_barang
AFTER INSERT ON barang
FOR EACH ROW
EXECUTE FUNCTION check_expired_barang();






