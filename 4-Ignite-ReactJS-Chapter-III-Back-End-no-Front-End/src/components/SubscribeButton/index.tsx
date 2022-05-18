import { useSession, signIn } from 'next-auth/react';
import { api } from '../../../services/api';
import { getStripeJS } from '../../../services/stripe.js';
import styles from './styles.module.scss';


interface SubscribeButtonProps {
  priceId: string;
}

export function SubscribeButton({ priceId }: SubscribeButtonProps) {

  const session = useSession();

  async function handleSubscribe() {

    if (!session.data) {
      signIn('github');
      return;
    }

    try {
      //Nome da rota é subscribe porque esse é o nome do arquivo.
      const response = await api.post('/subscribe');
      const { sessionId } = response.data;
      const stripe = await getStripeJS();
      await stripe.redirectToCheckout({ sessionId });
    } catch (err) {
      alert(err.message);
    }

  }

  return (
    <button
      type="button"
      onClick={handleSubscribe}
      className={styles.subscribeButton}
    >
      Subscribe now
    </button>
  );
}


