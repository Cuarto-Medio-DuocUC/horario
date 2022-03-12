package table

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

var (
	header      = []string{"Día", "Inicio", "Fin", "Ramo", "Código", "Sede", "Aula"}
	headerColor = tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor}
)

func headerColors() []tablewriter.Colors {
	var colors []tablewriter.Colors
	for i := 0; i < len(header); i++ {
		colors = append(colors, headerColor)
	}

	return colors
}

type Table struct {
	table *tablewriter.Table
}

func NewTable(writer io.Writer) Table {
	table := tablewriter.NewWriter(writer)
	table.SetAutoWrapText(false)
	table.SetHeader(header)
	table.SetHeaderColor(headerColors()...)
	table.SetAutoMergeCellsByColumnIndex([]int{0, 3, 4, 5, 6})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetRowLine(true)

	return Table{table}
}

func (t Table) Render() {
	t.table.Render()
}

func (t Table) Append(row []string) {
	t.table.Append(row)
}
