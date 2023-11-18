# Online Course

Online Course API with Echo & MySQL

## 🔥 Showcase

- [Postman Docs](https://documenter.getpostman.com/view/25042327/2s93ebSVop)

## 💻 Built with

- [Echo](https://github.com/labstack/echo) for web framework
- [JWT](https://github.com/golang-jwt/jwt) for authentication and authorization
- [Cloudinary](https://github.com/cloudinary/cloudinary-go) for cloud files
- [Docker](https://github.com/docker) for deployment

## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/rfauzi44/online-course.git
```
2. Create .env file (copy and set from .env-example)

```bash
# DATABASE
MYSQL_USER=root
MYSQL_PASSWORD=password
MYSQL_ROOT_PASSWORD=password
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_DBNAME=online-course

# APP
APP_PORT=3001
JWT_KEY=secret

# CLOUDINARY
CLOUD_NAME=
CLOUD_APIKEY=
CLOUD_SECRET=
```

2. Install dependencies

```bash
go get -u ./...
```

3. Setup Database using restore or migrate CLI

```bash
migrate -database "mysql://root:password@tcp(localhost:3306)/dbname" -path db/migrations up
```

5. Run the app

```bash
go run .
```

🌟 You are all set!
