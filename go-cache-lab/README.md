# Go Cache Lab

Este diretório faz parte do projeto **go-cache-lab**, um laboratório prático em Golang que demonstra **todas as principais estratégias de cache usadas em sistemas reais**, com foco em **concorrência, performance, consistência e resiliência**.

O objetivo não é apenas mostrar *como* implementar cache, mas **quando usar cada abordagem, quais problemas ela resolve e quais riscos introduz**.

---

## 📌 Princípios Gerais

Antes de entrar nos módulos, alguns princípios guiam todo o projeto:

- Cache **não é thread-safe por padrão**
- Mutex, RWMutex e singleflight são tão importantes quanto TTL
- Cache resolve latência, mas **cria problemas de consistência**
- Thundering herd é um problema inevitável sem coordenação
- Cache bom é **observável, resiliente e testável**

Cada pasta abaixo explora um desses pontos de forma isolada e prática.

---

## 🗂️ Estrutura do Projeto

```text
go-cache-lab/
├── l1-memory/
│   ├── map-mutex/
|   ├── map-no-lock/
│   ├── map-rwmutex/
│   ├── sync-map/
│   ├── ttl/
│   └── eviction/
│
├── l2-redis/
│   ├── basic/
│   ├── mutex-key/
│   └── pubsub-invalidation/
│
├── multilayer/
│   └── l1-l2-singleflight/
```

---

## 🧠 l1-memory — Cache em Memória (In-Process)

Cache local dentro do processo Go (L1). Extremamente rápido, mas **limitado ao pod/processo** e altamente sensível a problemas de concorrência.

### map-mutex/
- Implementação básica usando `map + sync.Mutex`
- Demonstra:
  - Proteção mínima para concorrência
  - Gargalos com lock global
- Quando usar:
  - Baixa concorrência
  - Código simples

### map sem lock (exemplo incorreto)
- Implementação básica usando `map` sem qualquer mecanismo de sincronização
- Demonstra:
  - Falta de segurança em ambientes concorrentes
  - Race conditions
  - Panic em runtime (`concurrent map writes`)
- Quando usar:
  - Nunca em produção
  - Para demonstrar a necessidade de mecanismos de sincronização

### map-rwmutex/
- Usa `sync.RWMutex` para otimizar leituras
- Demonstra:
  - Read-heavy workloads
  - Trade-off entre reads e writes
- Importante:
  - RWMutex pode piorar performance se houver muitas escritas

### sync-map/
- Uso de `sync.Map`
- Demonstra:
  - Cenários onde `sync.Map` é eficiente
  - Casos onde ele **não** é recomendado
- Quando usar:
  - Muitas leituras e poucas escritas
  - Chaves relativamente estáveis
  - Cache simples sem necessidade de controle fino de locking


### ttl/
- Cache com expiração por tempo
- Demonstra:
  - Lazy expiration
  - Goroutine de cleanup
  - Risco de memory leak

### eviction/
- Estratégias de remoção:
  - LRU
  - LFU
  - FIFO
- Demonstra:
  - Limitação de memória
  - Impacto no hit rate

---

## 🌐 l2-redis — Cache Distribuído com Redis

Cache compartilhado entre múltiplos pods (L2). Resolve consistência entre instâncias, mas adiciona latência de rede.

### basic/
- Cache aside clássico
- Redis como key-value store

### mutex-key/
- Uso de mutex por chave no lado da aplicação
- Demonstra:
  - Evitar múltiplas goroutines acessando Redis ao mesmo tempo
  - Proteção local vs lock distribuído

### pubsub-invalidation/
- Invalidação ativa usando Redis Pub/Sub
- Demonstra:
  - Consistência entre pods
  - Cache invalidation em tempo real

---

## 🧩 multilayer — Cache Multi-Camadas (L1 + L2)

Combinação de cache local e distribuído.

### l1-l2-singleflight/
- Singleflight no L1
- Demonstra:
  - Melhor abordagem para cache miss caro
  
---
