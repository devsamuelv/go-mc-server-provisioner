package main

import (
	"fmt"
	"main/commands"
	"main/drives"
	"main/file"
	"main/util"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configInfo := file.ReadConfigFile(".config.json")
	PORT := configInfo.Server.Port
	server := fiber.New(fiber.Config{
		Prefork: true,
	})

	if configInfo.DB.Key == "" || configInfo.DB.URL == "" {
		panic("Please Enter db Credentials")
	}

	if configInfo.MountedLocation == "" {
		cmd := exec.Command("/bin/sh", "-c", "mkdir Mounted")
		cmd.Start()

		path, err := os.Getwd()

		if err != nil {
			fmt.Printf("[ERROR] %s \n", err)
		}

		path += "/Mounted"

		configInfo.MountedLocation = path
		file.WriteConfigFile(".config.json", string(`"mounted_location": "",`), string(`"mounted_location": "`+path+`",`))

		fmt.Printf("[INFO] Mounted Location Set Defualt To %s \n", path)
	}

	if PORT == "" {
		PORT = "8545"
	}

	fmt.Print("=== Storage Server Online === \n")

	// SECTION Request Commands
	commands.GetDrives(server)
	commands.CreateDrive(server)
	commands.DisableDrive(server)
	// !SECTION

	util.PrintInit("db url", configInfo.DB.URL)
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
