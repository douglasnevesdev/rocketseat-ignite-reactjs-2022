import { Flex, Text, Box, Avatar } from '@chakra-ui/react';

interface ProfileProps {
  showProfileData?: boolean;
}

export function Profile({ showProfileData = true }: ProfileProps) {
  return (
    <Flex
      alignItems='center'
    >
      {showProfileData && (
        <Box
          marginRight="4"
          textAlign="right"
        >
          <Text>Douglas Neves</Text>
          <Text
            color="gray.300"
            fontSize="small"
          >d88.neves@gmail.com</Text>
        </Box>
      )}
      <Avatar size="md" name="Douglas Neves" src="https://github.com/douglasneves-net.png" />
    </Flex>

  );
}