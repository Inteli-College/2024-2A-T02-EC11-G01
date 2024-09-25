---
title: Requisitos funcionais e não funcionais
sidebar_position: 1
slug: /requisitos
---

# Requisitos funcionais e não funcionais

Para garantir que o projeto atenda às necessidades do cliente foram elicitados requisitos funcionais e não funcionais. Os requisitos funcionais definem o que o sistema deve fazer, enquanto os não funcionais estabelecem as metas de desempenho e as qualidades que o sistema deve atingir.

## Requisitos Funcionais

Os requisitos funcionas são funcionalidades que o sistema deve possuir para atingir que o sistema seja capaz de desempenhar suas funções e atenda às necessidades do cliente. Eles descrevem as ações que o sistema deve ser capaz de realizar, como a captura de imagens, processamento de dados e exibição de resultados. Abaixo estão listados os requisitos funcionais elicitados.

| Categoria                  | Requisito                                                                                      |
|----------------------------|------------------------------------------------------------------------------------------------|
| Captura de Imagens          | O sistema deve permitir que o drone capture imagens aéreas das áreas de floresta. |
| Processamento de Imagens    | O sistema deve processar as imagens capturadas localmente no drone para identificar e contar o número de árvores. |
| Processamento de Imagens    | O sistema deve ser capaz de tomar a decisão de tirar outra foto baseado no resultado do modelo embarcado no drone |
| Transmissão de Dados        | O sistema deve transmitir os resultados da contagem de árvores e a imagem obtida, sempre que houver conectividade. |
| Armazenamento de Dados      | O sistema deve armazenar de forma segura todos os dados capturados e processados, incluindo imagens e resultados da contagem. |
| Processamento de Imagens    | o resultado inicial do modelo embarcado deve ser validado por um modelo na nuvem mais preciso |
| Segurança                   | O sistema deve permitir que apenas usuários autenticados acessem os dados e controlem os drones, utilizando o Firebase para autenticação. |
| Usabilidade                 | O sistema deve fornecer uma interface intuitiva para que os usuários possam visualizar os dados de contagem de árvores e os resultados processados. |

## Requisitos Não Funcionais

Os requisitos não funcionais abordam aspectos como desempenho, segurança, usabilidade e escalabilidade. Eles definem padrões para a operação, manutenção e evolução do sistema, influenciando diretamente a qualidade do produto final, a experiência do usuário e a facilidade de integração com outras tecnologias. Abaixo estão listados os requisitos não funcionais elicitados.

| Categoria | Requisito | Métrica | Meta | 
|-------------|-------------|-------------|-------------|
| Desempenho | O sistema deve garantir baixa latência na captura e processamento das imagens, especialmente durante a identificação das árvores utilizando o modelo. | Tempo médio de processamento por imagem. | Tempo de processamento por imagem após o recebimento na edge layer não deve exceder 2 segundos. | 
| Segurança | O acesso ao sistema deve ser restrito a usuários autorizados, utilizando Firebase para gerenciar a autenticação. A API Gateway deve implementar regras de autorização para garantir que apenas usuários e serviços autorizados possam acessar determinados recursos. | Percentual de tentativas de acesso não autorizado bloqueadas. | 100% de tentativas de acesso não autorizado devem ser bloqueadas.| 
| Escalabilidade | O sistema deve ser capaz de escalar horizontalmente, suportando um aumento no número de drones em operação e no volume de imagens processadas. | Número de imagens processadas por segundo. | Aumento linear do número de imagens processadas conforme novos recursos são adicionados. | 
| Usabilidade | A interface desenvolvida deve ser intuitiva, permitindo que usuários com diferentes níveis de habilidade técnica possam interpretar os resultados. | Tempo médio para completar uma tarefa na interface. | Tempo necessário para completar tarefas comuns deve ser inferior a 2 minutos. | 
| Eficiência | O sistema deve ser otimizado para o uso eficiente de recursos computacionais, especialmente no dispositivo embarcado. Isso inclui a execução do modelo de forma que minimize o consumo de memória e CPU. | Consumo médio de CPU e memória por imagem processada. | Consumo de CPU não deve exceder 70% e de memória 200 MB por imagem processada. | 
