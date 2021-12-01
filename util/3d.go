package util

// yes I'm too lazy to manually calculate this out.
// running at initialization time isn't too terribad.
var AdjacentVector3D = func() []Vector3D {
	seenMap := make(map[Vector3D]bool)

	for x := int64(-1); x <= 1; x++ {
		for y := int64(-1); y <= 1; y++ {
			for z := int64(-1); z <= 1; z++ {
				seenMap[Vector3D{
					X: x,
					Y: y,
					Z: z,
				}] = true
			}
		}
	}

	delete(seenMap, Vector3D{})

	out := make([]Vector3D, len(seenMap))
	i := 0
	for vec := range seenMap {
		out[i] = vec
		i++
	}

	return out
}()

type Vector3D struct {
	X, Y, Z int64
}

func (p Vector3D) Add(p2 Vector3D) Vector3D {
	return Vector3D{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
		Z: p.Z + p2.Z,
	}
}

func (p Vector3D) To4D(w int64) Vector4D {
	return Vector4D{
		X: p.X,
		Y: p.Y,
		Z: p.Z,
		W: w,
	}
}