# gRPC

Criando uma aplicação em go usando o gRPC como forma de comunicação entre sistemas.

# Comandos Úteis

-> Do lado de server:

Compilar o proto: protoc --go_out=. --go-grpc_out=. proto/course_category.proto

Iniciar o server: go run cmd/grpcServer/main.go

-> Do lado do Evans (outro terminal):

Em outro terminal, abrir o evans: evans -r repl

Conectar ao pacote: package pb

Conectar ao serviço: service CategoryService

Executar um serviço: call service_name

Obs: o service_name acima pode ser algum dos que foram implementados: 
-CreateCategory
-CreateCategoryStream
-CreateCategoryStreamBidirectional
-ListCategories
-GetCategory

-> Do lado do sqlite (outro terminal):

Conectar no sqlite: sqlite3 db.sqlite

Obs: caso a tabela não tenha sido criada, crie-a:
create table categories (id string, name string, description string);
