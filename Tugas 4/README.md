# Auto Reload Data Project

Buatlah sebuah microservice untuk meng-update sebuah file json setiap 15 detik dengan angka random antara 1-100untuk valuewater dan wind. Seperti berikut:

```json
{
    "status": {
        "wind":78,
        "water":55
    }
}
```

Kemudian tampilkan data tersebut di path `/` . Selain itu kalian harus menentukan statuswater dan wind tersebut. Dengan ketentuan:
- jika water dibawah 5 maka status aman
- jika water antara 6 - 8 maka status siaga
- jika water diatas 8 maka status bahaya
- jika wind dibawah 6 maka status aman
- jika wind antara 7 - 15 maka status siaga
- jika wind diatas 15 maka status bahaya
- value water dalam satuan meter●value wind dalam satuan meter per detik