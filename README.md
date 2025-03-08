# Aplikasi Web Ping Sederhana dengan Gin

Aplikasi web sederhana ini menggunakan framework Gin untuk membuat endpoint `/ping` yang mengembalikan respons JSON.

## Prasyarat

* Go 1.22 atau lebih baru
* Docker (opsional, untuk menjalankan dalam container)

## Cara Menjalankan

### Tanpa Docker

1.  Pastikan Go sudah terinstall.
2.  Clone repositori ini:

    ```bash
    git clone [go-docker-railway](https://github.com/py2rx/go-docker-railway)
    cd go-docker-railway
    ```

3.  Download dependensi:

    ```bash
    go mod download
    go mod tidy
    ```

4.  Jalankan aplikasi:

    ```bash
    go run main.go
    ```

    Aplikasi akan berjalan di `http://localhost:8080`. Anda dapat mengubah port dengan variabel environment `PORT`:

    ```bash
    PORT=9000 go run main.go
    ```

5.  Akses endpoint `/ping`:

    ```bash
    curl http://localhost:8080/ping
    ```

    atau

    ```bash
    curl http://localhost:9000/ping
    ```

    (sesuaikan dengan port yang Anda gunakan).

### Dengan Docker

1.  Pastikan Docker sudah terinstall.
2.  Clone repositori ini.
3.  Build image Docker:

    ```bash
    docker build -t ping-app .
    ```

4.  Jalankan container Docker:

    ```bash
    docker run -p 8080:8080 ping-app
    ```

    atau jika ingin menggunakan port lain.

    ```bash
    docker run -p 9000:8080 ping-app
    ```

5.  Akses endpoint `/ping`:

    ```bash
    curl http://localhost:8080/ping
    ```

    atau

    ```bash
    curl http://localhost:9000/ping
    ```

    (sesuaikan dengan port yang Anda gunakan).

## Struktur Proyek

```sh
.
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
└── README.md
```

* `main.go`: Kode sumber aplikasi Go.
* `go.mod` dan `go.sum`: File dependensi Go.
* `Dockerfile`: Konfigurasi Docker untuk build image.
* `docker-compose.yml`: Konfigurasi Docker Compose (jika digunakan).
* `README.md`: Dokumentasi proyek.

## Endpoint

* `GET /ping`: Mengembalikan respons JSON dengan status, pesan, dan hasil.

## Contoh Respons JSON

```json
{
  "status": 200,
  "message": "pong",
  "result": "success"
}