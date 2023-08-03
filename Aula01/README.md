## Aula 01 - 02/08/2023


Panorama Geral da Disciplina

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
