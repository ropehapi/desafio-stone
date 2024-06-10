# API Árvore genealógica
*Antes de mais nada, gostaria de agradecer imensamente à equipe da stone pela
oportunidade e por me proprocionar esse desafio que fez meus olhos brilharem por
alguns dias.*

## Índice
1. Visão geral
2. Instalação
3. Uso
   - Endpoints
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
TODO: Escrever documentação instalação

## Uso
Para usar a aplicação basta instalar e rodar o projeto conforme o tópico anterior,
ter em mãos um client http e importar a collection de requisições deixada na raiz do projeto.
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
  
# Checklist durante o desenvolvimento
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
- [ ] Implementar testes
  - [x] Domínio aplicação
  - [x] Camada de persistencia
  - [ ] Usecases
- [x] Documentar a aplicação com swagger
- [ ] Conteinerizar a aplicação
- [ ] Documentar o projeto
- [ ] Documentar o desenvolvimento (vídeo)