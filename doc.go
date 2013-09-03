package cfg

/*
	package main

	import (
		"bitbucket.org/gotamer/cfg"
		"flag"
		"log"
		"os"
	)

	var Cfg struct {
		Debug     uint8
		AdminName string // Admin user name
		AdminPass string // Admin Password
		Srv       struct {
			ListenAddr string // :8080
			IpAddr     string // 127.0.0.1
			Hostname   string // www.example.com
		}
	}

	var (
		cfg_file = flag.String("c", "", "Config file (*.json)")
		help     = flag.Bool("h", false, "Display help")
	)

	func main() {
		flag.Parse()
		if *help == true || *cfg_file == "" {
			flag.PrintDefaults()
			os.Exit(0)
		}
		if err := cfg.Load(*cfg_file, &Cfg); err != nil {
			if err = cfg.Save(*cfg_file, &Cfg); err != nil {
				log.Println("cfg.Save error: ", err.Error())
				os.Exit(0)
			} else {
				log.Println("Please edit your config file at:\n\n\t", *cfg_file)
			}
		}

	}

*/
