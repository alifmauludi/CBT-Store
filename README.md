## Show all product

**localhost:8080/api/v1/products**

```javascript
{
    "meta": {
        "message": "List of products",
        "code": 200,
        "status": "success",
        "total": 100
    },
    "data": [
        {
            "id": 24,
            "name": "Pigtail SC-MM2",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp23.000,00",
            "discount": "Rp0,00",
            "photos": "localhost:8080/photos/474A80FF-294D-408C-9569-D63869C6403E.jpeg",
            "rank": 883,
            "updated_date": "2022-03-19T17:47:59+07:00"
        },
        {
            "id": 27,
            "name": "Pigtail LC OM3 1 meter",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp23.000,00",
            "discount": "Rp0,00",
            "photos": "",
            "rank": 158,
            "updated_date": "2022-03-19T17:47:59+07:00"
        },
        {
            "id": 28,
            "name": "Patch Cord LC to LC MM OM3 1 meter",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp85.000,00",
            "discount": "Rp0,00",
            "photos": "localhost:8080/photos/lc_lc_om3-removebg-preview.png",
            "rank": 461,
            "updated_date": "2022-03-19T17:47:59+07:00"
        },
        ......
       {
            "id": 143,
            "name": "Panduit NETKEY OTB 12 Port SC SM NKFD1W12BUDSCZ",
            "category": "OTB OPTICAL TERMINATION BOX",
            "brand": "PANDUIT",
            "price": "Rp8.755.890,00",
            "discount": "Rp0,00",
            "photos": "localhost:8080/photos/NKFD1W12BUDSCZ.jpg",
            "rank": 114,
            "updated_date": "2022-03-19T17:47:59+07:00"
        }
    ]
}        
```

**Parameter:**
|Input Parameter|Type    |           |
|---------------|--------|-----------|
|limit          |(int)   |Optional   |
|page           |(int)   |Optional   |
|category       |(string)|Optional   |
|brand          |(string)|Optional   |

**localhost:8080/api/v1/products?limit=2&page=2**

```javascript
{
    "meta": {
        "message": "List of products",
        "code": 200,
        "status": "success",
        "total": 2
    },
    "data": [
        {
            "id": 28,
            "name": "Patch Cord LC to LC MM OM3 1 meter",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp85.000,00",
            "discount": "Rp0,00",
            "photos": "localhost:8080/photos/lc_lc_om3-removebg-preview.png",
            "rank": 461,
            "updated_date": "2022-03-19T17:47:59+07:00"
        },
        {
            "id": 29,
            "name": "Patch Cord LC to LC MM OM3 2 meter",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp125.000,00",
            "discount": "Rp0,00",
            "photos": "localhost:8080/photos/lc_lc_om3-removebg-preview.png",
            "rank": 423,
            "updated_date": "2022-03-19T17:47:59+07:00"
        }
    ]
}
```

**localhost:8080/api/v1/products?category=patch-cord-fiber-optic**

```javascript
{
    "meta": {
        "message": "List of products",
        "code": 200,
        "status": "success",
        "total": 100
    },
    "data": [
        {
            "id": 24,
            "name": "Pigtail SC-MM2",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp23.000,00",
            "discount": "Rp0,00",
            "photos": "localhost:8080/photos/474A80FF-294D-408C-9569-D63869C6403E.jpeg",
            "rank": 883,
            "updated_date": "2022-03-19T17:47:59+07:00"
        },
        {
            "id": 27,
            "name": "Pigtail LC OM3 1 meter",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp23.000,00",
            "discount": "Rp0,00",
            "photos": "",
            "rank": 158,
            "updated_date": "2022-03-19T17:47:59+07:00"
        },
        {
            "id": 28,
            "name": "Patch Cord LC to LC MM OM3 1 meter",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp85.000,00",
            "discount": "Rp0,00",
            "photos": "localhost:8080/photos/lc_lc_om3-removebg-preview.png",
            "rank": 461,
            "updated_date": "2022-03-19T17:47:59+07:00"
        },
        ......
        {
            "id": 187,
            "name": "Netviel Patch Cord MTRJ-ST, Duplex Multimode 62,5/125um, 1 meter NVL-DPMM1-MTRJ-ST-001",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "NETVIEL",
            "price": "Rp377.000,00",
            "discount": "Rp0,00",
            "photos": "localhost:8080/photos/patch-cord-multimode-mtrj-mtrj-om1-om2.png",
            "rank": 157,
            "updated_date": "2022-03-19T17:47:59+07:00"
        }
    ]
}                
```

**localhost:8080/api/v1/products?category=patch-cord-fiber-optic&brand=cbt-net**

```javascript
{
    "meta": {
        "message": "List of products",
        "code": 200,
        "status": "success",
        "total": 31
    },
    "data": [
        {
            "id": 24,
            "name": "Pigtail SC-MM2",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp23.000,00",
            "discount": "Rp0,00",
            "photos": "localhost:8080/photos/474A80FF-294D-408C-9569-D63869C6403E.jpeg",
            "rank": 883,
            "updated_date": "2022-03-19T17:47:59+07:00"
        },
        {
            "id": 27,
            "name": "Pigtail LC OM3 1 meter",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp23.000,00",
            "discount": "Rp0,00",
            "photos": "",
            "rank": 158,
            "updated_date": "2022-03-19T17:47:59+07:00"
        },
        ......
        {
            "id": 57,
            "name": "Patch Cord SC to LC SM 20 meter",
            "category": "PATCH CORD FIBER OPTIC",
            "brand": "CBT-NET",
            "price": "Rp180.000,00",
            "discount": "Rp0,00",
            "photos": "localhost:8080/photos/sc_lc__1_-removebg-preview.png",
            "rank": 34,
            "updated_date": "2022-03-19T17:47:59+07:00"
        }
    ]
}        
```

## Show detail product

```
localhost:8080/api/v1/product/(id)
```

**localhost:8080/api/v1/product/57**

```javascript
{
    "meta": {
        "message": "Product detail",
        "code": 200,
        "status": "success",
        "total": 1
    },
    "data": {
        "id": 57,
        "name": "Patch Cord SC to LC SM 20 meter",
        "category": "PATCH CORD FIBER OPTIC",
        "brand": "CBT-NET",
        "spec": "Patch Cord SC-LC, Duplex Singlemode 9/125um, 20 meter (LSZH)\r\n",
        "price": "Rp180.000,00",
        "discount": "Rp0,00",
        "rank": 34,
        "updated_date": "2022-03-19T17:47:59+07:00",
        "photos": [
            {
                "file_name": "localhost:8080/photos/sc_lc__1_-removebg-preview.png",
                "is_cover": true
            }
        ]
    }
}
```