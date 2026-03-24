package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func calcTax(zvE float64, year int, splitting bool) float64 {
	tax := 0.0
	var err error
	if splitting == true {
		// Beim Splittingtarif wird das Einkommen halbiert, die Steuer berechnet und dann verdoppelt
		halfIncome := zvE / 2.0
		taxHalf, err := calculateTax(halfIncome, year)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		tax = taxHalf * 2.0
	} else {
		tax, err = calculateTax(zvE, year)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	return tax
}

// calculateTax berechnet die Einkommensteuer basierend auf dem Tarif
func calculateTax(income float64, year int) (float64, error) {

	zvE := income

	// Routinglogik für verschiedene Steuerjahre
	switch year {
	case 2022:
		return calculateTariff2022(zvE)
	case 2023:
		return calculateTariff2023(zvE)
	case 2024:
		return calculateTariff2024(zvE)
	case 2025:
		return calculateTariff2025(zvE)
	case 2026:
		return calculateTariff2026(zvE)
	}

	return 0.0, fmt.Errorf("Berechnung für das Jahr %d wird nicht unterstützt.\n", year)
}

// getBasicAllowance gibt den Grundfreibetrag für das Jahr zurück
func getBasicAllowance(year int) float64 {
	switch year {
	case 2023:
		return 11_000.0
	case 2024:
		return 11_600.0
	case 2025:
		return 11_600.0
	case 2026:
		return 11_600.0
	default:
		// Für zukünftige Jahre: 11.600€
		return 11_600.0
	}
}

func kidsAllowance(year int) float64 {
	// Kinderfreibetrag pro Kind
	switch year {
	case 2023: // Kinderfreibetrag 2023: 6.024€ + 2.928€ (für Betreuung, Erziehung, Ausbildung)
		return 8_952.0
	case 2024: // Kinderfreibetrag 2024: 6.612€ + 2.928€ (für Betreuung, Erziehung, Ausbildung)
		return 9_540.0
	case 2025: // Kinderfreibetrag 2025: 6.672€ + 2.928€ (für Betreuung, Erziehung, Ausbildung)
		return 9_600.0
	case 2026: // Kinderfreibetrag 2026: 6.828€ + 2.928€ (für Betreuung, Erziehung, Ausbildung)
		return 9_756.0
	default:
		// Für zukünftige Jahre:
		return 9_756.0
	}
}

// calculateTariff2023 berechnet Steuer nach dem 2023er Tarif basierend auf zvE (nach Grundfreibetrag)
func calculateTariff2023(zvE float64) (float64, error) {
	return 0.0, fmt.Errorf("Berechung für 2023 ist nicht implementiert.")
}

// calculateTariff2022 berechnet Steuer nach dem 2022er Tarif basierend auf zvE (nach Grundfreibetrag)
func calculateTariff2022(zvE float64) (float64, error) {
	return 0.0, fmt.Errorf("Berechung für 2022 ist nicht implementiert.")
}

// calculateTariff2024 berechnet Steuer nach dem 2024er Tarif basierend auf zvE (nach Grundfreibetrag)
func calculateTariff2024(zvE float64) (float64, error) {
	return 0.0, fmt.Errorf("Berechung für 2024 ist nicht implementiert.")
}

// calculateTariff2025 berechnet Steuer nach dem 2025er Tarif basierend auf zvE (nach Grundfreibetrag)
func calculateTariff2025(zvE float64) (float64, error) {
	return 0.0, fmt.Errorf("Berechung für 2025 ist nicht implementiert.")
}


// calculateTariff2026 berechnet Steuer nach dem 2026er Tarif basierend auf zvE (nach Grundfreibetrag)
// Der Tarif bleibt aktuell identisch mit 2024/2025
func calculateTariff2026(zvE float64) (float64, error) {
	zvE = math.Floor(zvE)

	basicAllowance := 12_348.0 // Grundfreibetrag für 2026
	tarifZone1 := 17_799.0     // Beginn der Zone 2 (ab zvE > 17.799€)
	tarifZone2 := 69_878.0     // Beginn der Zone 3 (ab zvE > 69.878€)
	tarifZone3 := 277_825.0    // Beginn der Zone 4 (ab zvE > 277.825€)

	tax := 0.0
	if zvE <= basicAllowance {
		tax = 0.0
	} else if zvE <= tarifZone1 {
		y := (zvE - basicAllowance) / 10_000.0
		tax = (914.51*y + 1_400) * y
	} else if zvE <= tarifZone2 {
		z := (zvE - tarifZone1) / 10_000.0
		tax = (173.10*z+2_397.0)*z + 1_034.87
	} else if zvE <= tarifZone3 {
		tax = 0.42*zvE - 11_135.63
	} else {
		tax = 0.45*zvE - 19_470.38
	}

	return math.Floor(tax), nil
}

func main() {
	income := flag.Float64("income", 0, "Zu versteuerndes Einkommen (in Euro)")
	year := flag.Int("year", 2026, "Steuerjahr (Standard: 2026)")
	splitting := flag.Bool("s", false, "Splittingtarif")
	kids := flag.Int("kids", 0, "Anzahl der Kinder (für Kinderfreibetrag, optional)")
	pv := flag.Float64("p", 0, "Lohnersatzleistung, die dem Progressionsvorbehalt unterliegt (optional)")

	flag.Parse()

	if *income <= 0 {
		fmt.Fprintf(os.Stderr, "Fehler: Das Einkommen muss größer als 0 sein.\n\n")
		fmt.Fprintf(os.Stderr, "Verwendung: est-cli -income <Betrag> [-year <Jahr>]\n\n")
		fmt.Fprintf(os.Stderr, "Optionen:\n")
		fmt.Fprintf(os.Stderr, "  -income <Betrag>  Zu versteuerndes Bruttoeinkommen in Euro (erforderlich)\n")
		fmt.Fprintf(os.Stderr, "  -year <Jahr>      Steuerjahr (Standard: 2026)\n")
		fmt.Fprintf(os.Stderr, "  -kids <Anzahl>    Anzahl der Kinder für Kinderfreibetrag (optional)\n")
		fmt.Fprintf(os.Stderr, "  -s 				Splittingtarif)\n")
		fmt.Fprintf(os.Stderr, "  -p <Betrag>		Lohnersatzleitung, die dem Progressionsvorbehalt unterliegt\n")
		fmt.Fprintf(os.Stderr, "\nHinweis: Das Einkommen sollte das zu versteuernde Einkommen sein, also nach Abzug von Werbungskosten, Sonderausgaben etc.\n")
		fmt.Fprintf(os.Stderr, "\nBeispiele:\n")
		fmt.Fprintf(os.Stderr, "  est-cli -income 50000         # Berechne Steuer für 50.000€ im Jahr 2024\n")
		fmt.Fprintf(os.Stderr, "  est-cli -income 70000 -year 2023  # Berechne Steuer für 70.000€ im Jahr 2023\n")
		os.Exit(1)
	}

	tax := 0.0
	zvE := *income - kidsAllowance(*year)*float64(*kids)

	tax = calcTax(zvE, *year, *splitting)
	if *pv > 0 {
		// Berechnung des Progressionsvorbehalts
		// Das Einkommen aus Lohnersatzleistungen wird zum zu versteuernden Einkommen hinzugerechnet, um den Steuersatz zu bestimmen, aber die Steuer wird nur auf das reguläre Einkommen berechnet.
		zvEWithPV := zvE + *pv
		taxWithPV := calcTax(zvEWithPV, *year, *splitting)
		taxDeductionPV := calcTax(*pv, *year, *splitting)
		tax = taxWithPV - taxDeductionPV
	}

	effectiveRate := (tax / *income) * 100.0

	fmt.Printf("\n=== Einkommensteuer-Berechnung ===\n")
	fmt.Printf("Steuerjahr:              %d\n", *year)
	fmt.Printf("Bruttoeinkommen:         %.2f €\n", *income)
	fmt.Printf("Kinderfreibetrag:       -%.2f €\n", kidsAllowance(*year)*float64(*kids))
	fmt.Printf("Zu versteuerndes Eink.:  %.2f €\n", zvE)
	fmt.Printf("---\n")
	fmt.Printf("Einkommensteuer:         %.2f €\n", tax)
	fmt.Printf("Effektiver Steuersatz:   %.2f %%\n", effectiveRate)
	fmt.Printf("Nettoeinkommen:          %.2f €\n", *income-tax)
	fmt.Printf("===================================\n\n")
}
