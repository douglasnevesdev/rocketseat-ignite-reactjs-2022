import { Text } from '@chakra-ui/react';

export function Logo() {
  return (
    <Text
      fontSize={["2x1", "3x1"]}
      fontWeight="bold"
      letterSpacing="tight"
      width='64'
    >
      DashBoard
      <Text
        as='span'
        marginLeft='1'
        color="red.500"
      >.</Text>
    </Text>
  );
}