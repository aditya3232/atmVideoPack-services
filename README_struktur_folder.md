.
├── go.mod
|   # file go.mod dipergunakan oleh go module (jika go mod diaktifkan).
|
├── config
|   # folder config isinya adalah file .env dan function load config.
|
├── connection
|   # folder connection berisi inisialisasi koneksi ke database, storage, cache, search engine, message queue.
|
├── constant
|   # folder constant berisi constanta yang digunakan di seluruh aplikasi.
|
├── cron
|   # folder cron berisi cron job yang akan dijalankan oleh aplikasi.
|
├── handler
|   # folder handler berisi handler untuk masing-masing route.
|   # handler yang akan menangani input, dan servicenya, dari servis nanti ada balikan response
|
├── helper
|   # folder helper berisi helper yang digunakan di seluruh aplikasi.
|
├── library
|   # folder library berisi library yang digunakan di seluruh aplikasi.
|   # atau tempat paket-paket (libraries) yang digunakan oleh bagian lain dari projek.
|   # bisa library apa aja. misal mino, jwt, dll
|
├── log
|   # folder log  berisi fungsi untuk menyimpan log
|   # bisa di file atau disimpan di elastic
|
├── middleware
|   # folder middleware berisi middleware yang digunakan di seluruh aplikasi.
|   # atau fungsi dari library middleware yang di-reuse oleh seluruh aplikasi,
|   # seperti middleware untuk auth jwt, api-key, cors.
|   # hampir sama seperti folder library, cuma disini khusus library middleware
|
├── model/
|   # folder model berisi source kode utama aplikasi. 
|   # atau bisa disebut model utama aplikasi
|   # didalam salah satu folder disini, berisi file-file seperti service, repository, entity, input, formatter
|   # service berfungsi mengatur bisnis logic utama
|   # repository berfungsi untuk mengatur crud data dari database
|   # entity berfungsi untuk mengatur struktur data dari database
|   # input (dto, data transfer object) berfungsi untuk mengatur input dari user
|   # formatter berfungsi untuk mengatur output dari aplikasi, atau memformat response
|   |
│   ├── your_app_1/
│   ├── your_app_2/
│   ├── your_app_3/
│   └── ...
|
├── routes
|   # folder routes adalah tempat untuk mengatur route yang digunakan oleh aplikasi.
|
└── ...

