package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/a631807682/ddltransform"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "ddltrans",
		Version: "0.0.1",
		Commands: []*cli.Command{
			generate(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func generate() *cli.Command {
	return &cli.Command{
		Name:    "start",
		Aliases: []string{"s"},
		Usage:   "parse data definition language and generate to code",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "path",
				Aliases:  []string{"p"},
				Usage:    "path for ddl file",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "parser",
				Aliases:     []string{"ps"},
				DefaultText: "mysql",
				Usage:       "use parser for parse ddl file",
			},
			&cli.StringFlag{
				Name:        "transformer",
				Aliases:     []string{"tf"},
				DefaultText: "gorm",
				Usage:       "use transformer for code generate",
			},
		},
		Action: func(c *cli.Context) (err error) {
			argPath := c.String("path")
			argParser := c.String("parser")
			argTransformer := c.String("transformer")

			ddlbytes, err := ioutil.ReadFile(argPath)
			if err != nil {
				err = fmt.Errorf("load file failed. path:%s err:%v", argPath, err)
				return
			}

			parserType := ddltransform.Mysql
			if argParser != "" && argParser != "mysql" {
				if argParser == "sqlite" {
					parserType = ddltransform.Sqlite
				} else {
					err = fmt.Errorf("not support parser:'%s'", argParser)
					return
				}
			}

			transformerType := ddltransform.Gorm
			if argTransformer != "" && argTransformer != "gorm" {
				err = fmt.Errorf("not support transformer:'%s'", argTransformer)
				return
			}

			code, err := ddltransform.Transform(string(ddlbytes), ddltransform.Config{
				ParserType:      parserType,
				TransformerType: transformerType,
			})

			if err != nil {
				err = fmt.Errorf("transform err:%v", err)
				return
			}

			fmt.Println(code)
			return
		},
	}
}
