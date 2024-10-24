package main

import (
	"MobileProject/AppInit"
	"MobileProject/HandlerFunc"
	"MobileProject/MyValidator"
	"MobileProject/Structs"
	"MobileProject/connGorm"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

func main() {

	router := gin.Default()
	//---------------gorm 業務------------------------------------------------------------
	connGorm.GormConnDatabase()
	migrate_err := connGorm.DB.AutoMigrate(&Structs.User{})
	if migrate_err != nil {
		log.Fatal("migrate fail:%s", migrate_err)
	}
	//---------------validator業務--------------------------------------------------------
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		v.RegisterValidation("myurlvalid", MyValidator.MyUrlValid)
	}
	//---------------router業務-----------------------------------------------------------
	v1 := router.Group("/v1/Menu")
	v1.Use(HandlerFunc.MustLogin()) //導層的api封裝
	{
		v1.GET("/query", HandlerFunc.GetQueryItem)          //http://localhost:8080/v1/Menu/query?token=123
		v1.GET("/user/get", HandlerFunc.GetUser)            //http://localhost:8080/v1/Menu/user/get?token=123
		v1.GET("vegetables/get", HandlerFunc.GetVegetables) //http://localhost:8080/v1/Menu/vegetables/get?token=123
		v1.GET("/users/get", HandlerFunc.GetUsers)          //http://localhost:8080/v1/Menu/users/get?token=123
		v1.POST("/user", HandlerFunc.CreateUser)            //http://localhost:8080/v1/Menu/user?token=123
		v1.POST("/users", HandlerFunc.CreateMultiUsers)     //http://localhost:8080/v1/Menu/users?token=123
	}
	//--------------router業務2-------------------------------------------------------------------------
	v2 := router.Group("/v2/Menu")
	v2.Use(HandlerFunc.MustLogin())
	{
		v2.POST("/user/m_create", HandlerFunc.CreateMultiUsers)
	}
	//--------------起動server ------------------------------------------------------------
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go (func() {
		serverErr := server.ListenAndServe()
		if serverErr != nil {
			log.Fatal("伺服器啟中斷", serverErr)
		}
	})()
	//-------------err關閉伺服器善後--------------------------------------------------------------

	AppInit.ServerNotify()                                                     //如果收到os.Interrupt 的話 執行下列工作
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second*5) // 善後工作
	defer cancel()

	shutDownErr := server.Shutdown(ctx)
	if shutDownErr != nil {
		log.Fatal("伺服器關閉")
	}
	log.Println("伺服器優雅退出")
}
