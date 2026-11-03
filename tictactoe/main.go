package main
func tictactoe(moves [][]int) string {
    movesA:= [][]int{}
	movesB:=[][]int{}

   
    
	for i,move:=range moves{
		if i%2==0{
			movesA=append(movesA, move)
		}else{
			movesB=append(movesB, move)
		}
	}
     
resultA:=reviewColumns(movesA)||reviewDiagonals(movesA)||reviewRows(movesA)
resultB:=reviewColumns(movesB)||reviewDiagonals(movesB)||reviewRows(movesB)
if resultA{
	return "A"
}
if resultB {
	return "B"
	
}
    if len(movesA)<3 || len(movesB)<3{
        return "Pending"
    }
    if(len(moves)==9){
        return "Draw"}
    return "Pending"
    
}

func reviewDiagonals(moves [][]int) bool {
	mainDiagonal := 0
	otherDiagonal := 0

	for _, move := range moves {
		if move[0] == move[1] {
			mainDiagonal++
		}
		if move[0]+move[1] == 2 {
			otherDiagonal++
		}
	}

	return mainDiagonal == 3 || otherDiagonal == 3
}
func reviewRows(moves [][]int) bool {
	for row := 0; row < 3; row++ {
		count := 0

		for _, move := range moves {
			if move[0] == row {
				count++
			}
		}

		if count == 3 {
			return true
		}
	}

	return false
}


func reviewColumns(moves [][]int) bool {
	for column := 0; column < 3; column++ {
		count := 0

		for _, move := range moves {
			if move[1] == column {
				count++
			}
		}

		if count == 3 {
			return true
		}
	}

	return false
}


