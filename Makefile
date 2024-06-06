# Nome do comando de execução
RUN_CMD = go run main.go wire_gen.go

# Diretório do servidor
SERVER_DIR = cmd/server

# Regra padrão para rodar o servidor
.PHONY: run
run:
	cd $(SERVER_DIR) && $(RUN_CMD)

# Regra para limpar (não é obrigatória mas pode ser útil)
.PHONY: clean
clean:
    # Aqui você pode adicionar comandos para limpar build artifacts, logs, etc.
	@echo "Clean is not implemented yet."

# Regra para instalar dependências (opcional)
.PHONY: deps
deps:
	go mod tidy

# Regra para testar (opcional)
.PHONY: test
test:
	go test ./...

# Regra para build (opcional)
.PHONY: build
build:
	cd $(SERVER_DIR) && go build -o server .

# Regra para rodar o build
.PHONY: start
start: build
	./$(SERVER_DIR)/server
