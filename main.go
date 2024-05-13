package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Board [9]Player

type Game struct {
	board Board
}

type NeuralNetwork struct {
	inputSize             int
	hiddenSize            int
	outputSize            int
	inputToHiddenWeights  [][]int
	hiddenToOutputWeights [][]int
}

type Player int

const (
	PlayerHuman       Player = 1
	PlayerAI          Player = -1
	PositionNotPlayed Player = 0
)

func NewNeuralNetwork(inputSize, hiddenSize, outputSize int) *NeuralNetwork {
	neuralNetwork := &NeuralNetwork{
		inputSize:  inputSize,
		hiddenSize: hiddenSize,
		outputSize: outputSize,
	}
	neuralNetwork.initializeWeights()
	return neuralNetwork
}

func (nn *NeuralNetwork) initializeWeights() {
	nn.inputToHiddenWeights = make([][]int, nn.inputSize)
	nn.hiddenToOutputWeights = make([][]int, nn.hiddenSize)

	// Weights connecting the input layer with the hidden Layer initializing with random number between 0 - 99
	for i := 0; i < nn.inputSize; i++ {
		nn.inputToHiddenWeights[i] = make([]int, nn.hiddenSize)
		for j := 0; j < nn.hiddenSize; j++ {
			nn.inputToHiddenWeights[i][j] = rand.Intn(100)
		}
	}

	// Weights connecting the hidden layer with the output layer initializing with random number between 0 - 99
	for i := 0; i < nn.hiddenSize; i++ {
		nn.hiddenToOutputWeights[i] = make([]int, nn.outputSize)
		for j := 0; j < nn.outputSize; j++ {
			nn.hiddenToOutputWeights[i][j] = rand.Intn(100)
		}
	}
}

func NewGame() *Game {
	var board Board
	for i := range board {
		board[i] = PositionNotPlayed
	}
	return &Game{board: board}
}

func (g *Game) printBoard() {
	for i, v := range g.board {
		if v == 1 {
			fmt.Print(" X ")
		} else if v == -1 {
			fmt.Print(" O ")
		} else {
			fmt.Print(" # ")
		}

		if i == 2 || i == 5 || i == 8 {
			fmt.Println()
		}
	}
}

/*
MakeMove represents a move made by a player in the game.

position: represents the position in the board.
player: 0 to represent Human and 1 to represent IA
*/

func (g *Game) makeMove(position int, player Player) bool {
	isOutOfBoard := position < 0 || position > 8
	if isOutOfBoard {
		return false
	}
	if g.board[position] == PositionNotPlayed {
		g.board[position] = player
		return true
	}
	return false
}

func (g *Game) isFull() bool {
	howManyIsFilled := 0
	for _, value := range g.board {
		if value != 0 {
			howManyIsFilled++
		}
	}
	if howManyIsFilled == 9 {
		return true
	}
	return false
}

func (g *Game) checkWinner() Player {
	// Verify the columns
	for i := 0; i < 3; i++ {
		if g.board[i] == PositionNotPlayed {
			break
		}
		if g.board[i] == g.board[i+3] && g.board[i] == g.board[i+6] {
			//fmt.Printf("%d was the winner\n", g.board[i])
			return g.board[i]
		}
	}

	// Verify the lines
	for i := 0; i < 7; i = i + 3 {
		if g.board[i] == PositionNotPlayed {
			break
		}
		if g.board[i] == g.board[i+1] && g.board[i] == g.board[i+2] {
			fmt.Printf("%d was the winner\n", g.board[i])
			return g.board[i]
		}
	}

	// Verify the diagonals
	if g.board[0] != PositionNotPlayed && g.board[2] != PositionNotPlayed {
		if g.board[0] == g.board[4] && g.board[4] == g.board[8] {
			fmt.Printf("%d was the winner\n", g.board[0])
			return g.board[0]
		}
		if g.board[2] == g.board[4] && g.board[4] == g.board[6] {
			//fmt.Printf("%d was the winner\n", g.board[2])
			return g.board[2]
		}
	}

	return PositionNotPlayed
}

/**
 * Activation Function
 */
func sigmoid(x float64) float64 {
	return 1.0 / (1 + math.Exp(-x))
}

//TODO: Implement derivative

func main() {
	println("Game Starting...")
	game := NewGame()
	for game.checkWinner() == 0 && !game.isFull() {
		isPLayerPositionValid := false
		//Force user to play in the right position
		var playerMove int
		for !isPLayerPositionValid {
			fmt.Print("Enter your move (0-8): ")
			_, err := fmt.Scan(&playerMove)
			if err != nil {
				return
			}
			isPLayerPositionValid = game.makeMove(playerMove, PlayerHuman)
			if !isPLayerPositionValid {
				fmt.Println("Position Invalid")
			}
		}
		game.printBoard()
		fmt.Println("---- IA Has Played ----")
		var AIMove int
		IsAIPositionValid := false
		for !IsAIPositionValid && !game.isFull() && game.checkWinner() == 0 {
			AIMove = rand.Intn((9 - 0) + 0)
			IsAIPositionValid = game.makeMove(AIMove, PlayerAI)
			//TODO
		}
		game.printBoard()

		if game.checkWinner() == PlayerHuman {
			fmt.Println("Human Rules!!")
		} else if game.checkWinner() == PlayerAI {
			fmt.Println("AI Is Coming Baby!!")
		}

	}

}
