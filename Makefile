WAY = ./cmd/todo
all: build list

build: clean
	@go build $(WAY)

list: build
	@./todo -list

add: build
	@printf "\033[94mEnter your task: \033[0m"
	@./todo -add $(read)

complete: build
	@printf "\033[92mEnter the completed task number: \033[0m"
	@read task_number; ./todo -complete $$task_number

del: build
	@printf "\033[91mEnter the number to delete: \033[0m"
	@read task_number; ./todo -del $$task_number

clean:
	@rm -rf todo