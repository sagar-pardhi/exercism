package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0

	for k, v := range cb {
		if k == file {
			for _, val := range v {
				if val {
					count += 1
				}
			}
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0

	l := len(cb)

	if rank >= 1 && rank <= l {
		for _, f := range cb {
			for i, v := range f {
				if i+1 == rank && v {
					count += 1
				}
			}
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0

	for _, v := range cb {
		for range v {
			count += 1
		}
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0

	for _, v := range cb {
		for _, x := range v {
			if x {
				count += 1
			}
		}
	}

	return count
}
