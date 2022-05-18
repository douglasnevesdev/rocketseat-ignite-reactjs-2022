import { useQuery, UseQueryOptions, UseQueryResult } from 'react-query';
import { api } from '../../services/api';

type User = {
  id: string;
  name: string;
  email: string;
  createdAt: string;
};

type GetUsersResponse = {
  users: User[];
  totalCount: number;
};

//Para melhorar a organização, separamos o metodo que faz a comunicação e formata as informações.
export async function getUsers(page: number): Promise<GetUsersResponse> {
  const { data, headers } = await api.get('/users', {
    params: {
      page,
    },
  });

  const totalCount = Number(headers['x-total-count']);


  //Formatação dos dados retornados da API.
  const users = data.users.map((user: User) => {
    return {
      id: user.id,
      name: user.name,
      email: user.email,
      createdAt: new Date(user.createdAt).toLocaleDateString('pt-BR', {
        day: '2-digit',
        month: 'long',
        year: 'numeric',
      }),
    };
  });

  return {
    users,
    totalCount,
  };
}

//useUsers é o hook que criamos para retornar os usuarios.
export function useUsers(page: number, options?: UseQueryOptions) {
  //useQuery é um Hook fornecido pelo react-query, ele recebe como parametro uma chave, função que deve executar e um objeto aonde podemos definir o tempo para o react-query buscar os dados novamente na API.
  //Quando fazer uma requisição, os dados dessa requisição sera armazenado no cache do front-end.
  //Lembre-se que precisa do provider em torno da aplicação para funcionar o React-Query.
  //Observe que adicionei um array ['users', page], porque dessa forma crio um cache de cada pagina.
  return useQuery(['users', page], () => getUsers(page), {
    //Tempo para o react-query buscar novos dados, neste caso 5 segundos.
    //staleTime: 1000 * 5,
    //Tempo para o react-query buscar novos dados, neste caso 10 minutos.
    staleTime: 1000 * 60 * 10,
    //Se passamos no options um initialData o useQuery não fara a requisição, ou seja, a requisição não vai acontecer pelo cliente, somente pelo servidor.
    ...options
  })
}