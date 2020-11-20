package main

import (
	"fmt"
	"main/util"
	"main/file"
	"main/drives"
	// "github.com/gofiber/fiber/v2"
)


func main() {
	// PORT := ":8506"
	// server := fiber.New()
	
	fmt.Println("=== Storage Server Online ===")

	sysdrives := file.ReadConfigFile("config.json");

	for i := 0; i < len(sysdrives); i++ {
		info := drives.ReadSpace(sysdrives[i])
	
		fmt.Println(util.ByteCountSI(int64(info.All)))
	}

	// server.Listen(PORT)
}