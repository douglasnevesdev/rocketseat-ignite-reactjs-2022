import { screen, render } from '@testing-library/react'
import Preview, { getStaticProps } from '../../pages/posts/preview/[slug]'
import { useSession } from 'next-auth/client'
import { mocked } from 'ts-jest/utils'
import { getPrismicClient } from '../../services/prismic'

jest.mock('next-auth/client')
jest.mock('next/router')
jest.mock('../../services/prismic')

const post = {
    slug: 'fake-slug-1',
    title: 'fake-title-1',
    content: 'fake-content-1',
    updatedAt: 'fake-data-1'
}

describe('Preview', () => {
    it('Should render the Preview page correctly', () => {

        const useSessionMocked = mocked(useSession)
        useSessionMocked.mockReturnValueOnce([null, false])

        render(
            <Preview
                post={post}
            />
        )
        expect(screen.getByText('Wanna continue reading?')).toBeInTheDocument()
        expect(screen.getByText('fake-title-1')).toBeInTheDocument()
    })

 
    it('Loads initial data from getServerSideProps', async () => {
        const getPrismicClientMocked = mocked(getPrismicClient)
        getPrismicClientMocked.mockReturnValueOnce({
            getByUID: jest.fn().mockResolvedValueOnce({
                data: {
                    title: [{ type: 'heading', text: 'fake-title-1' }],
                    content: [{ type: 'paragraph', text: 'fake-content-1' }]
                },
                last_publication_date: '01-01-2021'
            })
        } as any)

        const response = await getStaticProps({
            params: { slug: 'fake-slug-1' }
        })

        expect(response).toEqual(
            expect.objectContaining({
                props: {
                    post: {
                        slug: 'fake-slug-1',
                        title: 'fake-title-1',
                        content: '<p>fake-content-1</p>',
                        updatedAt: '01 de janeiro de 2021'
                    }
                }
            })
        )

    })
})