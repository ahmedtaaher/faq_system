package handler

import (
	"faq_sys_go/models"
	"faq_sys_go/repository"
	"faq_sys_go/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userRepo  *repository.UserRepository
	storeRepo *repository.StoreRepository
}

func NewAuthHandler(userRepo *repository.UserRepository, storeRepo *repository.StoreRepository) *AuthHandler {
	return &AuthHandler{
		userRepo:  userRepo,
		storeRepo: storeRepo,
	}
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	UserType string `json:"user_type" binding:"required,oneof=customer merchant admin"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  *models.User `json:"user"`
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	existingUser, _ := h.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		utils.ErrorResponse(c, http.StatusConflict, "User with this email already exists")
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := &models.User{
		Email:    req.Email,
		Password: hashedPassword,
		UserType: models.UserType(req.UserType),
	}

	if user.UserType == models.UserTypeMerchant {
		store := &models.Store{
			Name:        strings.Split(req.Email, "@")[0] + "'s Store",
			Description: "Auto-generated store for merchant",
		}

		if err := h.storeRepo.Create(store); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create store")
			return
		}

		user.StoreID = &store.ID
	}

	if err := h.userRepo.Create(user); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", AuthResponse{
		User:  user,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userRepo.FindByEmail(req.Email)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email, string(user.UserType), user.StoreID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login successful", AuthResponse{
		Token: token,
		User:  user,
	})
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID := c.GetUint("userID")

	user, err := h.userRepo.FindByID(userID)
	if err != nil {
    utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile retrieved successfully", user)
}