# Respostas da Lista da Unidade 1 de Algoritmos de Estrutura de dados 1

## Gerais

### Questão 1

Uma TAD seria o nome da estrutura e um conjunto de operacoes pre definidas mas sem implementacao de de nada, uma ED, seria efetivamente a implementacao de uma TAD, com suas peculiaridades e seguindo as operacoes que foram definidas na TAD

### Questão 2

A)

### Questão 3

C)

### Questão 4

A)

## Listas

### Questão 1

implementacao no arquivos:

- array_list.go (x)
- linked_list.go (x) revisar
- doubly_linked_list.go

### Questão 2

| Operação                         | ArrayList |             | LinkedList |             | DoublyLinkedList |             |
| -------------------------------- | --------- | ----------- | ---------- | ----------- | ---------------- | ----------- |
|                                  | Pior Caso | Melhor Caso | Pior Caso  | Melhor Caso | Pior Caso        | Melhor Caso |
| Add(value int)                   | O(N)      | Omega(1)    | O(N)       | Omega(N)    | O(1)             | O(1)        |
| AddOnIndex(value int, index int) | O(N)      | Omega(1)    | O(N)       | Omega(1)    | O(N)             | O(1)        |
| RemoveOnIndex(index int)         | O(N)      | Omega(1)    | O(N)       | Omega(1)    | O(N)             | O(1)        |
| Get(index int)                   | O(1)      | Omega(1)    | O(N)       | Omega(1)    | O(N)             | O(1)        |
| Set(value int, index int)        | O(1)      | Omega(1)    | O(N)       | Omega(1)    | O(N)             | O(1)        |
| Size()                           | O(1)      | Omega(1)    | O(1)       | Omega(1)    | O(1)             | O(1)        |

### Questão 3

Vantagem - Acesso O(1) em qualquer lugar da lista
Desvantagem - Na maioria do tempo tera espaço alocado sem uso

### Questão 4

Vantagem - Acesso O(1) nas extremidades
Desvantagem - Na lista duplamente ligada tempos mais espaço alocado, pois tempos que em um mesmo elemento sera guardado 2 ponteiros, o do anterior e o do proximo

### Questão 5

implementacao no arquivos:

- reverse_array_list.go

### Questão 6

implementacao no arquivos:

- reverse_linked_list.go

### Questão 7

implementacao no arquivos:

- reverse_doubly_linked_list.go

### Questão 8

Pois em cada operacao que alterasse o final da lista, teria que se atualizar sempre a tail, fazendo assim as operacoes ficarem mais complexas

## Pilhas

### Questão 1

implementacao no arquivos:

- array_stack.go
- linked_stack.go

### Questão 2

| Operação        | ArrayStack |             | LinkedStack |             |
| --------------- | ---------- | ----------- | ----------- | ----------- |
|                 | Pior Caso  | Melhor Caso | Pior Caso   | Melhor Caso |
| Push(value int) | O(N)       | Omega(1)    | O(N)        | Omega(1)    |
| Pop()           | O(1)       | Omega(1)    | O(N)        | Omega(1)    |
| Peek(index int) | O(1)       | Omega(1)    | O(N)        | Omega(1)    |
| Size()          | O(1)       | Omega(1)    | O(1)        | Omega(1)    |

### Questão 3

implementacao no arquivos:

- detect_parenteses.go

### Questão 4

C)

## Filas

### Questão 1

- Fila de processos gerenciados pela CPU
- Enfileramento de dados usando RabbitMQ

### Questão 2

implementacao no arquivos:

- array_queue.go
- linked_queue.go

### Questão 3

| Operação              | ArrayQueue |             | LinkedQueue |             |
| --------------------- | ---------- | ----------- | ----------- | ----------- |
|                       | Pior Caso  | Melhor Caso | Pior Caso   | Melhor Caso |
| Enqueue(value int)    | O(1)       | Omega(1)    | O(1)        | Omega(1)    |
| Dequeue() (int,error) | O(1)       | Omega(1)    | O(1)        | Omega(1)    |
| Front() (int,error)   | O(1)       | Omega(1)    | O(1)        | Omega(1)    |
| Size()                | O(1)       | Omega(1)    | O(1)        | Omega(1)    |

### Questão 4

implementacao no arquivos:

- size_func_queue.go

## Deque

### Questão 1

- Pode ser usado para cache, dados mais antigos ficam em uma ponta, e dados mais novos em outra ponta
- Janela dislizante

### Questão 2

implementacao no arquivos:

- array_deque.go
- linked_deque.go

### Questão 3

| Operação                   | ArrayDeque |             | LinkedDeque |             |
| -------------------------- | ---------- | ----------- | ----------- | ----------- |
|                            | Pior Caso  | Melhor Caso | Pior Caso   | Melhor Caso |
| EnqueueFront(value int)    | O(N)       | Omega(1)    | O(1)        | Omega(1)    |
| EnqueueRear(value int)     | O(N)       | Omega(1)    | O(1)        | Omega(1)    |
| DequeueFront() (int,error) | O(1)       | Omega(1)    | O(1)        | Omega(1)    |
| DequeueRear() (int,error)  | O(1)       | Omega(1)    | O(1)        | Omega(1)    |
| Front() (int,error)        | O(1)       | Omega(1)    | O(1)        | Omega(1)    |
| Rear() (int,error)         | O(1)       | Omega(1)    | O(1)        | Omega(1)    |
| IsEmpty()                  | O(1)       | Omega(1)    | O(1)        | Omega(1)    |
| Size()                     | O(1)       | Omega(1)    | O(1)        | Omega(1)    |

## Algoritmo de Busca

### Questão 1
