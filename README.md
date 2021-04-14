# Go with Go

### About

A Progressive Web App for playing the game of [Go](<https://en.wikipedia.org/wiki/Go_(game)>)! **Go with Go** is heavily inspired by [Cosumi](https://www.cosumi.net/en/). I love playing Go on Cosumi, but I dislike the design of the AI. It is slow to act and requires the client to have an active internet connection. My primary goal for this app is to create an enjoyable offline single-player Go experience. My secondary goal is to bring the game of Go to a wider audience. I plan to accomplish that goal with new variants of traditional Go and a _snazzy_ UI.

### Tech :construction:

- ReactJs
- Create-React-App
- Progressive Web App (PWA)
- Golang
- Web Assembly
- AWS
- TinyGo

---

# Release Notes

## v2.0 - Online Multiplayer

#### v1.3 - Scoring Breakdown

- [ ] Grid visualization
- [ ] Number breakdown

#### v1.3 - Smart Game Format
  - [ ] Export to .sgf

#### v1.2 - Game Replay

#### v1.1 - TinyGo

- [ ] Compile Golang to WASM using TinyGo

## v1.0 - Single Player

- [ ] Publish on Google Play
- [ ] Offline AI
- [ ] Select Board Size (5x5, 9x9, 13x13, 19x19)
- [ ] Mobile friendly controls
  - [ ] Scrollable board
  - [ ] Zoomable board
  - [ ] Double tap to place stone
- [ ] Main Menu
- [ ] Logo

## v0.0 - Local Multiplayer :construction:

- **0. Tech Stack Setup**
  - [x] PWA - Create-React-App
  - [x] Compile Golang to WASM
  - [ ] Automate local compilation
  - [ ] Host NodeJs server on ???
  - [ ] CICD Pipeline with ???
- **1. Game Logic (Japanese ruleset)**
  - [x] Start New Game
  - [ ] Scoring
    - Territory
    - Seki
    - Komi
  - [x] Pass turn
  - [x] Place stone
    - [x] Empty check
    - [x] Capture check
    - [x] Liberty check
    - [x] Ko check
    - [ ] Triple Ko
    - [ ] Forbid suicide
- **2. UI**
  - [x] Buttons
    - [x] Pass
    - [x] New game
  - [ ] Score screen
  - [x] Turn indicator
  - [x] Color scheme
  - [x] Board
  - [x] Stones
