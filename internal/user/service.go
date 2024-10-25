package user

import (
	"context"
	"initial_project_go/pkg/config"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo UserRepositoryInterface
}

func NewUserService(userRepo UserRepositoryInterface) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) ServiceAddUsers(ctx context.Context, request *RequestAddUsers) (any, int, error) {

	// -----------------------> Start use API, hapus jika tidak menggunakan api <-----------------------
	// urlInfo := config.GetConfig("api.test.testApi1")
	// result, err := utils.Get(ctx, &urlInfo, 1)
	// if err != nil {
	// 	return nil, http.StatusBadRequest, fmt.Errorf("1-%s", err.Error())
	// }

	// results, err := utils.InterfaceToStruct[PersonalInfoResponse](result)
	// if err != nil {
	// 	return nil, http.StatusBadRequest, fmt.Errorf("2-%s", err.Error())
	// }

	// fmt.Println("data API ====>>>> ", results)
	// fmt.Println("contoh ambil value dari key API ====>>>> ", results.City)
	// -----------------------> End use API <-----------------------

	users := Users{
		Id:        uuid.New().String(),
		Name:      request.Name,
		CreatedAt: time.Now(),
	}

	// insert users ke database
	status, errAdd := s.userRepo.AddUsers(ctx, &users)
	if errAdd != nil {
		return nil, status, errAdd
	}

	return "Add Users Success", status, nil
}

func (s *UserService) ServiceGetUsers(ctx context.Context, request *RequestGetUsers) (any, int, error) {
	limit, _ := strconv.Atoi(config.GetConfig("limit"))
	offset := 0

	if request.Limit == nil {
		request.Limit = &limit
	}

	if request.Offset == nil {
		request.Offset = &offset
	}

	// insert users ke database
	data, count, status, errAdd := s.userRepo.GetUsers(ctx, request)
	if errAdd != nil {
		return nil, status, errAdd
	}

	dataResp := ResponseGetUsers{
		Records: data,
		Total:   count,
	}

	return dataResp, status, nil
}
