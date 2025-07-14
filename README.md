# Rock Paper Scissors Lizard Spock Game (Go)

Welcome to the **Rock Paper Scissors Lizard Spock** game implemented in Go! This is a command-line game where you play against the computer (bot). The first to reach 5 wins is the overall winner.

This game is inspired by the popular variation of the classic Rock Paper Scissors game, made famous by *The Big Bang Theory* TV show, adding two more choices — Lizard and Spock — to reduce ties and increase fun.

---

## Features

- Play against a computer opponent with random choices.
- Clear instructions and numbered options for easy input.
- Score tracking for player wins, bot wins, and draws.
- First to 5 wins ends the game.
- Exit anytime by typing `exit`.

---

## How to Install and Run

### Prerequisites

- You need to have [Go](https://golang.org/dl/) installed on your machine (version 1.13+ recommended).
- Basic familiarity with running commands in your terminal or command prompt.

### Installation Steps

1. **Clone the repository** (or download the source code):

```bash
git clone https://github.com/gismc123/rock-paper-scissors-lizard-spock.git
cd rock-paper-scissors-lizard-spock
```

2. **Build the game executable:**

```bash
go build -o rpsls
```

This will create an executable file named `rpsls` (or `rpsls.exe` on Windows).

3. **Run the game:**

```bash
./rpsls
```

On Windows, run:

```bash
rpsls.exe
```

---

## How to Play

- When prompted, enter the **number** corresponding to your choice:

```
0 : rock
1 : paper
2 : scissors
3 : lizard
4 : spock
```

- The bot will randomly pick its choice.
- The game will display who wins the round and update the scores.
- The first player (you or the bot) to reach 5 points wins the game.
- To quit the game at any time, type `exit`.

---

## Example Gameplay

```
Choose your pick. Enter the number corresponding to your choice.

0 : rock
1 : paper
2 : scissors
3 : lizard
4 : spock

> 2

You chose scissors
Bot chose paper
Scissors cuts paper.
You win!

Player score: 1
Bot score: 0
Draw score: 0
```

---

## Notes

- Input validation ensures only numbers 0-4 or the command `exit` are accepted.
- The game logic follows the official rules of Rock Paper Scissors Lizard Spock.
- The game runs in a loop until either player reaches 5 wins or the user exits.

---

## Contribution

Feel free to fork the repository and submit pull requests if you want to improve the game or add features!

---

## License

This project is open source and available under the MIT License.

---

Enjoy playing and sharing this fun twist on a classic game! 🎉
