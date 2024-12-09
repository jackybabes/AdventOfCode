package main

type xyMap struct {
	matrix [][]string
}

func (m *xyMap) get(coord Coord) string {
	line := m.matrix[len(m.matrix)-1-coord.y]
	point := line[coord.x]
	return point
}
func (m *xyMap) set(coord Coord, s string) {
	line := m.matrix[len(m.matrix)-1-coord.y]
	line[coord.x] = s
}

func (m *xyMap) print() {
	printMatrix(m.matrix)
}

func (m *xyMap) collectAntenna() map[string][]Coord {
	antennas := make(map[string][]Coord)
	for x := range m.matrix[0] {
		for y := range m.matrix {
			coord := Coord{x, y}
			value := m.get(coord)
			if value != "." {
				antennas[value] = append(antennas[value], coord)
			}
		}
	}
	return antennas
}

func (m *xyMap) checkInBounds(c Coord) bool {
	if c.x >= 0 && c.x < len(m.matrix[0]) && c.y >= 0 && c.y < len(m.matrix) {
		return true
	}
	return false
}

func (m *xyMap) countPositions(s string) int {
	var count int
	for x := range m.matrix[0] {
		for y := range m.matrix {
			if m.get(Coord{x, y}) == s {
				count++
			}
		}
	}
	return count
}
