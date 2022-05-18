import { createServer, Factory, Model, Response, ActiveModelSerializer } from 'miragejs';
import { faker } from '@faker-js/faker';


type User = {
  name: string;
  email: string;
  created_at: string;
}

export function makeServer() {

  //O Mirage espera a declaração das routes() e models().
  // models() é os dados falsos que vamos guarda no banco de dados falso que o miragejs cria.
  const server = createServer({


    serializers: {
      application: ActiveModelSerializer
    },

    models: {
      //Partial é fornecido pelo Typescript, seu objetivo é informar que para user precisa dos campos de User porem não necessariamente todos os campos. 
      User: Model.extend<Partial<User>>({})
    },

    //Com factories podemos criar muitos dados para iniciar o banco de dados, factories é utilizado pela função seed.
    factories: {
      user: Factory.extend({
        //Todo metodo que esta dentro de uma factory recebe um indice que pode ser utilizado, i = 0, i = 1 e assim por diante.
        name(i: number) {
          return faker.name.findName();
        },
        email() {
          //Usamos a biblioteca faker para gerar dados falsos.
          return faker.internet.email().toLocaleLowerCase();
        },
        //O Factor entende que é created_at.
        createdAt() {
          return faker.date.recent(10);
        }
      })
    },

    //Função que inicia os dados fakes.
    seeds(server) {
      //Na função createList passo o nome do factory e quantos usuarios criar.
      //Toda vez que iniciar o servidor do miragejs ele vai gerar 200 usuarios aleatorios e fake.
      server.createList('user', 200);
    },


    routes() {
      //Precisamos setar o caminho que a aplicação vai utilizar para acessar nossa API do miragejs.
      this.namespace = "api";
      //Isso é um shothands, ou seja, o mirage entender que o comando abaixo deseja uma lista do tipo user.
      this.get('/users', function (schema, request) {
        const { page = 1, per_page = 10 } = request.queryParams;
        const total = schema.all('user').length;
        const pageStart = (Number(page) - 1) * Number(per_page);
        const pageEnd = pageStart + Number(per_page);
        const users = this.serialize(schema.all('user')).users.slice(pageStart, pageEnd);
        return new Response(
          200,
          { 'x-total-count': String(total) },
          { users }
        )
      });
      //Tambem é um shothands, da mesma forma que acima.
      this.post('/users');
      this.get('/users/:id');
      //Quando o miragejs terminar de definir as rotas,ele vai voltar o namespace para não conflitar com a api que existe no next.js.
      this.namespace = "";
      //Faz com que qualquer chamada do miragejs demore alguns segundos para acontecer
      this.timing = 750;
      //Isso faz com que todas as chamadas passe pelo miragejs, se não for detectada, o miragejs passa para api da aplicação (Se existir).
      this.passthrough();
      ;
    }


  });

  return server;


}