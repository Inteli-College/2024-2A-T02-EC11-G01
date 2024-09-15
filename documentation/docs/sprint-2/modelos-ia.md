---
label: "Modelos de IA"
sidebar_position: 2
---

Neste documento, apresentamos uma análise comparativa dos modelos desenvolvidos durante a Sprint 2, sendo dois modelos de deep learning utilizados para a contagem automática de árvores em imagens de satélite e aéreas: YOLO (You Only Look Once) e DeepForest. A contagem precisa de árvores é essencial para diversas aplicações, como monitoramento ambiental, gestão florestal, e estimativas de biomassa. Cada modelo oferece vantagens distintas e apresenta diferentes desafios em termos de precisão, velocidade, e aplicabilidade. A seguir, detalharemos as características técnicas de ambos os modelos para orientar na escolha do mais adequado para o nosso cenário de uso.

## Modelo YOLO (You Only Look Once)
YOLO é uma das arquiteturas de deep learning mais utilizadas para tarefas de detecção de objetos em tempo real. Desenvolvido originalmente para detectar uma ampla variedade de objetos em imagens, YOLO utiliza uma abordagem de rede neural convolucional (CNN) unificada, onde a detecção é realizada em uma única passagem pela rede. Isso confere ao YOLO uma vantagem significativa em termos de velocidade de inferência.

### Características Técnicas:

- **Arquitetura Unificada**: YOLO divide a imagem de entrada em uma grade e prevê bounding boxes e as classes dos objetos para cada célula da grade em uma única etapa, o que reduz a complexidade e o tempo de processamento.
- **Processamento em Tempo Real**: Graças à sua eficiência, YOLO pode operar em tempo real, processando dezenas de frames por segundo em GPUs de alta performance, o que o torna ideal para aplicações onde a latência é crítica.
- **Versatilidade**: Originalmente projetado para detecção de múltiplas classes de objetos, YOLO pode ser adaptado para a contagem de árvores, embora isso possa exigir a customização do modelo e ajustes no treinamento com datasets específicos.

### Desafios:

- **Precisão em Objetos Pequenos**: Embora seja rápido, YOLO pode ter dificuldades em detectar objetos muito pequenos, como árvores jovens ou densamente agrupadas, especialmente em imagens de alta resolução com muitos detalhes.
- **Falsos Positivos**: Devido à sua abordagem de detecção única, YOLO pode gerar falsos positivos, o que pode comprometer a precisão total em cenários de contagem precisa de árvores.
- **Dependência de Hardware**: Para alcançar a máxima performance, o modelo requer hardware potente, como GPUs modernas, o que pode limitar sua aplicabilidade em ambientes com restrições de recursos computacionais.

## Modelo DeepForest
DeepForest é um modelo de deep learning especializado, desenvolvido especificamente para a detecção e contagem de árvores em imagens de satélite e aéreas. Ao contrário de modelos mais generalistas como o YOLO, DeepForest foi projetado com foco na ecologia, levando em consideração as características únicas das florestas e vegetação.

### Características Técnicas:

- **Especialização em Ecologia**: DeepForest foi treinado utilizando grandes datasets de imagens aéreas de florestas, com o objetivo de identificar árvores individuais e suas características específicas, como altura e diâmetro da copa.
- **Detecção em Múltiplas Escalas**: O modelo é capaz de operar em diferentes escalas de resolução, o que é crucial para detectar tanto árvores individuais quanto grandes áreas de floresta com precisão.
- **Arquitetura Baseada em CNNs**: DeepForest utiliza redes neurais convolucionais adaptadas para maximizar a precisão em tarefas de detecção vegetativa, diferenciando-se de modelos generalistas pela sua capacidade de entender padrões complexos de vegetação.

### Desafios:

- **Inferência Mais Lenta**: Devido à sua especialização e processamento em múltiplas escalas, DeepForest pode ser mais lento em comparação a modelos como YOLO, especialmente quando aplicado a grandes áreas ou imagens de alta resolução.
- **Aplicabilidade Limitada**: Enquanto YOLO é versátil e pode ser usado para diversas tarefas de detecção de objetos, DeepForest é altamente especializado e não é adequado para outras aplicações fora do contexto ecológico.
- **Complexidade no Treinamento**: Treinar e adaptar o DeepForest para novos tipos de vegetação ou ambientes florestais pode exigir datasets extensivos e especializados, além de ajustes na arquitetura do modelo.

## Conclusão
A análise comparativa entre YOLO e DeepForest levou em consideração a velocidade, versatilidade e precisão exigidas pelo projeto de contagem de árvores. Apesar da velocidade e flexibilidade oferecidas pelo YOLO, optamos por utilizar o DeepForest como o modelo principal. Essa escolha foi baseada em uma avaliação de imagens, como é possível observar abaixo, onde o DeepForest demonstrou maior assertividade na detecção e contagem de árvores, especialmente em cenários complexos. 

![Modelo Deep Forest](../../../static/img/imagem-modelo-deep-forest.png)
