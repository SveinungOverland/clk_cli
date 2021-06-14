
all: install autocomplete

install:
	go install clk

autocomplete:
	clk completion zsh > /usr/local/share/zsh/site-functions/_clk
