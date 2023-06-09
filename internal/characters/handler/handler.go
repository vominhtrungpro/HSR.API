package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
	if errvalid := validatecreate(request); err != nil {
		http.Error(context.Writer, errvalid.Error(), http.StatusBadRequest)
		return
	}
	var character model.Character
	character.Name = request.Name
	character.Rarity = request.Rarity
	character.Element = request.Element
	character.Path = request.Path
	character.Picture = make([]byte, 0)
	result := h.charController.Create(context, character)
	if result != nil {
		http.Error(context.Writer, result.Error(), http.StatusBadRequest)
		return
	}
	context.IndentedJSON(http.StatusOK, "Success")
}

func (h charHandler) Update(context *gin.Context) {
	var request charactermodel.UpdateRequest
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	if errvalid := validateupdate(request); err != nil {
		http.Error(context.Writer, errvalid.Error(), http.StatusBadRequest)
		return
	}
	result := h.charController.Update(context, request)
	if result != nil {
		http.Error(context.Writer, result.Error(), http.StatusBadRequest)
		return
	}
	context.IndentedJSON(http.StatusOK, "Success")
}

func (h charHandler) UpdateCharacterImage(context *gin.Context) {
	charname := context.Param("name")
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
	result := h.charController.UpdateCharacterImage(context, charname, image)
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

func validatecreate(input charactermodel.CreateRequest) error {
	if strings.TrimSpace(input.Name) == "" {
		return errInvalidName
	}
	return nil
}

func validateupdate(input charactermodel.UpdateRequest) error {
	if strings.TrimSpace(input.Id) == "" {
		return errInvalidId
	}
	if strings.TrimSpace(input.Name) == "" {
		return errInvalidName
	}
	return nil
}
