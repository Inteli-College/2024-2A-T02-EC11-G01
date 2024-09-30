# Documentação Teórica - Frontend de Monitoramento de Árvores e Carbono Capturado
## Visão Geral do Projeto

Este frontend é projetado para monitorar e exibir informações sobre a quantidade de árvores ao longo do tempo e a captura de carbono associada a essas árvores. A interface é baseada em Streamlit, uma ferramenta interativa voltada para a visualização de dados, oferecendo uma experiência acessível e intuitiva aos usuários interessados em gestão florestal e créditos de carbono.

## Funcionalidades
### Monitoramento Temporal da Quantidade de Árvores:
A interface permite visualizar a evolução do número de árvores em uma área monitorada ao longo do tempo. O usuário pode navegar por diferentes períodos, observando o crescimento, a manutenção ou a diminuição da cobertura florestal.

Gráficos e visualizações dinâmicas ajudam a representar tendências, permitindo uma análise comparativa entre diferentes períodos (semanas, meses ou anos).

### Visualização das Imagens Processadas:
O modelo de contagem de árvores gera imagens processadas, que são exibidas diretamente na interface. Essas imagens são derivadas de dados de satélite e ajustadas pelo modelo de contagem automática.

O usuário pode alternar entre as imagens originais e as processadas, observando onde o modelo identificou as árvores e como o ajuste de contagem foi realizado.

A interface oferece a opção de visualizar as imagens de diferentes períodos, facilitando a comparação visual das mudanças florestais.

### Cálculo de Carbono Capturado:
A quantidade de carbono capturado é estimada com base no número de árvores detectadas e na sua biomassa, utilizando dados médios de captura de carbono por árvore. A interface apresenta essas informações de forma simples e direta, com gráficos que mostram a quantidade total de carbono capturado ao longo do tempo.

O frontend também apresenta resumos numéricos, destacando as estimativas de captura de carbono por hectare, bem como projeções futuras com base no crescimento contínuo da floresta.

## Interatividade e Navegação:
A interface permite que os usuários ajustem os parâmetros de visualização, como a escolha de períodos específicos, áreas geográficas de interesse e diferentes categorias de árvores, se disponíveis.

Ferramentas de filtro permitem explorar dados específicos para uma análise mais detalhada. O usuário pode focar em uma região específica ou analisar florestas com diferentes densidades de árvores.

## Arquitetura da Interface
### Tela Principal:
Contém um painel de navegação que permite ao usuário escolher entre visualizar as imagens processadas, acompanhar o gráfico da quantidade de árvores ao longo do tempo ou observar o histórico de captura de carbono.

Uma área de visualização central exibe gráficos dinâmicos ou imagens conforme a escolha do usuário.
