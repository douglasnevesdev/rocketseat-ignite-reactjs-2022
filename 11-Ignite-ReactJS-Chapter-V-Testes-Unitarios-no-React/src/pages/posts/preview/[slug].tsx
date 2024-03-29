import { GetStaticProps } from "next"
import { useSession } from 'next-auth/client'
import {useRouter} from 'next/router'
import { getPrismicClient } from "../../../services/prismic"
import { RichText } from 'prismic-dom'
import Head from 'next/head'
import styles from '../post.module.scss'
import Link from 'next/link'
import { useEffect } from "react"
import router from "next/router"

type PostPreviewProps = {
    post: {
        slug: string;
        title: string;
        content: string;
        updatedAt: string;
    }
}

export default function Preview({ post }: PostPreviewProps) {


    const [session] = useSession()
    const Router = useRouter()

    useEffect(() => {
        if(session?.activeSubscription){
            Router.push(`/posts/${post.slug}`)
        }
    }, [session])

    return (
        <>
            <Head>
                <title>{post.title} | Ignews</title>
            </Head>
            <main className={styles.container}>
                <article className={styles.post}>
                    <h1>{post.title}</h1>
                    <time>{post.updatedAt}</time>
                    <div
                        dangerouslySetInnerHTML={{ __html: post.content }}
                        className={`${styles.previewContent} ${styles.postContent}`}
                    >
                    </div>
                    <div className={styles.continueReading}>
                        Wanna continue reading?
                        <Link href='/'>
                            <a>Subscribe now 🐱‍🏍</a>
                        </Link>
                    </div>
                </article>
            </main>
        </>
    )
}

export const getStaticPaths = () => {
    return {
        paths: [],
        fallback: 'blocking'
    }
}

export const getStaticProps: GetStaticProps = async ({ params }) => {

    const { slug } = params;

    const prismic = getPrismicClient()

    const response = await prismic.getByUID('post', String(slug), {})

    const post = {
        slug,
        title: RichText.asText(response.data.title),
        content: RichText.asHtml(response.data.content.splice(0, 4)),
        updatedAt: new Date(response.last_publication_date).toLocaleDateString('pt-BR', {
            day: '2-digit',
            month: 'long',
            year: 'numeric'
        })
    }

    return {
        props: {
            post
        },
        revalidate: 60 * 30
    }

}