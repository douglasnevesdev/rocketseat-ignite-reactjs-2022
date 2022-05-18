import { Flex, FlexProps } from "@chakra-ui/react";
import { ReactNode } from 'react';

interface IContainer extends FlexProps {
  children: ReactNode;
}

export function Container({ children, ...rest }: IContainer) {
  return (
    <Flex width="100%" maxWidth="1440px" align="center" {...rest}>
      {children}
    </Flex >
  );
}