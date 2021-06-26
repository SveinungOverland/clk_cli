CGO_ENABLED=1

all: install autocomplete prompt source

install:
	go install clk

autocomplete:
	clk completion zsh > /usr/local/share/zsh/site-functions/_clk

prompt:
	cp ./prompt_clk /usr/local/share/zsh/functions/prompt_clk
	cp ./prompt_clkcontext /usr/local/share/zsh/functions/prompt_clkcontext

source:
	source ~/.zshrc