import { Box, Flex, Heading, Button, Icon, Table, Thead, Tr, Th, Checkbox, Tbody, Td, Text, useBreakpointValue, Spinner, Link } from '@chakra-ui/react';
import { GetServerSideProps } from 'next';
import NextLink from 'next/link';
import { useState } from 'react';
import { RiAddLine, RiPencilLine } from 'react-icons/ri';
import { Header } from "../../components/Header";
import Pagination from '../../components/Pagination';
import { Sidebar } from "../../components/SideBar";
import { api } from '../../services/api';
import { useUsers, getUsers } from '../../services/hooks/useUsers';
import { queryClient } from '../../services/QueryClient';

//export default function UserList({ users, totalCount }) {
export default function UserList() {

  const [page, setPage] = useState(1);

  //Criamos um hook useUsers para reduzir a quantidade de codigo, nesse hook usamos o useQuery.
  //Se passamos no options um initialData o useQuery não fara a requisição, ou seja, a requisição não vai acontecer pelo cliente, somente pelo servidor, porem esse codigo apresentou problemas porque estamos usando o mirage como API, devido a isso deixamos comentado.
  const { data, isLoading, isFetching, error } = useUsers(page, /*{
    initialData: users
  }*/);


  const isWideVersion = useBreakpointValue({
    base: false,
    lg: true,
  });

  //Quando o usuario passa o mouse encima do nome que esta na lista um cache é criado com todas as informações, quando o usuario clica sobre o nome ele tem a impressão que os dados carregaram rapidamente.
  async function handlePrefectUser(userId: string) {
    //No primeiro parametro informamos a chave do cache e no segundo parametro a função que fara o fetch dos dados.
    await queryClient.prefetchQuery(['user', userId], async () => {
      const response = await api.get(`users/${userId}`);
      return response.data;
    }, {
      //Durante 10 minutos os dados não vão precisar ser recarregados novamente.
      staleTime: 1000 * 60 * 10
    });
  }


  return (
    <Box>
      <Header />
      <Flex width="100%" marginY="6" maxWidth={1480} marginX='auto' paddingX="6">
        <Sidebar />
        <Box flex='1' borderRadius={8} bg="gray.800" padding={["6", "8"]}>
          <Flex marginBottom="8" justify="space-between" align="center">
            <Heading size="lg" fontWeight="normal">
              Usuários
              {!isLoading && isFetching && <Spinner size="sm" color="gray.500"></Spinner>}
            </Heading>
            <NextLink href="/users/create" passHref>
              <Button as='a' size="sm" fontSize="sm" colorScheme="red" leftIcon={<Icon as={RiAddLine} fontSize="20" />}>
                Criar novo usuário
              </Button>
            </NextLink>
          </Flex>

          {isLoading ? (
            <Flex justifyContent="center"><Spinner></Spinner></Flex>
          ) : error ? (
            <Flex justifyContent="center">Falha ao obter dados dos usuarios</Flex>
          ) : (
            <>
              <Table colorScheme="whiteAlpha">
                <Thead>
                  <Tr>
                    <Th paddingX={["4", "4", "6"]} color="gray.300" width="8">
                      <Checkbox colorshema="red" />
                    </Th>
                    <Th>
                      Usuário
                    </Th>
                    {isWideVersion &&
                      <Th>
                        Data de Cadastro
                      </Th>
                    }
                    <Th width="8">
                    </Th>
                  </Tr>
                </Thead>
                <Tbody>
                  {data.users.map(
                    user => {
                      return (
                        <Tr key={user.id}>
                          <Td paddingX={["4", "4", "6"]}>
                            <Checkbox colorshema="red" />
                          </Td>
                          <Td paddingX={["4", "4", "6"]}>
                            <Box>
                              <Link color="purple.400" onMouseEnter={() => handlePrefectUser(user.id)}>
                                <Text fontWeight="bold">{user.name}</Text>
                              </Link>
                              <Text fontSize="sm" color="gray.300">{user.email}</Text>
                            </Box>
                          </Td>
                          <Td>
                            {isWideVersion && <Text>{user.createdAt}</Text>}
                          </Td>
                          <Td>
                            <Button as='a' size="sm" fontSize="sm" colorScheme="purple" iconSpacing={isWideVersion ? '1.5' : '-0.5'} leftIcon={<Icon as={RiPencilLine} fontSize="16" />}>
                              {isWideVersion && 'Editar'}
                            </Button>
                          </Td>
                        </Tr>
                      )
                    }
                  )}

                </Tbody>
              </Table>

              <Pagination totalCountOfRegisters={data.totalCount} currentPage={page} onPageChange={setPage} />

            </>
          )}

        </Box>
      </Flex >

    </Box >
  );
}

/*

export const getServerSideProps: GetServerSideProps = async () => {

  const { users, totalCount } = await getUsers(1);

  return {
    props: {
      users,
    }
  }
}
*/