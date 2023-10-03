# Grupo A

Dining Philosophers.

Solução com filósofo canhoto.

Por que esta solução não tem inanição?  (ou starvation)


## Sobre o problema

- Cada filósofo tem um garfo à sua esquerda e um à sua direita.
- Um filósofo pode comer se ele tiver ambos os garfos.
- Um filósofo pode pegar um garfo se ele não estiver sendo usado por outro filósofo.


## Solução

Uma possível solução que evite deadlock/starvation para o problema dos filósofos é a de haver filósofos canhotos. Nessa solução, os filósofos que são canhotos pegam primeiro o garfo da esquerda e depois o da direita, enquanto os filósofos que são destros pegam primeiro o garfo da direita e depois o da esquerda. Dessa forma, os filósofos canhotos e destros não competem pelos mesmos garfos, evitando assim o deadlock.


## Implementação

Para provar que esta solução é possível, optou-se por fazer uma prova por contradição:

- Suponhamos que há um deadlock.
- Sabendo que há, pelo menos, um filósofo destro e um filósofo canhoto.
- Se há um deadlock, então todos os filósofos estão com um e apenas um garfo.
- Como existe, pelo menos, uma dupla de filósofos que "compartilham" diretamente o garfo 'inicial' e 'final' (canhoto e destro)
  - Não é possível que todos os filósofos estejam com um e apenas um garfo.


## Integrantes

- Arthur Antunes de Souza Both
- Carolina Michel Ferreira
- Felipe Freitas Silva
- Gabriel Moraes Ferreira
- Maria Eduarda Wendel Maia
- Matheus Campos Caçabuena
