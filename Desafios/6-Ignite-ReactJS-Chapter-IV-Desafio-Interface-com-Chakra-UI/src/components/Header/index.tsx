import { Flex, Image } from '@chakra-ui/react';
import Link from 'next/link';

export function Header() {
  return (
    <Flex as="header" width='100%' justifyContent="center" align="center" height="6.25rem" background="gray.700">
      <Link href="/">
        <a>
          <Image src="/logo.svg" alt="WordTrip" />
        </a>
      </Link>
    </Flex>
  );
}