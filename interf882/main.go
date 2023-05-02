package main

import (
	"errors"
	"fmt"
)

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

type dimensionsInch struct {
	length Unit
	width  Unit
	height Unit
}

type dimensionsCM struct {
	length Unit
	width  Unit
	height Unit
}

type Dimensions interface {
	Length() (Unit, error)
	Width() (Unit, error)
	Height() (Unit, error)
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

func NewUnit(v float64, t UnitType) Unit {
	return Unit{Value: v, T: t}
}

func (d *dimensionsInch) Length() (Unit, error) {

	v, err := d.length.Get(Inch)
	if err != nil {
		return NewUnit(0, Inch), fmt.Errorf("Ошибка при вычислении длины: %w", err)
	}

	return NewUnit(v, Inch), nil
}

func (d *dimensionsInch) Width() (Unit, error) {

	v, err := d.width.Get(Inch)
	if err != nil {
		return NewUnit(0, Inch), fmt.Errorf("Ошибка при вычислении длины: %w", err)
	}

	return NewUnit(v, Inch), nil
}

func (d *dimensionsInch) Height() (Unit, error) {

	v, err := d.height.Get(Inch)
	if err != nil {
		return NewUnit(0, Inch), fmt.Errorf("Ошибка при вычислении длины: %w", err)
	}

	return NewUnit(v, Inch), nil
}

func (d *dimensionsCM) Length() (Unit, error) {

	v, err := d.length.Get(CM)
	if err != nil {
		return NewUnit(0, CM), fmt.Errorf("Ошибка при вычислении длины: %w", err)
	}

	return NewUnit(v, CM), nil
}

func (d *dimensionsCM) Width() (Unit, error) {

	v, err := d.width.Get(CM)
	if err != nil {
		return NewUnit(0, CM), fmt.Errorf("Ошибка при вычислении длины: %w", err)
	}

	return NewUnit(v, CM), nil
}

func (d *dimensionsCM) Height() (Unit, error) {

	v, err := d.height.Get(CM)
	if err != nil {
		return NewUnit(0, CM), fmt.Errorf("Ошибка при вычислении длины: %w", err)
	}

	return NewUnit(v, CM), nil
}

func (u Unit) Get(t UnitType) (float64, error) {
	value := u.Value

	if t != u.T {
		switch u.T {
		case Inch:
			value = convertInchСm(u.Value)
		case CM:
			value = convertСmInch(u.Value)
		default:
			return 0, errors.New("Недопустимый тип для конвертации")
		}
	}

	return value, nil
}

func convertInchСm(v float64) float64 {
	return v * 2.54
}

func convertСmInch(v float64) float64 {
	return v / 2.54
}

type baseAuto struct {
	brand       string
	model       string
	maxSpeed    int
	enginePower int
	dimensions  Dimensions
}

func (p *baseAuto) Brand() string {
	return p.brand
}

func (p *baseAuto) Model() string {
	return p.model
}

func (p *baseAuto) MaxSpeed() int {
	return p.maxSpeed
}

func (p *baseAuto) EnginePower() int {
	return p.enginePower
}

func (p *baseAuto) Dimensions() Dimensions {
	return p.dimensions
}

type bmwAuto struct {
	baseAuto
}

type mercedesAuto struct {
	baseAuto
}

type dodgeAuto struct {
	baseAuto
}

func NewDimensionsInch(l, w, h float64) dimensionsInch {
	d := dimensionsInch{length: NewUnit(l, Inch), width: NewUnit(l, Inch), height: NewUnit(l, Inch)}
	return d
}

func NewDimensionsCM(l, w, h float64) dimensionsCM {
	d := dimensionsCM{length: NewUnit(l, CM), width: NewUnit(l, CM), height: NewUnit(l, CM)}
	return d
}

func main() {

	bmw := bmwAuto{
		baseAuto: baseAuto{brand: "BMV",
			model:       "x7",
			maxSpeed:    270,
			enginePower: 350,
			dimensions:  NewDimensionsCM(500, 250, 200)},
	}
}
