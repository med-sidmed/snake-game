package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/med-sidmed/snake-game/common"
	"github.com/med-sidmed/snake-game/entity"
)

type Game struct {
	Snake     []entity.Point
	Food      entity.Point
	Direction entity.Point
	LastMove  time.Time
	GameOver  bool
}

func (g *Game) Update() error {
	if g.GameOver {
		return nil
	}
	if time.Since(g.LastMove) < 100*time.Millisecond {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Direction = common.DirUp
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Direction = common.DirDown
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Direction = common.DirLeft
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Direction = common.DirRight
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.GameOver = true
		return nil
	}

	g.LastMove = time.Now()

	g.UpdateSnake(&g.Snake, g.Direction)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0,
		common.ScreenWidth, common.ScreenHeight,
		color.White, true)

	for x := float32(0); x < common.ScreenWidth; x += common.GridSize {
		for y := float32(0); y < common.ScreenHeight; y += common.GridSize {
			vector.StrokeRect(screen, x, y, common.GridSize, common.GridSize,
				1, color.RGBA{200, 200, 200, 255}, true)
		}
	}

	for _, p := range g.Snake {
		vector.DrawFilledRect(screen,
			float32(p.X), float32(p.Y),
			common.GridSize, common.GridSize,
			color.RGBA{0, 255, 0, 255}, true)
	}

	vector.DrawFilledRect(screen,
		float32(g.Food.X), float32(g.Food.Y),
		common.GridSize, common.GridSize,
		color.RGBA{255, 0, 0, 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return common.ScreenWidth, common.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")

	g := &Game{
		Snake: []entity.Point{
			{X: common.ScreenWidth / 2, Y: common.ScreenHeight / 2},
		},
		Direction: entity.Point{X: common.GridSize, Y: 0},
		Food:      entity.Point{X: 100, Y: 100},
		GameOver:  false,
		LastMove:  time.Now(),
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) UpdateSnake(snake *[]entity.Point, direction entity.Point) {
	head := (*snake)[0]
	newHead := entity.Point{
		X: (head.X + direction.X + common.ScreenWidth) % common.ScreenWidth,
		Y: (head.Y + direction.Y + common.ScreenHeight) % common.ScreenHeight,
	}

	// Vérifier la collision
	if g.isCollision(newHead, *snake) {
		g.GameOver = true
		return
	}

	if newHead == g.Food {
		*snake = append([]entity.Point{newHead}, *snake...)
		g.SpawnFood()
	} else {
		*snake = append([]entity.Point{newHead}, (*snake)[:len(*snake)-1]...)
	}
}

func (g *Game) SpawnFood() {
	if common.GridSize == 0 {
		log.Fatal("GridSize cannot be zero")
	}
	for {
		newFood := entity.Point{
			X: rand.Intn(common.ScreenWidth/common.GridSize) * common.GridSize,
			Y: rand.Intn(common.ScreenHeight/common.GridSize) * common.GridSize,
		}
		if !g.isCollision(newFood, g.Snake) {
			g.Food = newFood
			break
		}
	}
}

func (g *Game) isCollision(p entity.Point, snake []entity.Point) bool {
	for _, segment := range snake {
		if p == segment {
			return true
		}
	}
	return false
}
