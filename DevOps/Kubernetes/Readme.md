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

# Comandos

Comando para aplicar as alterações de um arquivo:
```
kubectl apply -f <to/path/file.yaml>
```

Comando para pegar os Deployments criados:
```
kubectl get deployments
```

Comando para pegar os Pods criados:
```
kubectl get pods
```

Comando para pegar os Services criados:
```
kubectl get services
```

Comando que retorna os logs de um pod:
```
kubectl logs <pod-name>
```

Comando para criação de deployment:
```
kubectl create deployment <deployment-name> --image=<image-name>
```

Comando para criar um service expondo um deployment:
```
kubectl expose deployment <deployment-name> --type=<service-type> --port=<container-port>
```

Comando para testar *service*:
```
minikube service <service-name>
```

Comando para deletar todos os *deployments*:
```
kubectl delete deployments --all
```

Comando para deletar um pod, deployment ou service:
```
kubectl delete <pods|services|deployments> <name>
```

Comando para criar um 'segredo':
```
kubectl create secret generic <secret-name> --from-literal=password='<password>'
```

Comando que cria um *pod* e um *service*:
```
kubectl run <name> --image=<image> --requests=cpu=200m --expose --port 80
```

# Kubernetes no GCP

Após a criação de um cluster no GCP é necessário conectar o `kubectl` ao servidores do GCP.

## Upgrade do cluster
Quando um determinado cluster faltar recursos para subir os pods você pode fazer um upgrade nele adicionando mais `node pool`

> Um **node pool** é um grupo de nodes dentro de um cluster, onde todos tem a mesma configuração

## Gerenciamento de recursos
O Kubernetes possibilita a configuração de quanto um determinado pod vai consumir seja de CPU, memória, etc...
Essa configuração é informado nos arquivos yaml através de `request` (solitações).

## Limitando a utilização de recursos
Enquanto as request define o mínimo que um pod precisa para subir, os `limits` informa qual o máximo de recursos consumidos por um pod. É importante notar que quando o limite de *CPU* é alcançado por um pod o Kubernetes vai diminuir sua atividade para não passar o máximo configurado, entretanto, quando o limite de *Memória* é atingido, o Kubernetes irá matar o seu pod e subir denovo.

## Usando Autoscaler