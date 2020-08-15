<div align="center"> 

`tzone` is a rest API that has written in `golang` and it serves address information for Turkish developers

[![Build Status](https://img.shields.io/travis/73VW/TechnicalReport.svg?style=for-the-badge&label=build)](https://travis-ci.com/enesusta/tzone)
![Coveralls github](https://img.shields.io/coveralls/github/enesusta/tzone?style=for-the-badge)
[![CodeFactor](https://www.codefactor.io/repository/github/enesusta/tzone/badge?style=for-the-badge)](https://www.codefactor.io/repository/github/enesusta/tzone)
[![Go Report Card](https://goreportcard.com/badge/github.com/enesusta/tzone?style=for-the-badge)](https://goreportcard.com/report/github.com/enesusta/tzone)

</div>
<br/>

- Index
  - [Install](#-install) 
  - [Preface](#-preface)
  - [Raw data before parsing](#raw-data-before-parsing)
  - [The data parsed by tzone-parser](#the-data-parsed-by-tzone-parser)
  - [Endpoints](#endpoints)
    - [/provinces](#provinces)
      - [Sample Request](#sample-request-1)
      - [Sample Response](#sample-response-1)
    - [/provinces/{provinceName}](#provinces/{provinceName})
      - [Sample Request](#sample-request-2)
      - [Sample Response](#sample-response-2)
    - [/counties](#counties)
      - [Sample Request](#sample-request-3)
      - [Sample Response](#sample-response-3)
    - [/counties/{provinceName}](#countiesprovincename)
      - [Sample Request](#sample-request-4)
      - [Sample Response](#sample-response-4)
    - [/towns](#towns)
    - [/towns/{provinceName}](#townsprovincename)
    - [/towns/{provinceName}/{countyName}](#townsprovincenamecountyname)
    - [/villages](#villages)
    - [/villages/{provinceName}](#villagesprovincename)
    - [/villages/{provinceName}/{countyName}](#villagesprovincenamecountyname)
    - [/villages/{provinceName}/{countyName}/{townName}](#villagesprovincenamecountynametownname)


## 🕺 Install

```
go get -u github.com/enesusta/tzone
```

## 📯 Preface

I have needed a rest API that serves the address location to my ERP(enterprise resource planning) application. After searching I didn't find a rest API that fits for me and my requirements.

PTT already has presented [data that contains address information for Turkey](https://postakodu.ptt.gov.tr/). But the problem is: the data is not relational and its file format is xlsx(Microsoft Excel Open XML Spreadsheet) that doesn't fit to rest API.

What do I mean by '`it doesn't fit to rest API`' ?

Let's look at the raw data that has provided by PTT.


## 📥 Raw Data Before Parsing

<br />

<div align="center">


| Province | County | Town/State | Village/Neighbourhood | Zip Code |
|:-:|:-:|:-:|:-:|:-:|
|EDİRNE                        |ENEZ                          |ENEZ                          |ÇATALTEPE MAH                                                              |22700|
|EDİRNE                        |ENEZ                          |ENEZ                          |GAZİÖMERBEY MAH                                                            |22700|
|EDİRNE                        |ENEZ                          |ENEZ                          |YENİ MAH                                                                   |22700|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |ABDURRAHİM KÖYÜ                                                            |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |BÜYÜKEVREN KÖYÜ                                                            |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |ÇANDIR KÖYÜ                                                                |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |ÇAVUŞKÖY KÖYÜ                                                              |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |ÇERİBAŞI KÖYÜ                                                              |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |GÜLÇAVUŞ KÖYÜ                                                              |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |HASKÖY KÖYÜ                                                                |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |HİSARLI KÖYÜ                                                               |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |IŞIKLI KÖYÜ                                                                |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |KARAİNCİRLİ KÖYÜ                                                           |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |KOCAALİ KÖYÜ                                                               |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |KÜÇÜKEVREN KÖYÜ                                                            |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |SULTANİÇE KÖYÜ                                                             |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |SÜTÇÜLER KÖYÜ                                                              |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |ŞEHİTLER KÖYÜ                                                              |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |UMURBEY KÖYÜ                                                               |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |VAKIF KÖYÜ                                                                 |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |YAZIR KÖYÜ                                                                 |22750|
|EDİRNE                        |ENEZ                          |MERKEZKÖYLER                  |YENİCE KÖYÜ                                                                |22750|

</div>

<br/>
<br/>

Precisely for this reason, `I had to write a xlsx parser` that parses xlsx file to produce the JSON file which fits for API.

Thus, [tzone-parser](https://github.com/enesusta/tzone-parser) was written.

tzone-parser is not only parsed the file. It also `provides relational data` for each record.

## 📤 The data parsed by `tzone-parser`

```json
{
   "provinceName":"Edirne",
   "provinceCounties":[
      {
         "countyName":"Enez",
         "countyTowns":[ { 
               "townName":"Enez",
               "townVillages":[
                  { "villageName":"Gaziömerbey mah", "zipCode":22700 },
                  { "villageName":"Yeni mah", "zipCode":22700 },
                  { "villageName":"Çataltepe mah", "zipCode":22700 }
               ]
            },
            {
               "townName":"Merkezköyler",
               "townVillages":[
                  { "villageName":"Abdurrahim köyü", "zipCode":22750 },
                  { "villageName":"Büyükevren köyü", "zipCode":22750 },
                  { "villageName":"Gülçavuş köyü", "zipCode":22750 },
                  { "villageName":"Hasköy köyü", "zipCode":22750 },
                  { "villageName":"Hisarlı köyü", "zipCode":22750 },
                  { "villageName":"Işıklı köyü", "zipCode":22750 },
                  { "villageName":"Karaincirli köyü", "zipCode":22750 },
                  { "villageName":"Kocaali köyü", "zipCode":22750 },
                  { "villageName":"Küçükevren köyü", "zipCode":22750 },
                  { "villageName":"Sultaniçe köyü", "zipCode":22750 },
                  { "villageName":"Sütçüler köyü", "zipCode":22750 },
                  { "villageName":"Umurbey köyü", "zipCode":22750 },
                  { "villageName":"Vakıf köyü", "zipCode":22750 },
                  { "villageName":"Yazır köyü", "zipCode":22750 },
                  { "villageName":"Yenice köyü", "zipCode":22750 },
                  { "villageName":"Çandır köyü", "zipCode":22750 },
                  { "villageName":"Çavuşköy köyü", "zipCode":22750 },
                  { "villageName":"Çeribaşı köyü", "zipCode":22750 },
                  { "villageName":"Şehitler köyü", "zipCode":22750 }
               ]
            }
         ]
      }
   ]
}
```


## Endpoints

### `/provinces`

This endpoint returns all provinces and their names that Turkey has.

#### Sample Request 1

```http
http://localhost:8080/provinces
```

#### Sample Response 1

```json
[
    { "provinceName": "Adana" },
    { "provinceName": "Adıyaman" },
    { "provinceName": "Afyonkarahisar" },
    { "provinceName": "Ağrı" },
    { "provinceName": "Aksaray" },
    { "provinceName": "Amasya" },
    { "provinceName": "Ankara" },
    { "provinceName": "Antalya" },
    { "provinceName": "Ardahan" },
    { "provinceName": "Artvin" },
    { "provinceName": "Aydın" },
    { "provinceName": "Balıkesir" },
    { "provinceName": "Bartın" },
    { "provinceName": "Batman" },
    { "provinceName": "Bayburt" },
    { "provinceName": "Bilecik" },
    { "provinceName": "Bingöl" },
    { "provinceName": "Bitlis" },
    { "provinceName": "Bolu" },
    { "provinceName": "Burdur" },
    { "provinceName": "Bursa" },
    { "provinceName": "Çanakkale" },
    { "provinceName": "Çankırı" },
    { "provinceName": "Çorum" },
    { "provinceName": "Denizli" },
    { "provinceName": "Diyarbakır" },
    { "provinceName": "Düzce" },
    { "provinceName": "Edirne" },
    { "provinceName": "Elazığ" },
    { "provinceName": "Erzincan" },
    { "provinceName": "Erzurum" },
    { "provinceName": "Eskişehir" },
    { "provinceName": "Gaziantep" },
    { "provinceName": "Giresun" },
    { "provinceName": "Gümüşhane" },
    { "provinceName": "Hakkari" },
    { "provinceName": "Hatay" },
    { "provinceName": "Iğdır" },
    { "provinceName": "Isparta" },
    { "provinceName": "İstanbul" },
    { "provinceName": "İzmir" },
    { "provinceName": "Kahramanmaraş" },
    { "provinceName": "Karabük" },
    { "provinceName": "Karaman" },
    { "provinceName": "Kars" },
    { "provinceName": "Kastamonu" },
    { "provinceName": "Kayseri" },
    { "provinceName": "Kırıkkale" },
    { "provinceName": "Kırklareli" },
    { "provinceName": "Kırşehir" },
    { "provinceName": "Kilis" },
    { "provinceName": "Kocaeli" },
    { "provinceName": "Konya" },
    { "provinceName": "Kütahya" },
    { "provinceName": "Malatya" },
    { "provinceName": "Manisa" },
    { "provinceName": "Mardin" },
    { "provinceName": "Mersin" },
    { "provinceName": "Muğla" },
    { "provinceName": "Muş" },
    { "provinceName": "Nevşehir" },
    { "provinceName": "Niğde" },
    { "provinceName": "Ordu" },
    { "provinceName": "Osmaniye" },
    { "provinceName": "Rize" },
    { "provinceName": "Sakarya" },
    { "provinceName": "Samsun" },
    { "provinceName": "Siirt" },
    { "provinceName": "Sinop" },
    { "provinceName": "Sivas" },
    { "provinceName": "Şanlıurfa" },
    { "provinceName": "Şırnak" },
    { "provinceName": "Tekirdağ" },
    { "provinceName": "Tokat" },
    { "provinceName": "Trabzon" },
    { "provinceName": "Tunceli" },
    { "provinceName": "Uşak" },
    { "provinceName": "Van" },
    { "provinceName": "Yalova" },
    { "provinceName": "Yozgat" },
    { "provinceName": "Zonguldak" }
] 
``` 
### `/provinces/{provinceName}`

| Param | Type | Description |
| -- | -- | -- |
| provinceName | string | The name of the province that you would like to get information. |

<br/>

This endpoint returns a specific province `given by parameter.`

### Sample Request 2

```http
http://localhost:8080/provinces/edirne
```

### Sample Response 2

```json
{ "provinceName": "Edirne" }
```

### `/counties`

This endpoint returns all counties and their names that Turkey has.

### Sample Request 3

```http
http://localhost:8080/counties
```

### Sample Response 3

Responded data bigger is than the sample.

```json
[
    {
        "provinceName": "Adana",
        "provinceCounties": [
            { "countyName": "Aladağ" },
            { "countyName": "Ceyhan" },
            { "countyName": "Feke" },
            { "countyName": "Karaisalı" },
            { "countyName": "Karataş" },
            { "countyName": "Kozan" },
            { "countyName": "Pozantı" },
            { "countyName": "Saimbeyli" },
            { "countyName": "Sarıçam" },
            { "countyName": "Seyhan" },
            { "countyName": "Tufanbeyli" },
            { "countyName": "Yumurtalık" },
            { "countyName": "Yüreğir" },
            { "countyName": "Çukurova" },
            { "countyName": "İmamoğlu" }
        ]
    },
]
```


### `/counties/{provinceName}`

| Param | Type | Description |
| -- | -- | -- |
| provinceName | string | The name of the county that you would like to get information. |

### Sample Request 4

```http
http://localhost:8080/counties/istanbul
```

### Sample Response 4

```json
[
    { "countyName": "Adalar" },
    { "countyName": "Arnavutköy" },
    { "countyName": "Ataşehir" },
    { "countyName": "Avcılar" },
    { "countyName": "Bahçelievler" },
    { "countyName": "Bakırköy" },
    { "countyName": "Bayrampaşa" },
    { "countyName": "Bağcılar" },
    { "countyName": "Başakşehir" },
    { "countyName": "Beykoz" },
    { "countyName": "Beylikdüzü" },
    { "countyName": "Beyoğlu" },
    { "countyName": "Beşiktaş" },
    { "countyName": "Büyükçekmece" },
    { "countyName": "Esenler" },
    { "countyName": "Esenyurt" },
    { "countyName": "Eyüpsultan" },
    { "countyName": "Fatih" },
    { "countyName": "Gaziosmanpaşa" },
    { "countyName": "Güngören" },
    { "countyName": "Kadıköy" },
    { "countyName": "Kartal" },
    { "countyName": "Kağıthane" },
    { "countyName": "Küçükçekmece" },
    { "countyName": "Maltepe" },
    { "countyName": "Pendik" },
    { "countyName": "Sancaktepe" },
    { "countyName": "Sarıyer" },
    { "countyName": "Silivri" },
    { "countyName": "Sultanbeyli" },
    { "countyName": "Sultangazi" },
    { "countyName": "Tuzla" },
    { "countyName": "Zeytinburnu" },
    { "countyName": "Çatalca" },
    { "countyName": "Çekmeköy" },
    { "countyName": "Ümraniye" },
    { "countyName": "Üsküdar" },
    { "countyName": "Şile" },
    { "countyName": "Şişli" }
]
``` 

### /towns

### /towns/{provinceName}

### /towns/{provinceName}/{countyName}

### /villages

### /villages/{provinceName}

### /villages/{provinceName}/{countyName}

### /villages/{provinceName}/{countyName}/{townName}


List of libraries used in tzone which written by me personally.
