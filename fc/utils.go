package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/GZGavinZhao/libeopkg/archive"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

func sha1AndSize(path string, ctx context.Context) (string, int64, error) {
	fcctx, _ := fccontext.FromContext(ctx)
	// credentials := fcctx.Credentials
	logger := fcctx.GetLogger()

	f, err := os.Open(path)
	if err != nil {
		logger.Error(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		logger.Error(err)
	}

	stat, err := f.Stat()
	if err != nil {
		logger.Error(err)
	}

	return hex.EncodeToString(h.Sum(nil)[:]), stat.Size(), nil
}

func isEopkg(path string) bool {
	return filepath.Ext(path) == ".eopkg"
}

func isDelta(path string) bool {
	return len(path) >= 12 && path[len(path)-12:] == ".delta.eopkg"
}

// verifyPackageLocation verifies that the package is at its proper location
// defined by `eopkg`. Note that `path` and `repoDir` must be absolute paths.
func verifyPackageLocation(path string, pkg archive.Package, repoDir string) error {
	expectedLocation := filepath.Join(repoDir, pkg.GetPathComponent(), filepath.Base(path))
	if path != expectedLocation {
		return errors.New(fmt.Sprintf("The package %s is not at the right location: expecting %s, found at %s", pkg.Name, expectedLocation, path))
	}

	return nil
}