# Go with Go

### About

A Progressive Web App for playing the game of [Go](<https://en.wikipedia.org/wiki/Go_(game)>)! **Go with Go** is heavily inspired by [Cosumi](https://www.cosumi.net/en/). I love playing Go on Cosumi, but I dislike the design of the AI. It is slow to act and requires the client to have an active internet connection. My primary goal for this app is to create an enjoyable offline single-player Go experience. My secondary goal is to bring the game of Go to a wider audience. I plan to accomplish that goal with new variants of traditional Go and a _snazzy_ UI.

### Tech :construction:

- ReactJs
- Create-React-App
- Progressive Web App (PWA)
- Golang
- Web Assembly
- Google Cloud Platform
- TinyGo

---

# Release Notes

## v3.0 - Online Multiplayer

## v2.0 - Arcade Gamemode(s)

#### v1.2 - Chinese Ruleset

#### v1.1 - TinyGo

- [ ] Compile Golang to WASM using TinyGo

## v1.0 - Offline Single Player

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
  - [x] Golang --> WASM
  - [ ] Host on GCP
- **1. Game Logic (Japanese rules)**
  - [x] Start New Game
  - [ ] Scoring
    - Territory
    - Komi
  - [ ] Pass turn
  - [x] Place stone
    - [x] Empty check
    - [x] Capture check
    - [x] Liberty check
    - [x] Ko check
- **2. UI**
  - [x] Buttons
    - [x] Pass
    - [x] New game
  - [ ] Score screen
  - [x] Turn indicator
  - [x] Color scheme
  - [x] Board
  - [x] Stones
- **3. Other**
  - [ ] Refactor GoGame funcs to use Game obj
