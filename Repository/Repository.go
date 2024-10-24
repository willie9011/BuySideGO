package Repository

import "MobileProject/Structs"

// 定义 User 的 Repository 接口
type UserRepository interface {
	Create(user *Structs.User) error
	FindByID(id string) (*Structs.User, error)
	Update(user *Structs.User) error
	Delete(id string) error
}
