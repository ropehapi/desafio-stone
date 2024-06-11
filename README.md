# API Árvore genealógica
*Antes de mais nada, gostaria de agradecer imensamente à equipe da stone pela
oportunidade e por me proprocionar esse desafio que fez meus olhos brilharem por
alguns dias.*

## Índice
1. Visão geral
2. Instalação
3. Uso
- Endpoints
    - Documentação no Swagger
    - BREAD de pessoa
    - Criar relacionamento entre pai e filho
    - Listar relacionamentos ascendentes indivíduo
    - Listar relacionamentos descendentes indivíduo

## Visão geral
Desenvolvida como a solução de um case técnico para o processo seletivo da Stone,
a API de árvore genealógica aqui desenvolvida é uma aplicação desenvolvida para
gerenciar pessoas e seus relacionamentos de modo a trazer de forma gráfica toda
a árvore genealógica de certo indivíduo, tanto de forma ascendente quanto
descendente.

Funcionando como uma API REST, a aplicação expõe endpoints que permitem o cadastro
de pessoas e relacionamentos entre pai e filho. A partir disso, podemos utilizar 
os endpoints de listagem de relacionamentos que podem trazer tanto de forma ascendente
quanto descendente todos os relacionamentos de um indivíduo.

## Instalação
Para rodar a aplicação localmente você vai precisar apenas do docker e do docker-compose
instalados na sua máquina.

1. Clone e acesse o diretório do repositório
> git clone git@github.com:ropehapi/desafio-stone.git

> cd desafio-stone/

2. Configure as variáveis de ambiente conforme o desejado em `cmd/server/.env`

3. Suba o conteiner da aplicação e do banco de dados
> docker compose up --build

4. Consuma a aplicação através do swagger em http://localhost:8080/docs/index.html#/

## Uso
Para usar a aplicação basta instalar e rodar o projeto conforme o tópico anterior e utilizar os
endpoints disponibilizados no swagger, ou ter em mãos um client http e importar a collection de
requisições deixada dentro do diretório `misc`.
Ainda assim, descreverei abaixo todos os endpoints a fins de documentação.

## Endpoints
### Documentação OpenAPI
- **Endpoint**: /docs/index.html#/
- **Método**: GET
- **Descrição**: Documentação OpenAPI visual da aplicação.

### Criar pessoa
- **Endpoint**: `/person`
- **Método**: POST
- **Descrição**: Cria uma nova pessoa.
- **Corpo da requisição**:
```json
{
  "name":"Leopoldino"
}
```
- **Resposta**: `201 Created` Detalhes da pessoa.
```json
{
"id": "dfc2af8b-0b9f-4b28-b991-101dcfc47f43",
"name": "Leopoldino"
}
```

### Listar pessoas
- **Endpoint**: `/person`
- **Método**: GET
- **Descrição**: Lista todas as pessoas.
- **Resposta**: `200 OK` Lista de pessoas detalhadas.
```json
[
  {
    "id": "017ff627-4d83-410b-9168-a166a1305b48",
    "name": "Leonilda"
  },
  {
    "id": "4fab833f-06be-46ad-a316-c308994a0efa",
    "name": "Haruo"
  }
]
```

### Detalhar pessoa
- **Endpoint**: `/person/:id`
- **Método**: GET
- **Descrição**: Trás os dados de uma pessoa.
- **Resposta**: `200 OK` Detalhes da pessoa.
```json
{
"id": "dfc2af8b-0b9f-4b28-b991-101dcfc47f43",
"name": "Leopoldino"
}
```

### Atualizar pessoa
- **Endpoint**: `/person/:id`
- **Método**: PUT
- **Descrição**: Atualiza os dados de uma pessoa.
- **Corpo da requisição**:
```json
{
  "name":"Jorge"
}
```
- **Resposta**: `200 OK` Dados da pessoa atualizados.
```json
{
  "id": "dfc2af8b-0b9f-4b28-b991-101dcfc47f43",
  "name": "Jorge"
}
```

### Deletar pessoa
- **Endpoint**: `/person`
- **Método**: DELETE
- **Descrição**: Deleta uma pessoa.
- **Resposta**: `200 OK` Em caso de sucesso

### Cadastrar relacionamento
- **Endpoint**: `/relationship`
- **Método**: POST
- **Descrição**: Cria um relacionamento de pai e filho.
- **Corpo da requisição**:
```json
{
  "childrenId": "4fab833f-06be-46ad-a316-c308994a0efa",
  "parentId": "dfc2af8b-0b9f-4b28-b991-101dcfc47f43"
}
```
- **Resposta**: `201 OK` Relacionamento criado.

### Obter árvore genealógica ascendente
- **Endpoint**: `/relationship/:id/asc`
- **Método**: GET
- **Descrição**: Trás a árvore genealógica da pessoa de forma ascendente.
- **Resposta**: `200 OK` Árvore genealógica ascendente do indivíduo.
```json
{
  "person": {
    "id": "d44cdacb-b07f-4f22-999f-89b399e5e99c",
    "name": "Filho",
    "relationships": [
      {
        "parent": {
          "id": "4fab833f-06be-46ad-a316-c308994a0efa",
          "name": "Pai",
          "relationships": [
            {
              "parent": {
                "id": "7e1eede6-71ed-46ce-a55e-4200aa6aa051",
                "name": "Vô",
                "relationships": null
              }
            },
            {
              "parent": {
                "id": "940b52ee-e95c-4a4c-83af-04f4b3434b48",
                "name": "Vó",
                "relationships": null
              }
            }
          ]
        }
      },
      {
        "parent": {
          "id": "8a2b7099-97ef-48d1-97a1-948a2aa468ff",
          "name": "Mãe",
          "relationships": null
        }
      }
    ]
  }
}
```

### Obter árvore genealógica descendente
- **Endpoint**: `/relationship/:id/desc`
- **Método**: GET
- **Descrição**: Trás a árvore genealógica da pessoa de forma descendente.
- **Resposta**: `200 OK` Árvore genealógica descendente do indivíduo.
```json
{
  "person": {
    "id": "7e1eede6-71ed-46ce-a55e-4200aa6aa051",
    "name": "Vô",
    "relationships": [
      {
        "children": {
          "id": "4fab833f-06be-46ad-a316-c308994a0efa",
          "name": "Pai",
          "relationships": [
            {
              "children": {
                "id": "ba5b336e-5d2b-4eed-979d-b4835b868782",
                "name": "Filho 1",
                "relationships": null
              }
            },
            {
              "children": {
                "id": "d44cdacb-b07f-4f22-999f-89b399e5e99c",
                "name": "Filho 2",
                "relationships": null
              }
            }
          ]
        }
      }
    ]
  }
}
```
---  
# Nível de desenvolvimento
## Testes
Todas as três camadas da aplicação foram testadas, desde testes unitários na camada de domínio a testes
de integração nas camadas de persistência e aplicação. Por conta de tempo fiquei devendo um conteiner
para a execução de testes, mas caso deseje, você pode utilizar os comandos definidos no `makefile`
para rodar os testes em seu ambiente.
## Arquitetura empregada
Abaixo farei uma breve descrição de como as camadas da aplicação foram delimitadas e
como são seus funcionamentos.

### Domain layer
**Diretório**: `internal/entity`

Camada onde residem os domínios da nossa aplicação, no caso `person` e `relationship`, é a
camada responsável por prover toda a regra de negócio do nosso software para as demais camadas.
Seguindo a filosofia de uma boa arquitetura, optei por definir no domínio da aplicação as interfaces
dos repositórios de cada entidade, camada responsável pela comunicação com o banco, que devem ter
suas implementações específicas na camada de infraestrutura.

### Application layer
**Diretório**: `internal/application`

Camada onde foram implementados os usecases, que nada mais são os objetos que acessam e manipulam
nosso domínio de forma a representar as intenções do usuário. Essa camada pode ser chamada a partir
de qualquer forma de expor nossa aplicação na camada de infra (como REST, gRPC, CLI etc) e através de 
injeção de dependências, manipula os repositórios dos nossos domínios.

### Infrastucture layer
**Diretório**: `internal/infra`

Camada onde ficam as interfaces de comunicação com a WEB e com os  nossos bancos de dados, nas pastas 
`web` e `database` respectivamente. Como nossa aplicação segue uma boa arquitetura, a forma como vamos
expor nossa aplicação, o banco de dados em que faremos a persistência entre outros detalhes devem ser 
apenas isso, detalhes, por isso, nessa camada implementamos a persitência no driver de banco de dados
desejado seguindo a interface do nosso domínio, e também implementamos nossos handlers da web, responsáveis
por receber requisições HTTP e a partir disso fazer o consumo dos usecases.

## Conteiner de DI
Afim de facilitar a injeção das dependências, optei por utilizar o [wire](https://github.com/google/wire),
um conteiner de injeção de dependências.

## Checklist durante o desenvolvimento em ordem
- [x] Modelar banco de dados de acordo com a regra de negócio
- [x] Escrever a camada de domínio
    - [x] Implementar validações
      - [x] Filho obrigatóriamente deve ter pai e mãe
      - [x] Filho não pode ser seu próprio descendente
      - [x] Filho nao pode ser pai do seu irmão
- [x] Escrever a camada de persistência
- [x] Escrever a camada Web
- [x] Escrever os usecases
  - [x] Crud de pessoa
  - [x] Crud de relacionamento
  - [x] Get tree
- [x] Implementar testes
  - [x] Domínio aplicação
  - [x] Camada de persistencia
  - [x] Usecases
- [x] Documentar a aplicação com swagger
- [x] Conteinerizar a aplicação
- [x] Documentar o projeto
- [x] Documentar o desenvolvimento (vídeo)
- [ ] (Extra) Montar um container para a execução dos testes