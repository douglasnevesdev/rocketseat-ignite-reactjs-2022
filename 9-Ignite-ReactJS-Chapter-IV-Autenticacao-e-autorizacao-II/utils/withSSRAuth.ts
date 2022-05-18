import { GetServerSideProps, GetServerSidePropsContext, GetServerSidePropsResult } from "next";
import { parseCookies, destroyCookie } from "nookies";
import { AuthTokenError } from "../errors/AuthTokenError";
import decode from 'jwt-decode';
import { validateUserPermissions } from "./validateUserPermissions";

type WithSSRAuthOptions = {
  permissions?: string[];
  roles?: string[];
}


//Uma função que retorna outra função.
//Foi criada para permitir acessar rotas livres que não exige autenticação.
export function withSSRAuth<P>(fn: GetServerSideProps<P>, options: WithSSRAuthOptions) {

  return async (ctx: GetServerSidePropsContext): Promise<GetServerSidePropsResult<P>> => {

    const cookies = parseCookies(ctx);

    const token = cookies['nextauth.token'];

    //Verifico se não existe o cookie, se não existe redireciono para / aonde o usuario fara o login.
    if (!token) {
      return {
        redirect: {
          destination: '/',
          permanent: false,
        }
      }
    }

    if (options) {

      const user = decode<{ permissions: [], roles: string[] }>(token);
      const { permissions, roles } = options;

      const userHasValidPermissions = validateUserPermissions({
        user,
        permissions,
        roles
      });

      //Retorna para uma pagina aonde todos os usuarios tem permissão para acessar.
      //Valida se usuario tem permissoes e roles para uma determinada pagina.
      if (!userHasValidPermissions) {
        return {
          redirect: {
            destination: '/dashboard',
            permanent: false
          }
          //notFound: true,
        }
      }

    }



    //Se tiver o cookie, vamos executar a função padrão fn.
    //A função padrão é o corpo de withSSRAuth, ou seja, o export const getServerSideProps.
    //Ou seja nessa etapa abaixo ele ta executando o corpo da função que esta declaro em getServerSideProps.
    try {
      return await fn(ctx);
    }
    catch (error) {

      if (error instanceof AuthTokenError) {
        destroyCookie(ctx, 'nextauth.token');
        destroyCookie(ctx, 'nextauth.refreshToken');
        return {
          redirect: {
            destination: '/',
            permanent: false
          }
        }
      }

    }



  }

}