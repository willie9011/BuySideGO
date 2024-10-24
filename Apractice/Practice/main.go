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
	router := gin.Default()

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("TESTDB_NAME")
	dns := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", username, password, host, port, dbname)

	var err error
	DB, err = gorm.Open(sqlserver.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("can't connect to database", err)
	}
	log.Println("Connected to database")

	migrate_err := DB.AutoMigrate(&User{})
	if migrate_err != nil {
		log.Fatal("AutoMigrate failed: %s", migrate_err)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("myvalid", MyValid)
	}

	v1 := router.Group("/v1")
	v1.Use(MustLogin())
	{
		v1.GET("/getquery", GetQuery) //http://localhost:8080/v1/getquery?token=123&pagename=willie&pagesize=3
		v1.GET("/getuser", GetUser)   //http://localhost:8080/v1//getuser?token=123
		v1.GET("/vegetables", GetVagetables)
		v1.POST("/createuser", CreateUser)   //http://localhost:8080/v1/createuser?token=123
		v1.POST("/createusers", CreateUsers) //http://localhost:8080/v1/createusers?token=123
	}
	router.Run()
}

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, Status := c.GetQuery("token")
		if !Status {
			c.String(http.StatusBadRequest, "Missing token")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func GetQuery(c *gin.Context) {
	value := Page{}
	err := c.BindQuery(&value)
	if err != nil {
		c.String(http.StatusBadRequest, "err:%s", err.Error())
	} else {
		c.JSON(http.StatusOK, value)
	}
}

func GetUser(c *gin.Context) {
	users := Users{}.UserList
	err := DB.Find(&users).Error
	if err != nil {
		c.String(http.StatusBadRequest, "err:%v", err.Error())
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(c *gin.Context) {
	user := User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "err:%s", err)
		return
	}

	err2 := DB.Create(&user).Error
	if err2 != nil {
		c.String(http.StatusBadRequest, "err:%s", err2.Error())
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUsers：插入多個 User 資料
func CreateUsers(c *gin.Context) {
	users := Users{}
	err := c.BindJSON(&users)
	if err != nil {
		c.String(http.StatusBadRequest, "err:%s", err)
		return
	}

	// 使用 GORM 的批量插入功能，將多筆 User 資料插入到同一張表
	err2 := DB.Create(&users.UserList).Error
	if err2 != nil {
		c.String(http.StatusBadRequest, "err:%s", err2.Error())
		return
	} else {
		c.JSON(http.StatusOK, users.UserList)
	}
}

func MyValid(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	re := regexp.MustCompile("^[a-zA-Z]{3,}")
	return re.MatchString(value)
}

// 這個結構體會對應到資料表 "users"
type User struct {
	UserID   uint   `json:"userid" form:"userid" binding:"required" gorm:"primaryKey"`
	UserName string `json:"username" form:"username" binding:"required,myvalid" gorm:"column:username"`
	UserAge  uint   `json:"userage" form:"userage" binding:"required" gorm:"column:userage"`
}

// Users 結構體只是個容器，不會建立資料表
type Users struct {
	UserList []User `json:"userlist" binding:"dive"` // 不設置 foreignKey，因為不需要關聯到多張表
	Size     uint   `json:"size" form:"size" binding:"gt=1"`
}

type Page struct {
	PageName string `json:"pagename" form:"pagename" binding:"required"`
	PageSize uint   `json:"pagesize" form:"pagesize" binding:"required"`
}
