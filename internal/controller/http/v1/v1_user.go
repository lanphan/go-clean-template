package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ironsail/whydah-go-clean-template/internal/entity"
	"github.com/ironsail/whydah-go-clean-template/internal/usecase"
	"github.com/ironsail/whydah-go-clean-template/pkg/logger"
)

type userRoutes struct {
	uc *usecase.UserUseCase
}

func newUserRoutes(handler *gin.RouterGroup, uc *usecase.UserUseCase) {
	r := &userRoutes{uc}

	h := handler.Group("/users")
	{
		h.GET("/", r.list)
		h.POST("/", r.create)
	}
}

type userListResponse struct {
	User []entity.User `json:"user"`
}

// @Summary     List all users
// @Description Show all users
// @ID          list-user
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Success     200 {object} userListResponse
// @Failure     500 {object} response
// @Router      /users [get]
func (r *userRoutes) list(c *gin.Context) {
	users, err := r.uc.List(c.Request.Context())
	if err != nil {
		logger.Error("http - v1 - user", logger.ErrWrap(err))
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, userListResponse{users})
}

type createUserRequest struct {
	WalletAddress string `json:"address"       binding:"required"  example:"0x321233"`
	Reward        uint   `json:"reward"  binding:"required"  example:100`
}

// @Summary     Create user
// @Description Create new user
// @ID          create-user
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body createUserRequest true "Create user"
// @Success     200 {object} entity.User
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /users [post]
func (r *userRoutes) create(c *gin.Context) {
	var request createUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("http - v1 - create user", logger.ErrWrap(err))
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	user, err := r.uc.Create(
		c.Request.Context(),
		entity.User{
			WalletAddress: request.WalletAddress,
			Reward:        request.Reward,
			ClaimStatus:   false,
			CreatedAt:     time.Now().UTC(),
		},
	)
	if err != nil {
		logger.Error("http - v1 - create user", logger.ErrWrap(err))
		errorResponse(c, http.StatusInternalServerError, "user service problems")

		return
	}

	c.JSON(http.StatusOK, user)
}
