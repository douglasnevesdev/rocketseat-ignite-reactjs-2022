import { useDisclosure, UseDisclosureReturn } from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { createContext, ReactNode, useContext, useEffect } from 'react';


interface SidebarDrawerProviderProps {
  children: ReactNode;
}

type SidebarDrawerContextData = UseDisclosureReturn

const sidebarDrawerContext = createContext({} as SidebarDrawerContextData);

export function SidebarDrawerProvider({ children }: SidebarDrawerProviderProps) {

  //Hook fornecido pelo chakra que abre, fecha, modifica.
  const disclosure = useDisclosure();
  //Hook fornecido pelo next que obtem a rota atual.
  const router = useRouter();

  //Executa quando uma pagina é modificada.
  useEffect(() => {
    disclosure.onClose();
  }, [router.asPath]);

  return (
    <sidebarDrawerContext.Provider value={disclosure} >
      {children}
    </sidebarDrawerContext.Provider>
  )
}

export const useSidebarDrawer = () => useContext(sidebarDrawerContext);