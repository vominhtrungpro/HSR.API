package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vominhtrungpro/internal/elements"
	"github.com/vominhtrungpro/internal/elements/elementmodel"
	"github.com/vominhtrungpro/internal/model/model"
)

type eleHandler struct {
	eleController elements.Controller
}

// Characters handlers constructor
func NewElementHandlers(eleController elements.Controller) elements.Handlers {
	return &eleHandler{eleController: eleController}
}

// Create element handler
func (h eleHandler) Create(context *gin.Context) {
	var request elementmodel.CreateRequest
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	if errvalid := validatecreate(request); err != nil {
		http.Error(context.Writer, errvalid.Error(), http.StatusBadRequest)
		return
	}
	var element model.Element
	element.Name = request.Name
	element.Enname = request.Enname
	element.Picture = make([]byte, 0)
	result := h.eleController.Create(context, element)
	if result != nil {
		http.Error(context.Writer, result.Error(), http.StatusBadRequest)
		return
	}
	context.IndentedJSON(http.StatusOK, "Success")
}

// Create path handler
func (h eleHandler) CreatePath(context *gin.Context) {
	var request elementmodel.CreatePathRequest
	err := json.NewDecoder(context.Request.Body).Decode(&request)
	if err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	if errvalid := validatecreatepath(request); err != nil {
		http.Error(context.Writer, errvalid.Error(), http.StatusBadRequest)
		return
	}
	var path model.Path
	path.Name = request.Name
	path.Enname = request.Enname
	path.Picture = make([]byte, 0)
	result := h.eleController.CreatePath(context, path)
	if result != nil {
		http.Error(context.Writer, result.Error(), http.StatusBadRequest)
		return
	}
	context.IndentedJSON(http.StatusOK, "Success")
}

// Update element image handler
func (h eleHandler) UpdateElementImage(context *gin.Context) {
	elementname := context.Param("name")
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
	result := h.eleController.UpdateElementImage(context, elementname, image)
	if result != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	context.IndentedJSON(http.StatusOK, "Success")
}

// Update path image handler
func (h eleHandler) UpdatePathImage(context *gin.Context) {
	pathname := context.Param("name")
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
	result := h.eleController.UpdatePathImage(context, pathname, image)
	if result != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	context.IndentedJSON(http.StatusOK, "Success")
}

func validatecreate(input elementmodel.CreateRequest) error {
	if strings.TrimSpace(input.Name) == "" {
		return errInvalidName
	}
	if strings.TrimSpace(input.Enname) == "" {
		return errInvalidName
	}
	return nil
}

func validatecreatepath(input elementmodel.CreatePathRequest) error {
	if strings.TrimSpace(input.Name) == "" {
		return errInvalidName
	}
	if strings.TrimSpace(input.Enname) == "" {
		return errInvalidName
	}
	return nil
}
