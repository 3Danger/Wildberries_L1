package main

import "fmt"

/*
	Реализовать паттерн «адаптер» на любом примере.
*/

func main() {
	fmt.Println("Before:")
	colorRGB := RGB{196, 2, 255}
	fmt.Printf("colorRGB %+v\n", colorRGB)

	colorBW := BW{255}
	fmt.Printf("colorBW  %+v\n", colorBW)

	// Адаптеры в действии
	fmt.Println("\nAfter:")
	colorBW.SetColor(&RGBtoBWAdapter{&colorRGB})
	fmt.Printf("colorBW  %+v\n", colorBW)

	colorRGB.SetColorRGB(&BWtoRGBAdapter{&colorBW})
	fmt.Printf("colorRGB %+v\n", colorRGB)
}

// BW Допустим у нас есть пакет
// Отвечающий за черно-белые цвета
type BW struct{ b int16 }
type GettableColor interface{ GetColor() int16 }

func (b *BW) GetColor() int16          { return b.b }
func (b *BW) SetColor(g GettableColor) { b.b = g.GetColor() }

// RGB Допустим мы имеем свой пакет
// Отвечающий за цвета RGB
type RGB struct{ R, G, B int16 }
type GettableColorRGB interface{ GetColorRGB() (int16, int16, int16) }

func (c *RGB) GetColorRGB() (int16, int16, int16) { return c.R, c.G, c.B }
func (c *RGB) SetColorRGB(g GettableColorRGB)     { c.R, c.G, c.B = g.GetColorRGB() }

// RGBtoBWAdapter
// Следующий Адаптер для RGB to BW будет выглядеть следующим образом
type RGBtoBWAdapter struct {
	adapter *RGB
}

func (c *RGBtoBWAdapter) GetColor() int16 {
	r, g, b := c.adapter.GetColorRGB()
	return int16((int32(r) + int32(g) + int32(b)) / 3)
}

// BWtoRGBAdapter
// Следующий Адаптер для BW to RGB будет выглядеть следующим образом
type BWtoRGBAdapter struct {
	adapter *BW
}

func (c *BWtoRGBAdapter) GetColorRGB() (int16, int16, int16) {
	b := c.adapter.GetColor()
	return b, b, b
}
