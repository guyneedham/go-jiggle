# go-jiggle

Ever wish you could automate moving your mouse around?

./go-jiggle --period 2m

go-jiggle is a cross platform mouse moving app. It uses sdl2 to move the mouse pointer to a random location every time period.

## Installation
Mac:
```
brew install sdl2{,_image,_ttf,_mixer} pkg-config
go get github.com/veandco/go-sdl2/sdl
go build
```
