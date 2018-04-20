package login

import (
	"gofacilito/modelos"
	"gofacilito/conectarGeneral"
	"gofacilito/Autentificacion"
	"net/http"
	"encoding/json"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	)

func TraerUsuario(w http.ResponseWriter, r *http.Request, ConexionDB *sql.DB) {

	type Devolucion struct {
		Loginvalido bool `json:"login_valido"`
		Descripcion string `json:"descripcion"`
		Datos_user modelos.Usuario `json:"datos_user"`
		Token string `json:"token"`
	}
	var DevolucionLogin Devolucion
	var tokenNuevo string = " "
	var valor modelos.ValoresLogin
	log.Println("Host: ", r.Host)
	log.Println("Navegador: ", r.UserAgent())
	log.Println("Esto es Authorization ", r.Header.Get("Authorization"))
	err := json.NewDecoder(r.Body).Decode(&valor)
	if err != nil  {
		log.Println("algo salio mal en los parametros de envio")
		DevolucionLogin = Devolucion{false, "La peticion no poseia datos de consulta", modelos.Usuario{0, " "," "," "," ", 0, 0 ," " }, "sin token"}

	}else {
		var mostrarLogin modelos.EnvioDatos = conectarGeneral.Traerusuariologin(ConexionDB, &valor)
		if (!mostrarLogin.StatusDeRow) {
			DevolucionLogin = Devolucion{false, "No se realizo la validacion por que el usuario no existe en la base de datos", mostrarLogin.Datos_user, "sin token"}
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(mostrarLogin.Datos_user.User_pass), []byte(valor.Clave))
			if err != nil {
				DevolucionLogin = Devolucion{false, "La clave de validacion de usuario es invalida", mostrarLogin.Datos_user, "sin token"}
			} else {
				tokenNuevo = Autentificacion.CrearTOken(mostrarLogin.Datos_user)
				DevolucionLogin = Devolucion{true, "Todo salio perfecto", mostrarLogin.Datos_user, tokenNuevo}

			}
		}
	}
	var jsonEnvi []byte
	jsonEnvi, err= json.Marshal(DevolucionLogin)
	log.Println( "El json final de envio es : "+string(jsonEnvi))
	if err != nil{
		log.Println( "no se genero el json")
	}

	w.Header().Add("content-Type", "application/json; charset=utf-8; token="+tokenNuevo)
	w.Header().Add("Etag", ` "499fd34e-29ec-42f695ca96761;48fe7523cfcc1"`)
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
	log.Println( "esto es header", w.Header())
	w.WriteHeader(http.StatusOK)
	w.Write(jsonEnvi)

}