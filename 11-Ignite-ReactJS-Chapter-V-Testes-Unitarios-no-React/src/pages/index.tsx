import Head from 'next/head'
import styles from './home.module.scss'
import { SubscribeButton } from '../components/SubscribeButton'
import { GetStaticProps } from 'next'
import { stripe } from '../services/stripe'
import Image from 'next/image'

type PriceProps = {
  product: {
    amount: string
  }
}


export default function Home({ product }: PriceProps) {

  return (
    <>
      <Head>
        <title>Início | ig.news</title>
      </Head>
      <main className={styles.contentContainer} >
        <section className={styles.hero}>
          <span> 👋 Hey, welcome</span>
          <h1>News about the <span>React</span> world</h1>
          <p>
            Get access to all  publications <br />
            <span>For {product.amount} per month</span>
          </p>
          <SubscribeButton />
        </section>
        <Image src="/images/avatar.svg" alt="ig news" width={400} height={500} />
      </main>
    </>
  )
}


export const getStaticProps: GetStaticProps = async () => {
  const price = await stripe.prices.retrieve('price_1JbiwWIezH33hFSXEL8nbJiI')

  const product = {
    priceId: price.id,
    amount: new Intl.NumberFormat('en-US', {
      style: 'currency',
      currency: 'USD'
    }).format(price.unit_amount / 100)
  }

  return {
    props: {
      product
    },
    revalidate:  60 * 60 * 24 // 24h
  }
}
