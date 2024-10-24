package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"regexp"
)

var DB *gorm.DB

func main() {
	db_user := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("TESTDB_NAME")
	dns := fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v", db_user, db_password, db_host, db_port, db_name)

	var DB_err error

	DB, DB_err = gorm.Open(sqlserver.Open(dns), &gorm.Config{})
	if DB_err != nil {
		fmt.Printf("DB ERROR:%v", DB_err)
	} else {
		log.Println("DB connected successfully")
	}

	migrate_err := DB.AutoMigrate(&User{})
	if migrate_err != nil {
		fmt.Printf("line:33 Migration ERROR:%v", migrate_err)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mytag", Myvalid)
	} else {
		fmt.Println("RegisterValidation failed")
	}

	router := gin.Default()
	v1 := router.Group("v1")
	v1.Use(MustLogin())
	{
		v1.GET("/getuser", GetUser)          //http://localhost:8080/v1/getuser?token=123
		v1.GET("getusers", GetUser)          //http://localhost:8080/v1/getusers?token=123
		v1.GET("/getquery", GetQuery)        //http://localhost:8080/v1/getquery?token=123
		v1.POST("/createuser", CreateUser)   //http://localhost:8080/v1/createuser?token=123
		v1.POST("/createusers", CreateUsers) //http://localhost:8080/v1/createusers?token=123
	}
	router.Run()
}

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, Status := c.GetQuery("token")
		if !Status {
			c.String(http.StatusUnauthorized, "Missing Token!")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func Myvalid(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile("^[a-zA-Z0-9]{3,}")
	return re.MatchString(value)
}

func GetQuery(c *gin.Context) {
	query := Page{}
	err := c.BindQuery(&query)
	if err != nil {
		c.String(http.StatusBadRequest, "err:%v", err)
	} else {
		c.JSON(http.StatusOK, query)
	}
}

func CreateUser(c *gin.Context) {
	user := User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "err:%s", err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func CreateUsers(c *gin.Context) {
	users := Users{}
	err := c.BindJSON(&users)
	if err != nil {
		c.String(http.StatusBadRequest, "err:%s", err)
		return
	}
	err2 := DB.Create(&users.UserList).Error
	if err2 != nil {
		c.String(http.StatusBadRequest, "err:%v", err2)
		return
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetUser(c *gin.Context) {
	users := Users{}.UserList
	err := DB.Find(&users).Error
	if err != nil {
		log.Printf("line 94 DB Cant Find User:%v", err)
		c.String(http.StatusBadRequest, "err!:%v", err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

type Page struct {
	PageName string `json:"page_name" form:"page_name" bind:"request"`
	PageSize uint   `json:"page_size" form:"page_size" bind:"request"`
}

type User struct {
	UserID   uint   `json:"user_id" form:"user_id" bind:"request" gorm:"PRIMARY_KEY"`
	UserName string `json:"user_name" form:"user_name" bind:"request" gorm:"user_name"`
	UserAge  uint   `json:"user_age" form:"user_age" bind:"request" gorm:"user_age"`
}

type Users struct {
	UserList []User `json:"user_list" bind:"dive"`
	Size     uint   `json:"size" form:"size"`
}
