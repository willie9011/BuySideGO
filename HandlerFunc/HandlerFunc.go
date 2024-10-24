package HandlerFunc

import (
	"MobileProject/Structs"
	"MobileProject/connGorm"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 判斷登錄
func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, status := c.GetQuery("token"); !status {
			c.String(http.StatusUnauthorized, "MissingToken")
			c.Abort() //在此中斷
		} else {
			c.Next() //繼續下一個
		}
	}
}

func GetQueryItem(c *gin.Context) { //讀取url上面的參數
	query := Structs.UrlQuery{}
	err := c.BindQuery(&query)
	if err != nil {
		c.String(http.StatusBadRequest, "參數錯誤:%s", err.Error())
	} else {
		c.JSON(http.StatusOK, query)
	}
}

func GetUser(c *gin.Context) {
	user := Structs.User{}
	err := connGorm.DB.First(&user).Error
	if err != nil {
		c.String(http.StatusBadRequest, "err:%s", err.Error())
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUsers(c *gin.Context) {
	users := Structs.MultiUsers{}.UserList
	err := connGorm.DB.Find(&users).Error
	if err != nil {
		c.String(http.StatusBadRequest, "err:%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetVegetables(c *gin.Context) {
	vegetables := Structs.Vegetables{}.VegeList
	err := connGorm.DB.Find(&vegetables).Error
	if err != nil {
		c.String(http.StatusBadRequest, "err:%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, vegetables)
}

func CreateUser(c *gin.Context) {
	new_user := Structs.User{}
	err := c.BindJSON(&new_user)
	if err != nil {
		c.String(http.StatusBadRequest, "error:%s", err.Error())
		return
	}

	err2 := connGorm.DB.Create(&new_user).Error //確定是藉由Error返回
	if err2 != nil {
		c.String(http.StatusBadRequest, "error:%s", err2) //因為返回得已經是錯誤類型所以就不用加.Error()
		return
	} else {
		c.JSON(http.StatusOK, new_user)
	}

	//err2 := connGorm.DB.Create(&new_user).Error
	//if err2 != nil {
	//	c.String(http.StatusBadRequest, "error:%s", err.Error())
	//} else {
	//	c.JSON(http.StatusOK, new_user)
	//}

}

func CreateMultiUsers(c *gin.Context) {
	userlist := Structs.MultiUsers{}
	err := c.BindJSON(&userlist)
	if err != nil {
		c.String(http.StatusBadRequest, "error:%s", err.Error())
	}

	err2 := connGorm.DB.Create(&userlist.UserList).Error
	if err2 != nil {
		c.String(http.StatusBadRequest, "error:%s", err2)
	} else {
		c.JSON(http.StatusOK, userlist)
	}
}
