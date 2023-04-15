package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var ls []Legend
var rect = Rect{Left: 1200, Bottom: 50, Right: 4000, Top: 2850}
var pr = DrawPr{Left: rect.Left, Top: rect.Top, Scale: 0.37037037037037035, Mashtab: 100}

var STR1 = "asrgfsadf12421"
var STR2 = "asrgfsadf12321"

func main() {

	plan, _ := ioutil.ReadFile("primitives.json")

	err := json.Unmarshal(plan, &ls)
	if err != nil {
		fmt.Print("Не могу прочитать json ", err)
	}

	// fmt.Println(data)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World golang!")
	})

	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()
		xStr := query["x"][0]
		yStr := query["y"][0]

		x, _ := strconv.ParseFloat(xStr, 64)
		y, _ := strconv.ParseFloat(yStr, 64)

		x /= 100
		y /= 100

		resultLs := build(ls,
			DrawPr{Left: pr.Left + x, Top: pr.Top + y, Scale: pr.Scale, Mashtab: pr.Mashtab},
			Rect{
				Left:   rect.Left + x,
				Top:    rect.Top + y,
				Right:  rect.Right,
				Bottom: rect.Bottom})

		// b, err := json.Marshal(resultLs)

		// if err != nil {
		// 	fmt.Print("Не могу сохранить json ", err)
		// }

		// w.Header().Set("Content-Type", "application/json")
		// result := strconv.Itoa(len(b))

		// buf := make([]byte, len(result))
		// copy(buf, result)
		// w.Write(buf)

		// //w.Write(b)

		result := strconv.Itoa(len(resultLs))

		buf := make([]byte, len(result))
		copy(buf, result)
		w.Write(buf)
	})

	http.HandleFunc("/mapJSON", func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()
		xStr := query["x"][0]
		yStr := query["y"][0]

		x, _ := strconv.ParseFloat(xStr, 64)
		y, _ := strconv.ParseFloat(yStr, 64)

		x /= 100
		y /= 100

		resultLs := build(ls,
			DrawPr{Left: pr.Left + x, Top: pr.Top + y, Scale: pr.Scale, Mashtab: pr.Mashtab},
			Rect{
				Left:   rect.Left + x,
				Top:    rect.Top + y,
				Right:  rect.Right,
				Bottom: rect.Bottom})

		b, err := json.Marshal(resultLs)

		if err != nil {
			fmt.Print("Не могу сохранить json ", err)
		}

		w.Header().Set("Content-Type", "application/json")
		result := strconv.Itoa(len(b))

		buf := make([]byte, len(result))
		copy(buf, result)
		w.Write(buf)

		//w.Write(b)
	})

	http.HandleFunc("/naturalsort", func(w http.ResponseWriter, r *http.Request) {
		result := 0
		for i := 0; i < 10000; i++ {
			result += Compare(STR1+strconv.Itoa(i), STR2+strconv.Itoa(i))
		}

		fmt.Fprint(w, result)
	})

	http.ListenAndServe("localhost:4000", nil)
}
