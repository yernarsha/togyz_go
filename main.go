package main
import (
  "fmt"
  "time"
  "math/rand"
  )

func main() {
  rand.Seed(time.Now().Unix())
  fmt.Println("Welcome to the Togyz Kumalak world!\n")
  fmt.Println("Enter the mode (h - human play, m - random machine, r - against random AI): ")

  var input string
  fmt.Scanln(&input)
  if input == "h" {
    humanPlay()
  } else if input == "m" {
    machinePlay() // 20x faster than Python
  } else if input == "r" {
    randomPlay()
  }
}

func humanPlay() {
  var input int
  togyzBoard := TogyzBoard{}
  togyzBoard.init()

  for {
    fmt.Println("\nEnter your move (1-9, 0 - exit): ")
    fmt.Scanln(&input)
    if input == 0 {
      break
    }

    if input >=1 && input <=9 {
      togyzBoard.makeMove(input)
      togyzBoard.printNotation()
      togyzBoard.printPosition()
      if togyzBoard.finished {
        break
      }
    }
  }

  togyzBoard.printNotation()
  togyzBoard.printPosition()
  fmt.Println("\nGame over:", togyzBoard.getScore(), ". Result:", togyzBoard.gameResult)
}

func machinePlay() {
  var num int

  for {
    fmt.Println("\nEnter number of iterations (1-100000): ")
    fmt.Scanln(&num)
    if num >=1 && num <=100000 {
      break
    }
  }

  win := 0
  draw := 0
  loss := 0
  sta := time.Now()
  togyzBoard := TogyzBoard{}

  for i:=1; i<=num; i++ {
    togyzBoard.init()

    for !togyzBoard.finished {
      togyzBoard.makeRandomMove()
    }

    if num <= 5 {
      togyzBoard.printPosition()
      fmt.Println("\nGame over:", togyzBoard.getScore(), ". Result:", togyzBoard.gameResult, "\n")
    }

    if togyzBoard.gameResult == 1 {
      win += 1
    } else if togyzBoard.gameResult == -1 {
      loss += 1
    } else if togyzBoard.gameResult == 0 {
      draw += 1
    } else {
      fmt.Println("What??")
    }

  }

  fmt.Println("W:", win, ", D:", draw, ", L:", loss)
  fmt.Println("Elapsed:", time.Since(sta))
}

func randomPlay() {
  togyzBoard := TogyzBoard{}
  togyzBoard.init()
  currentColor := 0
  var color, move int
  var ai string

  for {
    fmt.Println("\nEnter your color (0 - white, 1 - black): ")
    fmt.Scanln(&color)
    if color == 0 || color == 1 {
      break
    }
  }

  for !togyzBoard.finished {
    if currentColor == color {
      for {
        fmt.Println("\nEnter your move (1-9, 0 - exit): ")
        fmt.Scanln(&move)

        if move >= 0 && color <= 9 {
          break
        }
      }

      if move == 0 {
        break
      }

      togyzBoard.makeMove(move)
      togyzBoard.printNotation()
      togyzBoard.printPosition()
      if color == 0 {
        color = 1
      } else {
        color = 0
      }

    } else {
      ai = togyzBoard.makeRandomMove()
      fmt.Println("AI move:", ai)
      togyzBoard.printPosition()
      if color == 0 {
        color = 1
      } else {
        color = 0
      }
    }

  }

  togyzBoard.printNotation()
  togyzBoard.printPosition()
  fmt.Println("\nGame over:", togyzBoard.getScore(), ". Result:", togyzBoard.gameResult)
}
