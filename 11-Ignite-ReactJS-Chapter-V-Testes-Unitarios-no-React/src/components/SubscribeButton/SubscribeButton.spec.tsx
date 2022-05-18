import { render, screen, fireEvent } from '@testing-library/react'
import { SubscribeButton } from '.'
import { signIn, useSession } from 'next-auth/client'
import { mocked } from 'ts-jest/utils'
import { useRouter } from 'next/router'

jest.mock('next-auth/client')
jest.mock('next/router')

describe('SubscribeButton', () => {

    it('Should render the SubscribeButton component correctly', () => {

        const useSessionMocked = mocked(useSession)
        useSessionMocked.mockReturnValueOnce([null, false])

        render(
            <SubscribeButton />
        )
        expect(screen.getByText('Subscribe now')).toBeInTheDocument()
    })

    it('Redirects user to sign in when not authenticated', () => {

        const useSessionMocked = mocked(useSession)
        useSessionMocked.mockReturnValueOnce([null, false])

        render(
            <SubscribeButton />
        )

        const subscribeButton = screen.getByText('Subscribe now')
        const signInMocked = mocked(signIn)

        fireEvent.click(subscribeButton)
        expect(signInMocked).toHaveBeenCalled()
    })

    it('Redirects user to posts pages when authenticated and subscribed', () => {

        const useRouterMocked = mocked(useRouter)
        const useSessionMocked = mocked(useSession)
        const pushMock = jest.fn()

        useRouterMocked.mockReturnValueOnce({
            push: pushMock('/posts')
        } as any)

        useSessionMocked.mockReturnValueOnce([
            { user: { 
                name: 'Jhon doe', email: 'jhon.doe@gmail.com'
             }, expires: 'fake-expires', activeSubscription: 'fake-subscription'
            }, false
        ])

        render(
            <SubscribeButton />
        )

        const subscribeButton = screen.getByText('Subscribe now')
        fireEvent.click(subscribeButton)

        expect(pushMock).toHaveBeenCalled()

    })


})

