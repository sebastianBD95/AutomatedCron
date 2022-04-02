# ⏰ Crons

## 💡 Description
This project solves the problem of have to manually call apis to do a function.
with this we can centralize all the calls in one app to call all the other apis in order 
to facilate the control and maintenance.
## 🧰 Tech Stack
- [mongo](https://www.mongodb.com/)
- [RxGo](https://reactivex.io/)
- [Golang](https://go.dev/)

## 🚀 Built
```
go run main.go
```
## 🚀 Built Docker

```
docker build -t <tag_container> .

docker run -d <tag_container>
```
## 🌿 Mongo Configuration

- make sure you have your database runing on localhost:27017

- You must create de cron database and the user "cronUser"
with pass "cronPass"
