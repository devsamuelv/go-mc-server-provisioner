package main

import (
	"fmt"
	"main/drives"
	"main/file"
	"main/util"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configInfo := file.ReadConfigFile(".config.json")
	PORT := configInfo.Server.Port
	server := fiber.New(fiber.Config{
		Prefork: true,
	})

	fmt.Print("=== Storage Server Online ===")

	util.PrintInit("db url", configInfo.Db.URL)
	util.PrintInit("db username", configInfo.Db.Username)
	util.PrintInit("Server Port", configInfo.Server.Port)
	util.PrintInit("Server Public Key", configInfo.Cert.Public)

	for i := 0; i < len(configInfo.Drives); i++ {
		info := drives.ReadSpace(configInfo.Drives[i])

		if info.All == 0 {
			fmt.Print("Drive: ", configInfo.Drives[i], " Not Found")
			return
		}

		fmt.Print("Drive Name: ", "'", configInfo.Drives[i], "'", " Drive Size GB: ", util.ByteCountSI(int64(info.All)), "\n")
		fmt.Print("Drive Name: ", "'", configInfo.Drives[i], "'", " Free Drive Space GB: ", util.ByteCountSI(int64(info.Free)), "\n")
	}

	server.Listen(":" + PORT)
}
