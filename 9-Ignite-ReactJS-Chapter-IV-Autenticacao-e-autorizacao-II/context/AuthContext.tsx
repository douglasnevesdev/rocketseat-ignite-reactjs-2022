import { createContext, ReactNode, useEffect, useState } from "react";
import Router from 'next/router';
import { setCookie, parseCookies, destroyCookie } from 'nookies';
import { api } from "../services/apiClient";

type User = {
  email: string;
  permissions: string[];
  roles: string[];
}


type SignInCredentials = {
  email: string;
  password: string;
}


type AuthContextData = {
  //Toda função assincrona é automaticamente uma promise.
  signIn: (credentials: SignInCredentials) => Promise<void>;
  signOut: () => void;
  user: User;
  isAuthenticated: boolean;
};


type AuthProviderProps = {
  children: ReactNode;
}


export const AuthContext = createContext({} as AuthContextData)


//Utilização do BrodcastChannel que desconecta o usuario de todas as janelas.
let authChannel: BroadcastChannel;



export function signOut() {
  destroyCookie(undefined, 'nextauth.token');
  destroyCookie(undefined, 'nextauth.refreshToken');
  authChannel.postMessage('signOut');
  Router.push('/');
}


export function AuthProvider({ children }: AuthProviderProps) {

  const [user, setUser] = useState<User>();

  //Valido se existe um usuario, se existir é true, para essa verificação usamos !!user;
  const isAuthenticated = !!user;

  useEffect(() => {
    authChannel = new BroadcastChannel('auth');
    authChannel.onmessage = (message) => {
      switch (message.data) {
        case 'signOut':
          signOut();
          authChannel.close();
          break;
        /*
      case 'signIn':
        () => window.location.reload()
        break;
        */

        default:
          break;
      }
    }

  }, []);

  //Executa toda vez para verificar se usuario ja esta logado.
  useEffect(() => {
    //Retorna todos os cookies
    const { 'nextauth.token': token } = parseCookies();

    if (token) {
      api.get('/me').then(response => {

        const { email, permissions, roles } = response.data;

        setUser({
          email,
          permissions,
          roles
        })

      }).catch(() => {
        signOut();
      });
    }


  }, []);

  async function signIn({ email, password }: SignInCredentials) {

    try {

      const response = await api.post('sessions', {
        email,
        password
      });

      const { token, refreshToken, permissions, roles } = response.data;


      setUser({
        email,
        permissions,
        roles
      });


      //Primeiro parametro é o cotexto da requisição que acontece no backend, como estamos usando no frontend fica undefined.
      // Segundo parametro é o nome do cookie.
      // Terceiro parametro é a variavel.
      setCookie(undefined, 'nextauth.token', token, {
        maxAge: 60 * 60 * 24 * 30, //30 dias
        path: '/' // Qualquer endereco da aplicação pode ter acesso, ou seja, é um cookie global.
      });

      setCookie(undefined, 'nextauth.refreshToken', refreshToken, {
        maxAge: 60 * 60 * 24 * 30, //30 dias
        path: '/' // Qualquer endereco da aplicação pode ter acesso, ou seja, é um cookie global.
      });

      //Adiciona o cabeçalho na requisição.
      api.defaults.headers['Authorization'] = `Baerer ${token}`;

      Router.push('/dashboard');

      //authChannel.postMessage('signIn');

    } catch (erro) {
      console.log(erro);
    }

  }


  return <AuthContext.Provider value={{ isAuthenticated, signIn, signOut, user }}>
    {children}
  </AuthContext.Provider>


}