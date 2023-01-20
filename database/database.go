package database
import (
 "log"
 "os"
 "gorm.io/driver/postgres"
 "github.com/Janitherange/go_crud/config"
 "github.com/Janitherange/go_crud/model"
 "gorm.io/gorm"
 "gorm.io/gorm/logger"
)
// Database instance
type Dbinstance struct {
 Db *gorm.DB
}
var DB Dbinstance
// Connect function
func Connect() {
 
 dsn := config.Config("DSN_URL")

 db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
  Logger: logger.Default.LogMode(logger.Info),
 })

 if err != nil {
  log.Fatal("Failed to connect to database. \n", err)
  os.Exit(2)
 }

 log.Println("Connected")

 db.Logger = logger.Default.LogMode(logger.Info)

 log.Println("running migrations")

 db.AutoMigrate(&model.User{})

 DB = Dbinstance{
  Db: db,
 }
}