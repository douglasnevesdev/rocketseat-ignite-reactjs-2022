import { Flex, Stack } from '@chakra-ui/react';
import { Header } from './../components/Header';
import { Banner } from './../components/Banner';
import Head from "next/head";
import { Travel } from '../components/Travel';

export default function Home() {
  return (
    <>
      <Head>
        <title>WordTrip</title>
      </Head>
      <Flex width="100%" flexDir="column" align="center">
        <Header />
        <Banner />
        <Travel />
      </Flex>
    </>
  )
}
