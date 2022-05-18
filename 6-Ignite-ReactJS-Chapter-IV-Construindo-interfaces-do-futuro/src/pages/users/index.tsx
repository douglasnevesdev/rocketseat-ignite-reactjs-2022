import { Box, Flex, Heading, Button, Icon, Table, Thead, Tr, Th, Checkbox, Tbody, Td, Text, useBreakpointValue } from '@chakra-ui/react';
import Link from 'next/link';
import { RiAddLine, RiPencilLine } from 'react-icons/ri';
import { Header } from "../../components/Header";
import Pagination from '../../components/Pagination';
import { Sidebar } from "../../components/SideBar";

export default function UserList() {

  const isWideVersion = useBreakpointValue({
    base: false,
    lg: true,
  });

  return (
    <Box>
      <Header />
      <Flex width="100%" marginY="6" maxWidth={1480} marginX='auto' paddingX="6">
        <Sidebar />
        <Box flex='1' borderRadius={8} bg="gray.800" padding={["6", "8"]}>
          <Flex marginBottom="8" justify="space-between" align="center">
            <Heading size="lg" fontWeight="normal">Usuários</Heading>
            <Link href="/users/create" passHref>
              <Button as='a' size="sm" fontSize="sm" colorScheme="red" leftIcon={<Icon as={RiAddLine} fontSize="20" />}>
                Criar novo usuário
              </Button>
            </Link>
          </Flex>
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
              <Tr>
                <Td paddingX={["4", "4", "6"]}>
                  <Checkbox colorshema="red" />
                </Td>
                <Td paddingX={["4", "4", "6"]}>
                  <Box>
                    <Text fontWeight="bold">Douglas Neves</Text>
                    <Text fontSize="sm" color="gray.300">d88.neves@gmail.com</Text>
                  </Box>
                </Td>
                <Td>
                  {isWideVersion && <Text>21 de Janeiro, 2022</Text>}
                </Td>
                <Td>
                  <Button as='a' size="sm" fontSize="sm" colorScheme="purple" iconSpacing={isWideVersion ? '1.5' : '-0.5'} leftIcon={<Icon as={RiPencilLine} fontSize="16" />}>
                    {isWideVersion && 'Editar'}
                  </Button>
                </Td>
              </Tr>
              <Tr>
                <Td paddingX={["4", "4", "6"]}>
                  <Checkbox colorshema="red" />
                </Td>
                <Td paddingX={["4", "4", "6"]}>
                  <Box>
                    <Text fontWeight="bold">Douglas Neves</Text>
                    <Text fontSize="sm" color="gray.300">d88.neves@gmail.com</Text>
                  </Box>
                </Td>
                <Td>
                  {isWideVersion && <Text>21 de Janeiro, 2022</Text>}
                </Td>
                <Td>
                  <Button as='a' size="sm" fontSize="sm" colorScheme="purple" iconSpacing={isWideVersion ? '1.5' : '-0.5'} leftIcon={<Icon as={RiPencilLine} fontSize="16" />}>
                    {isWideVersion && 'Editar'}
                  </Button>
                </Td>
              </Tr>
              <Tr>
                <Td paddingX={["4", "4", "6"]}>
                  <Checkbox colorshema="red" />
                </Td>
                <Td paddingX={["4", "4", "6"]}>
                  <Box>
                    <Text fontWeight="bold">Douglas Neves</Text>
                    <Text fontSize="sm" color="gray.300">d88.neves@gmail.com</Text>
                  </Box>
                </Td>
                <Td>
                  {isWideVersion && <Text>21 de Janeiro, 2022</Text>}
                </Td>
                <Td>
                  <Button as='a' size="sm" fontSize="sm" colorScheme="purple" iconSpacing={isWideVersion ? '1.5' : '-0.5'} leftIcon={<Icon as={RiPencilLine} fontSize="16" />}>
                    {isWideVersion && 'Editar'}
                  </Button>
                </Td>
              </Tr>
            </Tbody>
          </Table>

          <Pagination />

        </Box>
      </Flex>

    </Box>
  );
}