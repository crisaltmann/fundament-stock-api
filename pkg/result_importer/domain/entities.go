package result_importer_domain

import "time"

type Template struct {
	Trimestre int
	DRE       DRETemplate
	Balanco   BalancoTemplate
}

type DRETemplate struct {
	Name    					string
	ReceitaLiquida 				int
	ResultadoBruto				int
	DespesasVendas				int
	DespesasGeraisAdm			int
	LucroLiquido				int
}

type BalancoTemplate struct {
	Name    					string
	EmprestimoCirculante		int
	EmprestimoNaoCirculante		int
	Caixa						int
}

type ImporterResults struct {
	Results 	[]ImporterResult
	Codigo		string
}

type ImporterResult struct {
	Trimestre	Trimestre
	Balanco		Balanco
	DRE			DRE
}

type Trimestre struct {
	Y int
	Value string
	Date  time.Time
}

type Balanco struct {
	EmprestimoCirculante 		float32
	EmprestimoNaoCirculante 	float32
	Caixa                   	float32
}

type DRE struct {
	ReceitaLiquida 				float32
	ResultadoBruto				float32
	DespesasVendas				float32
	DespesasGeraisAdm			float32
	LucroLiquido				float32
}