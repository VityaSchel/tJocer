# tJocer

Trainer for The Joy of Creation: Reborn, written in Go.

## This project is abandoned

I have no time and no energy left to keep this project alive. I'm too exhausted after 2 months of trying to find a way to get base address of process and other simple tasks. Yes, that may sound simple enough for you, but don't tell me about EnumProcessModules, GetModuleInformation, modBaseAddr and shit, none of this works, at least when I tried to do it. Go ahead and try to fix mem.go, specifically, incorrect base address (also known as image_hash or preferred_load_address etc) that is returned by module handle.

## Features

Implemented features are checked, others are either in-progress or proposals. You can open PR any time and I will try to accept it in next few years :) The hardest job is to find static pointer to value in memory, then you have to experiment with it and try to write new value, then implement logic (+render for some features) in tJocer.

- [ ] Radar
  - [x] Show position of player
  - [ ] Show rotation of player's camera (in progress: pointer found)
  - [x] Show possible locations of objects
  - [ ] Show locations of objects spawned
  - [ ] Show location of enemy
  - [ ] Display warning icon if enemy sees you
  - [ ] Write pathfinder module to draw a line to nearest collectible
  - [ ] Zooming map while in game
  - [ ] Mode switcher (switch to fixed map like in PacMan)
- [ ] Radar levels
  - [ ] Freddy, Chica and Foxy Levels
  - [x] Bonnie
- [ ] Timer
  - [ ] Freeze timer
  - [ ] Reset timer/set to value
- [ ] Objects counter
  - [ ] Set to value
- [ ] Player's movement
  - [ ] Teleport
  - [ ] Noclip
  - [ ] Change walk/run speed
- [ ] Enemies
  - [ ] Freeze in place
- [ ] Visuals
  - [ ] Disable all effects
  - [ ] Wall-hack (see enemy through the walls)
  - [ ] Switch to wireframe-mode
- [ ] Gameplay
  - [ ] Slow down or speed up game time (without affecting player's movement, actions and FPS)

## Installation

Download latest release here.

## Build

