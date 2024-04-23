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

__GoHadits__ is a simple open-source API built using Golang to access a collection of Hadiths.

*Read this in other languages: [Bahasa Indonesia](README.md).*

## Prerequisite

This project is built using [**Go version 1.22.2**](https://go.dev/dl/), and it is expected to be developed using this specific version of Golang to ensure the desired outcome.
 
## How to Run

- Install project dependencies using the command `go mod download`.
- Run the service using the command `go run .` or `go run main.go`.

## How to Build

Execute the following command to build the binary:
```shell
go build -ldflags "-s -w" -o ./out .
```

Then you can run the service using the command `./out`.

## Usage

### [GET] /api/v1/hadith
Returning a list of available hadith books.

#### Success response
The response will be returned an JSON. For example:
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

#### Error response
The response will be returned as JSON in case of an error. For example:
```json
{
    "code": 500,
    "message": "Internal Server Error",
    "status": "error"
}
```

### [GET] /api/v1/hadith/{book}
Returning a list of hadiths found in the requested book.

#### Parameters
| Name | Required | Type | Description |
| ----:|:--------:|:----:| ----------- |
| `book` | required | param  | Book name. <br /> (`data[i].book` from `api/v1/hadith`) |
| `limit` | optional | query  | Specifies the number of hadith to fetch per request. <br/>The default is `50` and the maximum is `150`. |
| `offset` | optional | query  | Specifies the starting point (offset) from where to fetch the items. The default is `0`. |

#### Success response
The response will be returned an JSON. For example:
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

#### Error response
The response will be returned as JSON in case of an error. For example:
```json
{
  "code": 400,
  "message": "offset should not larger than total hadith available",
  "status": "error"
}
```

### [GET] /api/v1/hadith/{book}/{number}
Returning a specific hadith according to the given hadith number.

#### Parameters
| Name | Required | Type | Description |
| ----:|:--------:|:----:| ----------- |
| `book` | required | param  | Book name. <br /> (`data[i].book` from `api/v1/hadith`) |
| `number` | required | param  | Hadith number. |

#### Success response
The response will be returned an JSON. For example:
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

#### Error response
The response will be returned as JSON in case of an error. For example:
```json
{
  "code": 404,
  "message": "hadith no. 5000 from book of Abu Dawud is not found",
  "status": "error"
}
```

## License

GoHadith is open-sourced software licensed under the [MIT license](https://github.com/iqbaleff214/gohadits/blob/main/LICENSE).