import { NextApiRequest, NextApiResponse } from "next";
import { Readable } from "stream";
import Stripe from "stripe";
import { stripe } from "../../../services/stripe";
import { saveSubscription } from "./_lib/manageSubscription";


//O proximo codigo é um pouco confuso, foi obtido atraves de um exemplo.
async function buffer(readable: Readable) {

  const chunks = [];

  for await (const chunk of readable) {
    chunks.push(
      typeof chunk === "string" ? Buffer.from(chunk) : chunk
    );
  }
  return Buffer.concat(chunks);

}

//Por padrão o next.js informa na documentação que se esperamos stream temos que desativar o bodyParser.
//Assim ele não fica esperando um json ou um formulario.
export const config = {
  api: {
    bodyParser: false
  }
}

//Especificamos os eventos que desejamos receber do stripe.
const relevantEvents = new Set([
  'checkout.session.completed',
  'customer.subscription.updated',
  'customer.subscription.deleted',
]);

export default async function (req: NextApiRequest, res: NextApiResponse) {

  //Validamos se a requisição é do tipo POST, se não for, ignoramos.
  if (req.method === "POST") {

    const buf = await buffer(req);

    //Validamos o id do terceiro que vai enviar POST nessa rota.
    //Tudo conforme documentação do Stripe.
    const secret = req.headers['stripe-signature'];



    let event: Stripe.Event;

    try {
      event = stripe.webhooks.constructEvent(buf, secret, process.env.STRIPE_WEBHOOK_SECRET);

    } catch (err) {
      return res.status(400).send(`Webhook error: ${err.message}`);
    }

    const { type } = event;


    if (relevantEvents.has(type)) {
      try {
        switch (type) {

          case 'customer.subscription.updated':
          case 'customer.subscription.deleted':

            const subscription = event.data.object as Stripe.Subscription;

            await saveSubscription(
              subscription.id,
              subscription.customer.toString(),
              false
            );

            break;
          case 'checkout.session.completed':

            const checkoutSessions = event.data.object as Stripe.Checkout.Session;

            await saveSubscription(
              checkoutSessions.subscription.toString(),
              checkoutSessions.customer.toString(),
              true
            )

            break;
          default:
            throw new Error('Unhandled event.');
        }
      } catch (err) {
        return res.json({ error: 'WebHook handle failure.' });
      }
    }




    res.status(200).json({ received: true });
  } else {
    //Se não for post ele informa que o metodo não é suportado e informa que espera POST.
    res.setHeader('Allow', 'POST');
    res.status(405).end('Method not allowed');
  }


}