import { extendTheme } from '@chakra-ui/react';

export const theme = extendTheme({

  colors: {
    black: {
      "900": "#000000",
    },
    gray: {
      "900": "#47585B",
      "800": "#999999",
      "700": "#F5F8FA"
    },
    yellow: {
      "900": "#FFBA08",
    }
  },
  fonts: {
    heading: 'Poppins',
    body: 'Poppins'
  },
  styles: {
    global: {
      body: {
        bg: '#ffffff',
        color: 'gray.900'
      }
    }
  }
});