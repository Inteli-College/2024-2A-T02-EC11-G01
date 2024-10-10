---
id: Resultados de Testes
title: Resultados de Testes
sidebar_position: 2
---

Durante o desenvolvimento deste projeto, nossa equipe focou em criar uma solução eficiente para a contagem automatizada de árvores, utilizando técnicas avançadas de aprendizado de máquina. A partir da **Sprint 4**, aprendemos a selecionar e propor testes de desempenho que nos permitiram identificar melhorias e detectar falhas em nossa abordagem. O objetivo principal era otimizar a precisão da contagem e a detecção das árvores, bem como avaliar o desempenho computacional dos modelos escolhidos.

Abaixo, apresentamos uma tabulação dos resultados dos modelos testados, juntamente com uma análise de sua performance com base em critérios de qualidade estabelecidos.

## Critérios de Qualidade e Resultados Esperados

Para garantir que os modelos selecionados atendessem às necessidades do projeto, estabelecemos os seguintes critérios de qualidade:

1. **Precisão da Contagem de Árvores:** Medida pela proporção entre o número de árvores detectadas corretamente e o número real de árvores no dataset. Buscamos uma precisão acima de 0.90, para garantir confiabilidade na detecção em larga escala.

2. **Detecção de Árvores (Quantidade):** Verificação da quantidade de árvores detectadas por cada modelo, comparando o desempenho em relação ao mesmo conjunto de dados.

3. **Desempenho Computacional:** Avaliação do consumo de memória dos modelos durante a execução. O objetivo é identificar soluções que ofereçam alto desempenho com consumo de recursos otimizado, sendo preferível modelos com menor impacto no uso de memória, especialmente para ambientes embarcados ou em nuvem.

## Resultados dos Testes

Abaixo, apresentamos os resultados obtidos para dois modelos principais utilizados para a contagem de árvores: **YOLO** e **DeepForest**.

| **Métricas**                                | **YOLO** | **DeepForest** |
|---------------------------------------------|----------|----------------|
| **Precisão da Contagem de Árvores**         | 0.36     | 0.97           |
| **Detecção de Árvores (Quantidade)**        | 34       | 70             |
| **Desempenho Computacional (Uso de Memória)**| 1.2GB    | 527.1MB        |

### Análise dos Resultados

1. **Precisão da Contagem de Árvores:**
   O modelo **DeepForest** apresentou uma precisão significativamente superior à do **YOLO**, atingindo 0.97, próximo do valor ideal esperado. Isso indica que o modelo é muito mais eficaz na detecção correta de árvores em imagens de satélites e drones, o que o torna a melhor escolha para o nosso projeto.

2. **Detecção de Árvores (Quantidade):**
   O **DeepForest** também foi capaz de detectar um número maior de árvores no mesmo conjunto de dados (70 árvores) em comparação ao **YOLO** (34 árvores). Este desempenho superior é crucial para garantir que áreas florestais extensas sejam monitoradas de forma mais detalhada e precisa.

3. **Desempenho Computacional:**
   O **DeepForest** não só apresentou melhores resultados em termos de precisão e quantidade de árvores detectadas, como também teve um desempenho computacional mais eficiente, utilizando **527.1MB de memória** comparado a **1.2GB** consumidos pelo **YOLO**. Isso o torna uma solução mais viável para ser implementada em sistemas com restrições de recursos, como embarcados ou nuvens de baixo custo.

## Conclusão

A partir dos testes realizados, o modelo **DeepForest** mostrou-se superior ao **YOLO** em todos os critérios estabelecidos: precisão, quantidade de detecções e eficiência computacional. Essa solução se mostrou ideal para o projeto de contagem automatizada de árvores, sendo robusta o suficiente para ser utilizada em ambientes reais, especialmente em áreas extensas de monitoramento florestal.

Com base nos resultados, recomendamos fortemente o uso do **DeepForest** como a solução padrão para a contagem de árvores em nosso sistema, devido ao seu desempenho otimizado e capacidade de gerar resultados confiáveis.

## Próximos Passos

- **Otimização do modelo:** Embora o **DeepForest** tenha atendido aos critérios esperados, futuras otimizações podem ser feitas para melhorar ainda mais o desempenho computacional.
- **Implementação em tempo real:** Integrar a solução com uma interface em tempo real para facilitar o monitoramento e análise das áreas florestais.
- **Escalabilidade:** Expandir o modelo para áreas maiores e introduzir mais variáveis para aumentar a acurácia em diferentes tipos de vegetação.

---

