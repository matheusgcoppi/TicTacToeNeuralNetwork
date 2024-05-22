package main

import (
	"errors"
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
	inputToHiddenWeights  [][]float64
	hiddenToOutputWeights [][]float64
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

/*
* Forward Propagation: Calculate the output of the network, for a given input using sigmoid to activate
* Product from Input to Hidden Layer
* Product from Hidden Layer to Output Layer
 */
func (nn *NeuralNetwork) forward(input []float64) []float64 {
	hidden := dotProduct(input, nn.inputToHiddenWeights)
	hidden = applySigmoid(hidden)
	output := dotProduct(hidden, nn.hiddenToOutputWeights)
	output = applySigmoid(output)
	return output
}

/*
* back propagation:
* - Loss Function: Mean Squared Error(MSE)
* - Propagate the error backwards in the network
* - Updating the weights
 */
func backPropagation(input, target, output []float64) {

}

/*
* Mean Squared Error Loss Function
* MSE = (1/n) * Σ(y_i - ŷ_i)²
* Calculate how much the network got it wrong
 */
func meanSquaredError(output, target []float64) (float64, error) {
	// verify if the length is equal
	if len(output) != len(target) {
		return 0, errors.New("the length of output and target slices must be the same")
	}
	var sum float64
	n := float64(len(output))
	for i := range output {
		// calculate the difference between output and target value
		diff := output[i] - target[i]
		// sum all the squared differences
		sum += diff * diff
	}
	// divide the sum by the length of output
	mse := sum / n
	return mse, nil
}

func (nn *NeuralNetwork) initializeWeights() {
	nn.inputToHiddenWeights = make([][]float64, nn.inputSize)
	nn.hiddenToOutputWeights = make([][]float64, nn.hiddenSize)

	// Weights connecting the input layer with the hidden Layer initializing with random number between 0 - 99
	for i := 0; i < nn.inputSize; i++ {
		nn.inputToHiddenWeights[i] = make([]float64, nn.hiddenSize)
		for j := 0; j < nn.hiddenSize; j++ {
			nn.inputToHiddenWeights[i][j] = rand.Float64()
		}
	}

	// Weights connecting the hidden layer with the output layer initializing with random number between 0 - 99
	for i := 0; i < nn.hiddenSize; i++ {
		nn.hiddenToOutputWeights[i] = make([]float64, nn.outputSize)
		for j := 0; j < nn.outputSize; j++ {
			nn.hiddenToOutputWeights[i][j] = rand.Float64()
		}
	}
}

// This function returns the product of a vector and a matrix
func dotProduct(vector []float64, matrix [][]float64) []float64 {
	if len(vector) != len(matrix) {
		panic("Vector and matrix dimensions do not match")
	}

	result := make([]float64, len(matrix[0]))

	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(vector); j++ {
			result[i] += vector[j] * matrix[j][i]
		}
	}

	return result
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
 * Activation Function, sigmoid is ALWAYS give you a value between 0 and 1, based on the X
 */
func sigmoid(x float64) float64 {
	return 1.0 / (1 + math.Exp(-x))
}

func applySigmoid(vector []float64) []float64 {
	for i := range vector {
		vector[i] = sigmoid(vector[i])
	}
	return vector
}

//TODO: Implement derivative

func main() {
	var neuralNetwork = NewNeuralNetwork(9, 10, 9)
	println("Game Starting...")

	game := NewGame()

	for game.checkWinner() == 0 && !game.isFull() {
		//force user to play in right position
		playerTurn(game)

		game.printBoard()
		// AI play in a random position
		aiTurn(game, neuralNetwork)

		game.printBoard()
		//Status of The game
		checkWinner(game)
	}
}

func playerTurn(game *Game) {
	isPLayerPositionValid := false
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
}

func aiTurn(game *Game, network *NeuralNetwork) {
	// Input
	// Convert the input to Float64
	boardInput := make([]float64, len(game.board))

	for i, value := range game.board {
		boardInput[i] = float64(value)
	}

	moveProbabilities := network.forward(boardInput)
	var validMoves []int
	for i, value := range game.board {
		if value == PositionNotPlayed {
			validMoves = append(validMoves, i)
		}
	}
	var bestMove = -1
	bestMoveValue := -math.MaxFloat64
	for _, move := range validMoves {
		if moveProbabilities[move] > bestMoveValue {
			bestMoveValue = moveProbabilities[move]
			bestMove = move
		}
	}

	if bestMove == -1 {
		fmt.Println("---- IA Has Played Randomly----")
		var AIMove int
		IsAIPositionValid := false
		for !IsAIPositionValid && !game.isFull() && game.checkWinner() == 0 {
			AIMove = rand.Intn((9 - 0) + 0)
			IsAIPositionValid = game.makeMove(AIMove, PlayerAI)
			//TODO
		}
	} else {
		fmt.Println("---- IA Has Played Based on Probabilities ----")
		game.makeMove(bestMove, PlayerAI)
	}
}

func checkWinner(game *Game) {
	if game.checkWinner() == PlayerHuman {
		fmt.Println("Human Rules!!")
	} else if game.checkWinner() == PlayerAI {
		fmt.Println("AI Is Coming Baby!!")
	}
}
