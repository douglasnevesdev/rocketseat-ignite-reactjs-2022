import { Flex, Text, Image } from '@chakra-ui/react';
import { useRef } from 'react';
//import Image from 'next/image';
import { Container } from './../Container';

export function Banner() {


  return (
    <Flex
      width="100%"
      height={["150px", "250px", "335px"]}
      backgroundImage={'background.png'}
      backgroundRepeat="no-repeat"
      backgroundSize="cover"
      justifyContent="center"
    >


      <Container>

        <Flex flexDir="column" justify="center" left={["15px", "15px", "139px"]} position="absolute">
          <Text color="gray.100" fontSize={["1.5rem", "2.25rem"]} fontWeight="medium" >
            5 Continentes, <br /> infinitas possibilidades.
          </Text>
          <Text
            mt="20px"
            color="gray.100"
            fontSize={["0.75rem", "1.25rem"]}
            maxWidth="500px"
          >
            Chegou a hora de tirar do papel a viagem que você sempre sonhou.
          </Text>
        </Flex>


        <Image
          src="/airplane.svg"
          position="absolute"
          right="141px"
          mt="2rem"
          width={["100%", "100%", "100%", "30%"]}
          display={["none", "none", "none", "inline"]}
        ></Image>

      </Container>


    </Flex >
  );
}

