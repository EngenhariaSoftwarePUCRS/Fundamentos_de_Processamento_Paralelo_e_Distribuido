## Módulo 01 - Introdução à disciplina e aos conceitos de concorrência - 02, 07 e 09 de Agosto de 2023


## 02/08 - Panorama Geral da Disciplina

Apresentação de Materiais, Metodologia e Avaliação

## [Sistemas Paralelos](#sistemas-paralelos) vs [Sistemas Distribuídos](#sistemas-paralelos)

### Sistemas Distribuídos:

- Coleção de computadores independentes que se apresenta a seus usuários como um sistema único e coerente.
- Troca de mensagem, rede com atrasos, não dividem memórias
- Posssíveis falhas
- Objetivo: Permitir colaboração entre processos para prover um serviço

### Sistemas Paralelos:

- Tipicamente em uma máquina multi-core ou então em um cluster de computadores
- Carga controlada; problemas muito grandes podem ser divididos em partes menores em um único computador
- Se há falhas, repete-se a computação
- Extrair o máximo desempenho da arquitetura disponível

## Sistemas Concorrentes

- Não é competição por recursos
- concur: happen or occur at the same time
- coexistência, não competição

### Conceitos Básicos

- ### Processos
  - um conjunto de processos sequenciais
  - cadaum representa uma sequência ou linha de execução
  - coexistindo, ou seja, executando concorrentemente
  - conjunto pode ser dinâmico

- ### Sincroniação e Comunicação
  - coordenação temporal e passagem de dados
  - memória compartilhada
  - passagem de mensagens

### Sistemas Concorrentes

- ### Locais
  - Editor de texto
    - Um processo responsável pela interação com usuário
    - Outro processo responsável pela formatação do texto
  - Jogos com vários personagens
    - Um processo representa o usuário no jogo; interage com o controle
    - Outros processos representam os NPC's; interagem com o usuário e entre si

#### Avaliação

| Assunto | Tempo | Avaliação
| ------- | ----- | --------- |
| Concorrência<br><ul></ul> | 60~65% | T1 = Trabalho com Canais<br>T2 = Trabalhos Memória Compartilhada
| Tipos de sistemas | 40~35% | T3 = Trabalho Sistemas Distribuidos<br>T4 = Trabalho Sistemas Paralelos

G1 = $P + ((T1+T2) \times 0,6 +(T3+T4) \times 0,4) \over 2$

#### Frase da Aula

<blockquote>
    Aquilo que escuto, eu esqueço.<br>
    Aquilo que vejo, eu lembro.<br>
    Aquilo que faço, eu aprendo.
</blockquote>


## 07/08 - Conceitos Básicos

### Definições

**Um programa concorrente** consiste de um conjunto finito de processos *sequenciais*.

**Um processo** contém um conjunto finito de comandos atômicos.

Cada processo tem seu **control pointer** que indica o próximo comando que pode ser executado pelo processo.

Uma **computação** descreve uma execução possível do programa concorrente.

Uma computação é obtida por um **entrelaçamento (interleaving)** arbitrário dos comandos atômicos dos processos do programa.

O conjunto de todas as computações do programa é chamado de **comportamento**.

**Semântica de *interleaving*** é uma forma de representar a execução de um sistema concorrente.
É a mais utilizada na literatura, e adotada aqui.
Há outras semânticas para representar concorrência, como: true concurrency; semântica de eventos; traces de Mazukievicz.

*Interleaving* = a cada momento, um processo é escolhido para dar um passo

**Um estado de um programa concorrente** é uma tupla com:
  - um rótulo de cada processo
    - (*program counter* para o processo)
  - o valor de cada variável local ou global

Sejam s1 e s2 estados de um programa concorrente, existe uma **transição de s1 para s2** se executando um comando em s1 leva a s2.
O comando executado é um dos apontados pelos contadores de programa dos processos no estado s1.

Partindo-se de um estado inicial, **um diagrama de estados é criado indutivamente** pela aplicação das transições possíveis.
Este conjunto de estados é dito **conjunto de estados alcançáveis**.

**Uma computação** é **um caminho** dirigido no diagrama de estados, iniciando no estado inicial.

**Ciclos** no diagrama representam a possibilidade de computações infinitas.

Estados sem arestas de saída representam situações de **bloqueio** ou **terminação**.

**Justiça (Fairness)**: não faz sentido supor a possibilidade de os comandos de um processo *nunca* serem selecionados para execução.
Uma computação é justa (*weak fairness*) se um comando que está continuamente habilitado nela acaba por ser executado em um momento.

## 08/09 - Exemplos de Concorrência

Revisão com alguns exercícios
