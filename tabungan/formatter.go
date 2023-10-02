package tabungan

type TabunganFormatter struct {
	NoRekening	int `json:"no_rekening"`
}

type SaldoFormatter struct {
	Saldo	int `json:"saldo"`
}

type SaldoDetailFormatter struct {
	NoRekening	int 	`json:"no_rekening"`
	Saldo		int 	`json:"saldo"`
}

type TabunganDetailFormatter struct {
	NoRekening	int 					`json:"no_rekening"`
	Mutasi		[]MutasiDetailFormatter	`json:"mutasi"`
}

type MutasiDetailFormatter struct {
	Type       string    `json:"type"`
	Nominal    int       `json:"nominal"`
}

func FormatTabungan(tabungan Tabungan) TabunganFormatter {
	tabunganFormatter := TabunganFormatter{}
	tabunganFormatter.NoRekening = tabungan.NoRekening
	return tabunganFormatter
}

func FormatSaldo(Saldo int) SaldoFormatter {
	saldoFormatter := SaldoFormatter{}
	saldoFormatter.Saldo = Saldo
	return saldoFormatter
}

func FormatTabunganDetail(tabungan Tabungan) TabunganDetailFormatter {
	tabunganFormatter := TabunganDetailFormatter{}
	tabunganFormatter.NoRekening = tabungan.NoRekening
	
	mutasis := []MutasiDetailFormatter{}
	for _, mutasi := range tabungan.Mutasi {
		mutasiFormatter := MutasiDetailFormatter{}
		mutasiFormatter.Type = mutasi.Type
		mutasiFormatter.Nominal = mutasi.Nominal
		mutasis = append(mutasis, mutasiFormatter)
	}

	tabunganFormatter.Mutasi = mutasis
	return tabunganFormatter
}

func SaldoTabunganDetail(tabungan Tabungan) SaldoDetailFormatter {
	tabunganFormatter := SaldoDetailFormatter{}
	tabunganFormatter.NoRekening = tabungan.NoRekening
	tabunganFormatter.Saldo = tabungan.Saldo
	return tabunganFormatter
}
