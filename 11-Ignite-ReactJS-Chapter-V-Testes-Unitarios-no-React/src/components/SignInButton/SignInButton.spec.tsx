import {render, screen} from '@testing-library/react'
import {SignInButton} from '.'
import {useSession} from 'next-auth/client'
import {mocked} from 'ts-jest/utils'

jest.mock('next-auth/client')

describe('SignInButton', () => {
  it('should render the button to login if the user is not logged', () => {

    const useSessionMocked = mocked(useSession)

    useSessionMocked.mockReturnValueOnce([null, false])

    render(
        <SignInButton/>
    )
    expect(screen.getByText('SignIn with GitHub')).toBeInTheDocument()
  })

  it('should render the button to logout with the user named if user is logged', () => {
    const useSessionMocked = mocked(useSession)
    useSessionMocked.mockReturnValueOnce([
        {user: {name: 'John Doe', email: 'johndoe@eamil.com'}, expires: 'fake-expires'}
        , false])
      render(
          <SignInButton />
      )
      expect(screen.getByText('John Doe')).toBeInTheDocument()
  })

})