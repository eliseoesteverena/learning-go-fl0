package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	generate()
}
func generate() {
	pdf := gofpdf.New("P", "mm", "Legal", "") //<- config por default

	//indentación de acuerdo al page size:
	w, h := pdf.GetPageSize()
	//xIndent := w/10.0
	fmt.Printf("ancho=%v, alto=%v \n", w, h)

	pdf.AddPage()

	/* ::Cosas de texto basicas:: */
	pdf.MoveTo(0, 0) // opcional
	pdf.SetFont("times", "", 12)

	//-> Agregar imagen:
	pdf.ImageOptions("example/files/cd.jpg", 0, 0, 215.9, 355.6, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")

	pdf.SetTextColor(0, 0, 0) // rgb
	pdf.Text(25, 35, "Nombre")

	_, lineHeight := pdf.GetFontSize()
	pdf.SetXY(0, 0)
	pdf.SetXY(17, 158)

	pdf.MultiCell(182, lineHeight*1, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam vel iaculis dolor, et semper ipsum. Fusce ornare dictum posuere. Donec mauris lectus, aliquet sit amet eros sit amet, posuere accumsan arcu. Phasellus posuere malesuada commodo. Donec orci lorem, vulputate auctor sapien nec, blandit molestie lectus.Donec ullamcorper leo in urna gravida, vitae malesuada tellus imperdiet. Praesent vitae convallis est, eu dignissim ante. Sed nisi elit, porta sit amet elit in, placerat dictum odio. Pellentesque nec feugiat velit, sed dictum nisl. Curabitur vel risus in metus volutpat blandit. Vestibulum fermentum et felis tempor scelerisque. Integer id eros id leo consectetur commodo vel ut metus. Nullam molestie nisi suscipit elit dictum, ut varius nulla consectetur. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Vivamus purus nunc, fermentum eget semper facilisis, tincidunt non velit. Pellentesque porta consequat lacus in ullamcorper. Sed sagittis magna nec sapien euismod efficitur. Sed vitae lectus ut felis ullamcorper euismod nec ac est. Ut eu risus erat. Aenean lobortis pellentesque libero, eget fermentum eros pretium vitae. Quisque egestas turpis in vehicula rhoncus.", "", "", false)

	err := pdf.OutputFileAndClose("example/files/pdf" + setName() + ".pdf")
	if err != nil {
		panic(err)
	}
}
func setName() string {
	var now time.Time
	now = time.Now()
	date := strconv.Itoa(now.Second()) + strconv.Itoa(now.Minute()) + strconv.Itoa(now.Hour()) + strconv.Itoa(now.Day()) + strconv.Itoa(int(now.Month())) + strconv.Itoa(now.YearDay())
	return date
}
