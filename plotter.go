package rolling

import "gonum.org/v1/plot/plotter"

type Values struct {
	state
	data plotter.Values
}

func NewValues(size int) *Values {
	return &Values{
		state: state{size: size},
		data:  make(plotter.Values, size),
	}
}

var _ plotter.Valuer = &Values{}

func (v *Values) Len() int            { return v.length }
func (v *Values) Value(i int) float64 { return v.data.Value(v.at(i)) }
func (v *Values) Add(x float64)       { v.data[v.add()] = x }
func (v *Values) Reset()              { v.length = 0 }

type XYs struct {
	state
	data plotter.XYs
}

func NewXYs(size int) *XYs {
	return &XYs{
		state: state{size: size},
		data:  make(plotter.XYs, size),
	}
}

var _ plotter.XYer = &XYs{}

func (xy *XYs) Len() int                    { return xy.length }
func (xy *XYs) XY(i int) (float64, float64) { return xy.data.XY(xy.at(i)) }
func (xy *XYs) Add(x plotter.XY)            { xy.data[xy.add()] = x }
func (xy *XYs) Reset()                      { xy.length = 0 }

type XYZs struct {
	state
	data plotter.XYZs
}

func NewXYZs(size int) *XYZs {
	return &XYZs{
		state: state{size: size},
		data:  make(plotter.XYZs, size),
	}
}

var _ plotter.XYZer = &XYZs{}

func (xyz *XYZs) Len() int                              { return xyz.length }
func (xyz *XYZs) XY(i int) (float64, float64)           { return xyz.data.XY(xyz.at(i)) }
func (xyz *XYZs) XYZ(i int) (float64, float64, float64) { return xyz.data.XYZ(xyz.at(i)) }
func (xyz *XYZs) Add(x plotter.XYZ)                     { xyz.data[xyz.add()] = x }
func (xyz *XYZs) Reset()                                { xyz.length = 0 }
