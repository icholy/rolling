# Rolling - data windows for gonum/plot

* Rolling data for plots implemented as array backed ring buffers.
* Provides types that implement: `gonum.org/v1/plot/plotter.{Valuer,XYer,XYZr}`

### Example:

``` go
// only retain the last 1000 added values
xys := rolling.NewXYs(1000)
xys.Add(plotter.XY{X: 1, Y: 0})
xys.Add(plotter.XY{X: 0, Y: 1})
```
 


