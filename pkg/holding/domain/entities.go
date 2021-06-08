package holding_domain

type Holdings struct {
	Holdings []Holding
}

type Holding struct {
	Id              	int64
	Trimestre    		int64
	ReceitaLiquida	    int64
	Ebitda	     		int64
	MargemEbitda		float32
	LucroLiquido	    int64
	MargemLiquida		float32
	DividaLiquida	    int64
	DivEbitda			float32
}