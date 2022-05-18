//Executa somente quando usuario abre a tela pela primeira vez.
import axios, { AxiosError } from 'axios';
import { parseCookies, setCookie, } from 'nookies';
import { signOut } from '../context/AuthContext';
import { AuthTokenError } from '../errors/AuthTokenError';

//Variavel para verificar se a requisição ja foi solicitada.
let isRefreshing = false;
//Criando uma fila das requições
let failedRequestQueue = [];

export function setupAPIClient(ctx = undefined) {

  let cookies = parseCookies(ctx);

  const api = axios.create({
    baseURL: 'http://localhost:3333',
    headers: {
      Authorization: `Bearer ${cookies['nextauth.token']}`
    }
  });

  //Pode executar varias vezes durante sessão do usuario.
  api.interceptors.response.use(response => {
    return response;
  }, (error: AxiosError) => {
    if (error.response.status === 401) {
      if (error.response.data?.code === 'token.expired') {
        //Renovar Token
        cookies = parseCookies(ctx);
        const { 'nextauth.refreshToken': refreshToken } = cookies;
        //Dentro dessa variavel vamos ter toda a configuração da requisição que fizemos ao backend.  
        //Vamos ter toda as configurações que precisamos para repetir a requisição ao backend.
        //Vamos identificar qual rota, parametros entre outros itens.
        const originalConfig = error.config;

        if (!isRefreshing) {

          isRefreshing = true;

          //Verifica com API se o refresh token é valido e retorna o novo token + refreshToken
          api.post('/refresh', {
            refreshToken,
          }).then(response => {

            const { token } = response.data;

            setCookie(ctx, 'nextauth.token', token, {
              maxAge: 60 * 60 * 24 * 30, //30 dias
              path: '/' // Qualquer endereco da aplicação pode ter acesso, ou seja, é um cookie global.
            });

            setCookie(ctx, 'nextauth.refreshToken', response.data.refreshToken, {
              maxAge: 60 * 60 * 24 * 30, //30 dias
              path: '/' // Qualquer endereco da aplicação pode ter acesso, ou seja, é um cookie global.
            });

            //Adiciona o cabeçalho na requisição.
            api.defaults.headers['Authorization'] = `Baerer ${token}`;

            failedRequestQueue.forEach(request => request.onSuccess(token));

            failedRequestQueue = [];

          }).catch(error => {

            failedRequestQueue.forEach(request => request.onFailure(error));

            failedRequestQueue = [];

            if (process.browser) {
              signOut();
            }

          }).finally(() => {

            isRefreshing = false;

          });

        }

        //Unica forma que o axios aceita para ter algo assincrono.
        return new Promise((resolve, reject) => {
          failedRequestQueue.push({
            onSuccess: (token: string) => {

              if (!originalConfig?.headers) {
                return //Eu coloquei um return mas pode colocar algum erro ou um reject 
              }

              originalConfig.headers['Authorization'] = `Baerer ${token}`;

              resolve(api(originalConfig));

            },

            onFailure: (error: AxiosError) => {
              reject(error);
            }

          });
        });


      } else {
        if (process.browser) {

          signOut();

        } else {

          //Retornamos um erro para redirecionar o usuario em caso de ter sido feito o processo no backend.
          return Promise.reject(new AuthTokenError);

        }
      }

      return Promise.reject(error);
    }
  });

  return api;

}