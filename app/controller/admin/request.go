package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/repository"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// RequestController is request
type RequestController struct {
	logger             pkg.Logger
	service            service.RequestService
	userRepository     repository.UserRepository
	categoryRepository repository.CategoryRepository
	tagRepository      repository.TagRepository
	value              support.RequestValue
	upload             helper.UploadHelper
	config             pkg.Config
}

// NewRequestController is create request controller
func NewRequestController(
	logger pkg.Logger,
	service service.RequestService,
	userRepository repository.UserRepository,
	categoryRepository repository.CategoryRepository,
	tagRepository repository.TagRepository,
	value support.RequestValue,
	upload helper.UploadHelper,
	config pkg.Config,
) RequestController {
	return RequestController{
		logger:             logger,
		service:            service,
		userRepository:     userRepository,
		categoryRepository: categoryRepository,
		tagRepository:      tagRepository,
		value:              value,
		upload:             upload,
		config:             config,
	}
}

// Index render request list
func (ctl RequestController) Index(c *fiber.Ctx) error {
	var form form.RequestSearch
	response := ctl.value.GetResponseHelper(c)

	if err := c.QueryParser(&form); err != nil {
		return response.Error(err)
	}

	requests, err := ctl.service.Search(object.NewRequestSearchInput(
		form.Keyword,
		form.Page,
	), 20)
	if err != nil {
		return response.Error(err)
	}

	return response.View("request/index", helper.DataMap{
		"requests": requests,
	})
}

// Detail render request detail
func (ctl RequestController) Detail(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	request, findErr := ctl.service.FindWithSuggests(uint(id))
	if findErr != nil {
		return response.Error(findErr)
	}

	return response.View("request/detail", helper.DataMap{
		"content_footer": true,
		"request":        request,
	})
}

// Create render request create form
func (ctl RequestController) Create(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)

	userID, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	user, findErr := ctl.userRepository.GetOne(uint(userID))
	if findErr != nil {
		return response.Error(findErr)
	}

	tags, tagErr := ctl.tagRepository.GetAll()
	if tagErr != nil {
		return response.Error(tagErr)
	}

	categories, categoryErr := ctl.categoryRepository.GetAll()
	if tagErr != nil {
		return response.Error(categoryErr)
	}

	return response.View("request/create", helper.DataMap{
		"content_footer": true,
		"statuses":       model.ToRequestStatusArray(),
		"tags":           model.TagsToArray(tags),
		"categories":     model.CategoriesToArray(categories),
		"user":           user,
	})
}

// Store request store process
func (ctl RequestController) Store(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)

	userID, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	user, findErr := ctl.userRepository.GetOne(uint(userID))
	if findErr != nil {
		return response.Error(findErr)
	}

	var form form.Request

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(
		c,
		ctl.categoryRepository,
		ctl.tagRepository,
	); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	thumbnail, upErr := ctl.upload.UploadToPublic(c, "thumbnail", "request")
	if upErr != nil {
		return upErr
	}

	if _, err := ctl.service.Store(object.NewRequestInput(
		form.Title,
		form.Content,
		thumbnail,
		form.Status,
		form.TagIDs,
		form.CategoryID,
	), user.ID); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message(), Messages{})
	return response.Redirect(c, "system/request")
}

// Edit render request edit form
func (ctl RequestController) Edit(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	request, findErr := ctl.service.Find(uint(id))
	if findErr != nil {
		return response.ErrorWithCode(findErr, fiber.StatusNotFound)
	}

	tags, tagErr := ctl.tagRepository.GetAll()
	if tagErr != nil {
		return response.Error(tagErr)
	}

	categories, categoryErr := ctl.categoryRepository.GetAll()
	if tagErr != nil {
		return response.Error(categoryErr)
	}

	return response.View("request/edit", helper.DataMap{
		"content_footer": true,
		"statuses":       model.ToRequestStatusArray(),
		"tags":           model.TagsToArray(tags),
		"categories":     model.CategoriesToArray(categories),
		"request":        request,
	})
}

// Update request update process
func (ctl RequestController) Update(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	var form form.Request

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	uID := uint(id)

	if err := form.Validate(
		c,
		ctl.categoryRepository,
		ctl.tagRepository,
	); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	thumbnail, upErr := ctl.upload.UploadToPublic(c, "thumbnail", "request")
	if upErr != nil {
		return upErr
	}

	if _, err := ctl.service.Update(uID, object.NewRequestInput(
		form.Title,
		form.Content,
		thumbnail,
		form.Status,
		form.TagIDs,
		form.CategoryID,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message(), Messages{})
	return response.Back(c)
}

// Delete request delete process
func (ctl RequestController) Delete(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	if err := ctl.service.Delete(uint(id)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message(), Messages{})
	return response.Redirect(c, "system/request")
}
