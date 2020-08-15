<div align="center"> 

`tzone` is a rest API that wrote in `golang` and it serves address information for Turkey

[![CodeFactor](https://www.codefactor.io/repository/github/enesusta/tzone/badge?style=for-the-badge)](https://www.codefactor.io/repository/github/enesusta/tzone)
[![Go Report Card](https://goreportcard.com/badge/github.com/enesusta/tzone?style=for-the-badge)](https://goreportcard.com/report/github.com/enesusta/tzone)

</div>
<br/>

- Index
  - [Install](#-install) 
  - [Endpoints](#endpoints)
    - [/provinces](#provinces)
    - [/provinces/{provinceName}](#provinces/{provinceName})
    - [/counties](#counties)
    - [/counties/{provinceName}](#countiesprovincename)
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

## Endpoints

### /provinces

### /provinces/{provinceName}

### /counties

### /counties/{provinceName}

### /towns

### /towns/{provinceName}

### /towns/{provinceName}/{countyName}

### /villages

### /villages/{provinceName}

### /villages/{provinceName}/{countyName}

### /villages/{provinceName}/{countyName}/{townName}

I have needed a rest API that serves the address location to my ERP(enterprise resource planning) application. After searching I didn't find a rest API that fits for me and my requirements.

PTT already has presented [data that contains address information for Turkey](https://postakodu.ptt.gov.tr/). But the problem is: the data is not relational and its file format is xlsx(Microsoft Excel Open XML Spreadsheet) that doesn't fit to rest API.

What do I mean by '`it doesn't fit to rest API`' ?

Let's look at the raw data that has provided by PTT.


## Raw data before parsing

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

List of libraries used in tzone which written by me personally.
