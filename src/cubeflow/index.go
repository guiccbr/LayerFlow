package main

const (
	DirUp = 1 << iota
	DirDown
	DirLeft
	DirRight
	DirTop
	DirBottom
	DirsPlane = DirUp | DirDown | DirLeft | DirRight
	DirsAll   = DirsPlane | DirTop | DirBottom
)

type Dir int

type Size struct {
	Rows, Cols, Layers int
}

type Index struct {
	Row, Col, Layer int
}

func (idx *Index) Neighbour(dir Dir) (Index, bool) {
	switch dir {
	case DirUp:
		return Index{idx.Row - 1, idx.Col, idx.Layer}, true
	case DirDown:
		return Index{idx.Row + 1, idx.Col, idx.Layer}, true
	case DirLeft:
		return Index{idx.Row, idx.Col - 1, idx.Layer}, true
	case DirRight:
		return Index{idx.Row, idx.Col + 1, idx.Layer}, true
	case DirTop:
		return Index{idx.Row, idx.Col, idx.Layer - 1}, true
	case DirBottom:
		return Index{idx.Row, idx.Col, idx.Layer + 1}, true
	}
	return *idx, false
}

func Dirs(dir Dir) []Dir {
	dirs := []Dir{}
	for _, d := range []Dir{DirUp, DirDown, DirLeft, DirRight, DirTop, DirBottom} {
		if dir&d != 0 {
			dirs = append(dirs, d)
		}
	}
	return dirs
}

func (idx *Index) Neighbours(dir Dir) []Index {
	dirs := Dirs(dir)
	idxs := make([]Index, len(dirs))
	for i, d := range dirs {
		idxs[i], _ = idx.Neighbour(d)
	}
	return idxs
}
