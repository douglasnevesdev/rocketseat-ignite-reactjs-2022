import { render, screen } from '@testing-library/react'
import Home, { getStaticProps } from '../../pages'
import { stripe } from '../../services/stripe'
import { mocked } from 'ts-jest/utils'

jest.mock('../../services/stripe')

describe('Home', () => {
    it('Home page renders correctly', () => {
        render(
            <Home product={{ amount: 'R$10,00' }} />
        )
        expect(screen.getByText('For R$10,00 per month')).toBeInTheDocument()
    })

    it('Should return correctly data from getServerSideProps', async () => {
        const retriveStripePricesMocked = mocked(stripe.prices.retrieve)
        retriveStripePricesMocked.mockResolvedValueOnce({
            id: 'fake-price-id',
            unit_amount: 1000
        } as any)

        const response = await getStaticProps({})
        expect(response).toEqual(expect.objectContaining({
            props:{
                product: {
                    priceId: 'fake-price-id',
                    amount: '$10.00'
                }
            }
        }))
    })

})
