package app

import "github.com/urfave/cli"

// Gerar retorna a app pronta
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação cli"
	app.Usage = "Buscao e IPS e nomes de Servidor"

	return app
}
