import { Flex, Image, Text } from "@chakra-ui/react";

interface ItemProps {
  src: string;
  text: string;
}

export function Item({ src, text }: ItemProps) {
  return (
    <Flex direction="column" align="center" >
      <Image src={src} alt="text"></Image>
      <Text mt="5">{text}</Text>
    </Flex>
  );
}