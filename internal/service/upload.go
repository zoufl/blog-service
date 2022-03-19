package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(filetype upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileNmae := upload.GetFileName(fileHeader.Filename)
	uploadSagePath := upload.GetSavePath()
	dst := uploadSagePath + "/" + fileNmae

	if !upload.CheckContainExt(filetype, fileNmae) {
		return nil, errors.New("file suffix is not supported")
	}

	if upload.CheckSavePath(uploadSagePath) {
		if err := upload.CreateSavePath(uploadSagePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}

	if upload.CheckMaxSize(filetype, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	if upload.CheckPermission(uploadSagePath) {
		return nil, errors.New("insufficient file permissions")
	}

	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileNmae
	return &FileInfo{Name: fileNmae, AccessUrl: accessUrl}, nil
}
