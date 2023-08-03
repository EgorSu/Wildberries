package pattern

import "fmt"

type Builder interface {
	buildRoof() string
	buildWall() string
	buildFoundation() string
}

type House struct {
	builder Builder
}

func (h *House) build() {
	fmt.Println(h.builder.buildRoof())
	fmt.Println(h.builder.buildWall())
	fmt.Println(h.builder.buildFoundation())
}

type StoneBuilder struct {
}

func (sb *StoneBuilder) buildWall() string {
	return "stone wall"
}
func (sb *StoneBuilder) buildRoof() string {
	return "metal roof"
}
func (sb *StoneBuilder) buildFoundation() string {
	return "stone foundation"
}

type WoodBuilder struct {
}

func (wb *WoodBuilder) buildWall() string {
	return "wooden wall"
}
func (wb *WoodBuilder) buildRoof() string {
	return "wooden roof"
}
func (wb *WoodBuilder) buildFoundation() string {
	return "concrete foundation"
}
