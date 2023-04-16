package main

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"strconv"
)

type MapStructCoord struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

var ls []Legend
var rect = Rect{Left: 1200, Bottom: 50, Right: 4000, Top: 2850}
var pr = DrawPr{Left: rect.Left, Top: rect.Top, Scale: 0.37037037037037035, Mashtab: 100}

var STR1 = "asrgfsadf12421"
var STR2 = "asrgfsadf12321"

func main() {
	router := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	plan, _ := os.ReadFile("primitives.json")

	err := json.Unmarshal(plan, &ls)
	if err != nil {
		fmt.Print("Не могу прочитать json ", err)
	}

	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World golang!")
	})

	router.Get("/map", func(ctx *fiber.Ctx) error {
		var q MapStructCoord
		if err := ctx.QueryParser(&q); err != nil {
			return err
		}

		q.X /= 100
		q.Y /= 100

		resultLs := build(ls,
			DrawPr{Left: pr.Left + q.X, Top: pr.Top + q.Y, Scale: pr.Scale, Mashtab: pr.Mashtab},
			Rect{
				Left:   rect.Left + q.X,
				Top:    rect.Top + q.Y,
				Right:  rect.Right,
				Bottom: rect.Bottom})

		result := strconv.Itoa(len(resultLs))

		buf := make([]byte, len(result))
		copy(buf, result)
		return ctx.JSON(buf)
	})

	router.Get("/mapJSON", func(ctx *fiber.Ctx) error {
		var q MapStructCoord
		if err := ctx.QueryParser(&q); err != nil {
			return err
		}

		q.X /= 100
		q.Y /= 100

		resultLs := build(ls,
			DrawPr{Left: pr.Left + q.X, Top: pr.Top + q.Y, Scale: pr.Scale, Mashtab: pr.Mashtab},
			Rect{
				Left:   rect.Left + q.X,
				Top:    rect.Top + q.Y,
				Right:  rect.Right,
				Bottom: rect.Bottom})

		b, err := sonic.Marshal(resultLs)

		if err != nil {
			fmt.Print("Не могу сохранить json ", err)
		}

		result := strconv.Itoa(len(b))

		return ctx.SendString(result)
	})

	router.Get("/naturalsort", func(ctx *fiber.Ctx) error {
		result := 0
		for i := 0; i < 10_000; i++ {
			result += Compare(STR1+strconv.Itoa(i), STR2+strconv.Itoa(i))
		}

		fmt.Println(result)
		return ctx.JSON(result)
	})

	if err := router.Listen(":4000"); err != nil {
		log.Fatal(err)
	}
}
