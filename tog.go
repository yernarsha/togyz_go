package main
import (
  "fmt"
  "strconv"
  "math/rand"
  "strings"
  )

const (
  NUM_KUMALAKS = 9
  DRAW_GAME = 81
  TOTAL_KUMALAKS = 162
  TUZD = -1
)

type TogyzBoard struct {
  fields [23]int
  finished bool
  gameResult int
  moves []string
}

func (board *TogyzBoard) init() {
  for i:=0; i<23; i++ {
    if i<18 {
      board.fields[i] = NUM_KUMALAKS
    } else {
      board.fields[i] = 0
    }
  }

  board.finished = false
  board.gameResult = -2
  if len(board.moves) > 0 {
    board.moves = board.moves[:0]
  }
}

func (board TogyzBoard) getScore() string {
  return strconv.Itoa(board.fields[20]) + " - " + strconv.Itoa(board.fields[21])
}

func PadLeft(str, pad string, length int) string {
  for {
    str = pad + str
    if len(str) > length {
      return str[0: length+1]
    }
  }
}

func (board TogyzBoard) printPosition() {
  values := []string{}
  for i := range board.fields {
    text := strconv.Itoa(board.fields[i])
    values = append(values, text)
  }
  position := strings.Join(values, ",") // for debug

  position += "\n" + PadLeft(strconv.Itoa(board.fields[21]), " ", 6)

  for i:=17; i>8; i-- {
    if board.fields[i] == TUZD {
      position += PadLeft("X", " ", 4)
    } else {
      position += PadLeft(strconv.Itoa(board.fields[i]), " ", 4)
    }
  }

  position += "\n" + strings.Repeat(" ", 7)

  for j:=0; j<9; j++ {
    if board.fields[j] == TUZD {
      position += PadLeft("X", " ", 4)
    } else {
      position += PadLeft(strconv.Itoa(board.fields[j]), " ", 4)
    }
  }

  position += PadLeft(strconv.Itoa(board.fields[20]), " ", 6)
  fmt.Println(position)
}

func (board TogyzBoard) printNotation() {
  notation := ""
  for i, v := range board.moves {
    if i%2 == 0 {
      notation += strconv.Itoa(i/2 + 1) + ". " + v
    } else {
      notation += " " + v + "\n"
    }
  }

  fmt.Println(notation)
}

func (board *TogyzBoard) checkPosition() {
  color := board.fields[22]
  numWhite := 0
  for i:=0; i<9; i++ {
    if board.fields[i] > 0 {
      numWhite += board.fields[i]
    }
  }

  numBlack := TOTAL_KUMALAKS - numWhite - board.fields[20] - board.fields[21]

  if color == 0 && numWhite == 0 {
    board.fields[21] += numBlack
  } else if color == 1 && numBlack == 0 {
    board.fields[20] += numWhite
  }

  if board.fields[20] > DRAW_GAME {
    board.finished = true
    board.gameResult = 1
  } else if board.fields[21] > DRAW_GAME {
    board.finished = true
    board.gameResult = -1
  } else if board.fields[20] == DRAW_GAME && board.fields[21] == DRAW_GAME {
    board.finished = true
    board.gameResult = 0
  }
}

func (board *TogyzBoard) makeMove(move int) string {
  var sow int
  tuzdCaptured := false
  color := board.fields[22]

  madeMove := strconv.Itoa(move)
  move = move + (color * 9) - 1
  num := board.fields[move]

  if num == 0 || num == TUZD {
    fmt.Println("Incorrect move!")
    return ""
  }

  if num == 1 {
    board.fields[move] = 0
    sow = 1
  } else {
    board.fields[move] = 1
    sow = num - 1
  }

  num = move

  for i:=1; i<=sow; i++ {
    num += 1
    if num > 17 {
      num = 0
    }

    if board.fields[num] == TUZD {
      if num < 9 {
        board.fields[21] +=1
      } else {
        board.fields[20] += 1
      }
    } else {
      board.fields[num] += 1
    }
  }

  if board.fields[num] % 2 == 0 {
    if color == 0 && num > 8 {
      board.fields[20] += board.fields[num]
      board.fields[num] = 0
    } else if color == 1 && num < 9 {
      board.fields[21] += board.fields[num]
      board.fields[num] = 0
    }
  } else if board.fields[num] == 3 {
    if color == 0 && board.fields[18] == 0 && num > 8 && num < 17 && board.fields[19] != num - 8 {
      board.fields[18] = num - 8
      board.fields[num] = TUZD
      board.fields[20] += 3
      tuzdCaptured = true
    } else if color == 1 && board.fields[19] == 0 && num < 8 && board.fields[18] != num + 1 {
      board.fields[19] = num + 1
      board.fields[num] = TUZD
      board.fields[21] += 3
      tuzdCaptured = true
    }
  }

  if color == 0 {
    board.fields[22] = 1
  } else {
    board.fields[22] = 0
  }

  if num < 9 {
    num = num + 1
  } else {
    num = num - 8
  }

  madeMove += strconv.Itoa(num)
  if tuzdCaptured {
    madeMove += "x"
  }

  board.moves = append(board.moves, madeMove)
  board.checkPosition()
  return madeMove
}

func (board *TogyzBoard) makeRandomMove() string {
  var move, num int
  possible := make([]int, 0)
  color := board.fields[22]

  for i:=1; i<=9; i++ {
    move = i + (color * 9) - 1
    num = board.fields[move]
    if num > 0 {
      possible = append(possible, i)
    }
  }

  if len(possible) == 0 {
    fmt.Println("No possible moves!")
    return ""
  }

  randomIndex := rand.Intn(len(possible))
  randMove := possible[randomIndex]
  madeMove := board.makeMove(randMove)
  return madeMove
}
