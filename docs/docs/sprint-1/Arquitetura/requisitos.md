---
title: Requisitos funcionais e não funcionais
sidebar_position: 1
slug: /requisitos
---

# Requisitos funcionais e não funcionais

Para garantir que o projeto atenda às necessidades do cliente foram elicitados requisitos funcionais e não funcionais. Os requisitos funcionais definem o que o sistema deve fazer, enquanto os não funcionais estabelecem as metas de desempenho e as qualidades que o sistema deve atingir.

## Requisitos Funcionais



## Requisitos Não Funcionais

Os requisitos não funcionais abordam aspectos como desempenho, segurança, usabilidade e escalabilidade. Eles definem padrões para a operação, manutenção e evolução do sistema, influenciando diretamente a qualidade do produto final, a experiência do usuário e a facilidade de integração com outras tecnologias. Abaixo estão listados os requisitos não funcionais elicitados.

| Categoria | Requisito | Métrica | Meta | 
|-------------|-------------|-------------|-------------|
| Desempenho | O sistema deve garantir baixa latência na captura e processamento das imagens, especialmente durante a identificação das árvores utilizando o modelo. | Tempo médio de processamento por imagem. | Tempo de processamento por imagem após o recebimento na edge layer não deve exceder 2 segundos. | 
| Segurança | O acesso ao sistema deve ser restrito a usuários autorizados, utilizando Firebase para gerenciar a autenticação. A API Gateway deve implementar regras de autorização para garantir que apenas usuários e serviços autorizados possam acessar determinados recursos. | Percentual de tentativas de acesso não autorizado bloqueadas. | 100% de tentativas de acesso não autorizado devem ser bloqueadas.| 
| Escalabilidade | O sistema deve ser capaz de escalar horizontalmente, suportando um aumento no número de drones em operação e no volume de imagens processadas. | Número de imagens processadas por segundo. | Aumento linear do número de imagens processadas conforme novos recursos são adicionados. | 
| Usabilidade | A interface desenvolvida deve ser intuitiva, permitindo que usuários com diferentes níveis de habilidade técnica possam interpretar os resultados. | Tempo médio para completar uma tarefa na interface. | Tempo necessário para completar tarefas comuns deve ser inferior a 2 minutos. | 
| Eficiência | O sistema deve ser otimizado para o uso eficiente de recursos computacionais, especialmente no dispositivo embarcado. Isso inclui a execução do modelo de forma que minimize o consumo de memória e CPU. | Consumo médio de CPU e memória por imagem processada. | Consumo de CPU não deve exceder 70% e de memória 200 MB por imagem processada. | 
