package services

type TgUserServiceInterface interface {
	CheckUserExist(tgId int64) bool
	//Create(request auth.RegisterRequest) (UserResponse.User, error)
	//HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error)
	//GetCachedUserById(userId int) (UserResponse.User, error)
}
