# Consultor de CEP

  
O Consultor de CEP é um microserviço que importa a base eDNE dos Correios e a partir de uma API REST retorna informações de um determinado endereço a partir de um CEP.


## Instalação

Obtenha o programa:

	git clone https://github.com/fsilva1985/consultor-de-cep

Faça a Build via docker:

	docker-compose up --build

## Como funciona

Durante a Build o importador faz a leitura o arquivo eDNE localizado dentro do projeto e importa para o banco de dados. Ambos estão como variavel de ambiente no arquivo **.env**

Caso queira atualizar os endereços, basta subir seu arquivo para dentro do projeto, fazer a alteração no **.env** e rodar novamente a Build.

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
## Como utilizar