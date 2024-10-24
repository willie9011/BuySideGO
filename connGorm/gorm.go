package connGorm

import (
	"MobileProject/AppInit"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB //全局變數  供給其他pack使用 DB變數

func validateConnection(db *gorm.DB) error {
	// 嘗試執行一個簡單的 SQL 查詢
	var result int64
	tx := db.Raw("SELECT 1").Scan(&result)
	if tx.Error != nil {
		return fmt.Errorf("database connection test failed: %v", tx.Error)
	}
	if result != 1 {
		return fmt.Errorf("unexpected result from test query")
	}
	return nil
}

func GormConnDatabase() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dns := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", username, password, host, port, dbname)

	var err error                                            // 要記設定err變數   因為:= 太方便如果是賦予只用=的情況記得檢查是否已經聲明變數
	DB, err = gorm.Open(sqlserver.Open(dns), &gorm.Config{}) // 使用赋值符号 = 而不是 := (因為是賦予已經創建的全局變數)
	if err != nil {
		AppInit.ShutDownServer(err)
	} //gorm.Config{} 允許你控制 GORM 的多種行為，適合不同場景下的需求，比如性能優化、日誌管理、命名策略等。

	log.Println("Database connected successfully!")
} //新版的gorm v2 不用defer db.Close() 它會自動關閉

// ------------------結構體示意圖----------------------------------
// gorm 會翻譯成 -> Users   (它會自動把它變成複數)
//type User struct {
//	Userid   int    `gorm:"primaryKey"`
//	UserName string `gorm:"size:100"`
//	Age      int    `gorm:"default:0"`
//}

func Query_Command() {
	//-----------------------
	// Get first record, order by primary key
	//db.First(&user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	// Get one record, no specfied order
	//db.Take(&user)
	//// SELECT * FROM users LIMIT 1;

	// Get last record, order by primary key
	//db.Last(&user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// Get all records
	//db.Find(&users)
	//// SELECT * FROM users;

	// Get record with primary key (only works for integer primary key)
	//db.First(&user, 10)
	//// SELECT * FROM users WHERE id = 10;
}
func EnvVar_Command() {
	//在shell 設定臨時環境變數
	//$env:DB_USERNAME="sa"
	//$env:DB_PASSWORD="123"
	//$env:DB_HOST="localhost"
	//$env:DB_PORT="1433"
	//$env:DB_NAME="Test01"

	//在shell設定永久環境變數(重複設定一次可以修改)
	//[System.Environment]::SetEnvironmentVariable("DB_USERNAME", "sa", "User")
	//[System.Environment]::SetEnvironmentVariable("DB_PASSWORD", "123", "User")
	//[System.Environment]::SetEnvironmentVariable("DB_HOST", "localhost", "User")
	//[System.Environment]::SetEnvironmentVariable("DB_PORT", "1433", "User")
	//[System.Environment]::SetEnvironmentVariable("DB_NAME", "Test01", "User")

	//刪除環境變數
	//[System.Environment]::SetEnvironmentVariable("你的變數名", $null, "User")

	//用os.Getenv() 來讀取環境變數
	//	  username := os.Getenv("DB_USERNAME")
	//    password := os.Getenv("DB_PASSWORD")
	//    host := os.Getenv("DB_HOST")
	//    port := os.Getenv("DB_PORT")
	//    dbname := os.Getenv("DB_NAME")

	//查看環境變數：
	//Get-ChildItem Env:

	//查看特定變數
	//$env:PATH

	//刷新環境變數
	//[System.Environment]::GetEnvironmentVariable("DB_NAME", "User")

	//組裝DNS（資料來源名稱

	//dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", username, password, host, port, dbname)

	//使用 GORM 連接資料庫
	//    db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	//    if err != nil {
	//        panic("failed to connect database")
	//    }
	//
	//    fmt.Println("Database connected successfully")
}
func DB_command() {
	//查詢相關函數
	//DB.Find(): 查詢所有符合條件的資料。
	//DB.Where(): 添加查詢條件。
	//DB.Not(): 查詢不符合條件的資料。
	//DB.Or(): 添加 OR 查詢條件。
	//DB.Limit(): 限定查詢結果數量。
	//DB.Offset(): 跳過指定數量的記錄。
	//DB.Order(): 指定排序方式。
	//DB.Group(): 分組查詢。
	//DB.Having(): 分組查詢後，對分組結果進行過濾。
	//新增、更新、刪除相關函數
	//DB.Create(): 新增一條記錄。
	//DB.Save(): 保存（新增或更新）一條記錄。
	//DB.Update(): 更新指定欄位。
	//DB.Delete(): 刪除記錄。
	//DB.Model(): 指定操作的模型。
	//關聯查詢
	//DB.Preload(): 預載入關聯的資料。
	//DB.Joins(): 手動加入 JOIN 查詢。
	//DB.Table():指定查詢某Table
	//其他常用函數
	//DB.Count(): 計算符合條件的記錄數量。
	//DB.Rows(): 獲取查詢結果的迭代器。
	//DB.Raw(): 執行原生 SQL。
	//DB.Transaction(): 開始一個事務。
}
