package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/middleware"
)

type Toastr string

const (
	ToastrError   Toastr = "error"
	ToastrSuccess Toastr = "success"
	ToastrStore   Toastr = "store"
	ToastrUpdate  Toastr = "update"
	ToastrDelete  Toastr = "delete"
)

func (t Toastr) String() string {
	return string(t)
}

func (t Toastr) Message() string {
	switch t {
	case ToastrError:
		return "入力内容を確認してください"
	case ToastrStore:
		return "追加しました"
	case ToastrUpdate:
		return "更新しました"
	case ToastrDelete:
		return "削除しました"
	}
	return ""
}

func SetToastr(c *fiber.Ctx, ttype Toastr, message string) error {
	return middleware.SetSessionMessages(c, middleware.SessionMessages{
		"toastr_type": string(ttype),
		"message":     message,
	})
}
