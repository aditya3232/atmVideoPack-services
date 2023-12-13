
# API Spec AtmVideoPackApp-Services

## 1 Auth

### 1.1 Login

Request :
- Method : POST
- URL : `{{local}}:3636/api/atmvideopack/v1/auth/login`
- Body (form-data) :
    - username : string, required
    - password : string, required
- Response :

```json 
{
    "meta": {
        "message": "Login berhasil",
        "code": 200
    },
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDUwNjcxNDcsInVzZXIiOnsiaWQiOjM3fX0.rtAkHtDN6YIwaIQ5fSv-lpCKgHQUvo7HFU1AHYMhZKQ",
        "expires": "2024-01-12T20:45:47.8116865+07:00"
    }
}
```
### 1.2 Logout

Request :
- Method : POST
- URL : `{{local}}:3636/api/atmvideopack/v1/auth/logout`
- Header : 
    - Authorization : string
- Response :

```json 
{
    "meta": {
        "message": "Logout berhasil",
        "code": 200
    },
    "data": null
}
```

## 2 Profile

### 2.1 Get Profile

Request :
- Method : GET
- URL : `{{local}}:3636/api/atmvideopack/v1/profile/getprofile`
- Header : 
    - Authorization : string
- Response :

```json 
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": {
        "user": {
            "id": 37,
            "role_id": 28,
            "name": "Maiden",
            "username": "maiden",
            "foto_profil": "",
            "created_at": "2023-12-13T12:02:57+07:00",
            "updated_at": "2023-12-13T20:45:48+07:00"
        },
        "role": {
            "id": 28,
            "name": "admin atm video pack",
            "created_at": "2023-10-27T02:06:42+07:00",
            "updated_at": "2023-10-27T02:06:42+07:00"
        },
        "permissions": [
            {
                "id": 36,
                "name": "sidebar parent dashboard",
                "created_at": "2023-10-27T02:03:17+07:00",
                "updated_at": "2023-10-27T02:03:17+07:00"
            },
            {
                "id": 37,
                "name": "sidebar parent setting admin",
                "created_at": "2023-10-27T02:03:30+07:00",
                "updated_at": "2023-10-27T02:03:30+07:00"
            },
            {
                "id": 39,
                "name": "sidebar parent human detection",
                "created_at": "2023-10-27T02:03:51+07:00",
                "updated_at": "2023-10-27T02:03:51+07:00"
            },
            {
                "id": 40,
                "name": "sidebar parent vandal detection",
                "created_at": "2023-10-27T02:03:59+07:00",
                "updated_at": "2023-10-27T02:03:59+07:00"
            },
            {
                "id": 41,
                "name": "sidebar parent streaming cctv",
                "created_at": "2023-10-27T02:04:09+07:00",
                "updated_at": "2023-10-27T02:04:09+07:00"
            },
            {
                "id": 42,
                "name": "sidebar parent download playback",
                "created_at": "2023-10-27T02:04:18+07:00",
                "updated_at": "2023-10-27T02:04:18+07:00"
            },
            {
                "id": 43,
                "name": "sidebar parent master device",
                "created_at": "2023-10-27T02:04:29+07:00",
                "updated_at": "2023-10-27T02:04:29+07:00"
            },
            {
                "id": 44,
                "name": "sidebar parent master location",
                "created_at": "2023-10-27T02:04:41+07:00",
                "updated_at": "2023-10-27T02:04:41+07:00"
            }
        ]
    }
}
```

## 3 Roles

### 3.1 Get One

Request :
- Method : GET
- URL : `{{local}}:3636/api/atmvideopack/v1/roles/getone/30`
- Header :
    - Authorization : string
- Response :
    
```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": {
        "id": 30,
        "name": "admin vandal detection",
        "created_at": "2023-11-08T07:38:12+07:00",
        "updated_at": "2023-11-08T01:56:37+07:00"
    }
}
```

### 3.2 Get All

Request :
- Method : GET
- URL : `{{local}}:3636/api/atmvideopack/v1/roles/getall
- Header :
    - Authorization : string
- Params :
    - limit 
    - page
    - sort
    - order
    - all column for search
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "pagination": {
        "page": 1,
        "limit": 10,
        "total": 5,
        "total_filtered": 5
    },
    "data": [
        {
            "id": 30,
            "name": "admin vandal detection",
            "created_at": "2023-11-08T07:38:12+07:00",
            "updated_at": "2023-11-08T01:56:37+07:00"
        },
        {
            "id": 29,
            "name": "admin human detection",
            "created_at": "2023-11-03T03:05:19+07:00",
            "updated_at": "2023-11-03T03:05:19+07:00"
        },
        {
            "id": 28,
            "name": "admin atm video pack",
            "created_at": "2023-10-27T02:06:42+07:00",
            "updated_at": "2023-10-27T02:06:42+07:00"
        },
        {
            "id": 14,
            "name": "Fresh",
            "created_at": "2023-07-18T01:58:55+07:00",
            "updated_at": "2023-07-18T02:03:03+07:00"
        },
        {
            "id": 2,
            "name": "admin",
            "created_at": "2023-05-30T01:57:46+07:00",
            "updated_at": "2023-05-30T01:57:46+07:00"
        }
    ]
}
```

### 3.3 Create

Request :
- Method : POST
- URL : `{{local}}:3636/api/atmvideopack/v1/roles/create`
- Header :
    - Authorization : string
- Body (form-data) :
    - name : string, required
- Response :

```json
{
    "meta": {
        "message": "Successfully created new data.",
        "code": 201
    },
    "data": {
        "id": 32,
        "name": "tes",
        "created_at": "2023-12-13T20:55:23.9219505+07:00"
    }
}
```

### 3.4 Update

Request :
- Method : PUT
- URL : `{{local}}:3636/api/atmvideopack/v1/roles/update/32`
- Header :
    - Authorization : string
- Body (form-data) :
    - name : string, required
- Response :

```json
{
    "meta": {
        "message": "Successfully updated data.",
        "code": 200
    },
    "data": {
        "id": 32,
        "name": "tes",
        "created_at": "2023-12-13T20:55:23.9219505+07:00",
        "updated_at": "2023-12-13T20:55:23.9219505+07:00"
    }
}
```

### 3.5 Delete

Request :
- Method : DELETE
- URL : `{{local}}:3636/api/atmvideopack/v1/roles/delete/32`
- Header :
    - Authorization : string
- Response : Status 204 No Content

## 4 Permissions

### 4.1 Get One
Request :
- Method : GET
- URL : `{{local}}:3636/api/atmvideopack/v1/permissions/getone/45`
- Response : 

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": {
        "id": 45,
        "name": "sidebar child permission",
        "created_at": "2023-10-27T02:04:49+07:00",
        "updated_at": "2023-10-27T02:04:49+07:00"
    }
}
```

### 4.2 Get All
Request :
- Method : GET
- URL : `{{local}}:3636/api/atmvideopack/v1/permissions/getall`
- Params :
    - limit 
    - page
    - sort
    - order
    - all column for search
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "pagination": {
        "page": 1,
        "limit": 10,
        "total": 10,
        "total_filtered": 10
    },
    "data": [
        {
            "id": 45,
            "name": "sidebar child permission",
            "created_at": "2023-10-27T02:04:49+07:00",
            "updated_at": "2023-10-27T02:04:49+07:00"
        },
        {
            "id": 44,
            "name": "sidebar parent master location",
            "created_at": "2023-10-27T02:04:41+07:00",
            "updated_at": "2023-10-27T02:04:41+07:00"
        },
        {
            "id": 43,
            "name": "sidebar parent master device",
            "created_at": "2023-10-27T02:04:29+07:00",
            "updated_at": "2023-10-27T02:04:29+07:00"
        },
        {
            "id": 42,
            "name": "sidebar parent download playback",
            "created_at": "2023-10-27T02:04:18+07:00",
            "updated_at": "2023-10-27T02:04:18+07:00"
        },
        {
            "id": 41,
            "name": "sidebar parent streaming cctv",
            "created_at": "2023-10-27T02:04:09+07:00",
            "updated_at": "2023-10-27T02:04:09+07:00"
        },
        {
            "id": 40,
            "name": "sidebar parent vandal detection",
            "created_at": "2023-10-27T02:03:59+07:00",
            "updated_at": "2023-10-27T02:03:59+07:00"
        },
        {
            "id": 39,
            "name": "sidebar parent human detection",
            "created_at": "2023-10-27T02:03:51+07:00",
            "updated_at": "2023-10-27T02:03:51+07:00"
        },
        {
            "id": 38,
            "name": "sidebar parent help",
            "created_at": "2023-10-27T02:03:39+07:00",
            "updated_at": "2023-10-27T02:03:39+07:00"
        },
        {
            "id": 37,
            "name": "sidebar parent setting admin",
            "created_at": "2023-10-27T02:03:30+07:00",
            "updated_at": "2023-10-27T02:03:30+07:00"
        },
        {
            "id": 36,
            "name": "sidebar parent dashboard",
            "created_at": "2023-10-27T02:03:17+07:00",
            "updated_at": "2023-10-27T02:03:17+07:00"
        }
    ]
}
```

### 4.3 Create

Request :
- Method : POST
- URL : `{{local}}:3636/api/atmvideopack/v1/permissions/create`
- Body (form-data) :
    - name : string, required
- Response :

```json
{
    "meta": {
        "message": "Successfully created new data.",
        "code": 201
    },
    "data": {
        "id": 46,
        "name": "tes",
        "created_at": "2023-12-13T21:00:27.9219505+07:00"
    }
}
```

### 4.4 Update

Request :
- Method : PUT
- URL : `{{local}}:3636/api/atmvideopack/v1/permissions/update/46`
- Body (form-data) :
    - name : string, required
- Response :

```json
{
    "meta": {
        "message": "Successfully updated data.",
        "code": 200
    },
    "data": {
        "id": 46,
        "name": "tes",
        "created_at": "2023-12-13T21:00:27.9219505+07:00",
        "updated_at": "2023-12-13T21:00:27.9219505+07:00"
    }
}
```

### 4.5 Delete

Request :
- Method : DELETE
- URL : `{{local}}:3636/api/atmvideopack/v1/permissions/delete/46`
- Response : Status 204 No Content

## 5 Users

### 5.1 Get One

Request :
- Method : GET
- URL : `{{local}}:3636/api/atmvideopack/v1/users/getone/35`
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": {
        "id": 35,
        "role_id": 14,
        "name": "Dita",
        "username": "dita32",
        "foto_profil": "",
        "created_at": "2023-12-12T09:58:43+07:00",
        "updated_at": null
    }
}
```

### 5.2 Get All

Request :
- Method : GET
- URL : `{{local}}:3636/api/atmvideopack/v1/users/getall`
- Params :
    - limit 
    - page
    - sort
    - order
    - all column for search
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "pagination": {
        "page": 1,
        "limit": 10,
        "total": 5,
        "total_filtered": 5
    },
    "data": [
        {
            "id": 35,
            "role_id": 14,
            "name": "Dita",
            "username": "dita32",
            "foto_profil": "",
            "created_at": "2023-12-12T09:58:43+07:00",
            "updated_at": null
        },
        {
            "id": 34,
            "role_id": 14,
            "name": "Dita",
            "username": "dita31",
            "foto_profil": "",
            "created_at": "2023-12-12T09:58:43+07:00",
            "updated_at": null
        },
        {
            "id": 33,
            "role_id": 14,
            "name": "Dita",
            "username": "dita30",
            "foto_profil": "",
            "created_at": "2023-12-12T09:58:43+07:00",
            "updated_at": null
        },
        {
            "id": 32,
            "role_id": 14,
            "name": "Dita",
            "username": "dita29",
            "foto_profil": "",
            "created_at": "2023-12-12T09:58:43+07:00",
            "updated_at": null
        },
        {
            "id": 31,
            "role_id": 14,
            "name": "Dita",
            "username": "dita28",
            "foto_profil": "",
            "created_at": "2023-12-12T09:58:43+07:00",
            "updated_at": null
        }
    ]
}
```

### 5.3 Create

Request :
- Method : POST
- URL : `{{local}}:3636/api/atmvideopack/v1/users/create`
- Body (form-data) :
    - role_id : int, required
    - name : string, required
    - username : string, required
    - password : string, required
    - foto_profil : string
- Response :

```json
{
    "meta": {
        "message": "Successfully created new data.",
        "code": 201
    },
    "data": {
        "id": 36,
        "role_id": 14,
        "name": "Dita",
        "username": "dita33",
        "foto_profil": "",
        "created_at": "2023-12-13T21:05:27.9219505+07:00"
    }
}
```

### 5.4 Update

Request :
- Method : PUT
- URL : `{{local}}:3636/api/atmvideopack/v1/users/update/36`
- Body (form-data) :
    - role_id : int, required
    - name : string, required
    - username : string, required
    - password : string, required
    - foto_profil : string
- Response :

```json
{
    "meta": {
        "message": "Successfully updated data.",
        "code": 200
    },
    "data": {
        "id": 36,
        "role_id": 14,
        "name": "Dita",
        "username": "dita33",
        "foto_profil": "",
        "created_at": "2023-12-13T21:05:27.9219505+07:00",
        "updated_at": "2023-12-13T21:05:27.9219505+07:00"
    }
}
```

### 5.5 Delete

Request :
- Method : DELETE
- URL : `{{local}}:3636/api/atmvideopack/v1/users/delete/36`
- Response : Status 204 No Content

## 6 Human Detection

### 6.1 Publisher Human Detection

Request :
- Method : POST
- URL : `{{local}}:3333/publisher/atmvideopack/v1/humandetection/create`
- Body (form-data) :
    - tid : string, required
    - date_time : string, required
    - person : string, required
    - file_name_capture_human_detection : string, required
- Response : 

```json
{
    "meta": {
        "message": "The request was processed successfully.",
        "code": 200
    },
    "data": {
        "tid": "160001",
        "date_time": "2023-12-13 08:00:00",
        "person": "0",
        "file_capture_human_detection": "Screenshot_2.jpg"
    }
}
```

### 6.2 Get All

Request :
- Method : GET
- URL : `{{local}}:3636/api/atmvideopack/v1/humandetection/getall`
- Body (form-data) :
    - id
    - tid_id
    - date_time
    - start_date
    - end_date
    - person
    - file_name_capture_human_detection
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": [
        {
            "id": "H6LaW4wB1iYbJPNCmksI",
            "tid": "160001",
            "date_time": "2023-12-04 08:00:00",
            "person": "0",
            "file_name_capture_human_detection": "2023-12-12-09-28-21-999.jpg"
        }
    ]
}
```

## 7 Vandal Detection

### 7.1 Publisher Vandal Detection

Request :
- Method : POST
- URL : `{{local}}:3434/publisher/atmvideopack/v1/vandaldetection/create`
- Body (form-data) :
    - tid : string, required
    - date_time : string, required
    - person : string, required
    - file_capture_vandal_detection : string, required
- Response :

```json
{
    "meta": {
        "message": "The request was processed successfully.",
        "code": 200
    },
    "data": {
        "tid": "160001",
        "date_time": "2023-12-13 08:00:00",
        "person": "0",
        "file_capture_vandal_detection": "Screenshot_2.jpg"
    }
}
```

### 7.2 Get All

Request :
- Method : GET
- URL : `{{local}}:3636/api/atmvideopack/v1/vandaldetection/getall`
- Body (form-data) :
    - id
    - tid_id
    - date_time
    - start_date
    - end_date
    - person
    - file_name_capture_vandal_detection
- Response :

```json
{
    "meta": {
        "message": "Data found.",
        "code": 200
    },
    "data": [
        {
            "id": "IKLgW4wB1iYbJPNCjkuk",
            "tid": "160001",
            "date_time": "2023-12-04 08:00:00",
            "person": "1",
            "file_name_capture_vandal_detection": "2023-12-12-09-34-52-306.jpg"
        }
    ]
}
```







