set allow-duplicate-recipes

[private]
default:
  @just --list

[private]
verify-asdf:
  @if ! command -v asdf >/dev/null 2>&1; then \
    just install-asdf ; \
  fi

  @if [ ! "${ASDF_DIR:-}" = '' ]; then \
    . "$ASDF_DIR/asdf.sh" ; \
  else \
    echo "Unable to find asdf instalation!" ; exit 1 ; \
  fi

[private, windows]
install-asdf:
  @echo "Please install asdf: https://asdf-vm.com/guide/getting-started.html" ; exit 1

[private, unix]
[confirm('Are you sure that you want to install asdf (y/N)?')]
install-asdf:
  @git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.14.1 && . "$HOME/.asdf/asdf.sh"
  @echo "Please see the configs in the official website: https://asdf-vm.com/guide/getting-started.html"

[group('dev')]
clean:
  @docker compose down -v 
  
infra option='':
  @if [ '{{option}}' = 'reset' ]; then \
    docker compose -f ./docker-compose.yml down -v postgres migrations rabbitmq ; \
  fi

  @docker compose -f ./docker-compose.yml up postgres migrations rabbitmq -d --build

[group('dev')]
install-deps:
  @if ! command -v rustc >/dev/null 2>&1; then \
    echo "Rust not found, installing..." ; \
    just verify-asdf && asdf plugin add rust && asdf install rust 1.81.0 && asdf global rust 1.81.0 ;  \
  fi

  @if [ ! $(rustc --version | cut -d' ' -f2) = "1.81.0" ]; then \
    echo "Rust found in the wrong version, updating..." ; \
    just verify-asdf && asdf plugin add rust && asdf install rust 1.81.0 && asdf global rust 1.81.0; \
  fi

  @if ! command -v node >/dev/null 2>&1; then \
    echo "Node not found, installing..." ; \
    just verify-asdf && asdf plugin add nodejs && asdf install node 20.17.0 && asdf global node 20.17.0 ;  \
  fi

  @if ! command -v cartesi >/dev/null 2>&1; then \
    echo "Cartesi cli not found, installing..." ; \
    npm i -g @cartesi/cli ; \
  fi

  @echo "All deps are updated!"

[group('dev')]
dapp:
  @docker build -t machine:latest -f ./backend/build/Dockerfile.dapp .
  @cartesi build --from-image machine:latest
  @cartesi run --verbose

[group('dev')]
dev: install-deps
  @cd ./backend/cmd/prover/lib && cargo build --release
  @docker compose up --build

[group('prod')]
prod:
  @cd ./backend/cmd/prover/lib && cargo build --release
  @docker compose up --build --profile prod
