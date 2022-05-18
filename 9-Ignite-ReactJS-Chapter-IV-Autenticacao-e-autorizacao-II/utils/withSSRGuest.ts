import { GetServerSideProps, GetServerSidePropsContext, GetServerSidePropsResult } from "next";
import { parseCookies } from "nookies";


//Uma função que retorna outra função.
//Foi criada para permitir acessar rotas livres que não exige autenticação.
export function withSSRGuest<P>(fn: GetServerSideProps<P>) {

  return async (ctx: GetServerSidePropsContext): Promise<GetServerSidePropsResult<P>> => {

    const cookies = parseCookies(ctx);

    //Verifico se existe cookie, se existe redireciono para /dashboard.
    if (cookies['nextauth.token']) {
      return {
        redirect: {
          destination: '/dashboard',
          permanent: false,
        }
      }
    }

    //Se não tiver cookie, vamos executar a função padrão fn.
    return await fn(ctx);

  }

}