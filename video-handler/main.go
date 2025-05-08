package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)


const (
	GetVideoByIDQuery string = "SELECT id, description, storage_link, title FROM video WHERE id = ?"

	CreateVideoQuery  string = "INSERT INTO video (id, description, storage_link, title) VALUES (?, ?, ?, ?)"
)


var (
	session *gocql.Session

	router = gin.Default()
)


type RestErr struct {
	ErrMessage string        `json:"message"`

	ErrStatus  int           `json:"status"`

	ErrError   string        `json:"error"`
	
	ErrCauses  []interface{} `json:"causes"`
}


func NewRestError(message string, status int, err string, causes []interface{}) *RestErr {
	return &RestErr{
		ErrMessage: message,

		ErrStatus:  status,

		ErrError:   err,

		ErrCauses:  causes,
	}
}


func NewRestErrorFromBytes(bytes []byte) (*RestErr, error) {
	var apiErr RestErr
	
	err := json.Unmarshal(bytes, &apiErr)
	
	if (err != nil) {
		return nil, errors.New("Invalid JSON")
	}
	
	return &apiErr, nil
}


func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		ErrMessage: message,

		ErrStatus:  http.StatusBadRequest,

		ErrError:   "bad_request",
	}
}


func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		ErrMessage: message,

		ErrStatus:  http.StatusNotFound,

		ErrError:   "not_found",
	}
}


func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		ErrMessage: message,

		ErrStatus:  http.StatusUnauthorized,

		ErrError:   "unauthorized",
	}
}


func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		ErrMessage: message,

		ErrStatus:  http.StatusInternalServerError,

		ErrError:   "internal_server_error",
	}
	
	if (err != nil) {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	
	return result
}


func NewDbRepository() DbRepository {
	return &dbRepository{}
}


type DbRepository interface {
	GetByID(videoID string) (*Video, *RestErr)

	Create(video Video) (*Video, *RestErr)
}


type dbRepository struct {}


func (repo *dbRepository) Create(video Video) (*Video, *RestErr) {
	err := GetSession().Query(CreateVideoQuery, video.ID, video.Description, video.StorageLink, video.Title).Exec()

	if (err != nil) {
		return nil, NewInternalServerError("Unable to insert video's metadata to Cassandra", errors.New(err.Error()))
	}

	return &video, nil
}


func (repo *dbRepository) GetByID(videoID string) (*Video, *RestErr) {
	var video Video

	err := GetSession().Query(GetVideoByIdQuery, videoID).Scan(&video.ID, &video.Description, &video.StorageLink, &video.Title)
	
	if (err != nil) {
		if (err.Error() == "Not found") {
			fmt.Println("Here")

			return nil, NewInternalServerError("No video for given video ID", errors.New(err.Error()))
		}

		return nil, NewInternalServerError("Unable to find video in Cassandra", errors.New(err.Error()))
	}

	return &video, nil
}


type Video struct {
	ID           string `json:"id"`

	Description  string `json:"description"`

	StorageLink  string `json:"storage_link"`

	Title        string `json:"title"`
}


func (video *Video) ValidateVideo() *RestErr {
	if (video.Title == "") {
		return NewInternalServerError("Title can't be blank", nil)
	}

	return nil
}


type VideoRepository interface {
	GetByID(videoID string) (*Video, *RestErr)

	Create(video Video) (*Video, *RestErr)
}


type VideoService interface {
	GetByID(videoID string) (*Video, *RestErr)

	Create(video Video) (*Video, *RestErr)
}


type videoService struct {
	repository VideoRepository
}


func (s *videoService) Create(video Video) (*Video, *RestErr) {
	err := video.ValidateVideo()

	if (err != nil) {
		return nil, err
	}

	return s.repository.Create(video)
}


func NewService(repository VideoRepository) VideoService {
	return &videoService{repository: repository}
}


func (s *videoService) GetByID(videoID string) (*Video, *RestErr) {
	videoID = strings.TrimSpace(videoID)

	if (len(videoID) == 0) {
		return nil, NewBadRequestError("Invalid video ID, it can't be blank")
	}

	video, err := s.repository.GetByID(videoID)

	if (err != nil) {
		videoNotFoundErr := fmt.Sprintf("Video not found for video ID %s", videoID)

		return nil, NewInternalServerError(videoNotFoundErr, errors.New("Here"))
	}

	return video, nil
}


func connectToCassandra() {
	cluster := gocql.NewCluster("127.0.0.1")

	cluster.Keyspace = "dream-stream"

	cluster.Consistency = gocql.Quorum

	var err error

	session, err = cluster.CreateSession()

	if (err != nil) {
		panic(err)
	}
}


func GetSession() *gocql.Session {
	return session
}


func Run()  {
	userHandler := NewHandler(NewService(NewDbRepository()))

	router.GET("/video/:id", videoHandler.GetById)

	router.POST("/video", videoHandler.Create)

	_ = router.Run(":8888")
}


type VideoHandler interface {
	GetById(ctx *gin.Context)

	Create(ctx *gin.Context)
}


type videoHandler struct {
	videoService videoService
}


func (videoHandler videoHandler) GetById(ctx *gin.Context) {
	videoID := strings.TrimSpace(ctx.Param("id"))

	video, err := videoHandler.videoService.GetByID(videoID)

	if (err != nil) {
		ctx.JSON(err.ErrStatus, err)

		return
	}

	ctx.JSON(http.StatusOK, video)
}


func NewHandler(videoService VideoService) VideoHandler {
	return &videoHandler{videoService: videoService}
}


func (videoHandler *videoHandler) Create(ctx *gin.Context)  {
	var video Video

	err := ctx.ShouldBindJSON(&video)

	if (err != nil) {
		restErr := NewBadRequestError("Invalid JSON body")

		ctx.JSON(restErr.ErrStatus, restErr)
	}

	_, videoErr := videoHandler.videoService.Create(video)

	if (videoErr != nil) {
		ctx.JSON(videoErr.ErrStatus, videoErr)

		return
	}

	ctx.JSON(http.StatusCreated, video)
}