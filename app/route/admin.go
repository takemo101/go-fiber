package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/controller/admin"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminRoute is struct
type AdminRoute struct {
	logger          pkg.Logger
	app             pkg.Application
	csrf            middleware.Csrf
	render          *helper.ViewRender
	adminController admin.AdminController
}

// Setup is setup route
func (r AdminRoute) Setup() {
	r.logger.Info("setup admin-route")

	app := r.app.App
	render := r.render

	admin := app.Group(
		"/admin",
		r.csrf.CreateHandler("form:csrf_token"),
		r.render.CreateHandler(r.ViewRenderCreateHandler),
	)
	{
		admin.Get("/", func(c *fiber.Ctx) error {
			render.SetJS(helper.DataMap{
				"Hello": "Yes",
			})
			return render.Render("home", helper.DataMap{
				"Title": "Dashboard",
			})
		})
		user := admin.Group("/admin")
		{
			user.Get("/", r.adminController.Index)
			user.Get("/create", r.adminController.Create)
			user.Post("/store", r.adminController.Store)
			user.Get("/:id/edit", r.adminController.Edit)
			user.Post("/:id/update", r.adminController.Update)
			user.Post("/:id/delete", r.adminController.Delete)
		}
	}
}

// NewAdminRoute create new admin route
func NewAdminRoute(
	logger pkg.Logger,
	app pkg.Application,
	csrf middleware.Csrf,
	render *helper.ViewRender,
	adminController admin.AdminController,
) AdminRoute {
	return AdminRoute{
		logger:          logger,
		app:             app,
		csrf:            csrf,
		render:          render,
		adminController: adminController,
	}
}

// ViewRenderCreateHandler middleware handler
func (r AdminRoute) ViewRenderCreateHandler(c *fiber.Ctx, vr *helper.ViewRender) {
	// load menu list
	adminlte, err := r.app.Config.Load("admin-lte")
	if err == nil {
		for k, v := range map[string]string{
			"adminlte_menus":   "menus",
			"adminlte_plugins": "plugins",
		} {
			if value, ok := adminlte[v]; ok {
				vr.SetData(helper.DataMap{
					k: value,
				})
			}
		}
	}

	// load meta
	meta, err := r.app.Config.Load("admin-meta")
	if err == nil {
		vr.SetData(helper.DataMap{
			"admin_meta": meta,
		})
	}

	// csrf-token
	csrfToken := middleware.GetCSRFToken(c)
	vr.SetData(helper.DataMap{
		"csrf_token": csrfToken,
	})

	// session errors
	errors, _ := middleware.GetSessionErrors(c)
	vr.SetData(helper.DataMap{
		"errors": errors,
	})

	// session inputs
	inputs, _ := middleware.GetSessionInputs(c)
	vr.SetData(helper.DataMap{
		"inputs": inputs,
	})
}
