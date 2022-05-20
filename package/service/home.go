package service

import (
	"io"
	"mime/multipart"
	"norbekov/configs"
	"norbekov/model"
	"norbekov/package/repository"
	"norbekov/util/logrus"
	"os"
)

type HomeService struct {
	repo repository.Home
}

func NewHomeService(repo repository.Home) *HomeService {
	return &HomeService{repo: repo}
}

func (s *HomeService) CreateHomePost(post model.HomePost, logrus *logrus.Logger) (int, error) {
	homeId, err := s.repo.CreateHomePost(post, logrus)
	if err != nil {
		return 0, err
	}
	return homeId, err
}
func (s *HomeService) UpdateHomeImage(homeID int, filePath string, logrus *logrus.Logger) (int64, error) {
	return s.repo.UpdateHomeImage(homeID, filePath, logrus)
}
func (s *HomeService) UpdateHome(id int, home model.HomePost, logrus *logrus.Logger) (int64, error) {
	return s.repo.UpdateHome(id, home, logrus)
}

func (service *HomeService) UploadImage(file multipart.File, header *multipart.FileHeader, logrus *logrus.Logger) (string, error) {

	filename := header.Filename
	folderPath := "public/"
	err := os.MkdirAll(folderPath, 0777)
	if err != nil {
		logrus.Errorf("ERROR: Failed to create folder %s: %v", folderPath, err)
		return "", err
	}
	err = os.Chmod(folderPath, 0777)
	if err != nil {
		logrus.Errorf("ERROR: Failed Access Permission denied %s", err)
		return "", err
	}
	filePath := folderPath + filename
	out, err := os.Create(filePath)
	if err != nil {
		logrus.Errorf("ERROR: Failed CreateFile: %v", err)
		return "", err
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		logrus.Errorf("ERROR: Failed copy %s", err)
		return "", err
	}
	configs, err := configs.InitConfig()
	logrus.Infof("configs %v", configs)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	imageURL := configs.ServiceHost + "/" + filePath
	return imageURL, nil
}

func (s *HomeService) GetHomeById(id string, logrus *logrus.Logger) (model.HomeFull, error) {
	return s.repo.GetHomeById(id, logrus)
}
func (s *HomeService) DeleteHome(id string, logrus *logrus.Logger) error {
	return s.repo.DeleteHome(id, logrus)
}
