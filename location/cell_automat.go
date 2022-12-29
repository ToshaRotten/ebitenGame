package location

type cell struct {
	aliveStatus       bool // Does cell alive
	arrounCellsNumber int  // How many cells allive arround this cell
}

func (c *cell) CountNeighbours() int {

}
