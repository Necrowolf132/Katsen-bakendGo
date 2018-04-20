package modelos
import "github.com/dgrijalva/jwt-go"

type ValoresLogin struct {
	Usuario string `json:"usuario"`
	Clave  string `json:"clave"`
}
type Usuario struct {
	Id_perfil  int    `json:"id_perfil"`
	Cedula     string    `json:"cedula"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Extra      string `json:"extra"`
	Id_usuario int    `json:"id_usuario"`
	Id_seccion int    `json:"id_seccion"`
	User_pass string  `json:"user_pass"`

}
type EnvioDatos struct {
	StatusDeRow bool
	Datos_user Usuario
}

type Token struct {
	Id_user int `json:"id_usuario"`
	Id_perfil int  `json:"id_perfil"`
	jwt.StandardClaims
}