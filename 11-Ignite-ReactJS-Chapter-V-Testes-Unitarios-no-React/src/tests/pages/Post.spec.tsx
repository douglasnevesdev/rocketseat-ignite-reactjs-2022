import { render, screen } from '@testing-library/react'
import { getSession } from 'next-auth/client'
import Post, { getServerSideProps } from '../../pages/posts/[slug]'
import { mocked } from 'ts-jest/utils'

jest.mock('next-auth/client')
jest.mock("../../services/prismic")

describe('Post', () => {

    it('Should render a post correctly', () => {
        render(
            <Post
                post={
                    {
                        slug: 'fake-slug-1',
                        title: 'fake-title-1',
                        content: 'fake-content-1',
                        updatedAt: 'fake-date-1',
                    }
                }
            />
        )
        expect(screen.getByText('fake-title-1')).toBeInTheDocument()
    })

    it('Redirects user if subscription is not found', async () => {

        const getSessionMocked = mocked(getSession)

        getSessionMocked.mockResolvedValueOnce({
            activeSubcription: null
        } as any)

        const response = await getServerSideProps({
            paramns: {
                slug: '/',
            }
        } as any)

        expect(response).toEqual(
            expect.objectContaining({
                redirect: expect.objectContaining({
                    destination: '/',
                    permanent: false
                })
            })
        )
    })
})