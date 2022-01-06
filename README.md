# Penjelasan Alur Aplikasi
Pengguna mengakses `/auth/google` pada aplikasi untuk mulai melakukan autentikasi aplikasi melalui akun google yang dimiliki. Kemudian aplikasi akan melakukan redirect ke halaman login milik Google. Autentikasi akun google ini sepenuhnya di*handle* oleh google sehingga aplikasi tidak akan mengetahui *password* google yang diinput oleh pengguna. Jika login google berhasil, google akan redirect ke halaman aplikasi `/auth/google/callback` denagn memberikan sejumlah informasi profile google yang baru saja sukses login.

Informasi yang diberikan Google pada aplikasi ini ialah:
- Nama Lengkap Profile
- Email
- Userid

Pengguna akan terus login dengan konsep `cookie` yang tersimpan di local device pengguna. Umur cookie diatur hanya 2 menit untuk saat ini hanya untuk keperluan testing. Pada kenyataanya umur cookie dapat berminggu-minggu.

Untuk mempertahankan status login user dapat menerapkan [JWT](https://jwt.io/) sehingga user tidak perlu login berkali-kali.

