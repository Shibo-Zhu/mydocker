package main

import (
	"os"
	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
)

func main(){
	urfaveCli()
}

// urfaveCli cli 包简单使用，具体可以参考官方文档
func urfaveCli() {
	app := cli.NewApp()

	// 指定全局参数
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}
	// 指定支持的命令列表
	app.Commands = []cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				log.Println("run command complete")
				for i, v := range c.Args() {
					log.Printf("args i:%v v:%v\n", i, v)
				}
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			// 每个命令下面还可以指定自己的参数
			Flags: []cli.Flag{cli.Int64Flag{
				Name:  "priority",
				Value: 1,
				Usage: "priority for the task",
			}},
			Usage: "add a task to the list",
			Action: func(c *cli.Context) error {
				log.Println("run command add")
				for i, v := range c.Args() {
					log.Printf("args i:%v v:%v\n", i, v)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
