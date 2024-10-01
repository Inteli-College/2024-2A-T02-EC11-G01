---
id: Análise Financeira
title: Análise FInanceira
sidebar_position: 2
---

# Análise Financeira

Análise financeira é o processo de avaliação da viabilidade, estabilidade e lucratividade de um projeto, empresa ou qualquer outra iniciativa que envolva o uso de recursos financeiros. Ela envolve o estudo detalhado das demonstrações financeiras (como balanços patrimoniais, demonstrações de resultados, fluxos de caixa, entre outros) para entender a saúde financeira de uma entidade ou projeto, identificar oportunidades de melhoria, prever desempenho futuro e tomar decisões estratégicas informadas.

No contexto do projeto Artemis, visamos entender os principais pontos de custo. Para isso, realizamos uma análise simples dos custos envolvidos no desenvolvimento do projeto, bem como uma projeção de receitas e despesas futuras. A seguir, apresentamos os principais resultados dessa análise.

Segundo o artigo da DroneShow robotics[^1], o tamanho de mercado global para contagem de árvores pode chegar à $227 bilhões. Isso pode ser um mercado promissor para o projeto Artemis. Pensando nisso, realizamos uma análise financeira simplificada para entender os custos envolvidos no desenvolvimento e manutenção do projeto, de forma a entender de forma mais clara sua viabilidade financeira em questão de custos. A partir dessa análise, é possível identificar oportunidades de redução de custos, otimização de recursos e melhoria da eficiência operacional de forma a entender qual deveria ser a margem de receita da operação para tornar o negócio viável.

## Custos

Pensando em uma visão simplificada, dividimos os custos do projeto em duas categorias principais: custos iniciais e custos recorrentes.

- Infraestrutura: custos iniciais com aquisição de servidores, licenças de software, equipamentos, entre outros.
- Manutenção de equipamentos: Drones, sensores, 
- Pessoal: custos com salários, benefícios, treinamentos, entre outros.

| **Categoria**                | **Produto/Serviço**                        | **Descrição**                                 | **Custo Estimado (R$)** |
|------------------------------|--------------------------------------------|-----------------------------------------------|-------------------------|
| **Infraestrutura**            | [Servidores EC2 (AWS)](https://aws.amazon.com/pt/ec2/pricing/on-demand/)                       | Deploy de aplicação                           | 1.200,00/mês            |
|                              | [Armazenamento S3 (AWS)](https://aws.amazon.com/pt/s3/pricing/)                     | Armazenamento de dados                        | 0,10/GB/mês             |
|                              | [Lambda (AWS)](https://aws.amazon.com/pt/lambda/pricing/)                               | Execução de funções sem servidor              | 0,20/milhão execuções    |
|                              | [Raspberry Pi 5](https://www.robocore.net/placa-raspberry-pi/raspberry-pi-5-8gb)                             | Computador embarcado                          | 975,90                  |
|                              | [Cartão de Memória 64GB](https://www.kingstonstore.com.br/products/sdcs2-64gb-cartao-de-memoria-microsd-de-64gb-canvas-select-plus-leitura-100mb-s-classe-10-com-adaptador-sd)                     | Armazenamento local                           | 60,00                   |
|                              | [Câmera USB](https://www.amazon.com.br/Webcam-Intelbras-CAM-720p-Preto/dp/B09Y736F6V/ref=sr_1_6?dib=eyJ2IjoiMSJ9.Nxfg_ir7qq-8C4QPnqUS6gWEoazUrBRTm79kI3vRUqMlvfSbV3fiQ2mg46Bc3Op-2E-ysMIcEM1hABqnlasX2Nzp18vgTpEm_1siSxGhYZXWyThXS4NVNXt84zrRzkZY2ImI9rHIArdaiGG-eL3b5lkbANUtzmhjkX5qdg2PggQZaLwi7xpgELJpcA4N4m6apBwiUY6Swi-1qRaoOnbbwiGnMN58FWuLzmcWOc9bdwBqjT1zMC3m7n_tzvUhzyIyFBuBlz3YPAcdSn_6RNexgi1k4OijBpOrIMaDuoXndug.9fMbkLUcbQzHkBK-ORGXcLngd9uZ7ygZ9033sFReT_c&dib_tag=se&keywords=camera+usb&qid=1727650796&sr=8-6&ufe=app_do%3Aamzn1.fos.6121c6c4-c969-43ae-92f7-cc248fc6181d)                                 | Periférico para captura de imagem             | 150,00                  |
|                              | Licenças de Software                       | Ferramentas e licenças de software            | 2.000,00/ano            |
| **Manutenção de Equipamentos**| Drones                                     | Manutenção e reparos                          | 500,00/mês              |
|                              | Sensores                                   | Troca e manutenção                            | 300,00/mês              |
| **Pessoal**                  | [Backend Developer (Go)](https://www.glassdoor.com.br/Sal%C3%A1rios/golang-developer-sal%C3%A1rio-SRCH_KO0,16.htm)                     | Salário + Benefícios                          | 10.000,00/mês           |
|                              | Desenvolvedor Embarcado (Rust)            | Salário + Benefícios                          | 10.000,00/mês           |
|                              | [Frontend Developer + Modelo (Python)](https://www.glassdoor.com.br/Sal%C3%A1rios/desenvolvedor-python-sal%C3%A1rio-SRCH_KO0,20.htm)       | Salário + Benefícios                          | 10.000,00/mês           |
|                              | Treinamentos e Cursos                      | Capacitação da equipe                         | 2.000,00/ano            |
| **Outros Custos**             | Energia Elétrica                           | Consumo de energia para infraestrutura        | 400,00/mês              |
|                              | Internet                                   | Conexão para servidores e equipe              | 500,00/mês              |
|                              | Aluguel de Escritório (se aplicável)        | Local físico para a equipe                    | 2.000,00/mês            |
|                              | Backup e Segurança de Dados                | Serviços para segurança de dados              | 300,00/mês              |
|                              | Marketing e Publicidade                    | Divulgação e atração de clientes              | 1.000,00/mês            |
|                              | Consultoria Especializada                  | Apoio técnico ou de gestão                    | 2.500,00/pontual        |

| **Total Mensal**                                                         | **R$ 49.710,10**                           |


## Profitabilidade e Receitas

Entendemos que a viabilidade financeira de um projeto depende não apenas dos custos envolvidos, mas também das receitas geradas. Nesse contexto, a profitabilidade e receita depende muito do modelo de negócio adotado e cabe a cada empresa ou projeto definir a melhor estratégia para gerar receitas. Recomendamos fortemente a realização de um estudo de viabilidade econômica e financeira mais detalhado para entender melhor a relação entre custos e receitas e identificar oportunidades de melhoria.

## Conclusão

A análise financeira é uma ferramenta essencial para entender a viabilidade de um projeto, empresa ou iniciativa. Ela permite avaliar a saúde financeira de uma entidade, identificar oportunidades de melhoria, prever desempenho futuro e tomar decisões estratégicas informadas. No contexto do projeto Artemis, a análise financeira nos permitiu entender os principais custos envolvidos no desenvolvimento e manutenção do projeto, bem como projetar receitas e despesas futuras. A partir desses resultados, é possível identificar oportunidades de redução de custos, otimização de recursos e melhoria da eficiência operacional, de forma a tornar o projeto mais viável financeiramente.

[^1]: [Drone Show robotics](https://droneshowla.com/contagem-automatica-de-arvores-em-imagens-obtidas-com-drones/)