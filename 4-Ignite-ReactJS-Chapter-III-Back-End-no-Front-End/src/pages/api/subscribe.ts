import { NextApiRequest, NextApiResponse } from "next";
import { getSession } from "next-auth/react";
import { fauna } from "../../../services/fauna";
import { stripe } from "../../../services/stripe";
import { query as q } from 'faunadb';

type User = {
  ref: {
    id: string;
  },
  data: {
    stripe_customer_id: string;
  }
}


export default async (req: NextApiRequest, res: NextApiResponse) => {

  if (req.method === 'POST') {

    //Função fornecida pelo next-auth que obtem as informações do cookies.
    const session = await getSession({ req });

    //Obtenho os dados do usuario no banco de dados faunadb
    const user = await fauna.query<User>(
      q.Get(
        q.Match(
          q.Index('user_by_email'),
          q.Casefold(session.user.email)
        )
      )
    );

    //Com LET e IF validamos se existe no FaubaDB um registro de compra.
    //Se não existir, criamos o usuario no Striper e no FaubaDB e retornamos o id do stripe.
    let customerId = user.data.stripe_customer_id

    if (!customerId) {

      //Cria usuario no striper.
      const stripeCustomer = await stripe.customers.create({
        email: session.user.email,
      });

      //Salva usuario no banco de dados.
      await fauna.query(
        q.Update(
          q.Ref(q.Collection('users'), user.ref.id),
          {
            data: {
              stripe_customer_id: stripeCustomer.id,
            }
          }
        )
      );

      customerId = stripeCustomer.id;

    }




    const checkoutCheckoutSession = await stripe.checkout.sessions.create({
      //Esse é o id do usuario no striper, ele foi gerado apos criação no passo anterior.
      //Se tudo ocorrer certo, o usuario sera redirecionado para pagar o valor informado.
      customer: customerId,
      payment_method_types: ['card'],
      billing_address_collection: 'required',
      line_items: [
        { price: 'price_1Jsse1FVP8EC7a9VmISrBY8N', quantity: 1 }
      ],
      mode: 'subscription',
      allow_promotion_codes: true,
      success_url: process.env.STRIPE_SUCCESS_URL,
      cancel_url: process.env.STRIPE_CANCEL_URL
    });

    return res.status(200).json({ sessionId: checkoutCheckoutSession.id });

  } else {
    res.setHeader('Allow', 'POST');
    res.status(405).end('Method not allowed');
  }

}