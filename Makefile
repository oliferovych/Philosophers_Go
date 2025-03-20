NAME        = philo
SRC_DIR     = ./src

all: $(NAME)

$(NAME):
	go mod init Philosophers_Go
	@cd $(SRC_DIR) && go build -o ../$(NAME)

clean:
	@echo "Cleaning up build artifacts"
	@rm -f go.mod

fclean: clean
	@rm -f $(NAME)

re: fclean all

.PHONY: all clean fclean re