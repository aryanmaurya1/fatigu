package main

import (
	"fmt"
	"strings"
)

type Metric struct {
	RoutineId   int64
	Start       int64
	End         int64
	Elasped     int64
	ElaspedInms int64
}

func (m Metric) String() string {
	repr := strings.Builder{}
	repr.Grow(300)

	repr.WriteString(strings.Repeat("-", 47))
	repr.WriteString("\n")

	// -------------------------------------------------------------------------------------
	repr.WriteString(fmt.Sprintf("| %-20s | %-20s |\n", "Metric", "Value"))
	// -------------------------------------------------------------------------------------

	repr.WriteString(strings.Repeat("-", 47))
	repr.WriteString("\n")

	// -------------------------------------------------------------------------------------
	repr.WriteString(fmt.Sprintf("| %-20s | %-20v |\n", "Metric", "Value"))
	repr.WriteString(fmt.Sprintf("| %-20s | %-20v |\n", "Routine ID", m.RoutineId))
	repr.WriteString(fmt.Sprintf("| %-20s | %-20v |\n", "Start Time", m.Start))
	repr.WriteString(fmt.Sprintf("| %-20s | %-20v |\n", "End Time", m.End))
	repr.WriteString(fmt.Sprintf("| %-20s | %-20v |\n", "Elasped Time (ns)", m.Elasped))
	// -------------------------------------------------------------------------------------

	repr.WriteString(strings.Repeat("-", 47))
	repr.WriteString("\n")

	// -------------------------------------------------------------------------------------
	repr.WriteString(fmt.Sprintf("| %-20s | %-20v |\n", "Elasped Time (ms)",
		fmt.Sprintf("%v ms", m.ElaspedInms)))
	// -------------------------------------------------------------------------------------

	repr.WriteString(strings.Repeat("-", 47))
	return repr.String()
}
