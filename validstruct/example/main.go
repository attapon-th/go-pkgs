// package main main go
package main

import (
	"encoding/json"
	"log"

	"github.com/attapon-th/go-pkgs/validstruct"
)

func main() {
	run()
}

// UserInfo test struct
type UserInfo struct {
	ID             int     `json:"id" validate:"required"`
	Username       string  `json:"username" validate:"required,max=10,lowercase"`
	FirstName      string  `json:"first_name" validate:"required"`
	LastName       string  `json:"last_name" validate:"required"`
	Age            uint8   `json:"age" validate:"gte=0,lte=130"`
	Email          string  `json:"email" validate:"required,email"`
	FavouriteColor string  `json:"favourite_color" validate:"iscolor"`
	Gender         string  `json:"gender" validate:"required,len=1"`
	Weigth         float32 `json:"weigth" validate:"numeric,gte=0.0,lte=300.0"`
	Heigth         float32 `json:"heigth" validate:"numeric,gte=0.0,lte=300.0"`
}

func run() {
	v := validstruct.New("json")
	user := UserInfo{
		ID:             -1,
		Username:       "asdfefasdfI",
		FirstName:      "test",
		LastName:       "last",
		Age:            99,
		Email:          "asdf",
		FavouriteColor: "",
		Gender:         "M",
		Weigth:         999.9,
		Heigth:         199.9,
	}
	if err := v.Struct(user); err != nil {
		log.Fatalln(err.Error())
	}
	printErr(v)
	// Output:
	// [
	//   {
	//     "field": "username",
	//     "valid": "max",
	//     "detail": "max",
	//     "param": "10",
	//     "massage": "Key: 'username' Error:Field validation for 'username' failed on the 'max' tag"
	//   },
	//   {
	//     "field": "email",
	//     "valid": "email",
	//     "detail": "email",
	//     "param": "",
	//     "massage": "Key: 'email' Error:Field validation for 'email' failed on the 'email' tag"
	//   },
	//   {
	//     "field": "favourite_color",
	//     "valid": "iscolor",
	//     "detail": "hexcolor|rgb|rgba|hsl|hsla",
	//     "param": "",
	//     "massage": "Key: 'favourite_color' Error:Field validation for 'favourite_color' failed on the 'iscolor' tag"
	//   },
	//   {
	//     "field": "weigth",
	//     "valid": "lte",
	//     "detail": "lte",
	//     "param": "300.0",
	//     "massage": "Key: 'weigth' Error:Field validation for 'weigth' failed on the 'lte' tag"
	//   }
	// ]

}

func printErr(v *validstruct.ValidStruct) {
	// json
	b, err := json.MarshalIndent(v.ValidErrs, "", "  ")
	log.Println(err)
	log.Fatalln(string(b))
}
