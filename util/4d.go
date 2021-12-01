package util

// yes I'm too lazy to manually calculate this out.
// running at initialization time isn't too terribad.
var AdjacentVector4D = func() []Vector4D {
	seenMap := make(map[Vector4D]bool)

	for w := int64(-1); w <= 1; w++ {
		for x := int64(-1); x <= 1; x++ {
			for y := int64(-1); y <= 1; y++ {
				for z := int64(-1); z <= 1; z++ {
					seenMap[Vector4D{
						X: x,
						Y: y,
						Z: z,
						W: w,
					}] = true
				}
			}
		}
	}

	delete(seenMap, Vector4D{})

	out := make([]Vector4D, len(seenMap))
	i := 0
	for vec := range seenMap {
		out[i] = vec
		i++
	}

	return out
}()

type Vector4D struct {
	X, Y, Z, W int64
}

func (p Vector4D) Add(p2 Vector4D) Vector4D {
	return Vector4D{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
		Z: p.Z + p2.Z,
		W: p.W + p2.W,
	}
}