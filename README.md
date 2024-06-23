# API Gudang Sederhana

## Deskripsi

Aplikasi sederhana berbasis API untuk CRUD data Gudang dan Barang.

Features:

1.  Struktur/Design Database :

        - Table gudang = Master data Customer
        - Table barang = Master data Rating


2.  Features :

        - Aplikasi memiliki fitur untuk melakukan GET, POST, PUT, dan DELETE pada tabel master

          
- - -



## Configuration

### Database Configuration

#### Connection
Edit value pada file `.env.example` lalu ubah nama file menjadi `.env` untuk mengatur koneksi ke database.
Pastikan mengganti nilai sesuai dengan konfigurasi database yang Anda gunakan.

#### Database & Tables
Eksekusi query yang dibutuhkan untuk menjalankan aplikasi. File query terdapat pada `/config/database/query/DDL.sql`



## API Spec

### Gudang API

#### Create Gudang

Request :

- Method : `POST`
- Endpoint : `/gudang`
- Body :

```json
{
    "kode": "string",
    "nama": "string"
}
```


#### Get List Gudang

Request :

- Method : GET
- Endpoint : `/gudang`


#### Update Gudang

Request :

- Method : PUT
- Endpoint : `/gudang`
- Body :

```json
{
    "kode": "string",
    "nama": "string"
}
```


#### Delete Gudang

Request :

- Method : DELETE
- Endpoint : `/gudang/:id`



### Barang API

#### Create Barang

Request :

- Method : `POST`
- Endpoint : `/barang`
- Body :

```json
{
    "kode": "string",
    "nama": "string",
    "harga":  int,
    "jumlah": int,
    "expired": "yyy-mm-dd",
    "kode_gudang": "string"
}
```


#### Get List Barang

Request :

- Method : GET
- Endpoint : `/barang`


#### Update Barang

Request :

- Method : PUT
- Endpoint : `/barang`
- Body :

```json
{
    "kode": "string",
    "nama": "string",
    "harga":  int,
    "jumlah": int,
    "expired": "yyy-mm-dd",
    "kode_gudang": "string"
}
```


#### Delete Barang

Request :

- Method : DELETE
- Endpoint : `/barang/:id`