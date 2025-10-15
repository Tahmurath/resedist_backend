package services

import (
	//UserRepository "resedist/internal/modules/user/repositories"
	//UserResponse "resedist/internal/modules/user/responses"
	tgUserRepository "resedist/internal/modules/tgminiapp/repositories"
	//UserResponse "resedist/internal/modules/user/responses"
)

type TgUserService struct {
	tgUserRepository tgUserRepository.TgUserRepositoryInterface
}

func New() *TgUserService {
	return &TgUserService{
		tgUserRepository: tgUserRepository.New(),
	}
}

func (TgUserService *TgUserService) CheckUserExist(tgId int64) bool {

	user := TgUserService.tgUserRepository.FindByTgID(tgId)

	if user.ID != 0 {
		return true
	}

	return false
}

//func (TgUserService *TgUserService) Create(request auth.RegisterRequest) (UserResponse.User, error) {
//
//	var response UserResponse.User
//	var user userModels.User
//
//	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
//	if err != nil {
//		log.Fatal("hash password error")
//		return response, errors.New("hash password error")
//	}
//
//	user.Name = request.Name
//	user.Email = request.Email
//	user.Password = string(hashPassword)
//
//	newUser := UserService.userRepository.Create(user)
//
//	if newUser.ID == 0 {
//		return response, errors.New("user create fail")
//	}
//
//	return UserResponse.ToUser(newUser), nil
//}

//func (UserService *UserService) HandleUserLogin(request auth.LoginRequest) (UserResponse.User, error) {
//
//	var response UserResponse.User
//	existUser := UserService.userRepository.FindByEmail(request.Email)
//
//	if existUser.ID == 0 {
//		return response, errors.New("invalid credentials")
//	}
//
//	err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(request.Password))
//	if err != nil {
//		return response, errors.New("invalid credentials")
//	}
//
//	return UserResponse.ToUser(existUser), nil
//}
//
//func (UserService *UserService) GetCachedUserById(userId int) (UserResponse.User, error) {
//
//	ctx := context.Background()
//	key := "user:" + strconv.Itoa(userId)
//	//tag := "users"
//	ttl := 10 * time.Minute
//
//	fetchFunc := func() (UserResponse.User, error) {
//		userRepo := UserRepository.New()
//		foundUser := userRepo.FindByID(userId)
//		if foundUser.ID == 0 {
//			return UserResponse.User{}, errors.New("user not found")
//		}
//		return UserResponse.ToUser(foundUser), nil
//	}
//	value, err := redis.GetOrSetJSON(ctx, key, ttl, fetchFunc)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return UserResponse.User{}, errors.New("user not found")
//	}
//	return value, nil
//}
