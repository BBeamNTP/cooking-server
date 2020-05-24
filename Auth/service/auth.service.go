package service

import (
	InAuth "bitbucket.org/BBeamnantapong/cooking-server/Auth"
	"bitbucket.org/BBeamnantapong/cooking-server/config"
	"bitbucket.org/BBeamnantapong/cooking-server/core"
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type AuthService struct {
	ctx core.IContext
}

func NewAuthrService(ctx core.IContext) InAuth.AuthInterface {
	return &AuthService{
		ctx: ctx,
	}
}

// Create User
func (s *AuthService) CreateUserSrv(postAuth *models.Userdata, c echo.Context, Xtoken string) interface{} {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	defer db.Close()
	userData := models.Userdata{}
	userDatas := models.Userdata{}
	db.Last(&userDatas)
	ID, err := strconv.ParseInt(userDatas.ID, 10, 64)
	M := ID + 1
	Str := strconv.FormatInt(M, 10)
	if postAuth.SigninMethod == "facebook.com" {
		err = db.Find(&userData, `email = ? && signin_method = "facebook.com"`, postAuth.Email).Error
		if err != nil {
			log.Print("Error Facebook : ", err)
		}
		log.Print("userData =============================== : ", userData)
		if userData.Email == "" {
			log.Println("Error : ====> ", err)
			log.Print("Fail : 11111111111")
			userData := models.Userdata{
				ID:           Str,
				Email:        postAuth.Email,
				Firstname:    postAuth.Firstname,
				Avatar:       postAuth.Avatar,
				UserId:       Str,
				AdminId:      "0",
				SigninMethod: postAuth.SigninMethod,
			}
			err = db.Create(&userData).Error
			if err != nil {
				log.Println("CreateUserSrv db.Create Error :", err)
				status := models.Status{Message: "Internal server error", Status: false,}
				return c.JSON(http.StatusOK, status)
			}
			Token := models.Token{
				UserId:       Str,
				Email:        postAuth.Email,
				Token:        Xtoken,
				SigninMethod: postAuth.SigninMethod,
			}
			err = db.Create(&Token).Error
			if err != nil {
				log.Println("Insert Token error :", err)
				status := models.Status{Message: "Internal server error", Status: false,}
				return c.JSON(http.StatusOK, status)
			}
			return c.JSON(http.StatusOK, userData)
		}
		log.Print("Xtoken : ", Xtoken)
		Token := models.Token{
			UserId:       userData.ID,
			Email:        userData.Email,
			Token:        Xtoken,
			SigninMethod: userData.SigninMethod,
		}
		err = db.Create(&Token).Error
		if err != nil {
			log.Println("Insert Token error :", err)
			status := models.Status{Message: "Internal server error", Status: false,}
			return c.JSON(http.StatusOK, status)
		}
		log.Print("Fail : 2222222222")
		return c.JSON(http.StatusOK, userData)
	}else if postAuth.SigninMethod == "google.com"{
		err = db.Find(&userData, `email = ? && signin_method = "google.com"`, postAuth.Email).Error
		if err != nil {
			log.Print("Error google.com : ", err)
		}
		log.Print("userData =============================== : ", userData)
		if userData.Email == "" {
			log.Println("Error : ====> ", err)
			log.Print("Fail : 11111111111")
			userData := models.Userdata{
				ID:           Str,
				Email:        postAuth.Email,
				Firstname:    postAuth.Firstname,
				Avatar:       postAuth.Avatar,
				UserId:       Str,
				AdminId:      "0",
				SigninMethod: postAuth.SigninMethod,
			}
			err = db.Create(&userData).Error
			if err != nil {
				log.Println("CreateUserSrv db.Create Error :", err)
				status := models.Status{Message: "Internal server error", Status: false,}
				return c.JSON(http.StatusOK, status)
			}
			Token := models.Token{
				UserId:       Str,
				Email:        postAuth.Email,
				Token:        Xtoken,
				SigninMethod: postAuth.SigninMethod,
			}
			err = db.Create(&Token).Error
			if err != nil {
				log.Println("Insert Token error :", err)
				status := models.Status{Message: "Internal server error", Status: false,}
				return c.JSON(http.StatusOK, status)
			}
			return c.JSON(http.StatusOK, userData)
		}
		log.Print("Xtoken : ", Xtoken)
		Token := models.Token{
			UserId:       userData.ID,
			Email:        userData.Email,
			Token:        Xtoken,
			SigninMethod: userData.SigninMethod,
		}
		err = db.Create(&Token).Error
		if err != nil {
			log.Println("Insert Token error :", err)
			status := models.Status{Message: "Internal server error", Status: false,}
			return c.JSON(http.StatusOK, status)
		}
		log.Print("Fail : Google.com 2222222222")
		return c.JSON(http.StatusOK, userData)
	} else {
		err = db.Find(&userData, `email = ? && signin_method = ""`, postAuth.Email).Error
		log.Println("CreateUserSrv db.Find Error :", err)
		if userData.Email == postAuth.Email {
			status := models.Status{Message: "This email is already used", Status: false,}
			return c.JSON(http.StatusOK, status)
		} else {
			hash, _ := HashPassword(postAuth.Password)
			log.Printf("password :%s hash :%s", postAuth.Password, hash)
			log.Println("userdata ======> ", postAuth)
			Data := models.Userdata{
				Email:     postAuth.Email,
				Password:  hash,
				Genderid:  postAuth.Genderid,
				Titleid:   postAuth.Titleid,
				Firstname: postAuth.Firstname,
				Lastname:  postAuth.Lastname,
				UserId:    Str,
				AdminId:   "0",
			}
			err = db.Create(&Data).Error
			if err != nil {
				log.Println("CreateUserSrv db.Create Error :", err)
				status := models.Status{Message: "Internal server error", Status: false,}
				return c.JSON(http.StatusOK, status)
			}
			status := models.Status{Message: "Register successful", Status: true,}
			return c.JSON(http.StatusOK, status)
		}
	}
	return nil
}

//Signin
func (s *AuthService) CheckUserSrv(postAuth *models.Userdata, c echo.Context) interface{} {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	defer db.Close()
	userData := models.Userdata{}
	status := models.Login{}
	err = db.Find(&userData, `email = ? && signin_method = ?`, postAuth.Email, postAuth.SigninMethod).Error
	if err != nil {
		log.Println("CheckUserSrv db.Find Error :", err)
		status := models.Status{Message: "Internal server error", Status: false,}
		return c.JSON(http.StatusOK, status)
	}
	match := CheckPasswordHash(postAuth.Password, userData.Password)
	if strings.ToLower(userData.Email) == strings.ToLower(postAuth.Email) && match == true {
		validToken, err := GenerateJWT(postAuth.Email)
		log.Println("GenerateJWT error :", err)
		//Token := models.Token{Email: postAuth.Email, Token: validToken,}
		Token := models.Token{
			UserId:       userData.ID,
			Email:        userData.Email,
			Token:        validToken,
			SigninMethod: userData.SigninMethod,
		}
		err = db.Create(&Token).Error
		log.Println("CheckUserSrv db.Create Error :", err)
		status = models.Login{AccStat: true, Message: "Login success", AccToken: validToken, User: userData.Firstname, Avatar: userData.Avatar, UserId: userData.UserId, AdminId: userData.AdminId}
		return c.JSON(http.StatusOK, status, )
	} else {
		status = models.Login{AccStat: false, Message: "E-mail or Password is incorrect"}
		return c.JSON(http.StatusOK, status, )
	}
	return nil
}

//Signout
func (s *AuthService) LogoutSrv(Xtoken string, c echo.Context) interface{} {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	defer db.Close()
	Token := models.Token{}
	status := models.Status{}
	err = db.Find(&Token, `token = ?`, Xtoken).Delete(&Token, `token = ?`, Xtoken).Error
	if err != nil {
		log.Println("LogoutSrv db.Delete Error :", err)
		status := models.Status{Message: "Internal server error", Status: false,}
		return c.JSON(http.StatusOK, status)
	}
	status = models.Status{Message: "Logout success", Status: true}
	return c.JSON(http.StatusOK, status, )
}

// Reset Password
func (s *AuthService) ResetPasswordSrv(postData *models.ResetPassword, c echo.Context) interface{} {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	defer db.Close()
	otp := models.Otps{}
	userData := models.Userdata{}
	err = db.Find(&otp, `email = ?`, postData.Email).Error
	err = db.Find(&userData, `email = ?`, postData.Email).Error
	TimeNow := time.Now()
	if otp.Otp == postData.OTP && inTimeSpan(otp.StartTime, otp.EndTime, TimeNow) {
		hash, _ := HashPassword(postData.NewPassword)
		err = db.Model(&userData).Where("email = ? && signin_method = ? ", userData.Email, "").Update("password", hash).Error
		log.Println("Error Change Password :", err)
		err = db.Find(&otp, `email = ?`, postData.Email).Delete(&otp, `email = ?`, postData.Email).Error
		log.Println("Error Delete OTP :", err)

		status := models.Status{Message: "Reset password success", Status: true,}
		return c.JSON(http.StatusOK, status)
	log.Print("*************************************123456********************************")
	} else {
		status := models.Status{Message: "Reset password Fail", Status: false,}
		return c.JSON(http.StatusOK, status)
	}
	return nil
}

func inTimeSpan(start, end time.Time, check time.Time) bool {
	start = start.Local()
	end = end.Local()
	log.Println("Time Start :", start)
	log.Println("Time End :", end)
	log.Println("Time Check :", check)
	return check.After(start) && check.Before(end)
}

//Send OTP And URL Reset Password With E-Mail
func (s *AuthService) SendOTPSrv(postAuth *models.Userdata, c echo.Context) interface{} {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	defer db.Close()
	min := 100000
	max := 999999
	OTP := strconv.FormatInt(int64(rand.Intn(max-min)+min), 10)
	OTPs := models.Otps{
		Email:     postAuth.Email,
		Otp:       OTP,
		StartTime: time.Now().In(time.Local),
		EndTime:   time.Now().Add(time.Hour*0 + time.Minute*5 + time.Second*0).In(time.Local),
	}
	log.Println("Time Start =============================== >", OTPs.StartTime)
	log.Println("Time End =============================== >", OTPs.EndTime)
	err = db.Model(&OTPs).Where("email =?", OTPs.Email).Updates(&OTPs).FirstOrCreate(&OTPs).Error
	log.Println("Err :", err)
	if err != nil {
		status := models.Status{Message: "This email doesn't exist in the system", Status: false,}
		return c.JSON(http.StatusOK, status)
	}
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Hostname)
	to := []string{postAuth.Email}

	msg := []byte("Subject:คำขอเปลี่ยนรหัสผ่าน \n\n" + "รหัสคำขอเปลี่ยนรหัสของคุณคือ " + OTP + " กรุณาใช้รหัสภายในเวลา 5 นาที" + "\r คุณสามารถเปลี่ยนรหัสผ่านได้ที่ " + "http://localhost:8080/#/pages/resetpassword \r\n" )
	err = smtp.SendMail("smtp.gmail.com:587", auth, "BanArHanService", to, msg)
	if err != nil {
		status := models.Status{Message: "Send OTP fail", Status: false,}
		return c.JSON(http.StatusOK, status)
		log.Println("Mail Err ====> ", err)
	}
	status := models.Status{Message: "Send OTP success", Status: true,}
	return c.JSON(http.StatusOK, status)
	return nil
}

//HashPassword ...
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash ...
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//GenerateJWT ...
func GenerateJWT(username string) (string, error) {
	MySigningKey := []byte("mysupersecretphrase")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = strings.ToLower(username)
	claims["exp"] = time.Now().Add(time.Second * 30).Unix()
	tokenString, err := token.SignedString(MySigningKey)
	if err != nil {
		log.Println("Error : ====> ", err)
	}
	return tokenString, nil
}

////ExtractClaims ...
//func ExtractClaims(tokenStr string) jwt.MapClaims {
//	hmacSecretString := "mysupersecretphrase"
//	hmacSecret := []byte(hmacSecretString)
//	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
//		// check token signing method etc
//		return hmacSecret, nil
//	})
//	if err != nil {
//		return nil
//	}
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		return claims
//	}
//	return nil
//}
