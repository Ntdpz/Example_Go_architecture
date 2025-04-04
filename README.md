# Project Structure (Go)

## Overview
โปรเจกต์นี้ใช้ภาษา Go และมีโครงสร้างที่แบ่งแยกส่วนต่างๆ อย่างชัดเจน เพื่อให้ง่ายต่อการดูแลรักษาและพัฒนา โครงสร้างนี้ถูกออกแบบมาเพื่อให้สามารถปรับปรุงได้ง่ายและสามารถขยายต่อได้ในอนาคต ด้วยการแยกส่วนต่างๆ ของโปรเจกต์ออกเป็นโมดูลที่ทำหน้าที่ต่างกัน

## Folder Structure
```
/project-root
│── /database
│   ├── db.go
│── /cmd
│   ├── main.go
│── /internal
│   ├── /repository
│   │   ├── user_repository.go
│   ├── /handlers
│   │   ├── user_handlers.go
│   ├── /service
│   │   ├── user_service.go
│── /config
│   ├── .env
│   ├── config.go
│── /routers
│   ├── router.go
│── /test
│   ├── test.go
│── /models
│   ├── user.go
│── /middlewares
│   ├── token.go
│── /example_databasse_Query
│   ├── create_table.txt
│   ├── insert_data.txt
│── docker-compose.yaml
│── Dockerfile
│── go.mod
│── go.sum
```


## Folder & File Explanation

### `/database`
- **db.go**: ไฟล์นี้ใช้จัดการการเชื่อมต่อกับฐานข้อมูล เช่น การเปิดและปิดการเชื่อมต่อ รวมถึงการกำหนดค่าและการเตรียมการในการเชื่อมต่อกับฐานข้อมูล เช่น PostgreSQL หรือ MySQL

### `/cmd`
- **main.go**: ไฟล์หลักที่ใช้ในการรันแอปพลิเคชัน เมื่อเรียกใช้งานโปรแกรมจะเริ่มต้นจากที่นี่ ซึ่งจะทำการโหลดการตั้งค่า (configuration), สร้าง router, และเริ่มต้น server

### `/internal`

#### `/repository`
- **user_repository.go**: ชั้นที่ทำหน้าที่เป็น Data Access Layer ซึ่งจะติดต่อกับฐานข้อมูล เช่น การ query ข้อมูลจากฐานข้อมูลและส่งผลลัพธ์กลับมา โดยเฉพาะข้อมูลของผู้ใช้ในกรณีนี้

#### `/handlers`
- **user_handlers.go**: จัดการ HTTP request และ response สำหรับ `User` เช่น การจัดการ API endpoints, การรับคำขอ (request) และส่งคำตอบ (response) กลับในรูปแบบ JSON

#### `/service`
- **user_service.go**: ชั้นที่จัดการกับ Business Logic สำหรับ `User` เช่น การตรวจสอบข้อมูล (validate), การประมวลผลข้อมูล, และการเรียกใช้งาน repository เพื่อจัดการกับข้อมูล

### `/config`
- **.env**: ไฟล์ที่ใช้สำหรับเก็บค่าคอนฟิกต่างๆ เช่น URL ของฐานข้อมูล, API keys, หรือ Secrets ที่ไม่ควรเผยแพร่ในโค้ด

- **config.go**: ไฟล์ที่ใช้ในการโหลดข้อมูลจากไฟล์ `.env` และจัดการกับ environment variables เพื่อให้สามารถใช้งานค่าต่างๆ ในโปรเจกต์ได้อย่างปลอดภัย

### `/routers`
- **router.go**: กำหนดเส้นทาง (routes) ต่างๆ ของ API เช่นการเชื่อมโยง routes ไปยัง handlers ต่างๆ ที่รับผิดชอบในแต่ละ request และ response

### `/test`
- **test.go**: ไฟล์นี้จะเก็บ unit tests หรือ integration tests เพื่อทดสอบฟังก์ชันการทำงานของแอปพลิเคชันให้แน่ใจว่าระบบทำงานได้ตามที่คาดหวัง

### `/models`
- **user.go**: โครงสร้าง (struct) ที่ใช้เก็บข้อมูลของผู้ใช้ เช่น `id`, `username`, `password`, หรือข้อมูลส่วนตัวอื่นๆ ของผู้ใช้

### `/middlewares`
- **token.go**: ใช้ในการจัดการการตรวจสอบ token เช่น การตรวจสอบ JWT (JSON Web Token) ว่าถูกต้องหรือไม่ เพื่อยืนยันตัวตนของผู้ใช้ในแต่ละ request

### `/example_databasse_Query`
- **create_table.txt**: คำสั่ง SQL สำหรับการสร้างตารางในฐานข้อมูล
- **insert_data.txt**: คำสั่ง SQL สำหรับการเพิ่มข้อมูลตัวอย่างลงในฐานข้อมูล

### Docker & Deployment
- **docker-compose.yaml**: ไฟล์นี้จะกำหนดว่าโปรเจกต์จะใช้งาน container อะไรบ้าง เช่น Go app, Database, Redis ฯลฯ โดยกำหนดค่าการตั้งค่าของแต่ละ container ในการใช้งาน
- **Dockerfile**: ใช้สำหรับสร้าง Docker image สำหรับโปรเจกต์ เพื่อให้สามารถนำไปใช้งานใน container ได้

### `go.mod` และ `go.sum`
- **go.mod**: ไฟล์ที่ใช้จัดการ dependencies ของโปรเจกต์ โดยจะระบุว่ามี package ใดบ้างที่โปรเจกต์ต้องการ
- **go.sum**: ไฟล์ที่เก็บค่าความถูกต้องของ dependencies ที่โปรเจกต์ใช้งานเพื่อความปลอดภัย

---

## How to Run

### Run with Go
```sh
# Install dependencies
go mod tidy

# Run the application
go run cmd/main.go
```
### Run with Docker
```sh
# Build and run with Docker
docker-compose up --build
```

# Project Setup Guide

## เริ่มการใช้งานโปรเจค

1. รันโปรเจคโดยใช้ Docker:
```docker-compose up --build```

2. หลังจากที่ Docker รันเสร็จแล้ว ให้เชื่อมต่อกับฐานข้อมูล PostgreSQL ผ่านโปรแกรมจัดการฐานข้อมูล เช่น **PGAdmin4**.

## การเชื่อมต่อฐานข้อมูล

- **Host:** localhost
- **User:** admin
- **Password:** admin
- **Maintenance Database:** gofiber_db

ค่าดังกล่าวสามารถดูได้จากไฟล์ `docker-compose.yml` ในโปรเจกต์ของคุณ.

## การใช้งานฐานข้อมูล

1. เมื่อเชื่อมต่อกับฐานข้อมูลสำเร็จแล้ว ให้ทำการรัน Script Query ของ PostgreSQL ที่อยู่ในโฟลเดอร์ `example_databasse_Query/`:

- **สร้างตาราง:** รัน Script ที่อยู่ในไฟล์ `create_table.txt`
- **เพิ่มข้อมูล:** รัน Script ที่อยู่ในไฟล์ `insert_data.txt`

ใช้คำสั่งใน PGAdmin หรือเครื่องมือ PostgreSQL อื่นๆ เพื่อรัน SQL Queries ที่มีอยู่ในไฟล์เหล่านี้.

---

**หมายเหตุ:** คำแนะนำนี้เป็นขั้นตอนเบื้องต้นในการเริ่มต้นใช้งานโปรเจกต์กับฐานข้อมูล PostgreSQL โดยใช้ Docker. หากมีข้อสงสัยหรือปัญหาในการใช้งาน, กรุณาตรวจสอบไฟล์ `docker-compose.yml` หรือเอกสารเพิ่มเติม.

---

# API Endpoint เพื่อลองทดสอบ

## GET /

### วัตถุประสงค์:
เมื่อผู้ใช้เข้าถึง URL หลัก ("/") ของแอปพลิเคชัน จะได้รับการตอบกลับเป็นข้อความ "Hello Gofiber!" เพื่อทดสอบว่าเซิร์ฟเวอร์ทำงานได้ปกติและตอบกลับ HTTP ได้หรือไม่

### การทำงาน:
เมื่อผู้ใช้เข้า URL นี้, เซิร์ฟเวอร์จะส่งข้อความ "Hello Gofiber!" ไปยังผู้ใช้.

---

## GET /testdb

### วัตถุประสงค์:
ใช้ทดสอบการเชื่อมต่อกับฐานข้อมูล (Database) เพื่อให้แน่ใจว่าเซิร์ฟเวอร์สามารถเชื่อมต่อกับฐานข้อมูลได้สำเร็จ.

### การทำงาน:
- ดึงการเชื่อมต่อกับฐานข้อมูลจาก GORM.
- ตรวจสอบการเชื่อมต่อกับฐานข้อมูล (Ping) ว่ามีการเชื่อมต่อหรือไม่.
- หากไม่สามารถเชื่อมต่อได้ หรือเกิดข้อผิดพลาด จะส่งกลับ HTTP status 500 พร้อมข้อความ "Database connection failed".
- หากเชื่อมต่อสำเร็จ จะส่งข้อความ "Successfully connected to the database!".

---

## API สำหรับการทำงานกับฐานข้อมูล

คุณจำเป็นต้องระบุ **token** ใน Header (หากทดสอบด้วย Postman) เพื่อระบุตัวตนของคุณในการใช้งาน API (ดูได้จากฐานข้อมูล)

### GET /users

#### วิธีทดสอบ:
- `/users` → คืนค่า ผู้ใช้ทั้งหมด
- `/users?token=token5` → คืนค่า ผู้ใช้ที่ตรงกับเงื่อนไข
- `/users?username=john_doe` → คืนค่า ผู้ใช้ที่ตรงกับเงื่อนไข
- `/users?username=john_doe&token=token5` → คืนค่า ผู้ใช้ที่ตรงกับเงื่อนไข
- `/users?Username=xxxx` (ถ้าไม่มีข้อมูล) → คืนค่า "User not found"

#### วัตถุประสงค์:
ดึงข้อมูลผู้ใช้ทั้งหมดจากระบบ (อ่านข้อมูลผู้ใช้).

#### การทำงาน:
- ใช้ `CheckTokenMiddleware` เพื่อตรวจสอบว่าในคำขอมี JWT Token ที่ถูกต้องหรือไม่.
- หากมีการตรวจสอบที่ถูกต้อง, จะเรียกฟังก์ชัน `GetUserHandler` จาก `userHandler` เพื่อดึงข้อมูลผู้ใช้ทั้งหมดจากระบบ.

---

### PUT /users

#### วิธีทดสอบ:
- `/users?Token=token5` ระบุ **body**: `Username`, `Password`, `Image` เพื่ออัพเดทข้อมูลของผู้ใช้ที่ตรงกับค่าของ Token หากไม่มีฟิลด์ใดส่งมา จะใช้ค่าเดิม.

#### วัตถุประสงค์:
อัปเดตข้อมูลผู้ใช้.

#### การทำงาน:
- ใช้ `CheckTokenMiddleware` เพื่อตรวจสอบ JWT Token ว่าถูกต้องหรือไม่.
- หากตรวจสอบผ่าน, จะเรียกฟังก์ชัน `UpdateUserHandler` จาก `userHandler` เพื่อทำการอัปเดตข้อมูลผู้ใช้ตามข้อมูลที่ได้รับ.

---

### DELETE /users

#### วิธีทดสอบ:
- `/users?Token=token5` ลบข้อมูลของผู้ใช้ที่ตรงกับค่าของ Token

#### วัตถุประสงค์:
ลบข้อมูลผู้ใช้จากระบบ.

#### การทำงาน:
- ใช้ `CheckTokenMiddleware` เพื่อตรวจสอบ JWT Token ว่าถูกต้องหรือไม่.
- หากตรวจสอบผ่าน, จะเรียกฟังก์ชัน `DeleteUserHandler` จาก `userHandler` เพื่อลบข้อมูลผู้ใช้.

---

### POST /users

#### วิธีทดสอบ:
- `/users` ระบุ **body**: `Username`, `Password`, `Image` เพื่อสร้างข้อมูลของผู้ใช้คนใหม่

#### วัตถุประสงค์:
สร้างผู้ใช้ใหม่ในระบบ.

#### การทำงาน:
- ใช้ `CheckTokenMiddleware` เพื่อตรวจสอบ JWT Token ว่าถูกต้องหรือไม่.
- หากตรวจสอบผ่าน, จะเรียกฟังก์ชัน `CreateUserHandler` จาก `userHandler` เพื่อสร้างผู้ใช้ใหม่ในระบบด้วยข้อมูลที่ได้รับ.

---
