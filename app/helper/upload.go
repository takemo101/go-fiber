package helper

import (
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/pkg"
)

type UploadHelper struct {
	Path pkg.Path
}

// NewUploadHelper is create UploadHelper
func NewUploadHelper(
	path pkg.Path,
) UploadHelper {
	return UploadHelper{
		Path: path,
	}
}

func (u UploadHelper) UploadToPublic(c *fiber.Ctx, key string, directory string) (string, error) {
	path, err := u.UploadTo(c, key, u.Path.Public(directory))
	if err != nil {
		return "", err
	}
	return u.Path.PublicSubstract(path), err
}

func (u UploadHelper) UploadToStorage(c *fiber.Ctx, key string, directory string) (string, error) {
	path, err := u.UploadTo(c, key, u.Path.Storage(directory))
	if err != nil {
		return "", err
	}
	return u.Path.StorageSubstract(path), err
}

func (u UploadHelper) UploadTo(c *fiber.Ctx, key string, directory string) (string, error) {
	file, err := c.FormFile(key)
	if err != nil {
		return "", nil
	}

	// mkedir
	if f, err := os.Stat(directory); os.IsNotExist(err) || !f.IsDir() {
		os.MkdirAll(directory, 0777)
	}

	filename := CreateUUID()
	ext := ExtractExt(file.Filename)

	name := path.Join(directory, filename+ext)
	saveErr := c.SaveFile(file, name)
	return name, saveErr
}
