# tJocer

Trainer for The Joy of Creation: Reborn, written in Go.

## This project is abandoned

I have no time and no energy left to keep this project alive. I'm too exhausted after 2 months of trying to find a way to get base address of process and other simple tasks. Yes, that may sound simple enough for you, but don't tell me about EnumProcessModules, GetModuleInformation, modBaseAddr and shit, none of this works, at least when I tried to do it. Go ahead and try to fix mem.go, specifically, incorrect base address (also known as image_hash or preferred_load_address etc) that is returned by module handle.

## Features

Implemented features are checked, others are either in-progress or proposals. You can open PR any time and I will try to accept it in next few years :) The hardest job is to find static pointer to value in memory, then you have to experiment with it and try to write new value, then implement logic (+render for some features) in tJocer.

- [x] Radar: show position of player
- [ ] Radar: show rotation of player's camera (in progress: pointer found)
- [x] Radar: Show possible locations of objects
- [ ] Radar: Show locations of objects spawned
- [ ] Radar: Show location of enemy
- [x] Radar: Bonnie level
- [ ] Radar: Freddy, Chica and Foxy Levels
- [ ] Timer: freeze timer
- [ ] Timer: reset timer/set to value
- [ ] Objects counter: set to value
- [ ] Player: teleport
- [ ] Player: noclip
- [ ] Player: change speed
- [ ] Enemies: freeze in place
- [ ] Visuals: disable all effects
- [ ] Visuals: wall-hack (see enemy through the walls)
- [ ] Visuals: Switch to wireframe-mode
- [ ] Gameplay: change speed of game excluding player's movement and actions and FPS