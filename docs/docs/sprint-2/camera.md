---
id: Embarcado
title: Embarcado
sidebar_position: 3
---

# Embarcado

## Utilizando a câmera no sistema embarcado

Pensando que a captura de imagens pode ocorrer através de um drone, buscamos uma solução que envolve o uso de um dispositivo embarcado equipado com uma câmera. Inicialmente, utilizamos um Raspberry Pi 5 junto com uma webcam USB para coletar as imagens.

## Ferramentas utilizadas

A device layer consiste no dispositivo embarcado, neste caso o drone, responsável pela captura das imagens para processamento posterior. Tendo em mente que essa etapa precisa ser executada com precisão e rapidez, optamos pela linguagem Rust. Rust é uma linguagem compilada que pode ser executada em praticamente qualquer ambiente, além de ser performática e segura.

Também utilizamos a biblioteca v4l (Video for Linux), que fornece suporte para captura de vídeo em tempo real em sistemas Linux, garantindo eficiência no processo de captura.

## Estrutura e execução

Os arquivos relacionados ao sistema embarcado estão organizados na pasta embedded, com a seguinte estrutura:

```bash
/embedded
├── Cargo.lock
├── Cargo.toml
├── images
├── README.md
└── src
    └── main.rs
```

    - Cargo.lock e Cargo.toml: Arquivos de configuração do projeto Rust.
    - images: Diretório onde as imagens capturadas são armazenadas.
    - src/main.rs: Código-fonte principal do projeto, onde ocorre a captura de imagens e a integração com a câmera.

1. Navegar até o diretório do projeto
```bash
cd /2024-2A-T02-EC11-G01/embedded
```

1. Compilar o projeto

Para compilar o projeto, você pode utilizar o comando:

```bash
cargo build
```

Esse comando gera o executável no diretório target/debug. Se desejar uma versão otimizada para produção, use:

```bash
cargo build --release
```

O executável será gerado no diretório target/release.

3. Executar o projeto

Após a compilação, você pode rodar o projeto diretamente com:

```bash
cargo run
```

Esse comando compila (se necessário) e executa o projeto em um único passo.

4. Executar diretamente o binário

Se preferir rodar o executável compilado diretamente, primeiro compile o projeto usando cargo build ou cargo build --release. Depois, execute o binário gerado:

    Para a versão de debug:

```bash
./target/debug/embedded
```

Este projeto foi desenvolvido visando fornecer uma solução eficiente para captura de imagens em tempo real em sistemas embarcados, permitindo futuras implementações em drones e outros dispositivos semelhantes.