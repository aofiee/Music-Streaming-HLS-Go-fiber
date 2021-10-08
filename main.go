package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/:id/stream/", streamHandler)
	app.Get("/:id/stream/file:segment.m4a", streamHandler)
	app.Listen(":3000")
}

func streamHandler(c *fiber.Ctx) error {
	musicID, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	if c.Params("segment") != "" {
		segment, err := c.ParamsInt("segment")
		if err != nil {
			return err
		}
		musicDir := getMusic(musicID)
		musicSegment := fmt.Sprintf("file%d.m4a", segment)
		log.Println("sending segment", musicDir+"/"+musicSegment)
		c.Set(fiber.HeaderContentType, "audio/m4a")
		return c.SendFile(musicDir + "/" + musicSegment)
	} else {
		musicDir := getMusic(musicID)
		m3u8 := fmt.Sprintf("music%d.m3u8", musicID)
		c.Set(fiber.HeaderContentType, "application/x-mpegURL")
		log.Println("music started ", musicDir+"/"+m3u8)
		return c.SendFile(musicDir + "/" + m3u8)
	}
}

func getMusic(musicID int) string {
	root := "music"
	return fmt.Sprintf("%s/%d", root, musicID)
}
