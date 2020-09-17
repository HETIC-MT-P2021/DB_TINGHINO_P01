DOCKER_COMPOSE ?= docker-compose
GO_WORKDIR ?= ""

.PHONY: help
help: ## Provides help information on available commands
	@printf "Usage: make <command>\n\n"
	@printf "Commands:\n"
	@awk -F ':(.*)## ' '/^[a-zA-Z0-9%\\\/_.-]+:(.*)##/ { \
	  printf "  \033[36m%-30s\033[0m %s\n", $$1, $$NF \
	}' $(MAKEFILE_LIST)
	@printf "\nNOTE: the / is interchangable with the : in target names\n"

.PHONY: compose/build
compose/build: ## Build all Docker images of the project
	@$(DOCKER_COMPOSE) build

.PHONY: compose/up
compose/up: ## Start all containers (in the background)
	@$(DOCKER_COMPOSE) up -d

.PHONY: compose/down
compose/down: ## Stops and deletes containers and networks created by "up".
	@$(DOCKER_COMPOSE) down

.PHONY: compose/restart
compose/restart: compose/down compose/up ## Restarts all containers

.PHONY: compose/start
compose/start: ## Starts existing containers for a service
	@$(DOCKER_COMPOSE) start

.PHONY: compose/stop
compose/stop: ## Stops containers without removing them
	@$(DOCKER_COMPOSE) stop

.PHONY: compose/purge/local
compose/purge/local:
	@$(DOCKER_COMPOSE) down -v --rmi local

.PHONY: compose/purge
compose/purge: compose/purge/local ## Stops and deletes containers, volumes, images (local) and networks created by "up".

.PHONY: compose/purge/all
compose/purge/all: ## Stops and deletes containers, volumes, images (all) and networks created by "up".
	@$(DOCKER_COMPOSE) down -v --rmi all

.PHONY: compose/rebuild
compose/rebuild: compose/down compose/build compose/up ## Rebuild the project

.PHONY: lint
lint: ## Run golangci-lint (All-In-One config)
	@docker run --rm -v ${PWD}${GO_WORKDIR}:/app -w /app golangci/golangci-lint golangci-lint run --out-format tab | \
	awk -F '[[:space:]][[:space:]]+' '{ \
		error_file = $$1 ; \
		linter_name = $$2 ; \
		error_message = $$3 ; \
		split(error_file, error_file_info, ":") ; \
		error_file_path = sprintf(".%s/%s", ${GO_WORKDIR}, error_file_info[1]) ; \
		error_line_number = error_file_info[2] ; \
		error_col_number = error_file_info[3] ; \
		\
		dashed_line_length = 80 ; \
		dashed_line = sprintf("%*s", dashed_line_length, ""); gsub(/ /, "-", dashed_line) ; \
		\
		printf "\n\033[36m-- %s %0.*s %s\033[m", toupper(linter_name), dashed_line_length - length($$1) - length($$2), dashed_line, error_file ; \
		printf "\n\n\033[1mLine %s, Column %s", error_line_number, error_col_number ; \
		printf "\n\n\033[1m%s", error_message ; \
		\
		cmd_read_error_line = sprintf("sed -n %sp %s | sed -e \"s/\t/ /g\"", error_line_number, error_file_path) ; \
		cmd_read_error_line | getline error_line ; close(cmd_read_error_line) ; \
		printf "\n\n\033[33m%s|\033[m %s", error_line_number, error_line ; \
		\
		printf "\n\033[31m\033[1m%*s\033[m", error_col_number + length(error_line_number) + 2, "^" ; \
	} END { printf "\n\033[31m%s errors detected\n", NR	}'


%:
	@$(MAKE) -s $(subst :,/,$@)