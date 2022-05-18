import { loadStripe } from '@stripe/stripe-js';

export async function getStripeJS() {
  //Como sera utilizado pelo frontend informamos a chave publica do stripe.
  const stripeJS = await loadStripe(process.env.NEXT_PUBLIC_STRIPE_PUBLIC_KEY);
  return stripeJS;

}