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

// MenuController is menu
type MenuController struct {
	logger             pkg.Logger
	service            service.MenuService
	userRepository     repository.UserRepository
	categoryRepository repository.CategoryRepository
	tagRepository      repository.TagRepository
	value              support.RequestValue
}

// NewMenuController is create menu controller
func NewMenuController(
	logger pkg.Logger,
	service service.MenuService,
	userRepository repository.UserRepository,
	categoryRepository repository.CategoryRepository,
	tagRepository repository.TagRepository,
	value support.RequestValue,
) MenuController {
	return MenuController{
		logger:             logger,
		service:            service,
		userRepository:     userRepository,
		categoryRepository: categoryRepository,
		tagRepository:      tagRepository,
		value:              value,
	}
}

// Index render menu list
func (ctl MenuController) Index(c *fiber.Ctx) error {
	var form form.MenuSearch
	response := ctl.value.GetResponseHelper(c)

	if err := c.QueryParser(&form); err != nil {
		return response.Error(err)
	}

	menus, err := ctl.service.Search(object.NewMenuSearchInput(
		form.Keyword,
		form.Page,
	), 20)
	if err != nil {
		return response.Error(err)
	}

	return response.View("menu/index", helper.DataMap{
		"menus": menus,
	})
}

// Detail render menu detail
func (ctl MenuController) Detail(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	menu, findErr := ctl.service.Find(uint(id))
	if findErr != nil {
		return response.Error(findErr)
	}

	return response.View("menu/detail", helper.DataMap{
		"content_footer": true,
		"menu":           menu,
	})
}

// Create render menu create form
func (ctl MenuController) Create(c *fiber.Ctx) error {
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

	return response.View("menu/create", helper.DataMap{
		"content_footer": true,
		"processes":      model.ToMenuProcessArray(),
		"statuses":       model.ToMenuStatusArray(),
		"tags":           model.TagsToArray(tags),
		"categories":     model.CategoriesToArray(categories),
		"user":           user,
	})
}

// Store menu store process
func (ctl MenuController) Store(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)

	userID, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	user, findErr := ctl.userRepository.GetOne(uint(userID))
	if findErr != nil {
		return response.Error(findErr)
	}

	var form form.Menu

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(
		ctl.categoryRepository,
		ctl.tagRepository,
	); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	if _, err := ctl.service.Store(object.NewMenuInput(
		form.Title,
		form.Content,
		form.Process,
		form.Status,
		form.TagIDs,
		form.CategoryID,
	), user.ID); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message(), Messages{})
	return response.Redirect(c, "system/menu")
}

// Edit render menu edit form
func (ctl MenuController) Edit(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	menu, findErr := ctl.service.Find(uint(id))
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

	return response.View("menu/edit", helper.DataMap{
		"content_footer": true,
		"processes":      model.ToMenuProcessArray(),
		"statuses":       model.ToMenuStatusArray(),
		"tags":           model.TagsToArray(tags),
		"categories":     model.CategoriesToArray(categories),
		"menu":           menu,
	})
}

// Update menu update process
func (ctl MenuController) Update(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	var form form.Menu

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	uID := uint(id)

	if err := form.Validate(ctl.categoryRepository, ctl.tagRepository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	if _, err := ctl.service.Update(uID, object.NewMenuInput(
		form.Title,
		form.Content,
		form.Process,
		form.Status,
		form.TagIDs,
		form.CategoryID,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message(), Messages{})
	return response.Back(c)
}

// Delete menu delete process
func (ctl MenuController) Delete(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	if err := ctl.service.Delete(uint(id)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message(), Messages{})
	return response.Redirect(c, "system/menu")
}
