package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar retorna a app pronta
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação cli"
	app.Usage = "Buscao e IPS e nomes de Servidor"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIPs,
		},
	}

	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")

	//pacote net
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}
