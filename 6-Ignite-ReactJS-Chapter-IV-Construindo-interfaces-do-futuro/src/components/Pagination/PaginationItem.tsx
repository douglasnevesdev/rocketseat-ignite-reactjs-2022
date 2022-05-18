import { HStack, Button, Stack, Box } from '@chakra-ui/react';

interface PaginationitemProps {
  isCurrent?: boolean;
  number: number;
}

export function PaginationItem({ isCurrent = false, number }: PaginationitemProps) {

  if (isCurrent) {
    return (
      <Button
        size="sm"
        fontSize="xs"
        width="4"
        colorScheme="red"
        disabled
        _disabled={{
          bgColor: 'red.500',
          cursor: 'default'
        }}
      >{number}</Button>
    );
  }

  return (
    <Button
      size="sm"
      fontSize="xs"
      width="4"
      bgColor="gray.700"
      _hover={{
        bgColor: 'gray.500'
      }}
    >{number}</Button>
  );

}