import { Flex, useBreakpointValue, IconButton, Icon } from '@chakra-ui/react';
import { RiMenuLine } from 'react-icons/ri';
import { useSidebarDrawer } from '../../context/SidebarDrawerContext';
import { Logo } from '../Header/Logo';
import { NotificationsNav } from '../Header/NotificationsNav';
import { Profile } from '../Header/Profile';
import { SearchBox } from '../Header/SearchBox';

export function Header() {

  const { onOpen } = useSidebarDrawer();

  //Hook fornecido pelo chakra para tratar de responsividade
  const isWideVersion = useBreakpointValue({
    base: false,
    lg: true,
  });

  return (
    <Flex
      as="header"
      width='100%'
      maxWidth={1480}
      height='20'
      marginX='auto'
      marginTop="4"
      paddingX="6"
      alignItems='center'
    >

      {!isWideVersion && (<IconButton icon={
        <Icon
          as={RiMenuLine}
        />}
        aria-label="Open Navigation"
        fontSize="24"
        variant="unstyled"
        onClick={onOpen}
        marginRight="2"
      >
      </IconButton>)}

      <Logo />

      {isWideVersion && <SearchBox />}

      <Flex
        alignItems='center'
        marginLeft='auto'
      >
        <NotificationsNav />

        <Profile showProfileData={isWideVersion} />

      </Flex>

    </Flex>
  );
}