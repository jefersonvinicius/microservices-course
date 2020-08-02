# **Indrodução**

## O que é?
Kubernetes é uma ferramenta open-source que ajuda no gerenciamento de aplicativos em containers

## Pontos importantes
- Ele é disponibilizado através de um conjunto de APIs
- Essas APIs são acessadas via um CLI chamado: *kubectl*
- Tudo é baseado em estado. Você configura o estado de cada objeto

## Dinâmica superficial
**Clusters**  
- É um conjunto de máquinas (Nodes)  
- Cada máquina possui uma quantidade de vCPU e Memória

**Pods**  
- Unidade que contém os containers provisionados  
- O Pod representa os processos rodando o cluster

**Deployment**  
Um *deployment* provê uma forma de como uma determinada aplição vai ser escalada em um cluster de acordo com a *ReplicaSets* especificada.

> **ReplicaSets** tem o objetivo de manter estável um determinado conjunto de Pods executando por um determinado tempo. Geralmente usado para garantir a disponibilidade de um número idêntico de Pods.

# **Services**
É uma forma de unir um conjunto de Pods e definir uma forma de acesso a eles.

## ClusterIP (Padrão)
Funciona através de um proxy. O tráfego vai para um proxy e este sabe exatamente para qual service encaminhar e funciona de forma interna

## NodePort
Ele funciona através de Nodes, onde existe um porta configurada em cada um para receber o tráfego e posteriormente ser redirecionado para algum service.

## LoadBalancer (Queridinho)
Expõe um IP externo para receber o tráfego. Esse tráfego é controlado por um LoadBalancer que vai redicionar para os services de acordo com a "carga"

---

> Para fazer um conjunto de Pods é necessário o uso de selectors. Ou seja, um Pod vai ter um selector e quando for criado o LoadBalancer poderá ser setado para que tipo de aplicação ele devera encaminhar o tráfego

