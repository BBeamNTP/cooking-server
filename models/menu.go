package models

type DataMenu struct {
	Menu []Menu `json:"Menu"`
}
type Menu struct {
	Id           string      `json:"id"`
	MenuName     string      `json:"menuName"`
	CategoryId   string      `json:"categoryId"`
	Point        float64     `json:"point"`
	UserId       string      `json:"userId"`
	AdminId      string      `json:"adminId"`
	Method       string      `json:"methods"`
	MenuCalories float64     `json:"menuCalories"`
	CreatedDate  string     `json:"createdDate"`
	ImgPath      []Foodsimg  `gorm:"-"  json:"imgPath"`
	Userdata     []Usersdata `gorm:"-"  json:"userdata"`
}

type Usersdata struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Avatar    string `json:"avatar"`
	UserId    string `json:"userId"`
	AdminId   string `json:"adminId"`
}
type ArrayMenu struct {
	Menu []Menu `json:"menu"`
}
type ArrayFoods struct {
	Food []Food `json:"food"`
}
type Food struct {
	Id              string `json:"id"`
	MenuId          string `json:"menuId"`
	IngredientsId   string `json:"ingredientsId"`
	IngredientsName string `json:"ingredientsName"`
}

type Foods struct {
	Id                  string  `gorm:"-" json:"-"`
	MenuId              string  `json:"-"`
	MenuName            string  `gorm:"-" json:"-"`
	IngredientsId       string  `json:"ingredientsId"`
	IngredientsName     string  `gorm:"-" json:"ingredientsName"`
	IngredientsCalories float64 `gorm:"-" json:"ingredientsCalories"`
	IngredientsType     string  `gorm:"-" json:"ingredientsType"`
	Quantity            float64 `json:"quantity"`
}
type DataFoods struct {
	ID           string        `json:"id"`
	MenuId       string        `gorm:"-" json:"menuId"`
	MenusName    string        `json:"menuName"`
	CategoryId   string        `json:"categoryId"`
	Point        float64       `json:"point"`
	UserId       string        `json:"userId"`
	AdminId      string        `json:"adminId"`
	Method       string        `json:"methods"`
	Name         string        `json:"name"`
	MenuCalories float64       `json:"menuCalories"`
	Ingredients  []Ingredients `json:"ingredients"`
	ImgPath      []Foodsimg    `json:"imgPath"`
	Databases    []Database    `gorm:"-" json:"databases"`
}
type Foodsimg struct {
	ID     string `json:"id"`
	MenuId string `json:"menuId"`
	Href   string `json:"href"`
}

type DataIngredients struct {
	ID           string        `json:"id"`
	MenuId       string        `gorm:"-" json:"menuId"`
	MenusName    string        `json:"menuName"`
	CategoryId   string        `json:"categoryId"`
	Point        float64       `json:"point"`
	UserId       string        `json:"userId"`
	AdminId      string        `json:"adminId"`
	MenuCalories float64       `json:"menuCalories"`
	Method       string        `json:"methods"`
	Name         string        `gorm:"-" json:"name"`
	Ingredients  []Ingredients `json:"ingredients"`
}
type ArrayIngredients struct {
	Ingredients []Ingredients `json:"ingredients"`
}
type Ingredients struct {
	Id                  string   `json:"id"`
	IngredientsName     string   `json:"ingredientsName"`
	IngredientsCalories float64  `json:"ingredientsCalories"`
	IngredientsType     string   `json:"ingredientsType"`
	Quantity            float64  `json:"quantity"`
	Type                string   `json:"type"`
	Database            Database `gorm:"-" json:"database"`
}
type Points struct {
	Id     string  `json:"id"`
	MenuId string  `json:"menuId"`
	UserId string  `json:"userId"`
	Point  float64 `json:"point"`
}
type DataPoints struct {
	Points []Points `json:"points"`
}

type Database struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
