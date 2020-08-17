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
  - [How you can consume this API?](#how-you-can-consume-this-api)
  - [CORS Options](#cors-options)
  - [Endpoints](#endpoints)
    - [/](#)
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
      - [Sample Request](#sample-request-5)
      - [Sample Response](#sample-response-5)
    - [/towns/{provinceName}](#townsprovincename)
      - [Sample Request](#sample-request-6)
      - [Sample Response](#sample-response-6)
    - [/towns/{provinceName}/{countyName}](#townsprovincenamecountyname)
      - [Sample Request](#sample-request-7)
      - [Sample Response](#sample-response-7)
    - [/villages](#villages)
    - [/villages/{provinceName}](#villagesprovincename)
      - [Sample Request](#sample-request-8)
      - [Sample Response](#sample-response-8)
    - [/villages/{provinceName}/{countyName}](#villagesprovincenamecountyname)
      - [Sample Request](#sample-request-9)
      - [Sample Response](#sample-response-9)
    - [/villages/{provinceName}/{countyName}/{townName}](#villagesprovincenamecountynametownname)
      - [Sample Request](#sample-request-10)
      - [Sample Response](#sample-response-10)
  - [Licence](#licence)
    - All code found in this repository is licensed under **CC BY-NC-ND**


![](https://raw.githubusercontent.com/enesusta/assets-host-for-github-pages/assets/tzone/tzone-5.gif)

## 🕺 Install

tzone runs on `12071 HTTP` port.

### Locally

If you're running tzone on your machine, after instruction that listed below go to your browser and type:
http://localhost:12071/provinces

```
go get -u github.com/enesusta/tzone
cd $GOPATH/src/github/enesusta/tzone
go build
./tzone.exe
```

### With Docker

> If you're not familiar with docker, you can skip this section.

tzone has already been containerized and has deployed to hub.docker.com. Check that [link](https://hub.docker.com/repository/docker/enesusta/tzone)

You can mapping the ports whatever you want. The most important thing is tzone runs on `12071 HTTP port`. You must consider this during configure your mapping.

-  With CLI

```bash
docker run --name tzone -d -p 8080:12071 enesusta/tzone:1.2
```

-  With docker-compose

Let's look at this sample.

`docker-compose.yml`
```yml
version: '3'

services:
  tzone:
    container_name: tzone
    image: enesusta/tzone:1.2
    ports:
      - 8080:12071
```

Then:

```bash
docker-compose up -d
```


## 📯 Preface

I have needed a rest API that serves the address location to my ERP(enterprise resource planning) application. After searching I didn't find a rest API that fits for me and my requirements.

PTT already has presented [data that contains address information for Turkey](https://postakodu.ptt.gov.tr/). But the problem is: the data is not relational and its file format is xlsx(Microsoft Excel Open XML Spreadsheet) that doesn't fit to rest API.

To better understand what this application does let's look at this picture that shows `how much record exists in data` that have provided by PTT.

![](https://raw.githubusercontent.com/enesusta/assets-host-for-github-pages/assets/tzone/tzone-1.png)


> This illustration illustrates how tzone works.

![](https://raw.githubusercontent.com/enesusta/assets-host-for-github-pages/assets/tzone/tzone2.svg)

What do I mean by '`it doesn't fit to rest API`' ?

Let's look at the raw data that have provided by PTT.


## Raw Data Before Parsing

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

## The data parsed by `tzone-parser`

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

### How you can consume this API?

There is a React application that demonstrates how you can do this. It has written by me. 

You can check it from [this link](https://github.com/enesusta/tzone-react)


#### CORS Options

Tzone has configuration that listed below.

If would you like to configure tzone to your server/website. Change AllowedOrigins value.

```go
c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
})
```

For example:

```go
c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://enesusta.tech"}, // Only accessible from enesusta.tech
		AllowedMethods: []string{"GET"}, // Allowing only get, just an example
})
```

## Endpoints

### `/`

This endpoint return a data that README.md has.

### `/provinces`

> This endpoint returns all provinces and their names that Turkey has.

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

> This endpoint returns a specific province `by the given parameter.`

### Sample Request 2

```http
http://localhost:8080/provinces/edirne
```

### Sample Response 2

```json
{ "provinceName": "Edirne" }
```

### `/counties`

> This endpoint returns all counties and their names that Turkey has.

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
    /**.
      .
      .
      .
      .
      other provinces and countries that turkey has.
    */
]
```


### `/counties/{provinceName}`

| Param | Type | Description |
| -- | -- | -- |
| provinceName | string | The name of the province that you would like to get information about its counties. |

> This endpoint returns all counties and their names by the given parameter..

#### Sample Request 4

```http
http://localhost:8080/counties/istanbul
```

#### Sample Response 4

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

### `/towns`

> This endpoint returns all towns and their names that Turkey has.

#### Sample Request 5

```http
http://localhost:8080/towns
```

#### Sample Response 5
```json
[
   {
        "provinceName": "Edirne",
        "provinceCounties": [
            {
                "countyName": "Enez",
                "countyTowns": [
                    { "townName": "Enez" },
                    { "townName": "Merkezköyler" }
                ]
            },
            {
                "countyName": "Havsa",
                "countyTowns": [
                    { "townName": "Hasköy" },
                    { "townName": "Havsa" },
                    { "townName": "Merkezköyler" }
                ]
            },
            {
                "countyName": "Keşan",
                "countyTowns": [
                    { "townName": "Keşan" },
                    { "townName": "Mecidiye" },
                    { "townName": "Merkezköyler" },
                    { "townName": "Yerlisu" }
                ]
            },
            {
                "countyName": "Lalapaşa",
                "countyTowns": [
                    { "townName": "Lalapaşa" },
                    { "townName": "Merkezköyler" }
                ]
            },
            {
                "countyName": "Meriç",
                "countyTowns": [
                    { "townName": "Küplü" },
                    { "townName": "Meriç" },
                    { "townName": "Merkezköyler" }
                ]
            },
            {
                "countyName": "Merkez",
                "countyTowns": [
                    { "townName": "Edirne" },
                    { "townName": "Merkezköyler" }
                ]
            },
            {
                "countyName": "Süloğlu",
                "countyTowns": [
                    { "townName": "Süloğlu" }
                ]
            },
            {
                "countyName": "Uzunköprü",
                "countyTowns": [
                    { "townName": "Aşçıoğlu" },
                    { "townName": "Hamidiye" },
                    { "townName": "Kırcasalih" },
                    { "townName": "Merkezköyler" },
                    { "townName": "Uzunköprü" },
                    { "townName": "Çöpköy" }
                ]
            },
            {
                "countyName": "İpsala",
                "countyTowns": [
                    { "townName": "Merkezköyler" },
                    { "townName": "İpsala" }
                ]
            }
        ]
    },
    /**
         .
         .
         .  
         Other provinces and their datas.
         */
]
```


### `/towns/{provinceName}`

| Param | Type | Description |
| -- | -- | -- |
| provinceName | string | The name of the province that you would like to get information about its counties and towns. |

> This endpoint returns all towns and their names by the given parameter..

#### Sample Request 6

```http
http://localhost:8080/towns/bolu
```

#### Sample Response 6

```json
[
    {
        "countyName": "Dörtdivan",
        "countyTowns": [
            { "townName": "Dörtdivan" },
            { "townName": "Merkezköyler" }
        ]
    },
    {
        "countyName": "Gerede",
        "countyTowns": [
            { "townName": "Gerede" },
            { "townName": "Merkezköyler" }
        ]
    },
    {
        "countyName": "Göynük",
        "countyTowns": [
            { "townName": "Göynük" },
            { "townName": "Merkezköyler" }
        ]
    },
    {
        "countyName": "Kıbrıscık",
        "countyTowns": [
            { "townName": "Kıbrıscık" },
            { "townName": "Merkezköyler" }
        ]
    },
    {
        "countyName": "Mengen",
        "countyTowns": [
            { "townName": "Gökçesu" },
            { "townName": "Mengen" },
            { "townName": "Merkezköyler" }
        ]
    },
    {
        "countyName": "Merkez",
        "countyTowns": [
            { "townName": "Aşağısoku" },
            { "townName": "Karacasu" },
            { "townName": "Karamanlı" },
            { "townName": "Merkezköyler" },
            { "townName": "İhsaniye" }
        ]
    },
    {
        "countyName": "Mudurnu",
        "countyTowns": [
            { "townName": "Mudurnu" }
        ]
    },
    {
        "countyName": "Seben",
        "countyTowns": [
            { "townName": "Merkezköyler" },
            { "townName": "Seben" }
        ]
    },
    {
        "countyName": "Yeniçağa",
        "countyTowns": [
            { "townName": "Merkezköyler" },
            { "townName": "Yeniçağa" }
        ]
    }
]
```

### `/towns/{provinceName}/{countyName}`

| Param | Type | Description |
| -- | -- | -- |
| provinceName | string | The name of the province that you would like to get information about its counties and towns. |
| countyName | string | The name of the county that you would ike to get information about its towns |

#### Sample Request 7

```http
http://localhost:8080/towns/istanbul/adalar
```

#### Sample Response 7

```json
{
    "countyName": "Adalar",
    "countyTowns": [
        { "townName": "Burgazada" },
        { "townName": "Büyükada" },
        { "townName": "Heybeliada" },
        { "townName": "Kınalıada" }
    ]
}
```

### `/villages`

This endpoint returns all villages with their name and their zip-code that Turkey has.

### `/villages/{provinceName}`

| Param | Type | Description |
| -- | -- | -- |
| provinceName | string | The name of the province that you would like to get information about its counties, towns and villages. |


> This endpoint returns all villages with their name and their zip-code `by the given parameter.`

#### Sample Request 8

```http
http://localhost:8080/villages/edirne
```

#### Sample Response 8

```json
{
    "provinceName": "Edirne",
    "provinceCounties": [
        {
            "countyName": "Enez",
            "countyTowns": [
                {
                    "townName": "Enez",
                    "townVillages": [
                        { "villageName": "Gaziömerbey mah", "zipCode": 22700 },
                        { "villageName": "Yeni mah", "zipCode": 22700 },
                        { "villageName": "Çataltepe mah", "zipCode": 22700 }
                    ]
                },
                {
                    "townName": "Merkezköyler",
                    "townVillages": [
                        { "villageName": "Abdurrahim köyü", "zipCode": 22750 },
                        { "villageName": "Büyükevren köyü", "zipCode": 22750 },
                        { "villageName": "Gülçavuş köyü", "zipCode": 22750 },
                        { "villageName": "Hasköy köyü", "zipCode": 22750 },
                        { "villageName": "Hisarlı köyü", "zipCode": 22750 },
                        { "villageName": "Işıklı köyü", "zipCode": 22750 },
                        { "villageName": "Karaincirli köyü", "zipCode": 22750 },
                        { "villageName": "Kocaali köyü", "zipCode": 22750 },
                        { "villageName": "Küçükevren köyü", "zipCode": 22750 },
                        { "villageName": "Sultaniçe köyü", "zipCode": 22750 },
                        { "villageName": "Sütçüler köyü", "zipCode": 22750 },
                        { "villageName": "Umurbey köyü", "zipCode": 22750 },
                        { "villageName": "Vakıf köyü", "zipCode": 22750 },
                        { "villageName": "Yazır köyü", "zipCode": 22750 },
                        { "villageName": "Yenice köyü", "zipCode": 22750 },
                        { "villageName": "Çandır köyü", "zipCode": 22750 },
                        { "villageName": "Çavuşköy köyü", "zipCode": 22750 },
                        { "villageName": "Çeribaşı köyü", "zipCode": 22750 },
                        { "villageName": "Şehitler köyü", "zipCode": 22750 }
                    ]
                }
            ]
        },
        {
            "countyName": "Havsa",
            "countyTowns": [
                {
                    "townName": "Hasköy",
                    "townVillages": [
                        { "villageName": "Arpaç köyü", "zipCode": 22530 },
                        { "villageName": "Habiller köyü", "zipCode": 22530 },
                        { "villageName": "Hasköy köyü", "zipCode": 22530 },
                        { "villageName": "Musulca köyü", "zipCode": 22530 },
                        { "villageName": "Söğütlüdere köyü", "zipCode": 22530 }
                    ]
                },
                {
                    "townName": "Havsa",
                    "townVillages": [
                        { "villageName": "Cumhuriyet mah", "zipCode": 22500 },
                        { "villageName": "Hacıgazi mah", "zipCode": 22500 },
                        { "villageName": "Hacıisa mah", "zipCode": 22500 },
                        { "villageName": "Helvacı mah", "zipCode": 22500 },
                        { "villageName": "Varoş mah", "zipCode": 22500 },
                        { "villageName": "Yeni mah", "zipCode": 22500 }
                    ]
                },
                {
                    "townName": "Merkezköyler",
                    "townVillages": [
                        { "villageName": "Abalar mah (abalar köyü)", "zipCode": 22532 },
                        { "villageName": "Abalar merkez mah (abalar köyü)", "zipCode": 22532 },
                        { "villageName": "Azatlı köyü", "zipCode": 22532 },
                        { "villageName": "Bakışlar köyü", "zipCode": 22532 },
                        { "villageName": "Bostanlı köyü", "zipCode": 22532 },
                        { "villageName": "Kabaağaç köyü", "zipCode": 22532 },
                        { "villageName": "Kulubalık köyü", "zipCode": 22532 },
                        { "villageName": "Kuzucu köyü", "zipCode": 22532 },
                        { "villageName": "Köseömer köyü", "zipCode": 22532 },
                        { "villageName": "Naipyusuf köyü", "zipCode": 22532 },
                        { "villageName": "Necatiye köyü", "zipCode": 22532 },
                        { "villageName": "Osmanlı köyü", "zipCode": 22532 },
                        { "villageName": "Oğulpaşa köyü", "zipCode": 22532 },
                        { "villageName": "Oğulpaşa mah (oğulpaşa köyü)", "zipCode": 22532 },
                        { "villageName": "Tahal köyü", "zipCode": 22532 },
                        { "villageName": "Taptık köyü", "zipCode": 22532 },
                        { "villageName": "Yolageldi köyü", "zipCode": 22532 },
                        { "villageName": "Çukurköy köyü", "zipCode": 22532 },
                        { "villageName": "Şerbettar köyü", "zipCode": 22532 }
                    ]
                }
            ]
        },
        {
            "countyName": "Keşan",
            "countyTowns": [
                {
                    "townName": "Keşan",
                    "townVillages": [
                        { "villageName": "Aşağı zaferiye mah", "zipCode": 22800 },
                        { "villageName": "Büyük cami mah", "zipCode": 22800 },
                        { "villageName": "Cumhuriyet mah", "zipCode": 22800 },
                        { "villageName": "Mustafa kemal paşa mah", "zipCode": 22800 },
                        { "villageName": "Yeni mah", "zipCode": 22800 },
                        { "villageName": "Yenimescit mah", "zipCode": 22800 },
                        { "villageName": "Yukarı zaferiye mah", "zipCode": 22800 },
                        { "villageName": "İspat cami mah", "zipCode": 22800 },
                        { "villageName": "İstasyon mah", "zipCode": 22800 }
                    ]
                },
                {
                    "townName": "Mecidiye",
                    "townVillages": [
                        { "villageName": "Akhoca köyü", "zipCode": 22410 },
                        { "villageName": "Atatürk mah (mecidiye köyü)", "zipCode": 22410 },
                        { "villageName": "Barağı köyü", "zipCode": 22410 },
                        { "villageName": "Beyköy köyü", "zipCode": 22410 },
                        { "villageName": "Cumhuriyet mah (mecidiye köyü)", "zipCode": 22410 },
                        { "villageName": "Danişment köyü", "zipCode": 22410 },
                        { "villageName": "Danişment mah (danişment köyü)", "zipCode": 22410 },
                        { "villageName": "Dişbudak köyü", "zipCode": 22410 },
                        { "villageName": "Erikli köyü", "zipCode": 22410 },
                        { "villageName": "Erikli mah (erikli köyü)", "zipCode": 22410 },
                        { "villageName": "Koruklu köyü", "zipCode": 22410 },
                        { "villageName": "Kızkapan köyü", "zipCode": 22410 },
                        { "villageName": "Orhaniye köyü", "zipCode": 22410 },
                        { "villageName": "Suluca köyü", "zipCode": 22410 },
                        { "villageName": "Yayla mah (yaylaköy köyü)", "zipCode": 22410 },
                        { "villageName": "Yaylaköy köyü", "zipCode": 22410 },
                        { "villageName": "Yeşilköy köyü", "zipCode": 22410 },
                        { "villageName": "Çelebi köyü", "zipCode": 22410 },
                        { "villageName": "Çeltik köyü", "zipCode": 22410 },
                        { "villageName": "Şabanmera köyü", "zipCode": 22410 }
                    ]
                },
                {
                    "townName": "Merkezköyler",
                    "townVillages": [
                        { "villageName": "Akçeşme köyü", "zipCode": 22880 },
                        { "villageName": "Altıntaş köyü", "zipCode": 22880 },
                        { "villageName": "Atatürk mah (yenimuhacir köyü)", "zipCode": 22880 },
                        { "villageName": "Aydoğan mah (beğendik köyü)", "zipCode": 22880 },
                        { "villageName": "Beykent mah (beğendik köyü)", "zipCode": 22880 },
                        { "villageName": "Boztepe köyü", "zipCode": 22880 },
                        { "villageName": "Cumhuriyet mah (yenimuhacir köyü)", "zipCode": 22880 },
                        { "villageName": "Gündüzler köyü", "zipCode": 22880 },
                        { "villageName": "Karacaali köyü", "zipCode": 22880 },
                        { "villageName": "Karahisar köyü", "zipCode": 22880 },
                        { "villageName": "Karasatı köyü", "zipCode": 22880 },
                        { "villageName": "Kemalpaşa mah (beğendik köyü)", "zipCode": 22880 },
                        { "villageName": "Kozköy köyü", "zipCode": 22880 },
                        { "villageName": "Küçükdoğanca köyü", "zipCode": 22880 },
                        { "villageName": "Kılıçköy köyü", "zipCode": 22880 },
                        { "villageName": "Lalacık köyü", "zipCode": 22880 },
                        { "villageName": "Maltepe köyü", "zipCode": 22880 },
                        { "villageName": "Paşayiğit mah", "zipCode": 22880 },
                        { "villageName": "Siğilli köyü", "zipCode": 22880 },
                        { "villageName": "Türkmen köyü", "zipCode": 22880 },
                        { "villageName": "Yeni mah (beğendik köyü)", "zipCode": 22880 },
                        { "villageName": "Yeniceçiftlik köyü", "zipCode": 22880 },
                        { "villageName": "Çobançeşmesi köyü", "zipCode": 22880 },
                        { "villageName": "İnönü mah (yenimuhacir köyü)", "zipCode": 22880 },
                        { "villageName": "İstiklal mah (yenimuhacir köyü)", "zipCode": 22880 },
                        { "villageName": "İzzetiye mah", "zipCode": 22880 }
                    ]
                },
                {
                    "townName": "Yerlisu",
                    "townVillages": [
                        { "villageName": "Atatürk mah (çamlıca köyü)", "zipCode": 22272 },
                        { "villageName": "Bahçeköy köyü", "zipCode": 22272 },
                        { "villageName": "Büyükdoğanca köyü", "zipCode": 22272 },
                        { "villageName": "Fevzipaşa mah (çamlıca köyü)", "zipCode": 22272 },
                        { "villageName": "Gökçetepe köyü", "zipCode": 22272 },
                        { "villageName": "Gökçetepe mah (gökçetepe köyü)", "zipCode": 22272 },
                        { "villageName": "Kadıköy köyü", "zipCode": 22272 },
                        { "villageName": "Karlı köyü", "zipCode": 22272 },
                        { "villageName": "Mahmutköy köyü", "zipCode": 22272 },
                        { "villageName": "Mercan köyü", "zipCode": 22272 },
                        { "villageName": "Pırnar köyü", "zipCode": 22272 },
                        { "villageName": "Sazlıdere köyü", "zipCode": 22272 },
                        { "villageName": "Sazlıdere mah (sazlıdere köyü)", "zipCode": 22272 },
                        { "villageName": "Seydiköy köyü", "zipCode": 22272 },
                        { "villageName": "Yerlisu köyü", "zipCode": 22272 },
                        { "villageName": "Şükrüköy köyü", "zipCode": 22272 }
                    ]
                }
            ]
        },
        {
            "countyName": "Lalapaşa",
            "countyTowns": [
                {
                    "townName": "Lalapaşa",
                    "townVillages": [
                        { "villageName": "Merkez mah", "zipCode": 22950 }
                    ]
                },
                {
                    "townName": "Merkezköyler",
                    "townVillages": [
                        { "villageName": "Büyünlü köyü", "zipCode": 22970 },
                        { "villageName": "Demirköy köyü", "zipCode": 22970 },
                        { "villageName": "Dombay köyü", "zipCode": 22970 },
                        { "villageName": "Doğanköy köyü", "zipCode": 22970 },
                        { "villageName": "Hacıdanişment köyü", "zipCode": 22970 },
                        { "villageName": "Hacılar köyü", "zipCode": 22970 },
                        { "villageName": "Hamzabeyli köyü", "zipCode": 22970 },
                        { "villageName": "Hanlıyenice köyü", "zipCode": 22970 },
                        { "villageName": "Hüseyinpınar köyü", "zipCode": 22970 },
                        { "villageName": "Kalkansöğüt köyü", "zipCode": 22970 },
                        { "villageName": "Kavaklı köyü", "zipCode": 22970 },
                        { "villageName": "Küçünlü köyü", "zipCode": 22970 },
                        { "villageName": "Ortakçı köyü", "zipCode": 22970 },
                        { "villageName": "Saksağan köyü", "zipCode": 22970 },
                        { "villageName": "Sarıdanişment köyü", "zipCode": 22970 },
                        { "villageName": "Sinanköy köyü", "zipCode": 22970 },
                        { "villageName": "Süleymandanişment köyü", "zipCode": 22970 },
                        { "villageName": "Taşlımüsellim köyü", "zipCode": 22970 },
                        { "villageName": "Tuğlalık köyü", "zipCode": 22970 },
                        { "villageName": "Uzunbayır köyü", "zipCode": 22970 },
                        { "villageName": "Vaysal köyü", "zipCode": 22970 },
                        { "villageName": "Yünlüce köyü", "zipCode": 22970 },
                        { "villageName": "Çallıdere köyü", "zipCode": 22970 },
                        { "villageName": "Çatma köyü", "zipCode": 22970 },
                        { "villageName": "Çömlek köyü", "zipCode": 22970 },
                        { "villageName": "Çömlekakpınar köyü", "zipCode": 22970 },
                        { "villageName": "Ömeroba köyü", "zipCode": 22970 }
                    ]
                }
            ]
        },
        {
            "countyName": "Meriç",
            "countyTowns": [
                {
                    "townName": "Küplü",
                    "townVillages": [
                        { "villageName": "Adasarhanlı köyü", "zipCode": 22610 },
                        { "villageName": "Büyükaltıağaç köyü", "zipCode": 22610 },
                        { "villageName": "Gazi mah (küplü köyü)", "zipCode": 22610 },
                        { "villageName": "Küçükaltıağaç köyü", "zipCode": 22610 },
                        { "villageName": "Merkez mah (küplü köyü)", "zipCode": 22610 },
                        { "villageName": "İzzet paşa mah (subaşı köyü)", "zipCode": 22610 },
                        { "villageName": "Şehit mehmet aydın mah (subaşı köyü)", "zipCode": 22610 }
                    ]
                },
                {
                    "townName": "Meriç",
                    "townVillages": [
                        { "villageName": "B.doğanca mah", "zipCode": 22600 },
                        { "villageName": "K. doğanca mah", "zipCode": 22600 }
                    ]
                },
                {
                    "townName": "Merkezköyler",
                    "townVillages": [
                        { "villageName": "Akçadam köyü", "zipCode": 22680 },
                        { "villageName": "Akıncılar köyü", "zipCode": 22680 },
                        { "villageName": "Alibey köyü", "zipCode": 22680 },
                        { "villageName": "Hasırcıarnavutköy köyü", "zipCode": 22680 },
                        { "villageName": "Kadıdondurma köyü", "zipCode": 22680 },
                        { "villageName": "Karahamza köyü", "zipCode": 22680 },
                        { "villageName": "Karayusuflu köyü", "zipCode": 22680 },
                        { "villageName": "Kavaklı köyü", "zipCode": 22680 },
                        { "villageName": "Küpdere köyü", "zipCode": 22680 },
                        { "villageName": "Nasuhbey köyü", "zipCode": 22680 },
                        { "villageName": "Olacak köyü", "zipCode": 22680 },
                        { "villageName": "Paşayenice köyü", "zipCode": 22680 },
                        { "villageName": "Rahmanca köyü", "zipCode": 22680 },
                        { "villageName": "Saatağacı köyü", "zipCode": 22680 },
                        { "villageName": "Serem köyü", "zipCode": 22680 },
                        { "villageName": "Umurca köyü", "zipCode": 22680 },
                        { "villageName": "Yakupbey köyü", "zipCode": 22680 },
                        { "villageName": "Yenicegörüce köyü", "zipCode": 22680 }
                    ]
                }
            ]
        },
        {
            "countyName": "Merkez",
            "countyTowns": [
                {
                    "townName": "Edirne",
                    "townVillages": [
                        { "villageName": "1.murat mah", "zipCode": 22100 },
                        { "villageName": "Abdurrahman mah", "zipCode": 22100 },
                        { "villageName": "Babademirtaş mah", "zipCode": 22100 },
                        { "villageName": "Barutluk mah", "zipCode": 22100 },
                        { "villageName": "Dilaverbey mah", "zipCode": 22100 },
                        { "villageName": "Fatih mah", "zipCode": 22100 },
                        { "villageName": "Karaağaç mah", "zipCode": 22100 },
                        { "villageName": "Koca sinan mah", "zipCode": 22100 },
                        { "villageName": "Medrese ali bey mah", "zipCode": 22100 },
                        { "villageName": "Menzilahır mah", "zipCode": 22100 },
                        { "villageName": "Meydan mah", "zipCode": 22100 },
                        { "villageName": "Mithat paşa mah", "zipCode": 22100 },
                        { "villageName": "Nişancıpaşa mah", "zipCode": 22100 },
                        { "villageName": "Sabuni mah", "zipCode": 22100 },
                        { "villageName": "Sarıcapaşa mah", "zipCode": 22100 },
                        { "villageName": "Talatpaşa mah", "zipCode": 22100 },
                        { "villageName": "Umurbey mah", "zipCode": 22100 },
                        { "villageName": "Yancıkçı şahin mah", "zipCode": 22100 },
                        { "villageName": "Yeniimaret mah", "zipCode": 22100 },
                        { "villageName": "Yıldırım beyazıt mah", "zipCode": 22100 },
                        { "villageName": "Yıldırım hacı sarraf mah", "zipCode": 22100 },
                        { "villageName": "Çavuşbey mah", "zipCode": 22100 },
                        { "villageName": "İstasyon mah", "zipCode": 22100 },
                        { "villageName": "Şükrüpaşa mah", "zipCode": 22100 }
                    ]
                },
                {
                    "townName": "Merkezköyler",
                    "townVillages": [
                        { "villageName": "Ahı köyü", "zipCode": 22130 },
                        { "villageName": "Avarız köyü", "zipCode": 22130 },
                        { "villageName": "Bosna köyü", "zipCode": 22130 },
                        { "villageName": "Budakdoğanca köyü", "zipCode": 22130 },
                        { "villageName": "Büyükdöllük köyü", "zipCode": 22130 },
                        { "villageName": "Büyükismailçe köyü", "zipCode": 22130 },
                        { "villageName": "Demirhanlı köyü", "zipCode": 22130 },
                        { "villageName": "Değirmenyeni köyü", "zipCode": 22130 },
                        { "villageName": "Doyran köyü", "zipCode": 22130 },
                        { "villageName": "Ekmekçi köyü", "zipCode": 22130 },
                        { "villageName": "Elçili köyü", "zipCode": 22130 },
                        { "villageName": "Eskikadın köyü", "zipCode": 22130 },
                        { "villageName": "Hacıumur köyü", "zipCode": 22130 },
                        { "villageName": "Hasanağa köyü", "zipCode": 22130 },
                        { "villageName": "Hatip köyü", "zipCode": 22130 },
                        { "villageName": "Hıdırağa köyü", "zipCode": 22130 },
                        { "villageName": "Karabulut köyü", "zipCode": 22130 },
                        { "villageName": "Karakasım köyü", "zipCode": 22130 },
                        { "villageName": "Karayusuf köyü", "zipCode": 22130 },
                        { "villageName": "Kayapa köyü", "zipCode": 22130 },
                        { "villageName": "Kemal köyü", "zipCode": 22130 },
                        { "villageName": "Korucu köyü", "zipCode": 22130 },
                        { "villageName": "Köşençiftliği köyü", "zipCode": 22130 },
                        { "villageName": "Küçükdöllük köyü", "zipCode": 22130 },
                        { "villageName": "Menekşesofular köyü", "zipCode": 22130 },
                        { "villageName": "Muratçalı köyü", "zipCode": 22130 },
                        { "villageName": "Musabeyli köyü", "zipCode": 22130 },
                        { "villageName": "Orhaniye köyü", "zipCode": 22130 },
                        { "villageName": "Sarayakpınar köyü", "zipCode": 22130 },
                        { "villageName": "Sazlıdere köyü", "zipCode": 22130 },
                        { "villageName": "Suakacağı köyü", "zipCode": 22130 },
                        { "villageName": "Tayakadın köyü", "zipCode": 22130 },
                        { "villageName": "Uzgaç köyü", "zipCode": 22130 },
                        { "villageName": "Yenikadın köyü", "zipCode": 22130 },
                        { "villageName": "Yolüstü köyü", "zipCode": 22130 },
                        { "villageName": "Üyüklütatar köyü", "zipCode": 22130 },
                        { "villageName": "İskender köyü", "zipCode": 22130 }
                    ]
                }
            ]
        },
        {
            "countyName": "Süloğlu",
            "countyTowns": [
                {
                    "townName": "Merkezköyler",
                    "townVillages": [
                        { "villageName": "Akardere köyü", "zipCode": 22580 },
                        { "villageName": "Büyük gerdelli köyü", "zipCode": 22580 },
                        { "villageName": "Domurcalı köyü", "zipCode": 22580 },
                        { "villageName": "Geçkinli köyü", "zipCode": 22580 },
                        { "villageName": "Keramettin köyü", "zipCode": 22580 },
                        { "villageName": "Küküler köyü", "zipCode": 22580 },
                        { "villageName": "Sülecik köyü", "zipCode": 22580 },
                        { "villageName": "Tatarlar köyü", "zipCode": 22580 },
                        { "villageName": "Taşlısekban köyü", "zipCode": 22580 },
                        { "villageName": "Yağcılı köyü", "zipCode": 22580 }
                    ]
                },
                {
                    "townName": "Süloğlu",
                    "townVillages": [
                        { "villageName": "Cumhuriyet mah", "zipCode": 22560 },
                        { "villageName": "Merkez mah", "zipCode": 22560 },
                        { "villageName": "Organize sanayi bölgesi mah", "zipCode": 22560 },
                        { "villageName": "Yeni mah", "zipCode": 22560 },
                        { "villageName": "Şafak mah", "zipCode": 22560 }
                    ]
                }
            ]
        },
        {
            "countyName": "Uzunköprü",
            "countyTowns": [
                {
                    "townName": "Aşçıoğlu",
                    "townVillages": [
                        { "villageName": "Atatürk mah", "zipCode": 22300 },
                        { "villageName": "Aşçıoğlu mah", "zipCode": 22300 },
                        { "villageName": "Demirtaş mah", "zipCode": 22300 },
                        { "villageName": "Kavak mah", "zipCode": 22300 },
                        { "villageName": "Muradiye mah", "zipCode": 22300 } ]
                },
                {
                    "townName": "Hamidiye",
                    "townVillages": [
                        { "villageName": "Altınyazı köyü", "zipCode": 22402 },
                        { "villageName": "Alıç köyü", "zipCode": 22402 },
                        { "villageName": "Atatürk mah (kurtbey köyü)", "zipCode": 22402 },
                        { "villageName": "Balaban köyü", "zipCode": 22402 },
                        { "villageName": "Balabankoru köyü", "zipCode": 22402 },
                        { "villageName": "Dereköy köyü", "zipCode": 22402 },
                        { "villageName": "Fevzi çakmak mah (kurtbey köyü)", "zipCode": 22402 },
                        { "villageName": "Hamidiye köyü", "zipCode": 22402 },
                        { "villageName": "Harmanlı köyü", "zipCode": 22402 },
                        { "villageName": "Hasanpınar köyü", "zipCode": 22402 },
                        { "villageName": "Kadıağılı köyü", "zipCode": 22402 },
                        { "villageName": "Kadıköy köyü", "zipCode": 22402 },
                        { "villageName": "Karapınar köyü", "zipCode": 22402 },
                        { "villageName": "Kavakayazma köyü", "zipCode": 22402 },
                        { "villageName": "Kırköy köyü", "zipCode": 22402 },
                        { "villageName": "Maksutlu köyü", "zipCode": 22402 },
                        { "villageName": "Namık kemal mah (kurtbey köyü)", "zipCode": 22402 },
                        { "villageName": "Süleymaniye köyü", "zipCode": 22402 },
                        { "villageName": "Türkobası köyü", "zipCode": 22402 },
                        { "villageName": "Çavuşlu köyü", "zipCode": 22402 },
                        { "villageName": "Çobanpınar köyü", "zipCode": 22402 },
                        { "villageName": "İnönü mah (kurtbey köyü)", "zipCode": 22402 }
                    ]
                },
                {
                    "townName": "Kırcasalih",
                    "townVillages": [
                        { "villageName": "Aslıhan köyü", "zipCode": 22260 },
                        { "villageName": "Balaban mah (kırcasalih köyü)", "zipCode": 22260 },
                        { "villageName": "Cumhuriyet mah (kırcasalih köyü)", "zipCode": 22260 },
                        { "villageName": "Meşeli köyü", "zipCode": 22260 },
                        { "villageName": "Okullar mah (kırcasalih köyü)", "zipCode": 22260 },
                        { "villageName": "Sazlımalkoç köyü", "zipCode": 22260 },
                        { "villageName": "İstiklal mah (kırcasalih köyü)", "zipCode": 22260 }
                    ]
                },
                {
                    "townName": "Merkezköyler",
                    "townVillages": [
                        { "villageName": "Değirmenci köyü", "zipCode": 22360 },
                        { "villageName": "Eskiköy köyü", "zipCode": 22360 },
                        { "villageName": "Gemici köyü", "zipCode": 22360 },
                        { "villageName": "Hamitli köyü", "zipCode": 22360 },
                        { "villageName": "Karayayla köyü", "zipCode": 22360 },
                        { "villageName": "Kavacık köyü", "zipCode": 22360 },
                        { "villageName": "Kiremitçisalih köyü", "zipCode": 22360 },
                        { "villageName": "Kurdu köyü", "zipCode": 22360 },
                        { "villageName": "Kurttepe köyü", "zipCode": 22360 },
                        { "villageName": "Kırkkavak köyü", "zipCode": 22360 },
                        { "villageName": "Malkoç köyü", "zipCode": 22360 },
                        { "villageName": "Salarlı köyü", "zipCode": 22360 },
                        { "villageName": "Saçlımüsellim köyü", "zipCode": 22360 },
                        { "villageName": "Sığırcılı köyü", "zipCode": 22360 },
                        { "villageName": "Çakmak köyü", "zipCode": 22360 },
                        { "villageName": "Çalı köyü", "zipCode": 22360 },
                        { "villageName": "Çiftlik köyü", "zipCode": 22360 }
                    ]
                },
                {
                    "townName": "Uzunköprü",
                    "townVillages": [
                        { "villageName": "Büyük şehsuvarbey mah", "zipCode": 22200 },
                        { "villageName": "Cumhuriyet mah", "zipCode": 22200 },
                        { "villageName": "Habib hoca mah", "zipCode": 22200 },
                        { "villageName": "Halise hatun mah", "zipCode": 22200 },
                        { "villageName": "Küçük şehsuvarbey mah", "zipCode": 22200 },
                        { "villageName": "Mareşal fevzi çakmak mah", "zipCode": 22200 },
                        { "villageName": "Mescit mah", "zipCode": 22200 },
                        { "villageName": "Muradiye cami mah", "zipCode": 22200 },
                        { "villageName": "Rızaefendi mah", "zipCode": 22200 },
                        { "villageName": "Yeniköy mah", "zipCode": 22200 },
                        { "villageName": "Çöpköy mah", "zipCode": 22200 }
                    ]
                },
                {
                    "townName": "Çöpköy",
                    "townVillages": [
                        { "villageName": "Bayramlı köyü", "zipCode": 22270 },
                        { "villageName": "Başağıl köyü", "zipCode": 22270 },
                        { "villageName": "Beykonak köyü", "zipCode": 22270 },
                        { "villageName": "Bıldır köyü", "zipCode": 22270 },
                        { "villageName": "Danişment köyü", "zipCode": 22270 },
                        { "villageName": "Elmalı köyü", "zipCode": 22270 },
                        { "villageName": "Gazimehmet köyü", "zipCode": 22270 },
                        { "villageName": "Karabürçek köyü", "zipCode": 22270 },
                        { "villageName": "Muhacırkadı köyü", "zipCode": 22270 },
                        { "villageName": "Sipahi köyü", "zipCode": 22270 },
                        { "villageName": "Sultanşah köyü", "zipCode": 22270 },
                        { "villageName": "Turnacı köyü", "zipCode": 22270 },
                        { "villageName": "Yağmurca köyü", "zipCode": 22270 },
                        { "villageName": "Ömerbey köyü", "zipCode": 22270 }
                    ]
                }
            ]
        },
        {
            "countyName": "İpsala",
            "countyTowns": [
                {
                    "townName": "Merkezköyler",
                    "townVillages": [
                        { "villageName": "Aliçopehlivan köyü", "zipCode": 22490 },
                        { "villageName": "Atatürk mah (esetçe köyü)", "zipCode": 22490 },
                        { "villageName": "Atatürk mah (kocahıdır köyü)", "zipCode": 22490 },
                        { "villageName": "Atatürk mah (sultan köyü)", "zipCode": 22490 },
                        { "villageName": "Balabancık köyü", "zipCode": 22490 },
                        { "villageName": "Cumhuriyet mah (esetçe köyü)", "zipCode": 22490 },
                        { "villageName": "Cumhuriyet mah (kocahıdır köyü)", "zipCode": 22490 },
                        { "villageName": "Fevzi çakmak mah (yenikarpuzlu köyü)", "zipCode": 22490 },
                        { "villageName": "Gazi mah (yenikarpuzlu köyü)", "zipCode": 22490 },
                        { "villageName": "Gemici mah (ibriktepe köyü)", "zipCode": 22490 },
                        { "villageName": "Gündoğan mah (hacı köyü)", "zipCode": 22490 },
                        { "villageName": "Hamidiye mah (ibriktepe köyü)", "zipCode": 22490 },
                        { "villageName": "Hıdırköy köyü", "zipCode": 22490 },
                        { "villageName": "Karaağaç köyü", "zipCode": 22490 },
                        { "villageName": "Korucu köyü", "zipCode": 22490 },
                        { "villageName": "Koyuntepe köyü", "zipCode": 22490 },
                        { "villageName": "Kumdere köyü", "zipCode": 22490 },
                        { "villageName": "Kurtuluş mah (esetçe köyü)", "zipCode": 22490 },
                        { "villageName": "Küçükdoğanca köyü", "zipCode": 22490 },
                        { "villageName": "Menderes mah (sultan köyü)", "zipCode": 22490 },
                        { "villageName": "Pazardere köyü", "zipCode": 22490 },
                        { "villageName": "Paşaköy köyü", "zipCode": 22490 },
                        { "villageName": "Sarpdere köyü", "zipCode": 22490 },
                        { "villageName": "Sarıcaali köyü", "zipCode": 22490 },
                        { "villageName": "Selanik mah (hacı köyü)", "zipCode": 22490 },
                        { "villageName": "Selanik mah (ibriktepe köyü)", "zipCode": 22490 },
                        { "villageName": "Tevfikiye köyü", "zipCode": 22490 },
                        { "villageName": "Turpçular köyü", "zipCode": 22490 },
                        { "villageName": "Yapıldak köyü", "zipCode": 22490 },
                        { "villageName": "Zaferiye mah (hacı köyü)", "zipCode": 22490 },
                        { "villageName": "İnönü mah (ibriktepe köyü)", "zipCode": 22490 },
                        { "villageName": "İnönü mah (yenikarpuzlu köyü)", "zipCode": 22490 }
                    ]
                },
                {
                    "townName": "İpsala",
                    "townVillages": [
                        { "villageName": "Bayrambey mah", "zipCode": 22400 },
                        { "villageName": "Bozkurt mah", "zipCode": 22400 },
                        { "villageName": "Fatih mah", "zipCode": 22400 },
                        { "villageName": "Kapucu mah", "zipCode": 22400 },
                        { "villageName": "Köprü mah", "zipCode": 22400 },
                        { "villageName": "Saraç ilyas mah", "zipCode": 22400 }
                    ]
                }
            ]
        }
    ]
}
```

### `/villages/{provinceName}/{countyName}`

| Param | Type | Description |
| -- | -- | -- |
| provinceName | string | The name of the province that you would like to get information about its counties, towns and villages. |
| countyName | string | The name of the county that you would like to get information about its towns and villages. |


> This endpoint returns all villages with their name and their zip-code `by the given parameter.`

#### Sample Request 9

```http
http://localhost:8080/villages/edirne/merkez
```

#### Sample Response 9

```json
{
    "countyName": "Merkez",
    "countyTowns": [
        {
            "townName": "Edirne",
            "townVillages": [
                { "villageName": "1.murat mah", "zipCode": 22100 },
                { "villageName": "Abdurrahman mah", "zipCode": 22100 },
                { "villageName": "Babademirtaş mah", "zipCode": 22100 },
                { "villageName": "Barutluk mah", "zipCode": 22100 },
                { "villageName": "Dilaverbey mah", "zipCode": 22100 },
                { "villageName": "Fatih mah", "zipCode": 22100 },
                { "villageName": "Karaağaç mah", "zipCode": 22100 },
                { "villageName": "Koca sinan mah", "zipCode": 22100 },
                { "villageName": "Medrese ali bey mah", "zipCode": 22100 },
                { "villageName": "Menzilahır mah", "zipCode": 22100 },
                { "villageName": "Meydan mah", "zipCode": 22100 },
                { "villageName": "Mithat paşa mah", "zipCode": 22100 },
                { "villageName": "Nişancıpaşa mah", "zipCode": 22100 },
                { "villageName": "Sabuni mah", "zipCode": 22100 },
                { "villageName": "Sarıcapaşa mah", "zipCode": 22100 },
                { "villageName": "Talatpaşa mah", "zipCode": 22100 },
                { "villageName": "Umurbey mah", "zipCode": 22100 },
                { "villageName": "Yancıkçı şahin mah", "zipCode": 22100 },
                { "villageName": "Yeniimaret mah", "zipCode": 22100 },
                { "villageName": "Yıldırım beyazıt mah", "zipCode": 22100 },
                { "villageName": "Yıldırım hacı sarraf mah", "zipCode": 22100 },
                { "villageName": "Çavuşbey mah", "zipCode": 22100 },
                { "villageName": "İstasyon mah", "zipCode": 22100 },
                { "villageName": "Şükrüpaşa mah", "zipCode": 22100 }
            ]
        },
        {
            "townName": "Merkezköyler",
            "townVillages": [
                { "villageName": "Ahı köyü", "zipCode": 22130 },
                { "villageName": "Avarız köyü", "zipCode": 22130 },
                { "villageName": "Bosna köyü", "zipCode": 22130 },
                { "villageName": "Budakdoğanca köyü", "zipCode": 22130 },
                { "villageName": "Büyükdöllük köyü", "zipCode": 22130 },
                { "villageName": "Büyükismailçe köyü", "zipCode": 22130 },
                { "villageName": "Demirhanlı köyü", "zipCode": 22130 },
                { "villageName": "Değirmenyeni köyü", "zipCode": 22130 },
                { "villageName": "Doyran köyü", "zipCode": 22130 },
                { "villageName": "Ekmekçi köyü", "zipCode": 22130 },
                { "villageName": "Elçili köyü", "zipCode": 22130 },
                { "villageName": "Eskikadın köyü", "zipCode": 22130 },
                { "villageName": "Hacıumur köyü", "zipCode": 22130 },
                { "villageName": "Hasanağa köyü", "zipCode": 22130 },
                { "villageName": "Hatip köyü", "zipCode": 22130 },
                { "villageName": "Hıdırağa köyü", "zipCode": 22130 },
                { "villageName": "Karabulut köyü", "zipCode": 22130 },
                { "villageName": "Karakasım köyü", "zipCode": 22130 },
                { "villageName": "Karayusuf köyü", "zipCode": 22130 },
                { "villageName": "Kayapa köyü", "zipCode": 22130 },
                { "villageName": "Kemal köyü", "zipCode": 22130 },
                { "villageName": "Korucu köyü", "zipCode": 22130 },
                { "villageName": "Köşençiftliği köyü", "zipCode": 22130 },
                { "villageName": "Küçükdöllük köyü", "zipCode": 22130 },
                { "villageName": "Menekşesofular köyü", "zipCode": 22130 },
                { "villageName": "Muratçalı köyü", "zipCode": 22130 },
                { "villageName": "Musabeyli köyü", "zipCode": 22130 },
                { "villageName": "Orhaniye köyü", "zipCode": 22130 },
                { "villageName": "Sarayakpınar köyü", "zipCode": 22130 },
                { "villageName": "Sazlıdere köyü", "zipCode": 22130 },
                { "villageName": "Suakacağı köyü", "zipCode": 22130 },
                { "villageName": "Tayakadın köyü", "zipCode": 22130 },
                { "villageName": "Uzgaç köyü", "zipCode": 22130 },
                { "villageName": "Yenikadın köyü", "zipCode": 22130 },
                { "villageName": "Yolüstü köyü", "zipCode": 22130 },
                { "villageName": "Üyüklütatar köyü", "zipCode": 22130 },
                { "villageName": "İskender köyü", "zipCode": 22130 }
            ]
        }
    ]
}
```


### `/villages/{provinceName}/{countyName}/{townName}`

| Param | Type | Description |
| -- | -- | -- |
| provinceName | string | The name of the province that you would like to get information about its counties, towns and villages. |
| countyName | string | The name of the county that you would like to get information about its towns and villages. |
| townName | string | The name of the town that you would like to get information about its villages |

> This endpoint returns all villages with their name and their zip-code `by the given parameter.`

#### Sample Request 10

```http
http://localhost:8080/villages/edirne/merkez/edirne
```

#### Sample Response 10

```json
{
    "townName": "Edirne",
    "townVillages": [
        { "villageName": "1.murat mah", "zipCode": 22100 },
        { "villageName": "Abdurrahman mah", "zipCode": 22100 },
        { "villageName": "Babademirtaş mah", "zipCode": 22100 },
        { "villageName": "Barutluk mah", "zipCode": 22100 },
        { "villageName": "Dilaverbey mah", "zipCode": 22100 },
        { "villageName": "Fatih mah", "zipCode": 22100 },
        { "villageName": "Karaağaç mah", "zipCode": 22100 },
        { "villageName": "Koca sinan mah", "zipCode": 22100 },
        { "villageName": "Medrese ali bey mah", "zipCode": 22100 },
        { "villageName": "Menzilahır mah", "zipCode": 22100 },
        { "villageName": "Meydan mah", "zipCode": 22100 },
        { "villageName": "Mithat paşa mah", "zipCode": 22100 },
        { "villageName": "Nişancıpaşa mah", "zipCode": 22100 },
        { "villageName": "Sabuni mah", "zipCode": 22100 },
        { "villageName": "Sarıcapaşa mah", "zipCode": 22100 },
        { "villageName": "Talatpaşa mah", "zipCode": 22100 },
        { "villageName": "Umurbey mah", "zipCode": 22100 },
        { "villageName": "Yancıkçı şahin mah", "zipCode": 22100 },
        { "villageName": "Yeniimaret mah", "zipCode": 22100 },
        { "villageName": "Yıldırım beyazıt mah", "zipCode": 22100 },
        { "villageName": "Yıldırım hacı sarraf mah", "zipCode": 22100 },
        { "villageName": "Çavuşbey mah", "zipCode": 22100 },
        { "villageName": "İstasyon mah", "zipCode": 22100 },
        { "villageName": "Şükrüpaşa mah", "zipCode": 22100 }
    ]
}
```

## Licence 

All code found in this repository is licensed under CC BY-NC-ND

This license is the most restrictive of our six main licenses, only allowing others to download your works and share them with others as long as they credit you, `but they can’t change them in any way or use them commercially.`

[More Information](https://creativecommons.org/licenses/by-nc-nd/4.0/)

Copyright 2020 Enes Usta