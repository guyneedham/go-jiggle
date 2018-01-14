package main

import "flag"
import "fmt"
import "time"
import "math/rand"

import "github.com/veandco/go-sdl2/sdl"


func warpMouse(window *sdl.Window, max_x int, max_y int) {
    window.WarpMouseInWindow(int32(rand.Intn(max_x)), int32(rand.Intn(max_y)))
}

func getWindow(max_x int32, max_y int32) *sdl.Window {

   window, err := sdl.CreateWindow("jiggle", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
        max_x, max_y, sdl.WINDOW_HIDDEN)
   if err != nil {
      panic(err)
   }
   return window
}

func jiggle() {

    max_x := 800
    max_y := 600

    window := getWindow(int32(max_x), int32(max_y))

    defer window.Destroy()
    x, y := window.GetSize()
    warpMouse(window, int(x), int(y))
    warpMouse(window, int(x), int(y))
}

func main() {

    period := flag.String("period", "5s", "duration to wait until jiggling. Valid time units: {ns, us (or µs), ms, s, m, h}.")
    flag.Parse()
    fmt.Println("period: ", *period)

    time_period, err := time.ParseDuration(*period)

    if err != nil {
        panic(err)
    }

    for {
        time.Sleep(time_period)
        jiggle()        
    }

}
