# Consultor de CEP

O Consultor de CEP é um microserviço que importa a base eDNE dos Correios e a partir de uma API REST retorna informações de um determinado endereço a partir de um CEP.

## Como utilizar

Requisite esse endpoint:

http://localhost:1981/api/address/22430041

  A resposta da requisição:

	{
		"data": {
			"type": "Avenida",
			"name": "Borges de Medeiros",
			"neighborhood": "Leblon",
			"city": "Rio de Janeiro",
			"state": "Rio de Janeiro"
		}
	}

Resposta caso o CEP estiver errado ou não encontrado.

	{
		"message": "zipcode is not valid. || 'zipcode is not found."
	}
