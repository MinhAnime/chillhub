package minio

import (
	"path"
	"strings"
)

func buildObjectPath(basePath, filename string) string {
	basePath = strings.Trim(basePath, "/")
	return path.Join(basePath, filename)
}
