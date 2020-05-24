package service

import (
	"bitbucket.org/BBeamnantapong/cooking-server/core"
	InMenu "bitbucket.org/BBeamnantapong/cooking-server/foodsmenu"
	"bitbucket.org/BBeamnantapong/cooking-server/models"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type MenuSrv struct {
	ctx core.IContext
}

func NewMenuService(ctx core.IContext) InMenu.MenuInterface {
	return &MenuSrv{
		ctx: ctx,
	}
}

var token = models.Token{}
var userData = models.Userdata{}

// เมนูทั้งหมด
func (s *MenuSrv) GetMenu(Xtoken string, c echo.Context) interface{} {
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
	Menu := models.Menu{}
	Foodsimg := models.Foodsimg{}
	UserData := models.Usersdata{}
	ArrayMenus := models.ArrayMenu{}

	rows, err := db.Order(` point DESC`).Table("menus").Select(`*`).Rows() // มากไปน้อย
	defer rows.Close()
	if err != nil {
		return err
		log.Println(err)
	}
	Loop := 0
	for rows.Next() {
		if Loop <= 9 { // เอาแค่ 10 เมนูแรก
			Loop = Loop + 1
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

			rows, err = db.Order(`id asc`).Table("userdata").Where("user_id = ?", Menu.UserId).
				Select(`*`).
				Rows()
			defer rows.Close()
			if err != nil {
				log.Println(err)
			}
			for rows.Next() {
				db.ScanRows(rows, &UserData)
				Menu.Userdata = append(Menu.Userdata, UserData)
			}
			ArrayMenus.Menu = append(ArrayMenus.Menu, Menu)

		}

	}


	return c.JSON(http.StatusOK, ArrayMenus.Menu)
}
// รายละเอียดอาหาร
func (s *MenuSrv) GetFood(Xtoken string, c echo.Context, postMenu *models.DataIngredients) interface{} {
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
	Menus := models.Menu{}
	Ingredients := models.Ingredients{}
	DataFoods := models.DataFoods{}
	Foodsimg := models.Foodsimg{}

	err = db.Where("id = ?", postMenu.ID).Select("*").Model(&Menus).Scan(&Menus).Error
	log.Print("Error :", err)
	log.Print("postMenu.ID :", postMenu.ID)

	rows, err := db.Order(`id asc`).Table("foods").Where("menu_id = ?", postMenu.ID).
		Select(`ingredients.id, ingredients.ingredients_name, ingredients.ingredients_calories, ingredients.ingredients_type, foods.quantity, ingredients.type`).
		Joins("JOIN cooking_server.ingredients ON foods.ingredients_id = ingredients.id").
		Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		db.ScanRows(rows, &Ingredients)
		DataFoods.Ingredients = append(DataFoods.Ingredients, Ingredients)
	}
	rows, err = db.Order(`id asc`).Table("foodsimg").Where("menu_id = ?", postMenu.ID).
		Select(`*`).
		Rows()
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		db.ScanRows(rows, &Foodsimg)
		DataFoods.ImgPath = append(DataFoods.ImgPath, Foodsimg)
	}

	log.Print("DataFood :", DataFoods)
	DataFoods = models.DataFoods{
		MenuId:       Menus.Id,
		MenusName:    Menus.MenuName,
		CategoryId:   Menus.CategoryId,
		Point:        Menus.Point,
		UserId:       Menus.UserId,
		AdminId:      Menus.AdminId,
		Method:       Menus.Method,
		MenuCalories: Menus.MenuCalories,
		Ingredients:  DataFoods.Ingredients,
		ImgPath:      DataFoods.ImgPath,
	}
	log.Print("DataFood : ", DataFoods)
	return c.JSON(http.StatusOK, DataFoods)
}
// ให้ดาว
func (s *MenuSrv) GetPoint(Xtoken string, c echo.Context, postMenu *models.Menu) interface{} {
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

	points := models.Points{}
	ArrayPoint := models.DataPoints{}
	point := models.Points{
		MenuId: postMenu.Id,
		UserId: postMenu.UserId,
		Point:  postMenu.Point,
	}
	err = db.Table("points").Where("user_id =? && menu_id = ?", postMenu.UserId, postMenu.Id).
		Select("*").Scan(&points).Error
	if err != nil {
		log.Println("Error : ====> ", err)
	}
	if points.UserId == postMenu.UserId && points.UserId != "" { ///// if ผิด แก้ด้วยยยยยยย
		if points.MenuId == postMenu.Id && points.MenuId != "" {
			err = db.Model(&point).Where("user_id =? && menu_id = ?", postMenu.UserId, postMenu.Id).Updates(&point).Error
			if err != nil {
				log.Print("Error Update 153 : ", err)
				status := models.Status{Message: "Update point error ", Status: false,}
				return c.JSON(http.StatusOK, status)
			}
			rows, err := db.Where("menu_id = ?", postMenu.Id).Order(`id asc`).Table("points").Select(`point`).Rows()
			defer rows.Close()
			if err != nil {
				return err
				log.Println(err)
			}
			for rows.Next() {
				db.ScanRows(rows, &point)
				ArrayPoint.Points = append(ArrayPoint.Points, point)
			}
			var X float64
			var loop float64
			for _, value := range ArrayPoint.Points {
				log.Print("sum point top : ", X)

				X = X + value.Point
				loop = loop + 1
				log.Print("sum point down : ", X)

			}
			X = X / loop
			Menu := models.Menu{
				Point: X,
			}
			log.Print("sum point : ", X)
			err = db.Where("id = ?", postMenu.Id).Model(&Menu).Update(&Menu).Error
			status := models.Status{Message: "Update point success ", Status: true,}
			return c.JSON(http.StatusOK, status)
		}
		log.Print("Error Update 160 : ", err)
		status := models.Status{Message: "Update point error ", Status: false,}
		return c.JSON(http.StatusOK, status)
	}
	err = db.Model(&point).Create(&point).Error
	if err != nil {
		status := models.Status{Message: "Create point error :", Status: false,}
		return c.JSON(http.StatusOK, status)
	}

	rows, err := db.Where("menu_id = ?", postMenu.Id).Order(`id asc`).Table("points").Select(`point`).Rows()
	defer rows.Close()
	if err != nil {
		return err
		log.Println(err)
	}
	for rows.Next() {
		db.ScanRows(rows, &point)
		ArrayPoint.Points = append(ArrayPoint.Points, point)
	}
	var X float64
	var loop float64
	for _, value := range ArrayPoint.Points {
		log.Print("sum point top : ", X)

		X = X + value.Point
		loop = loop + 1
		log.Print("sum point down : ", X)

	}
	X = X / loop
	Menu := models.Menu{
		Point: X,
	}
	log.Print("sum point : ", X)
	err = db.Where("id = ?", postMenu.Id).Model(&Menu).Update(&Menu).Error
	status := models.Status{Message: "Create point success ", Status: true,}
	return c.JSON(http.StatusOK, status)
}
// ประเภทเมนู
func (s *MenuSrv) GetCategory(Xtoken string, c echo.Context, postMenu *models.Menu) interface{} {
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
	log.Print("Post Menu : ", postMenu)
	if postMenu.AdminId == "1" {
		rows, err := db.Where("category_id = ? and admin_id = ?", postMenu.CategoryId, postMenu.AdminId).
			Order(`id asc`).Table("menus").Select(`*`).Rows()
		defer rows.Close()
		if err != nil {
			return err
			log.Println(err)
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
			rows, err = db.Order(`id asc`).Table("userdata").Where("user_id = ?", Menu.UserId).
				Select(`*`).
				Rows()
			defer rows.Close()
			if err != nil {
				log.Println(err)
			}
			for rows.Next() {
				db.ScanRows(rows, &UserData)
				Menu.Userdata = append(Menu.Userdata, UserData)
			}
			ArrayMenus.Menu = append(ArrayMenus.Menu, Menu)
		}
		log.Print("ArrayMenus.Menu : ", ArrayMenus.Menu)
	} else {

		log.Print("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*------------------ 2")
		rows, err := db.Where(" admin_id != 1").Order(`id asc`).Table("menus").Select(`*`).Rows()
		defer rows.Close()
		if err != nil {
			return err
			log.Println(err)
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
			rows, err = db.Order(`id asc`).Table("userdata").Where("user_id = ?", Menu.UserId).
				Select(`*`).
				Rows()
			defer rows.Close()
			if err != nil {
				log.Println(err)
			}
			for rows.Next() {
				db.ScanRows(rows, &UserData)
				Menu.Userdata = append(Menu.Userdata, UserData)
			}
			ArrayMenus.Menu = append(ArrayMenus.Menu, Menu)
		}

		//for j, value2 := range ArrayMenus.Menu {
		//	Menu.ImgPath = nil
		//	rows, err := db.Order(`id asc`).Table("foodsimg").Where("menu_id = ?", value2.Id).
		//		Select(`*`).
		//		Rows()
		//	defer rows.Close()
		//	if err != nil {
		//		log.Println(err)
		//	}
		//
		//	for rows.Next() {
		//		db.ScanRows(rows, &Foodsimg)
		//		Menu.ImgPath = append(Menu.ImgPath, Foodsimg)
		//	}
		//	log.Print(j, ": ArrayFoodsimg : ", Menu.ImgPath)
		//}
	}

	return c.JSON(http.StatusOK, ArrayMenus.Menu)
}
// ค้นหาเมนู
func (s *MenuSrv) Search(Xtoken string, c echo.Context, postIngredients *models.ArrayIngredients) interface{} {
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
	Food := models.Food{}
	ArrayFoods := models.ArrayFoods{}
	ArrayFoods2 := models.ArrayFoods{}

	ArrayDataFoods := []models.ArrayFoods{}
	X := 0
	for i, value := range postIngredients.Ingredients {
		X = X+1
		//ArrayFoods.Food = nil
		log.Print(i, ": Value ID : ", value.Id)
		log.Print(i, ": Value Name : ", value.IngredientsName)
		rows, err := db.Order(`id asc`).Table("foods").Where(`ingredients_id = ? `, value.Id).Select(`*`).Rows()
		defer rows.Close()
		if err != nil {
			return err
			log.Println(err)
		}
		if i == 0 {
			for rows.Next() {
				db.ScanRows(rows, &Food)
				log.Print("*********************************** 0")

				log.Print("Food : ", Food)
				ArrayFoods.Food = append(ArrayFoods.Food, Food)
			}
		} else {
			for rows.Next() {
				db.ScanRows(rows, &Food)
				log.Print("Food : ", Food)

				for i, value := range ArrayFoods2.Food {
					log.Print(i, ": *********************************** 2")
					log.Print(i, ": value.MenuId : ", value.MenuId)
					log.Print(i, ": Food.MenuId : ", Food.MenuId)
					if value.MenuId == Food.MenuId {
						log.Print("*********************************** 3")
						ArrayFoods.Food = append(ArrayFoods.Food, Food)
					} else {
						log.Print("Fail")
					}
					log.Print("*********************************** 4")
				}
				log.Print("*********************************** 5")
			}
		}
		log.Print("*********************************** 6")
		ArrayDataFoods = append(ArrayDataFoods, ArrayFoods)
		ArrayFoods2 = ArrayFoods
		ArrayFoods = models.ArrayFoods{}
	}

	log.Print("XXXXXXXXXXXXXXXXXXXXXXXXXXX :", X)
	log.Print(": ArrayFoods top : ", ArrayDataFoods)
	log.Print(": ArrayFoods down : ", ArrayDataFoods)
	X = X-1
	log.Print("ArrayDataFoods X :", ArrayDataFoods[X])
	Menu := models.Menu{}
	//ArrayMenu := []models.Menu{}
	Foodsimg := models.Foodsimg{}
	//ArrayMenuss := []models.ArrayMenu{}
	UserData := models.Usersdata{}
	ArrayMenus := models.ArrayMenu{}

	for _, value := range ArrayDataFoods[X].Food {

		rows, err := db.Where(`id = ?`, value.MenuId ).
			Order(`id asc`).Table("menus").Select(`*`).Rows()
		defer rows.Close()
		if err != nil {
			return err
			log.Println(err)
		}
		for rows.Next() {
			db.ScanRows(rows, &Menu)
			Menu.ImgPath = nil
			Menu.Userdata = nil
			rows, err := db.Order(`id asc`).Table("foodsimg").Where("menu_id = ?", value.MenuId).
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
			rows, err = db.Order(`id asc`).Table("userdata").Where("user_id = ?", Menu.UserId).
				Select(`*`).
				Rows()
			defer rows.Close()
			if err != nil {
				log.Println(err)
			}
			for rows.Next() {
				db.ScanRows(rows, &UserData)
				Menu.Userdata = append(Menu.Userdata, UserData)
			}
			ArrayMenus.Menu = append(ArrayMenus.Menu, Menu)
		}
		log.Print("ArrayMenus.Menu : ", ArrayMenus.Menu)

	}



	return c.JSON(http.StatusOK, ArrayMenus.Menu)

}
