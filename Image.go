package main

import (
	"net/http"
	"image"

	"github.com/fogleman/gg"
)

func PrintImage(respWr http.ResponseWriter, req *http.Request, d *Data, label string) {
    dc := gg.NewContext(X, Y)
    dc.SetRGB(0.9, 0.9, 0.9)
    dc.Clear()

    for _, c := range d.circles {
	    dc.DrawCircle(float64(c.x), float64(c.y), float64(c.r))
	    dc.SetRGB(float64(c.red), float64(c.green), float64(c.blue))
	    dc.FillPreserve()
	    dc.Stroke()
	}

    for _, r := range d.rectangles {
	    dc.DrawRectangle(float64(r.x), float64(r.y), float64(r.w), float64(r.h))
	    dc.SetRGB(float64(r.red), float64(r.green), float64(r.blue))
	    dc.FillPreserve()
	    dc.Stroke()
	}

    for _, t := range d.triangles {
	    dc.LineTo(float64(t.p1), float64(t.p2))
	    dc.LineTo(float64(t.p2), float64(t.p3))
	    dc.LineTo(float64(t.p3), float64(t.p1))
	    dc.SetRGB(float64(t.red), float64(t.green), float64(t.blue))
	    dc.FillPreserve()
	    dc.Stroke()
	}

	var img image.Image = dc.Image()
	writeImageWithTemplate(respWr, &img, label)
}
