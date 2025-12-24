package minio

import (
	"context"
	"mime/multipart"
	"net/http"
	"time"

	appErr "chillhub/internal/shared/error"

	"github.com/minio/minio-go/v7"
)

type Util struct {
	storage *Client
}

func NewUtil(client *Client) *Util {
	return &Util{storage: client}
}

func (u *Util) Upload(
	ctx context.Context,
	file *multipart.FileHeader,
	bucket string,
	basePath string,
) (*ObjectInfo, error) {

	src, err := file.Open()
	if err != nil {
		return nil, appErr.New(
			http.StatusInternalServerError,
			"file.open.failed",
		)
	}
	defer src.Close()

	object := buildObjectPath(basePath, file.Filename)

	_, err = u.storage.cli.PutObject(
		ctx,
		bucket,
		object,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return nil, appErr.Wrap(
			err,
			http.StatusInternalServerError,
			"minio.upload.failed",
		)
	}

	return &ObjectInfo{
		Bucket: bucket,
		Object: object,
		URL:    u.storage.baseURL + "/" + bucket + "/" + object,
	}, nil
}

func (u *Util) Delete(
	ctx context.Context,
	bucket string,
	object string,
) error {

	err := u.storage.cli.RemoveObject(
		ctx,
		bucket,
		object,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return appErr.Wrap(
			err,
			http.StatusInternalServerError,
			"minio.delete.failed",
		)
	}

	return nil
}

func (u *Util) PresignGet(
	ctx context.Context,
	bucket string,
	object string,
	expiry time.Duration,
) (string, error) {

	url, err := u.storage.cli.PresignedGetObject(
		ctx,
		bucket,
		object,
		expiry,
		nil,
	)
	if err != nil {
		return "", appErr.Wrap(
			err,
			http.StatusInternalServerError,
			"minio.presign.get.failed",
		)
	}

	return url.String(), nil
}

func (u *Util) PresignPut(
	ctx context.Context,
	bucket string,
	object string,
	expiry time.Duration,
) (string, error) {

	url, err := u.storage.cli.PresignedPutObject(
		ctx,
		bucket,
		object,
		expiry,
	)
	if err != nil {
		print("error: ",err)
		return "", appErr.Wrap(
			err,
			http.StatusInternalServerError,
			"minio.presign.put.failed",
		)
	}

	return url.String(), nil
}


func (u *Util) EnsureBucket(
	ctx context.Context,
	bucket string,
) error {

	exists, err := u.storage.cli.BucketExists(ctx, bucket)
	if err != nil {
		return appErr.Wrap(err, http.StatusInternalServerError, "minio.bucket.check.failed")
	}

	if exists {
		return nil
	}

	err = u.storage.cli.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
	if err != nil {
		return appErr.Wrap(err, http.StatusInternalServerError, "minio.bucket.create.failed")
	}

	return nil
}
