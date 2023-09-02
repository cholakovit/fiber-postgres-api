go mod init fiber-postgres-api

go get github.com/githubnemo/CompileDaemon
go get github.com/joho/godotenv
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

go get github.com/stretchr/testify/assert

CompileDaemon -command="./fiber-postgres-api"