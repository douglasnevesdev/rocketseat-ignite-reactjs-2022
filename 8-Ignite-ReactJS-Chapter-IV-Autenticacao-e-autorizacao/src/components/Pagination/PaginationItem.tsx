import { HStack, Button, Stack, Box } from '@chakra-ui/react';

interface PaginationitemProps {
  isCurrent?: boolean;
  number: number;
  onPageChange: (page: number) => void;
}

export function PaginationItem({ isCurrent = false, number, onPageChange }: PaginationitemProps) {

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
      onClick={() => onPageChange(number)}
    >{number}</Button>
  );

}