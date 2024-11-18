# Snake Game in Go

This is a simple implementation of the classic Snake Game using the Go programming language. The game runs in the terminal and utilizes the `termbox-go` library for handling keyboard input and rendering.

## Features

- **Snake Movement:** The snake moves continuously in the direction specified by the arrow keys.
- **Food Collection:** The snake grows in length each time it eats food, which is represented by an 'X' on the screen.
- **Scoring System:** The score increases by 1 each time the snake eats food. The current score is displayed at the bottom of the game area.
- **Borders:** The game area is surrounded by borders. If the snake collides with the borders, the game ends.
- **Game Over:** The game ends when the snake collides with the borders or itself. The final score is displayed in the terminal.

## Requirements

- Go programming language installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).
- The `termbox-go` library for terminal handling.

## Installation

1. **Install the `termbox-go` library:**

    Open your terminal and run the following command:

    ```bash
    $ go get github.com/nsf/termbox-go
    ```

2. **Create a new directory for the game:**
    ```bash
    $ mkdir snake_game
    $ cd snake_game
    ```

3. **Clone the repository:**
    ```bash
    $ git clone https://github.com/gdev-kingx/Snake.go.git
    ```

4. **Run the game:**
    1. Navigate to the directory where you saved `main.go`.
    2. Run the following command:
    
        ```bash
        $ go run main.go
        ```