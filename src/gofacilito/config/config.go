package config



type StructParametros struct {
	DatabaseLoc map[string]string
	DatabasePro map[string]string
	Modulo uint
}

	var  Parametros = StructParametros{
		DatabaseLoc:map[string]string{
			"host":"localhost",
			"port":"5432",
			"user":"uxqgrykpjvpzft",
			"clave":"bd4d5c4b43a1221b3405266acf774047cc17959bd22d85f143e96e1651f7f5dd",
			"database":"df3thhp6ihie67",

			},
		DatabasePro:map[string]string{
			"host":"ec2-54-247-95-125.eu-west-1.compute.amazonaws.com",
			"port":"5432",
			"user":"uxqgrykpjvpzft",
			"clave":"bd4d5c4b43a1221b3405266acf774047cc17959bd22d85f143e96e1651f7f5dd",
			"database":"df3thhp6ihie67",

			},
			Modulo:1,
			}
