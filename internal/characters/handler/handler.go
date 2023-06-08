package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vominhtrungpro/internal/characters"
	"github.com/vominhtrungpro/internal/characters/charactermodel"
	"github.com/vominhtrungpro/internal/model/model"
)

type charHandler struct {
	charController characters.Controller
}

// Characters handlers constructor
func NewCharsHandlers(charController characters.Controller) characters.Handlers {
	return &charHandler{charController: charController}
}

func (h charHandler) GetAll(context *gin.Context) {
	result, err := h.charController.GetAll(context)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	context.IndentedJSON(http.StatusOK, result)
}

func (h charHandler) Create(context *gin.Context) {
	var request charactermodel.CreateRequest
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	var character model.Character
	character.ID = request.ID
	character.Name = request.Name
	character.Rarity = request.Rarity
	character.Element = request.Element
	character.Path = request.Path
	character.Picture = make([]byte, 0)
	h.charController.Create(context, character)
	context.IndentedJSON(http.StatusOK, "Success")
}

func (h charHandler) UpdateCharacterImage(context *gin.Context) {
	charId := context.Param("id")
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	image, err := io.ReadAll(file)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := h.charController.UpdateCharacterImage(context, charId, image)
	if result != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	context.IndentedJSON(http.StatusOK, "Success")
}

func (h charHandler) GetImageById(context *gin.Context) {
	charId := context.Param("id")
	image, name, err := h.charController.GetImageById(context, charId)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	assetsurl := os.Getenv("ASSETs_URL")
	fileurl := fmt.Sprintf(assetsurl + name + ".png")
	err = os.WriteFile(fileurl, image, 0644)
	if err != nil {
		panic(fileurl)
	}
	filename := fmt.Sprintf("%v.png", name)
	context.FileAttachment(fileurl, filename)
}

func (h charHandler) Test(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "oke")
}
