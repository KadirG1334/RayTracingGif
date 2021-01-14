package main

import (
	"strconv"
	
	"fmt"

	"os"
	
	
	"time"
	"log"
	"image/gif"
	
	"image"
	
	"image/png"
	"io/ioutil"
	"image/color/palette"
    "image/draw"

)
var i = 0

const (
	fov      = 75.0
	width    = 500
	height   = 500
	samples  = 500
	aperture = 0.09
	cpus=12
	path ="images/")

func main() {
	lookFrom :=Vec3{X: 0, Y: 4, Z: 2}
	lookAt := Vec3{X: 0, Y: 0, Z: -6}
	
	camera := NewCamera(lookFrom, lookAt, fov, float64(width)/float64(height), aperture)

	start := time.Now()
	for i=0;i<4;{
	scene := Scene()
    i++
	
	imaged := Do(scene, camera, cpus, samples, width, height)
    s := strconv.Itoa(i) //for int to string
	f, err := os.Create("images/img"+s+".png")
	
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = png.Encode(f, imaged)}
	
	
	
	
	var files []os.FileInfo // for gif
	delay:=2
	
    files, err := ioutil.ReadDir(path)
    if err != nil {
        fmt.Println(err)
        return
	}
	

	anim := gif.GIF{}
    for _, info := range files {
        f, err := os.Open(path + "/" + info.Name())
        if err != nil {
            fmt.Printf("Could not open file %s. Error: %s\n", info.Name(), err)
            return
        }
        defer f.Close()
        img, _, _ := image.Decode(f)

        // Image has to be palleted before going into the GIF
        paletted := image.NewPaletted(img.Bounds(), palette.Plan9)
        draw.FloydSteinberg.Draw(paletted, img.Bounds(), img, image.ZP)

        anim.Image = append(anim.Image, paletted)
        anim.Delay = append(anim.Delay, delay*10)
    }

    f, _ := os.Create("examples/gif.gif")
    defer f.Close()
    gif.EncodeAll(f, &anim)

	


	fmt.Printf("\nDone. Elapsed: %v", time.Since(start))
	
	
	}



