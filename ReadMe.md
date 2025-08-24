
// import (
// 	"bytes"
// 	"image/color"
// 	"log"
// 	"math/rand"
// 	"time"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
// 	"github.com/hajimehoshi/ebiten/v2/text/v2"
// 	"github.com/hajimehoshi/ebiten/v2/vector"
// )

// var (
// 	dirUp           = Point{0, -1}
// 	dirDown         = Point{0, 1}
// 	dirLeft         = Point{-1, 0}
// 	dirRight        = Point{1, 0}
// 	mplusFaceSource *text.GoTextFaceSource
// )

// const (
// 	gameSpeed    = time.Second / 6
// 	screenWidth  = 640
// 	screenHeight = 480
// 	gridSize     = 20
// )

// type Point struct {
// 	x, y int
// }
// type Game struct {
// 	snake     []Point
// 	direction Point
// 	lastMove  time.Time
// 	food      Point
// 	gameOver  bool
// }

// func (g *Game) Update() error {
// 	if g.gameOver {
// 		return nil
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyW) {
// 		g.direction = dirUp
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyS) {
// 		g.direction = dirDown
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyA) {
// 		g.direction = dirLeft
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyD) {
// 		g.direction = dirRight
// 	}

// 	if time.Since(g.lastMove) < gameSpeed {
// 		return nil
// 	}
// 	g.lastMove = time.Now()
// 	g.UpdateSnake(&g.snake, g.direction)

// 	return nil
// }

// func (g *Game) UpdateSnake(snake *[]Point, direction Point) {
// 	head := (*snake)[0]
// 	newHead := Point{
// 		head.x + direction.x,
// 		head.y + direction.y,
// 	}

// 	if g.isCollision(newHead, *snake) {
// 		g.gameOver = true
// 		return
// 	} else {
// 		g.gameOver = false
// 	}

// 	if newHead == g.food {
// 		*snake = append([]Point{newHead}, *snake...)
// 		g.SpowenFood()
// 	} else {
// 		*snake = append([]Point{newHead}, (*snake)[:len(*snake)-1]...)
// 	}
// }

// func (g Game) isCollision(p Point, snake []Point) bool {
// 	if p.x < 0 || p.x >= screenWidth/gridSize || p.y < 0 || p.y >= screenHeight/gridSize {
// 		return true
// 	}
// 	for _, sp := range snake {
// 		if p == sp {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (g *Game) Draw(screen *ebiten.Image) {

// 	vector.DrawFilledRect(screen,
// 		float32(g.food.x*gridSize),
// 		float32(g.food.y*gridSize),
// 		float32(gridSize), float32(gridSize),
// 		color.RGBA{255, 255, 0, 255},
// 		true)

// 	if g.gameOver {
// 		face := &text.GoTextFace{
// 			Source: mplusFaceSource,
// 			Size:   48,
// 		}
// 		w, s := text.Measure("Game Over", face, face.Size)

// 		op := &text.DrawOptions{}
// 		op.GeoM.Translate(screenWidth/2-(w)/2, screenHeight/2-(s)/2)
// 		text.Draw(screen,
// 			"Game Over",
// 			face,
// 			op)
// 		return
// 	}

// 	for _, p := range g.snake {
// 		vector.DrawFilledRect(screen,
// 			float32(p.x*gridSize),
// 			float32(p.y*gridSize),
// 			float32(gridSize), float32(gridSize),
// 			color.White, true)
// 	}

// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return 640, 480
// }

// func (g *Game) SpowenFood() {
// 	g.food = Point{
// 		rand.Intn(screenWidth / gridSize),
// 		rand.Intn(screenHeight / gridSize),
// 	}
// }

// func main() {
// 	ebiten.SetWindowSize(screenWidth, screenHeight)

// 	s, err := text.NewGoTextFaceSource(
// 		bytes.NewReader(fonts.MPlus1pRegular_ttf),
// 	)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	mplusFaceSource = s

// 	g := &Game{
// 		snake: []Point{
// 			{
// 				x: screenWidth / gridSize / 2,
// 				y: screenHeight / gridSize / 2,
// 			},
// 		},
// 		direction: Point{1, 0},
// 	}

// 	g.SpowenFood()

// 	ebiten.SetWindowTitle("My First Game")
// 	if err := ebiten.RunGame(g); err != nil {
// 		log.Fatal(err)
// 	}
// }
