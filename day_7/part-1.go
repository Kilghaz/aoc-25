package day_7

import (
	"aoc25/slice"
	"aoc25/vec2"

	"github.com/samber/lo"
)

type Beam = vec2.Vector2i

func Part1(input string) {
	tachyon := parseTachyonManifold(input)

	var beams []*Beam

	splitCounter := 0

	splitBeam := func(beam *Beam) {
		index := lo.IndexOf(beams, beam)
		slice.Splice(&beams, index, 1)
		beams = append(beams, &Beam{X: beam.X - 1, Y: beam.Y})
		beams = append(beams, &Beam{X: beam.X + 1, Y: beam.Y})
		splitCounter += 1
	}

	mergeBeams := func() {
		beams = lo.UniqBy(beams, func(item *Beam) int {
			return item.X
		})
	}

	advanceBeam := func(beam *Beam) {
		beam.Y += 1

		_, hasSplitter := lo.Find(tachyon.splitter, func(splitter vec2.Vector2[int]) bool {
			return splitter.X == beam.X && splitter.Y == beam.Y
		})

		if hasSplitter {
			splitBeam(beam)
		}
	}

	beams = append(beams, &Beam{X: tachyon.start.X, Y: tachyon.start.Y})
	maxY := lo.Max(lo.Map(tachyon.splitter, func(splitter vec2.Vector2[int], _ int) int {
		return splitter.Y
	}))

	for {
		lo.ForEach(beams, func(item *Beam, index int) {
			advanceBeam(item)
		})

		if beams[0].Y == maxY {
			break
		}

		mergeBeams()
	}

	println(splitCounter)
}
