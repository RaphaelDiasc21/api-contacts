# API-CONTACTS

 ## Inicialização do projeto
    - Comando no terminal: docker-compose up
 ### Tecnologias utilizadas
    - Linguagem Golang
    - Biblioteca Gorilla mux para roteamento
    - Biblioteca Testify para assert de resultados de teste
    - Gorm, ORM para comunicação com banco de dados

### Arquitetura
    - Foi pensada em uma arquitetura em camadas, centralizando as regras de negócios na camada de 'services'
    - Os models fazem interface direta com o banco de dados
    - Os resources servem de comunicação e transferência de dados entre o cliente e a API em si

### Resolução do problema principal
    - Identifiquei que o maior problema seria a resolução de qual banco de dados trabalhar com base no cliente informado, por isso em todos endpoints é obrigatório informar o nome do cliente e com base nisso a lógica proposta no código se encarrega de insanticiar a conexão com os respectivos banco de dados correspondestes

### Pendencias
    - Acabei que não consegui implementar o Swagger UI no projeto, apenas consegui gerar o arquivo json de swagger que documenta os endpoints da API
    - Não consegui implementar os testes dos handlers, senti dificuldade em mockar alguns métodos específicos