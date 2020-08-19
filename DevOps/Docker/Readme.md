# Docker

## O que são containers?
É o encapsulamento da aplicação e todas as suas dependências, possibilitando que a aplicação seja executado em qualquer outra máquina.

## Evitar:
1. Na minha máquina funciona
2. Garantir o mesmo ambiente para todos
3. Rápido e leve para trabalharmos, tanto em dev como em produção

## O que é docker?
É um software que executa e organiza containers. Diferente de uma máquina virtual, quando subimos um container e é apenas um processo no sistema operacional.

## Pontos chaves do funcionamento do docker
- Namespace. Cada container possui seu processo e ele é isolado de todo restante dos processos do sitema opercional
- Cgroups. Controla os recursos do sistema operacional
- Sistema de arquivos - OFS (Overlay File System)

## Comandos básicos

```
// Executa um container
docker run -it ubuntu  /bin/bash 
// -it : modo interativo

docker run -p 8080:80 nginx
// -p : redireciona a porta 8080 do  maquina para a porta 80 do docker

docker run --rm -p 8080:80 nginx
// --rm : remove o container automaticamento ao sair
```

```
// Lista os container que estão sendo utilizados
docker ps 

// Lista os todos os containers
docker ps -a

// Lista apenas os ids de todos os containers
docker ps -a -q
```

```
// Para o container
docker stop <container-id>
```

```
//  Remove o container
docker rm <container -id> -f

// Remove todos os containers
sudo docker rm $(docker ps -a -q) -f

// -f : força a remoção do container
```

## Imagens
Imagens são templates do container que será gerado. Elas são imutaveis. O container através de uma imagem possui uma área de gravação, onde ficarão os arquivos gerados durante o uso docker, entretanto, eles serão removidos assim que o container também for.

Remoção de imagem:

```
docker rmi <image-id>
```

Remoção de todas as imagens:

```
docker rmi $(docker images -q) -f
```

## Dockerfile

É um arquivo que possui todo o processos e comandos para que uma determinada imagem sera gerada para a sua aplicação.

## Gerando um build

Dockerfile de exemplo para aplicação node:

```
FROM node:10

WORKDIR /usr/src/app

COPY package*.json ./

RUN yarn install

COPY . .

EXPOSE 3000

CMD [ "node", "server.js" ]
```

Gerando uma imagem

```
sudo docker build -t <user-name>/<image-name>:<version> .
```
*Obs*: o ponto no final indica onde está o dockerfile da aplicação, normalmente ele fica na pasta `root` do projeto.

## Publicação no Docker Hub

Fazendo login:

```
docker login
```  

Fazendo o push para Docker Hub:
```
docker push <image-name>
```

## Exercícios e Desafios

- Publicando imagem Laravel ([Link da Imagem](https://hub.docker.com/repository/docker/jefersonvinicius/laravel-environment))

- Desafio de Golang ([Link da Imagem](https://hub.docker.com/repository/docker/jefersonvinicius/codeeducation))