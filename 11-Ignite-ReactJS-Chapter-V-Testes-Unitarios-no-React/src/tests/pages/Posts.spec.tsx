import { render, screen } from '@testing-library/react'
import Posts, { getStaticProps } from '../../pages/posts'
import { getPrismicClient } from '../../services/prismic'
import { mocked } from 'ts-jest/utils'

const posts = [
    {
        slug: 'fake-slug-1',
        title: 'fake-title-1',
        excerpt: 'fake-excerpt-1',
        updatedAt: 'fake-date-1'
    }
]

jest.mock('../../services/prismic')

describe('Posts', () => {
    it('Should to render the page posts correctly', () => {
        render(
            <Posts
                posts={posts}
            />
        )
        expect(screen.getByText('fake-title-1')).toBeInTheDocument()
    })

    it('Loads initial data from getServerSideProps', async () => {
        const getPrismicClientMocked = mocked(getPrismicClient)
        getPrismicClientMocked.mockReturnValueOnce({
            query: jest.fn().mockResolvedValueOnce({
                results: [
                    {
                        uid: 'fake-slug-1',
                        data: {
                            title: [{ type: 'heading', text: 'fake-title-1' }],
                            content: [{ type: 'paragraph', text: 'fake-excerpt-1' }]
                        },
                        last_publication_date: '01-01-2021'
                    }
                ]
            })
        } as any)

        const response = await getStaticProps({})

        expect(response).toEqual(
            expect.objectContaining({
                props: {
                    posts: [{
                        slug: 'fake-slug-1',
                        title: 'fake-title-1',
                        excerpt: 'fake-excerpt-1',
                        updatedAt: '01 de janeiro de 2021'
                    }]
                }
            })
        )
        
    })
    
})