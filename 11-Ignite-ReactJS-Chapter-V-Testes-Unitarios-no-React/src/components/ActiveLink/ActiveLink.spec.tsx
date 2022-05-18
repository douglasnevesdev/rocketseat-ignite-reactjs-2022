import { render, screen } from '@testing-library/react'
import {ActiveLink} from '.'

jest.mock('next/router', () => {
    return {
        useRouter() {
            return {
                asPath: '/'
            }
        }
    }
})

describe('ActiveLink', () => {
    it('active link should renders correctly', () => {
        const { debug, getByText } = render(
            <ActiveLink
                activeClassName='active'
                href='/'
            >
                <p>OK!</p>
            </ActiveLink>
        )
        debug()
        expect(screen.getByText('OK!')).toBeInTheDocument()
    })
    
    it('active link is receiving the href attribute string', () => {
        render(
            <ActiveLink
                activeClassName='active'
                href='/'
            >
                <p>OK!</p>
            </ActiveLink>
        )
        expect(screen.getByText('OK!')).toHaveClass('active')
    })
    
})