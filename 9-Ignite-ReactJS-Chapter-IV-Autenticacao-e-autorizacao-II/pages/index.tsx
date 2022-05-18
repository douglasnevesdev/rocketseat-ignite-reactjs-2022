import styles from '../styles/Home.module.css'
import { FormEvent, useContext, useState } from 'react';
import { AuthContext } from '../context/AuthContext';
import { GetServerSideProps } from 'next';
import { parseCookies } from 'nookies';
import { withSSRGuest } from '../utils/withSSRGuest';

export default function Home() {


  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const { signIn } = useContext(AuthContext);

  async function handleSubmit(event: FormEvent) {
    event.preventDefault();
    const data = { email, password };
    await signIn(data);
  }

  return (
    <form onSubmit={handleSubmit} className={styles.container}>
      <input type="email" value={email} onChange={e => setEmail(e.target.value)} ></input>
      <input type="password" value={password} onChange={e => setPassword(e.target.value)} ></input>
      <button type="submit">Entrar</button>
    </form>
  )
}

//O next.js espera que o getServerSideProps retorne uma função.
//A função que criamos withSSRGuest é chamada e precisa retorna uma função para suprir a exigencia do next.js.
export const getServerSideProps = withSSRGuest(async (ctx) => {
  return {
    props: {}
  }
});
