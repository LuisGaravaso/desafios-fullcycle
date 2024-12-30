# Full Cycle Docker Image

Esta é a solução para o desafio de criar uma imagem Docker minimalista com menos de 2MB que imprime a mensagem **"Full Cycle Rocks!!"**.

A imagem foi publicada no Docker Hub e está disponível para uso. Este documento também ensina como buildar a imagem a partir do código-fonte.

---

## Como usar a imagem Docker

1. Certifique-se de que o Docker está instalado no seu sistema.
2. Execute o comando abaixo para baixar e rodar a imagem diretamente do Docker Hub:

```bash
docker run luisgaravaso/fullcycle
```

3. O resultado esperado é:

```plaintext
Full Cycle Rocks!!
```

---

## Buildando a imagem manualmente

Se você deseja clonar o repositório e construir a imagem manualmente, siga os passos abaixo:

1. Clone este repositório:

```bash
git clone https://github.com/LuisGaravaso/desafios-fullcycle.git
cd desafios-fullcycle/FC3-Docker-ImagemGoMinimalista
```

2. Certifique-se de que o Docker está instalado e ativo.

3. Execute o comando para construir a imagem:

```bash
docker build -t <seu-user>/fullcycle .
```

Substitua `<seu-user>` pelo seu nome de usuário do Docker Hub ou qualquer identificador que você deseje.

4. Para testar a imagem criada, execute:

```bash
docker run <seu-user>/fullcycle
```

Você deve ver a saída:

```plaintext
Full Cycle Rocks!!
```

---

## Publicando no Docker Hub

Se você deseja publicar sua imagem para que outros possam usá-la:

1. Autentique-se no Docker Hub:

```bash
docker login
```

2. Envie a imagem para sua conta:

```bash
docker push <seu-user>/fullcycle
```

Agora sua imagem estará disponível no Docker Hub!

---

## Estrutura do projeto

- **main.go**: Arquivo fonte em Go Lang que imprime a mensagem "Full Cycle Rocks!!".
- **Dockerfile**: Arquivo de configuração para construir a imagem Docker minimalista usando multi-stage build.

---

## Sobre a imagem

- Linguagem utilizada: [Go Lang](https://go.dev/)
- Tamanho da imagem: < 2MB
- Base minimalista: `scratch`

---

## Links relevantes

- Repositório Docker Hub: [luisgaravaso/fullcycle](https://hub.docker.com/r/luisgaravaso/fullcycle)
- Documentação do Go Lang: [https://go.dev/](https://go.dev/)
- Documentação do Docker: [https://docs.docker.com/](https://docs.docker.com/)

