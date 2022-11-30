package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibrat-muslim/students-api/api/models"
	"github.com/ibrat-muslim/students-api/storage/repo"
)

// @Router /students [post]
// @Summary Create a student
// @Description Create a student
// @Tags student
// @Accept json
// @Produce json
// @Param student body models.StudentsArray true "Student"
// @Success 201 {object} models.OKResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateStudent(ctx *gin.Context) {

	req := make(models.StudentsArray, 0)

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var s []*repo.Student

	for _, val := range req {
		s = append(s, &repo.Student{
			FirstName:   val.FirstName,
			LastName:    val.LastName,
			UserName:    val.UserName,
			Email:       val.Email,
			PhoneNumber: val.PhoneNumber,
		})
	}

	err = h.storage.Student().Create(s)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.OKResponse{
		Message: "successfully created",
	})
}

func validateGetAllStudentsParams(ctx *gin.Context) (*models.GetAllStudentsParams, error) {
	var (
		limit int64 = 10
		page  int64 = 1
		err   error
	)

	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return &models.GetAllStudentsParams{
		Limit:  int32(limit),
		Page:   int32(page),
		Search: ctx.Query("search"),
	}, nil
}

// @Router /students [get]
// @Summary Get students
// @Description Get students
// @Tags student
// @Accept json
// @Produce json
// @Param filter query models.GetAllStudentsParams false "Filter"
// @Success 200 {object} models.GetAllStudentsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetStudents(ctx *gin.Context) {
	request, err := validateGetAllStudentsParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := h.storage.Student().GetAll(&repo.GetAllStudentsParams{
		Limit:  request.Limit,
		Page:   request.Page,
		Search: request.Search,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getStudentsResponse(result))
}

func getStudentsResponse(data *repo.GetAllStudentsResult) *models.GetAllStudentsResponse {
	response := models.GetAllStudentsResponse{
		Students: make([]*models.Student, 0),
		Count:    data.Count,
	}

	for _, student := range data.Students {
		s := parseStudentToModel(student)
		response.Students = append(response.Students, &s)
	}

	return &response
}

func parseStudentToModel(student *repo.Student) models.Student {
	return models.Student{
		ID:          student.ID,
		FirstName:   student.FirstName,
		LastName:    student.LastName,
		UserName:    student.UserName,
		Email:       student.Email,
		PhoneNumber: student.PhoneNumber,
		CreatedAt:   student.CreatedAt,
	}
}
