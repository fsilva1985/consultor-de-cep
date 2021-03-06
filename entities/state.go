package entities

// State returns collection
type State struct {
	ID   uint
	Code string
	Name string
}

// GetStates returns collection
func GetStates() []State {
	var states = []State{
		{ID: 1, Code: "AC", Name: "Acre"},
		{ID: 2, Code: "AL", Name: "Alagoas"},
		{ID: 3, Code: "AP", Name: "Amapá"},
		{ID: 4, Code: "AM", Name: "Amazonas"},
		{ID: 5, Code: "BA", Name: "Bahia"},
		{ID: 6, Code: "CE", Name: "Ceará"},
		{ID: 7, Code: "DF", Name: "Distrito Federal"},
		{ID: 8, Code: "ES", Name: "Espírito Santo"},
		{ID: 9, Code: "GO", Name: "Goiás"},
		{ID: 10, Code: "MA", Name: "Maranhão"},
		{ID: 11, Code: "MT", Name: "Mato Grosso"},
		{ID: 12, Code: "MS", Name: "Mato Grosso do Sul"},
		{ID: 13, Code: "MG", Name: "Minas Gerais"},
		{ID: 14, Code: "PA", Name: "Pará"},
		{ID: 15, Code: "PB", Name: "Paraíba"},
		{ID: 16, Code: "PR", Name: "Paraná"},
		{ID: 17, Code: "PE", Name: "Pernambuco"},
		{ID: 18, Code: "PI", Name: "Piauí"},
		{ID: 19, Code: "RJ", Name: "Rio de Janeiro"},
		{ID: 20, Code: "RN", Name: "Rio Grande do Norte"},
		{ID: 21, Code: "RS", Name: "Rio Grande do Sul"},
		{ID: 22, Code: "RO", Name: "Rondônia"},
		{ID: 23, Code: "RR", Name: "Roraima"},
		{ID: 24, Code: "SC", Name: "Santa Catarina"},
		{ID: 25, Code: "SP", Name: "São Paulo"},
		{ID: 26, Code: "SE", Name: "Sergipe"},
		{ID: 27, Code: "TO", Name: "Tocantins"},
	}

	return states
}
