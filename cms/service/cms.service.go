package service

import (
	InCMS "bitbucket.org/BBeamnantapong/cooking-server/cms"
	"bitbucket.org/BBeamnantapong/cooking-server/config"
	"bitbucket.org/BBeamnantapong/cooking-server/core"
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type CMSSrv struct {
	ctx core.IContext
}

func NewCMSService(ctx core.IContext) InCMS.CMSInterface {
	return &CMSSrv{
		ctx: ctx,
	}
}

var token = models.Token{}
func (s *CMSSrv) GetIngredients(Xtoken string, c echo.Context) interface{} {
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

	Ingredients := models.Ingredients{}
	DataIngredients := models.DataIngredients{}
	ArrayDataIngredients := &[]models.DataIngredients{}
	rows, err := db.Order(`id asc`).Table("ingredients").Select(`*`).Rows()
	defer rows.Close()
	if err != nil {
		return err
		log.Println(err)
	}
	for rows.Next() {
		db.ScanRows(rows, &Ingredients)
		DataIngredients.Ingredients = append(DataIngredients.Ingredients, Ingredients)
	}
	DataIngredients.ID = "001"
	DataIngredients.Name = "วัตถุดิบ"
	*ArrayDataIngredients = append(*ArrayDataIngredients, DataIngredients)

	return c.JSON(http.StatusOK, DataIngredients)
}
func (s *CMSSrv) CMSCreateMenu(Xtoken string, c echo.Context, postMenu *models.DataIngredients) interface{} {
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
	userData := models.Userdata{}
	err = db.Table("userdata").Where("id = ?", token.UserId).Select("*").Scan(&userData).Error
	log.Print("Userdata : ", userData)
	log.Print("Post Menu Ingredients : ", postMenu.Ingredients)

	if postMenu.MenuId == "" || postMenu.MenuId == "0" {
		log.Print("****************************** Create Menu *************************************")
		Menu := models.Menu{}
		db.Last(&Menu)
		ID, err := strconv.ParseInt(Menu.Id, 10, 64)
		M := ID + 1
		Str := strconv.FormatInt(M, 10)
		postMenu.MenuId = Str

		log.Println("Err :", err)
		var X float64
		for j, value2 := range postMenu.Ingredients {
			log.Print("IngredientsName : ", value2.IngredientsName)
			log.Print("IngredientsCalories : ", value2.IngredientsCalories)
			log.Print("Quantity : ", value2.Quantity)
			X = X + (value2.IngredientsCalories * value2.Quantity)
			log.Print("X :", X)
			Foods := models.Foods{
				MenuId:        Str,
				IngredientsId: value2.Id,
				Quantity:      value2.Quantity,
			}
			log.Print("--------------------------------")
			log.Print("Loop : ", j)
			log.Print("Foods :", Foods)
			err = db.Model(&Foods).Create(&Foods).Error
		}
		log.Print("Sum X :", X)
		T := time.Now().In(time.Local).AddDate(543, 0, 0).Format("02-01-2006")
		Menus := models.Menu{
			Id:           Str,
			MenuName:     postMenu.MenusName,
			CategoryId:   postMenu.CategoryId,
			Point:        0,
			UserId:       userData.UserId,
			AdminId:      userData.AdminId,
			Method:       postMenu.Method,
			MenuCalories: X,
			CreatedDate:  T,
			ImgPath:      nil,
			Userdata:     nil,
		}

		err = db.Table("menus").Create(&Menus).Error
	} else {
		log.Print("****************************** Update Menu *************************************")

		//Menu := models.Menu{}
		Foods := models.Foods{}

		//err = db.Model(&Menu).Where("id = ?", postMenu.MenuId).Delete(&Menu).Error
		//if err != nil {
		//	log.Println("Error Delete Menu : ====> ", err)
		//	return err
		//}
		err = db.Model(&Foods).Where("Menu_id = ?", postMenu.MenuId).Delete(&Foods).Error
		if err != nil {
			log.Println("Error Delete Foods : ====> ", err)
			return err
		}

		var X float64
		for _, value2 := range postMenu.Ingredients {
			X = X + (value2.IngredientsCalories * value2.Quantity)
			log.Print("X :", X)
			Foods := models.Foods{
				MenuId:        postMenu.MenuId,
				IngredientsId: value2.Id,
				Quantity:      value2.Quantity,
			}
			log.Print("--------------------------------")
			err = db.Model(&Foods).Create(&Foods).Error
		}
		log.Print("Sum X :", X)
		Menus := models.Menu{
			Id:           postMenu.MenuId,
			MenuName:     postMenu.MenusName,
			CategoryId:   postMenu.CategoryId,
			Point:        0,
			MenuCalories: X,
			Method:       postMenu.Method,
		}
		err = db.Table("menus").Where(`id = ?`, postMenu.MenuId).Update(&Menus).Error
	}

	return c.JSON(http.StatusOK, postMenu)

}
func (s *CMSSrv) CMSDeleteMenu(Xtoken string, c echo.Context, postMenu *models.DataIngredients) interface{} {
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
	userData := models.Userdata{}
	FoodIMG := models.Foodsimg{}
	err = db.Table("userdata").Where("user_id = ?", token.UserId).Select("*").Scan(&userData).Error
	log.Print("Userdata : ", userData)
		Menu := models.Menu{}
		Foods := models.Foods{}
		err = db.Model(&Menu).Where("id = ?", postMenu.MenuId).Delete(&Menu).Error
		if err != nil {
			log.Println("Error Delete Menu : ====> ", err)
			return err
		}
		err = db.Model(&Foods).Where("menu_id = ?", postMenu.MenuId).Delete(&Foods).Error
		if err != nil {
			log.Println("Error Delete Foods : ====> ", err)
			return err
		}

	err = db.Table(`foodsimg`).Where("menu_id = ?", postMenu.MenuId).Delete(&FoodIMG).Error
	if err != nil {
		log.Println("Error Delete FoodsIMG : ====> ", err)
		return err
	}
	err = os.RemoveAll("./img/imgfood/"+"menuid"+postMenu.MenuId)  // Remove dir
	if err == nil || os.IsExist(err) {
		log.Print("No error Remove Dir. ")
	}
	log.Print("Error Remove Dir.", err)

	status := models.Status{Message: "Delete menu successful", Status: true,}
	log.Print("status :", status)
	return c.JSON(http.StatusOK, status)
}
func (s *CMSSrv) CMSGetDetailUpdateMenu(Xtoken string, c echo.Context, postMenu *models.DataIngredients) interface{} {
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
	log.Print("PostMenu.MenuID : ", postMenu.MenuId)
	Menus := models.Menu{}
	Ingredients := models.Ingredients{}
	DataFoods := models.DataFoods{}
	Foodsimg := models.Foodsimg{}
	Database := models.Database{
		Id:   "001",
		Name: "วัตถุดิบ",
	}
	err = db.Where("id = ?", postMenu.MenuId).Select("*").Model(&Menus).Scan(&Menus).Error
	log.Print("Error :", err)
	log.Print("postMenu.ID :", postMenu.MenuId)

	rows, err := db.Order(`id asc`).Table("foods").Where("menu_id = ?", postMenu.MenuId).
		Select(`ingredients.id, ingredients.ingredients_name, ingredients.ingredients_calories, ingredients.ingredients_type, foods.quantity, ingredients.type`).
		Joins("JOIN cooking_server.ingredients ON foods.ingredients_id = ingredients.id").
		Rows()
	defer rows.Close()
	if err != nil {
		log.Println("Error Foods : ", err)
	}

	for rows.Next() {
		db.ScanRows(rows, &Ingredients)
		Ingredients.Database = Database
		DataFoods.Ingredients = append(DataFoods.Ingredients, Ingredients)
	}
	rows, err = db.Order(`id asc`).Table("foodsimg").Where("menu_id = ?", postMenu.ID).
		Select(`*`).
		Rows()
	defer rows.Close()
	if err != nil {
		log.Println("Error img : ", err)
	}

	for rows.Next() {
		db.ScanRows(rows, &Foodsimg)
		DataFoods.ImgPath = append(DataFoods.ImgPath, Foodsimg)
	}
	DataFoods.Databases = append(DataFoods.Databases, Database)
	log.Print("DataFood :", DataFoods)
	DataFoods = models.DataFoods{
		ID:           "001",
		MenuId:       Menus.Id,
		MenusName:    Menus.MenuName,
		CategoryId:   Menus.CategoryId,
		Point:        Menus.Point,
		UserId:       Menus.UserId,
		AdminId:      Menus.AdminId,
		Method:       Menus.Method,
		Name:         "วัตถุดิบ",
		MenuCalories: Menus.MenuCalories,
		Ingredients:  DataFoods.Ingredients,
		ImgPath:      DataFoods.ImgPath,
		Databases:    DataFoods.Databases,
	}

	return c.JSON(http.StatusOK, DataFoods)
}
func (s *CMSSrv) CMSManageMenu(Xtoken string, c echo.Context, postMenu *models.Menu) interface{}  {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println("Error : ====> ", err)
		return err
	}
	defer db.Close()
	Menu := models.Menu{}
	//ArrayMenu := []models.Menu{}
	Foodsimg := models.Foodsimg{}
	//ArrayMenuss := []models.ArrayMenu{}
	UserData := models.Usersdata{}
	ArrayMenus := models.ArrayMenu{}

	err = db.Table("tokens").Where("token = ?", Xtoken).Select("*").Scan(&token).Error
	if err != nil {
		status := models.Status{Message: "invalid Token :", Status: false,}
		return c.JSON(http.StatusOK, status)
	}
	//err = db.Table("userdata").Where("user_id = ?", token.UserId).Select("*").Scan(&userData).Error // Get userData
	rows, err := db.Order(`id asc`).Table("userdata").Where("id = ?", token.UserId).
		Select(`*`).
		Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		db.ScanRows(rows, &UserData)
	}
	log.Print("Post Menu : ", postMenu)
	log.Print("postMenu.AdminId : ", postMenu.AdminId)
	log.Print("postMenu.UserId : ", postMenu.UserId)
	log.Print("UserData.AdminId : ", UserData.AdminId)
	if UserData.AdminId == "1"{
		log.Print("*************** Admin 1 ")
		rows, err = db.Order(`id asc`).Table("menus").Select(`*`).Rows()
		defer rows.Close()
		if err != nil {
			return err
		}
		for rows.Next() {
			db.ScanRows(rows, &Menu)
			Menu.ImgPath = nil
			Menu.Userdata = nil
			rows, err := db.Order(`id asc`).Table("foodsimg").Where("menu_id = ?", Menu.Id).
				Select(`*`).
				Rows()
			defer rows.Close()
			if err != nil {
				log.Println(err)
			}
			for rows.Next() {
				db.ScanRows(rows, &Foodsimg)
				Menu.ImgPath = append(Menu.ImgPath, Foodsimg)
			}
			Menu.Userdata = append(Menu.Userdata, UserData)
			ArrayMenus.Menu = append(ArrayMenus.Menu, Menu)
		}
	}else{
		log.Print("****************** User ")
		rows, err = db.Order(`id asc`).Table("menus").Where(" admin_id = ? and user_id = ?", UserData.AdminId, UserData.UserId).Select(`*`).Rows()
		defer rows.Close()
		if err != nil {
			return err
		}
		for rows.Next() {
			db.ScanRows(rows, &Menu)
			Menu.ImgPath = nil
			Menu.Userdata = nil
			rows, err := db.Order(`id asc`).Table("foodsimg").Where("menu_id = ?", Menu.Id).
				Select(`*`).
				Rows()
			defer rows.Close()
			if err != nil {
				log.Println(err)
			}
			for rows.Next() {
				db.ScanRows(rows, &Foodsimg)
				Menu.ImgPath = append(Menu.ImgPath, Foodsimg)
			}
			Menu.Userdata = append(Menu.Userdata, UserData)
			ArrayMenus.Menu = append(ArrayMenus.Menu, Menu)
		}
	}


	return c.JSON(http.StatusOK, ArrayMenus.Menu)
}
func (s *CMSSrv) UploadFile(Xtoken string, c echo.Context) interface{} {
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

	err = c.Request().ParseMultipartForm(200000) // grab the multipart form
	log.Print("Error : ", err)

	if err != nil {
		log.Println("Error ========================== ", err)
		status := models.Status{Message: "c.Request().ParseMultipartForm :", Status: false,}
		return c.JSON(http.StatusOK, status)
	}

	formdata := c.Request().MultipartForm // ok, no problem so far, read the Form data
	var MenuID string
	menuID := formdata.Value["menuID"] // Get Menuid
	if menuID == nil {
		log.Print("files : No Data")
	}
	log.Print("files ==========> ", menuID)
	for i, value := range menuID {
		log.Print(i, " MenuID :", value)
		MenuID = value
	}
	//get the *fileheaders
	log.Print("MenuID :", MenuID)
	err = os.RemoveAll("./img/imgfood/"+"menuid"+MenuID)  // Remove dir
	if err == nil || os.IsExist(err) {
		log.Print("No error Remove Dir. ")
	}
	log.Print("Error Remove Dir.", err)

	err = os.MkdirAll("./img/imgfood/"+"menuid"+MenuID+"/", os.ModeDir) // crate dir
	if err == nil || os.IsExist(err) {
		log.Print("No error Crate Dir.")
	} else {
		return err
	}
	files := formdata.File["myFile"] // grab the filenames
	if files == nil {
		log.Print("files : No Data")
	}
	log.Print("files ==========> ", files)
	FoodsIMG := models.Foodsimg{}
	err = db.Table("foodsimg").Where(`menu_id = ?`, menuID).Delete(&FoodsIMG).Error
	if err != nil {
		log.Print("Error Delete img : ", err)
	}

	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()

		if err != nil {
			log.Println(err)
			status := models.Status{Message: "Uploaded fail :", Status: false,}
			return c.JSON(http.StatusOK, status)
		}

		out, err := os.Create("./img/imgfood/" + "menuid" + MenuID + "/"+files[i].Filename)
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
		log.Println("Files uploaded successfully : ")
		log.Println(MenuID + "-" + strconv.Itoa(i) + ".png" + "\n")
		FoodsIMG := models.Foodsimg{
			MenuId: MenuID,
			Href:   config.Domain + "/img/imgfood/" + "menuid" + MenuID + "/"+files[i].Filename,
		}
		log.Print("FoodsIMG :", FoodsIMG)
		err = db.Table("foodsimg").Create(&FoodsIMG).Error
		if err != nil {
			log.Print("Error :", err)
		}
	}

	status := models.Status{Message: "Uploaded success", Status: true,}
	return c.JSON(http.StatusOK, status)
}
