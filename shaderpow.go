package main

import (
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed test.kage
	testShaderSrc []byte

	testShader *ebiten.Shader
)

func main() {
	var err error

	testShader, err = ebiten.NewShader(testShaderSrc)
	if err != nil {
		log.Fatal(err)
	}
}
