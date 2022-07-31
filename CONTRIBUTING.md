# Contributing

Hello and thank you for your time!

tJocer is built with Go, UI is built on giu (which is built on imgui-go wrapper for dear imgui).

## How to find pointers?

Follow the general recommendations on how to find addresses in Cheat Engine such as looks for coordinates in vector when found one of coordinates and use hotkeys for faster search.

I used pointerscans (not pointermaps!) to find static pointers

- There is a shared coordinate address for all levels, though you may think there is a separate coordinate system for each level. It has really different address than other coordinates and you will usually see 2 addresses with weird prefix and 10 or so more with usual prefix (1 or 2) for coordinates. Look at these two addresses, go through different levels. One of them will be zero, other is what you're looking for. None of them are writable. It takes 3 or 4 restarts of game to find static pointer for coordinates.