<div align="center">
    <p>
        <a href="https://github.com/404NotFoundIndonesia/" target="_blank">
            <img src="https://avatars.githubusercontent.com/u/87377917?s=200&v=4" width="200" alt="404NFID Logo">
        </a>
    </p>

 [![GitHub Stars](https://img.shields.io/github/stars/iqbaleff214/gohadits.svg)](https://github.com/iqbaleff214/gohadits/stargazers)
 [![GitHub license](https://img.shields.io/github/license/iqbaleff214/gohadits)](https://github.com/iqbaleff214/gohadits/blob/main/LICENSE)
 
</div>

# GoHadits

__GoHadits__ adalah API _open-source_ sederhana yang dibangun menggunakan Golang untuk mengakses kumpulan hadits.

*Baca dengan bahasa lain: [English](README.en.md).*

## Prasyarat

Proyek ini dibangun menggunakan [**Go version 1.22.2**](https://go.dev/dl/), dan diharapkan untuk dikembangkan menggunakan versi Golang yang serupa untuk mendapatkan hasil sesuai harapan.
 
## Cara Menjalankan

- Instalasikan dependensi proyek menggunakan perintah `go mod download`.
- Jalankan proyek dengan perintah `go run .` atau `go run main.go`.

## Cara Membuat _File_ Biner

Jalankan perintah berikut untuk membuat _file_ biner:
```shell
go build -ldflags "-s -w" -o ./out .
```

Kemudian jalankan menggunakan perintah `./out`.

## Penggunaan

### [GET] /api/v1/hadith
Mengembalikan daftar kitab hadits yang tersedia.

#### Respon sukses
Respon akan dikembalikan dalam bentuk JSON. Contohnya:
```json
{
  "code": 200,
  "data": [
    {
      "book": "bukhari",
      "name": "Bukhari",
      "total": 6638
    },
    {
      "book": "muslim",
      "name": "Muslim",
      "total": 4930
    }
  ],
  "message": "Hadith books successfully retrieved.",
  "status": "success"
}
```

#### Respon eror
Respon akan dikembalikan dalam bentuk JSON juga jika terdapat eror. Contohnya:
```json
{
    "code": 500,
    "message": "Internal Server Error",
    "status": "error"
}
```

### [GET] /api/v1/hadith/{book}
Mengembalikan daftar hadits yang terdapat pada kitab yang diminta.

#### Parameters
| Nama | Keberadaan | Tipe | Deskripsi |
| ----:|:--------:|:----:| ----------- |
| `book` | wajib | _param_  | Nama kitab hadits. <br /> (diambil dari `data[i].book` yang ada di `api/v1/hadith`) |
| `limit` | opsional | _query_  | Jumlah hadits yang diambil dalam sekali _request_. <br/>Bawaannya adalah `50` dan maksimal `150`. |
| `offset` | opsional | _query_  | Indeks awal hadits diambil. Bawaannya adalah `0`. |

#### Respon sukses
Respon akan dikembalikan dalam bentuk JSON. Contohnya:
```json
{
  "code": 200,
  "data": {
    "contents": [
      {
        "number": 1,
        "arab": "حَدَّثَنَا عَبْدُ اللَّهِ بْنُ مَسْلَمَةَ بْنِ قَعْنَبٍ الْقَعْنَبِيُّ حَدَّثَنَا عَبْدُ الْعَزِيزِ يَعْنِي ابْنَ مُحَمَّدٍ عَنْ مُحَمَّدٍ يَعْنِي ابْنَ عَمْرٍو عَنْ أَبِي سَلَمَةَ عَنْ الْمُغِيرَةِ بْنِ شُعْبَةَأَنَّ النَّبِيَّ صَلَّى اللَّهُ عَلَيْهِ وَسَلَّمَ كَانَ إِذَا ذَهَبَ الْمَذْهَبَ أَبْعَدَ",
        "id": "Telah menceritakan kepada kami [Abdullah bin Maslamah bin Qa'nab al Qa'nabi] telah menceritakan kepada kami [Abdul Aziz yakni bin Muhammad] dari [Muhammad yakni bin Amru] dari [Abu Salamah] dari [Al Mughirah bin Syu'bah] bahwasanya Nabi shallallahu 'alaihi wasallam apabila hendak pergi untuk buang hajat, maka beliau menjauh."
      }
    ],
    "limit": 1,
    "name": "Abu Dawud",
    "offset": 0,
    "total": 4419
  },
  "message": "Hadith collection of Abu Dawud successfully retrieved.",
  "status": "success"
}
```

#### Respon eror
Respon akan dikembalikan dalam bentuk JSON juga jika terdapat eror. Contohnya:
```json
{
  "code": 400,
  "message": "offset should not larger than total hadith available",
  "status": "error"
}
```

### [GET] /api/v1/hadith/{book}/{number}
Mengembalikan hadits spesifik berdasarkan nomor hadits.

#### Parameters
| Nama | Keberadaan | Tipe | Deskripsi |
| ----:|:--------:|:----:| ----------- |
| `book` | wajib | _param_  | Nama kitab hadits. <br /> (diambil dari `data[i].book` yang ada di `api/v1/hadith`) |
| `number` | wajib | _param_  | Nomor hadits. |

#### Respon sukses
Respon akan dikembalikan dalam bentuk JSON. Contohnya:
```json
{
  "code": 200,
  "data": {
    "content": {
      "number": 5,
      "arab": "حَدَّثَنَا عَمْرُو بْنُ مَرْزُوقٍ أَخْبَرَنَا شُعْبَةُ عَنْ قَتَادَةَ عَنْ النَّضِرِ بْنِ أَنَسٍ عَنْ زَيْدِ بْنِ أَرْقَمَعَنْ رَسُولِ اللَّهِ صَلَّى اللَّهُ عَلَيْهِ وَسَلَّمَ قَالَ إِنَّ هَذِهِ الْحُشُوشَ مُحْتَضَرَةٌ فَإِذَا أَتَى أَحَدُكُمْ الْخَلَاءَ فَلْيَقُلْ أَعُوذُ بِاللَّهِ مِنْ الْخُبُثِ وَالْخَبَائِثِ",
      "id": "Telah menceritakan kepada kami [Amru bin Marzuq] telah mengabarkan kepada kami [Syu'bah] dari [Qatadah] dari [an Nadhr bin Anas] dari [Zaid bin Arqam] dari Rasulullah shallallahu 'alaihi wasallam, beliau bersabda: \"Sesungguhnya tempat buang hajat itu dihadiri oleh setan-setan, maka apabila salah seorang dari kalian mendatangi WC, hendaklah dia mengucapkan; 'Aku berlindung kepada Allah dari setan jantan dan setan betina'.\""
    },
    "name": "Abu Dawud",
    "total": 4419
  },
  "message": "Hadith no. 5 from book of Abu Dawud successfully retrieved.",
  "status": "success"
}
```

#### Respon eror
Respon akan dikembalikan dalam bentuk JSON juga jika terdapat eror. Contohnya:
```json
{
  "code": 404,
  "message": "hadith no. 5000 from book of Abu Dawud is not found",
  "status": "error"
}
```

## License

GoHadith adalah perangkat lunak _open-source_ yang dilisensikan di bawah lisensi [MIT license](https://github.com/iqbaleff214/gohadits/blob/main/LICENSE).