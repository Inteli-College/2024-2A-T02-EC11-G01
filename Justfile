set allow-duplicate-recipes

[private]
default:
  @just --choose

clean:
  @docker compose down -v
  
infra option='':
  @if [ '{{option}}' = 'reset' ]; then \
    docker compose -f ./docker-compose.yml down -v postgres migrations rabbitmq ; \
  fi

  @docker compose -f ./docker-compose.yml up postgres migrations rabbitmq -d --build

[windows]
install-asdf:
  @echo "Please install asdf: https://asdf-vm.com/guide/getting-started.html"

[unix]
[confirm('Are you sure that you want to install asdf (y/N)?')]
install-asdf:
  @git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.14.1 && . "$HOME/.asdf/asdf.sh"
  @echo "Please see the configs in the official website: https://asdf-vm.com/guide/getting-started.html"

install-deps:
  @if ! command -v asdf >/dev/null 2>&1; then \
    just install-asdf ; \
  fi

  @if ! command -v rust >/dev/null 2>&1; then \
    . "$HOME/.asdf/asdf.sh" && asdf plugin add rust && asdf install rust 1.81.0 && asdf global rust 1.81.0 ;  \
  fi

  @if ! command -v node >/dev/null 2>&1; then \
    . "$HOME/.asdf/asdf.sh" && asdf plugin add nodejs && asdf install node 20.17.0 && asdf global node 20.17.0 ;  \
  fi

  @if ! command -v cartesi >/dev/null 2>&1; then \
     npm i -g @cartesi/cli ; \
  fi

  @echo "All deps are updated!"

dapp:
  @docker build -t machine:latest -f ./backend/build/Dockerfile.dapp .
  @cartesi build --from-image machine:latest
  @cartesi run

dev:
  @cd ./backend/cmd/prover/lib && cargo build --release
  @docker compose up --build
