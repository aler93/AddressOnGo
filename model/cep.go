package model

type Cep struct {
	//Id          uint64 `json:id`
	Cep         string `json:cep`
	Rua         string `json:rua`
	Bairro      string `json:bairro`
	Cidade      string `json:cidade`
	Uf          string `json:uf`
	Unidade     string `json:unidade`
	Regiao      string `json:regiao`
	Complemento string `json:complemento`
	Ibge        string `json:ibge`
	Ddd         string `json:ddd`
	Siafi       string `json:siafi`
	Origem		string `json:origem`
	Local 		bool   `json:local`
	AutoUpdate  bool   `json:autoupdate`
}

type ViaCep struct {
	Cep         string `json:cep`
	Logradouro  string `json:logradouro`
	Bairro      string `json:bairro`
	Localidade  string `json:localidade`
	Uf          string `json:uf`
	Unidade     string `json:unidade`
	Regiao      string `json:regiao`
	Complemento string `json:complemento`
	Ibge        string `json:ibge`
	Ddd         string `json:ddd`
	Siafi       string `json:siafi`
}

type OpenCep struct {
	Cep         string `json:cep`
	Logradouro  string `json:logradouro`
	Bairro      string `json:bairro`
	Localidade  string `json:localidade`
	Uf          string `json:uf`
	Ibge        string `json:ibge`
	Complemento string `json:complemento`
}

type BrasilApi struct {
	Cep          string `json:cep`
	State        string `json:state`
	City         string `json:city`
	Neighborhood string `json:neighborhood`
	Street       string `json:street`
}

