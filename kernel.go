package main

func build(ls []Legend, pr DrawPr, rect Rect) []Layer {
	result := make([]Layer, len(ls))
	mashtab := 1 / pr.Scale

	for index, l := range ls {

		if l.MashtabRange.Min > pr.Mashtab || l.MashtabRange.Max < pr.Mashtab {
			continue
		}

		mas := make([][]float64, 0)

		for _, obraz := range clipPrimitives(&l, &pr, &rect) {
			csOpt := optimize(obraz.Coords, mashtab)
			translate(csOpt, pr)
			mas = append(mas, csOpt)
		}

		result[index] = Layer{LegendId: l.Id, Coords: mas}
	}

	return result
}

func clipPrimitives(l *Legend, pr *DrawPr, rect *Rect) []ResultClip {
	result := make([]ResultClip, 0)

	for _, g := range l.Primitives {
		if g.Rect.Left >= rect.Left &&
			g.Rect.Bottom >= rect.Bottom &&
			g.Rect.Right <= rect.Right &&
			g.Rect.Top <= rect.Top {
			// Целиком лежит внутри прямоугольника

			coords := make([]float64, len(g.Coords))
			copy(coords, g.Coords)

			result = append(result, ResultClip{Coords: coords})

		} else if g.Rect.Left < rect.Right &&
			g.Rect.Bottom < rect.Top &&
			g.Rect.Right > rect.Left &&
			g.Rect.Top > rect.Bottom {
			// Необходимо отсекать
			switch l.Type {
			case 1:
				for _, cs := range clipPolyline(&g, rect) {
					result = append(result, ResultClip{Coords: cs})
				}
			case 2:
				cs := clipPolygon(&g, rect)
				if len(cs) > 0 {
					result = append(result, ResultClip{Coords: cs})
				}
			}
		}
	}
	return result
}
