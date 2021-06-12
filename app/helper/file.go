package helper

import (
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/takemo101/go-fiber/pkg"
)

type FileHelper struct {
	Path pkg.Path
}

// NewUploadHelper is create FileHelper
func NewFileHelper(
	path pkg.Path,
) FileHelper {
	return FileHelper{
		Path: path,
	}
}

func (f FileHelper) RemovePublic(p string) error {
	public := f.Path.Public(p)
	return os.Remove(public)
}

func (f FileHelper) RemoveStorage(p string) error {
	storage := f.Path.Storage(p)
	return os.Remove(storage)
}

func CreateUUID() string {
	uuidV4 := uuid.New()
	return uuidV4.String()
}

func ExtractExt(name string) string {
	pos := strings.LastIndex(name, ".")

	return name[pos:]
}
