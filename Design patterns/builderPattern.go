package main

import (
	"fmt"
	"log"
)

func main() {

	log.SetPrefix("main: ")
	log.SetFlags(0)
	builder := NewBuilder()
	doorFrame := NewFrame(100, 100, "blue")
	windowFrame := NewFrame(200, 800, "green")

	doors := []Door{}
	doors = append(doors, *NewDoor(*doorFrame))

	windows := []Window{}
	windows = append(windows, *NewWindow(*windowFrame))

	builder, err := builder.WithDoors(doors)
	if err != nil {
		log.Fatal(err)
	}
	builder, err = builder.WithWindows(windows)
	if err != nil {
		log.Fatal(err)
	}

	builder, err = builder.WithPrice(154000)
	if err != nil {
		log.Fatal(err)
	}

	builder, err = builder.WithNumberOfOccupants(5)
	if err != nil {
		log.Fatal(err)
	}
	builder, err = builder.WithNumberOfStoreys(2)
	if err != nil {
		log.Fatal(err)
	}

	builder, err = builder.WithStyle("Nordic")
	if err != nil {
		log.Fatal(err)
	}

	house, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The house has: %d doors and %d windows, with %d occupants", len(house.Doors), len(house.Windows), house.Occupants)

}
func (pn *PositiveNumber) ToString() string {
	return fmt.Sprintf("%d", pn.number)
}

type PositiveNumber struct {
	number uint
}

type Frame struct {
	height uint16
	width  uint16
	color  string
}

func NewFrame(height, width uint16, color string) *Frame {
	return &Frame{height: height, width: width, color: color}
}

type Window struct {
	frame Frame
}

func NewWindow(Frame Frame) *Window {
	return &Window{frame: Frame}
}

type Door struct {
	frame Frame
}

func NewDoor(Frame Frame) *Door {
	return &Door{frame: Frame}
}

type House struct {
	Doors     []Door
	Windows   []Window
	Storeys   uint8
	Occupants uint8
	Style     string
	Price     uint32
}

type IsBuilder interface {
	WithDoors(Doors []Door) (IsBuilder, error)
	WithWindows(Windows []Window) (IsBuilder, error)
	WithNumberOfStoreys(numberOfStoreys uint8) (IsBuilder, error)
	WithNumberOfOccupants(numberOfOccupants uint8) (IsBuilder, error)
	WithStyle(style string) (IsBuilder, error)
	WithPrice(price uint32) (IsBuilder, error)
	Build() (House, error)
}

func (builder Builder) Build() (House, error) {
	return builder.House, nil
}

func (builder *Builder) WithPrice(Price uint32) (IsBuilder, error) {
	builder.House.Price = Price
	return builder, nil
}

func (builder *Builder) WithStyle(Style string) (IsBuilder, error) {
	if len(Style) <= 0 {
		return nil, fmt.Errorf("Style cannot be empty")
	}
	builder.House.Style = Style
	return builder, nil
}

func (builder *Builder) WithNumberOfStoreys(numberOfStoreys uint8) (IsBuilder, error) {
	builder.House.Storeys = numberOfStoreys
	return builder, nil
}

func (builder *Builder) WithWindows(Windows []Window) (IsBuilder, error) {
	if len(Windows) == 0 {
		return nil, fmt.Errorf("No windows provided")
	}
	builder.House.Windows = Windows
	return builder, nil
}

func (builder *Builder) WithNumberOfOccupants(numberOfOccupants uint8) (IsBuilder, error) {
	builder.House.Occupants = numberOfOccupants
	return builder, nil
}

func (builder *Builder) WithDoors(Doors []Door) (IsBuilder, error) {
	if len(Doors) == 0 {
		return nil, fmt.Errorf("No doors provided")
	}
	builder.House.Doors = Doors
	return builder, nil
}

type Builder struct {
	House House
}

func NewBuilder() IsBuilder {
	return &Builder{}
}
