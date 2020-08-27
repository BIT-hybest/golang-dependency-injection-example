package mvcExample

type UserService struct {
	RedisClient CacheClient
	UserModel *UserModel
}

func NewUserService(client CacheClient, model *UserModel) *UserService {
	return &UserService{
		RedisClient: client,
		UserModel:   model,
	}
}