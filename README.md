Esse diretório contém um exemplo de web api feita utilizando Go
- Para manipular rotas foi utilizado o Fiber (github.com/gofiber/fiber/v2)
- Para autenticação foi utilizada golang-jwt (github.com/golang-jwt/jwt/v5)

Como se trata apenas de um exemplo, o código possui usuário e senha fixos.

### Para testar ###
1 - clone o repositório <br>
2 - abra no vscode <br>
3 - no terminal digite ```go run main.go``` <br>
4 - abrindo outro terminal digite <br>
  ```
    curl --data "user=angelo&password=123" http://localhost:3000/login
  ```
- Ao rodar o curl, um token deve ser retornado no terminal.
- Caso seja passado algum outro dado no usuário ou senha, é retornado status 401
