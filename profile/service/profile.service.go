package service

import (
	AuthSrv "bitbucket.org/BBeamnantapong/cooking-server/Auth/service"
	"bitbucket.org/BBeamnantapong/cooking-server/config"
	"bitbucket.org/BBeamnantapong/cooking-server/core"
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	InProfile "bitbucket.org/BBeamnantapong/cooking-server/profile"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"os"
)

type ProfileSrv struct {
	ctx core.IContext
}

func NewUserService(ctx core.IContext) InProfile.ProfileInterface {
	return &ProfileSrv{
		ctx: ctx,
	}
}

var token = models.Token{}
var userData = models.Userdata{}

func (s *ProfileSrv) GetUserProfile(Xtoken string, c echo.Context) interface{} {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	defer db.Close()
	err = db.Table("tokens").Where("token = ?", Xtoken).Select("*").Scan(&token).Error
	if err != nil {
		status := models.Status{Message: "invalid Token :", Status: false,}
		return c.JSON(http.StatusOK, status)
	}
	log.Print("token.UserId : ", token.UserId)
	err = db.Table("userdata").Where("userdata.id = ?", token.UserId).
		Select("userdata.id, userdata.email, userdata.genderid, genders.gender, userdata.titleid, " +
			"titlenames.titlename, userdata.firstname, userdata.lastname, userdata.avatar, userdata.created_date").
		Joins("JOIN cooking_server.genders ON userdata.genderid = genders.id").
		Joins("JOIN cooking_server.titlenames ON userdata.titleid = titlenames.id").Scan(&userData).Error
	userData = models.Userdata{
		ID:        userData.ID,
		Email:     userData.Email,
		Password:  "",
		Genderid:  userData.Genderid,
		Gender:    userData.Gender,
		Titleid:   userData.Titleid,
		Titlename: userData.Titlename,
		Firstname: userData.Firstname,
		Lastname:  userData.Lastname,
		Avatar:    userData.Avatar,
	}
	if err != nil {
		log.Println("Get Profile DB Error :", err)
		return c.JSON(http.StatusOK, &userData)
	}
	return c.JSON(http.StatusOK, &userData)
}
func (s *ProfileSrv) Upload(Xtoken string, c echo.Context) interface{} {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	defer db.Close()
	err = db.Table("tokens").Where("token = ?", Xtoken).Select("*").Scan(&token).Error
	if err != nil {
		log.Println("DB err : ", err)
		status := models.Status{Message: "Invalit token", Status: false,}
		return c.JSON(http.StatusOK, status)
	}
	err = db.Table("userdata").Where("userdata.id = ?", token.UserId).
		Select("*").Scan(&userData).Error /// เพื่อ เอาข้อมูล userData
	if err != nil {
		log.Println("userdata Error :", err)
		return c.JSON(http.StatusOK, &userData)
	}
	ID := userData.ID                            // กันตัวแปล หาย
	err = c.Request().ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		status := models.Status{Message: "c.Request().ParseMultipartForm :", Status: false,}
		return c.JSON(http.StatusOK, status)
	}
	// Delete old img avatar
	err = os.Remove("./img/avatar/" + ID + ".png")
	if err != nil {
		log.Println("Delete Error : ", err)
	}
	formdata := c.Request().MultipartForm // ok, no problem so far, read the Form data
	//get the *fileheaders
	files := formdata.File["myFile"] // grab the filenames
	for i, _ := range files {        // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Println(err)
			status := models.Status{Message: "Uploaded fail :", Status: false,}
			return c.JSON(http.StatusOK, status)
		}
		out, err := os.Create("./img/avatar/" + ID + ".png")
		defer out.Close()
		if err != nil {
			log.Println(err)
			status := models.Status{Message: "Uploaded fail :", Status: false,}
			return c.JSON(http.StatusOK, status)
		}
		_, err = io.Copy(out, file) // file not files[i] !
		if err != nil {
			log.Println(err)
			status := models.Status{Message: "Uploaded fail :", Status: false,}
			return c.JSON(http.StatusOK, status)
			return err
		}
		userData := models.Userdata{
			ID:    userData.ID,
			Email: userData.Email,

			Avatar: config.Domain + "/img/avatar/" + ID + ".png",
		}
		log.Println("User Data ========> ", userData.Email)
		err = db.Model(&userData).Where("id = ?", userData.ID).Updates(&userData).Error // Update path imgs
		if err != nil {
			log.Println("err = db.Model(&userData) :", err)
			status := models.Status{Message: "db.Model fail :", Status: false,}
			return c.JSON(http.StatusOK, status)
			return err
		}
	}
	status := models.Status{Message: "Uploaded success", Status: true,}
	return c.JSON(http.StatusOK, status)
}

func (s *ProfileSrv) UpdateUserProfile(Xtoken string, postData *models.Data, c echo.Context) interface{} {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	defer db.Close()
	if postData == nil {
	}
	err = db.Table("tokens").Where("token = ?", Xtoken).Select("*").Scan(&token).Error
	if err != nil {
		log.Println("DB err : ", err)
		status := models.Status{Message: "Invalit token", Status: false,}
		return c.JSON(http.StatusOK, status)
	}
	err = db.Table("userdata").Where("userdata.id = ?", token.UserId).
		Select("userdata.id, userdata.email, userdata.password").Scan(&userData).Error
	ID := userData.ID
	log.Print("ID : ", ID)
	log.Print("Error : ", err)
	if err != nil {
		log.Println("Token Error :", err)
		userData = models.Userdata{
			ID:        userData.ID,
			Email:     userData.Email,
			Password:  "",
			Genderid:  userData.Genderid,
			Gender:    userData.Gender,
			Titleid:   userData.Titleid,
			Titlename: userData.Titlename,
			Firstname: userData.Firstname,
			Lastname:  userData.Lastname,
			Avatar:    userData.Avatar,
		}
		log.Print("---------------------- 0 userData: ",userData)

		return c.JSON(http.StatusOK, &userData)
	}
	//// Check password
	if postData.Password.OldPass == "" && postData.Password.NewPass1 == "" {
		userData = models.Userdata{
			Email:     userData.Email,
			Genderid:  postData.Userdata.Genderid,
			Titleid:   postData.Userdata.Titleid,
			Firstname: postData.Userdata.Firstname,
			Lastname:  postData.Userdata.Lastname,
		}
		log.Print("---------------------- 1 userData: ",userData)

		//err = db.Model(&userData).Where("userdata.email = ?", userData.Email).Updates(&userData).Error
		err = db.Table("userdata").Where("id = ?", ID).Updates(&userData).Error
		log.Println("Error 1: ", err)
		status := models.Status{Message: "Update Profile success and Password not change", Status: true,}
		return c.JSON(http.StatusOK, status)
	}
	match := AuthSrv.CheckPasswordHash(postData.Password.OldPass, userData.Password)
	if match == true {
		hash, _ := AuthSrv.HashPassword(postData.Password.NewPass1) // แปลง Password
		userData = models.Userdata{
			Email:     userData.Email,
			Password:  hash,
			Genderid:  postData.Userdata.Genderid,
			Titleid:   postData.Userdata.Titleid,
			Firstname: postData.Userdata.Firstname,
			Lastname:  postData.Userdata.Lastname,
		}

		log.Print("---------------------- 2 userData: ",userData)
		//err = db.Model(&userData).Where("userdata.emaill = ?", userData.Email).Updates(&userData).Error
		err = db.Table("userdata").Where("id = ?", ID).Updates(&userData).Error
		status := models.Status{Message: "Update Profile and Password success", Status: true,}
		return c.JSON(http.StatusOK, status)
	}
	status := models.Status{Message: "Can't update password , Password worng.", Status: false,}
	return c.JSON(http.StatusOK, status)
}
