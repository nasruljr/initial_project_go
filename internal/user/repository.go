package user

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) AddUsers(ctx context.Context, request *Users) (int, error) {
	// Coba langsung membuat user, jika gagal karena duplikasi nama, handle errornya
	err := r.DB.WithContext(ctx).Create(&request).Error
	if err != nil {
		// Cek apakah error disebabkan oleh pelanggaran constraint unique
		if strings.Contains(err.Error(), "duplicate key value") {
			return http.StatusBadRequest, errors.New("user with this name already exists")
		}
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (r *UserRepository) GetUsers(ctx context.Context, request *RequestGetUsers) ([]Users, int64, int, error) {
	var users []Users
	var totalCount int64

	query := r.DB.WithContext(ctx).Model(users)

	// Jika ada keyword, tambahkan kondisi pencarian berdasarkan name
	if request.Keyword != "" {
		query = query.Where("name LIKE ?", "%"+request.Keyword+"%")
	}

	// Hitung total count
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, http.StatusInternalServerError, err
	}

	// Jika ada limit, tambahkan limit ke query
	if request.Limit != nil {
		query = query.Limit(*request.Limit)
	}

	// Jika ada offset, tambahkan offset ke query
	if request.Offset != nil {
		query = query.Offset(*request.Offset)
	}

	// Eksekusi query
	err := query.Order("name ASC").Find(&users).Error
	if err != nil {
		return nil, 0, http.StatusInternalServerError, err
	}

	return users, totalCount, http.StatusOK, nil
}
