// Package conectarGeneral contiene funciones de utilidad para trabajar con cadenas html por http
package conectarGeneral

import (
	"gofacilito/config"
	_"github.com/lib/pq"
	"gofacilito/modelos"
	"database/sql"
	//"backend/nucleo"
	"log"

)

var parametros *config.StructParametros = &config.Parametros
var parapoll string = "postgres://" + parametros.DatabaseLoc["user"] + ":" + parametros.DatabaseLoc["clave"] + "@" + parametros.DatabaseLoc["host"] + ":" + parametros.DatabaseLoc["port"] + "/" + parametros.DatabaseLoc["database"]
//var parapoll string = parametros.DatabasePro["user"] + ":" + parametros.DatabasePro["clave"] + "@tcp(" + parametros.DatabasePro["host"] + ":" + parametros.DatabasePro["port"] + ")/" + parametros.DatabasePro["database"]


func Crearpoll() *sql.DB {

	ConexionBD, err := sql.Open("postgres", parapoll)
	if err != nil {
		panic(err)
	}
	err = ConexionBD.Ping()
	if err != nil {
		log.Fatal("fallo el ping:"+err.Error())
	}
	return ConexionBD
}

func Traerusuariologin(ConexionDB *sql.DB, login *modelos.ValoresLogin) modelos.EnvioDatos {
	log.Println("Parametro pasado",login.Usuario)
	var envi modelos.EnvioDatos
	var m modelos.Usuario
	err := ConexionDB.QueryRow("SELECT id_perfil,cedula,nombre,apellido,extra,id_usuario,id_seccion,user_pass FROM usuario WHERE user_name='"+login.Usuario +"'").Scan(&m.Id_perfil, &m.Cedula, &m.Nombre, &m.Apellido, &m.Extra, &m.Id_usuario, &m.Id_seccion, &m.User_pass)

	if err != nil {
		log.Println(err, "entreeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
		m = modelos.Usuario{0, "", "", "", "", 0, 0, ""}
		envi = modelos.EnvioDatos{false, m}
	}else{

		/*if filas.Next(){

			defer filas.Close()
			var StatusDeRow bool = true
			for filas.Next() {
				var Id_usuario int
				var Nombre string
				var Apellido string
				var Cedula string
				var Extra string
				var Id_perfil int
				var Id_seccion int
				var User_pass string
				err := filas.Scan(&Id_perfil, &Cedula, &Nombre, &Apellido, &Extra, &Id_usuario, &Id_seccion, &User_pass)
				if err != nil {
					log.Println(err)
				} else {*/


					//m = modelos.Usuario{ true,Id_usuario, Nombre, Apellido, Cedula, Extra, Id_perfil, Id_seccion, User_pass}
					envi = modelos.EnvioDatos{true, m}
					log.Println("modelo usuario es", m)


					}

			/*}
			if err := filas.Err(); err != nil {
				log.Println(err)
			}

		}else{
			m = modelos.Usuario{false, 0, "", "", "", "", 0, 0, ""}
			//jsonEnvi, err = json.Marshal(m)
			//log.Println(string(jsonEnvi))
		}
	}*/
	return  envi
}