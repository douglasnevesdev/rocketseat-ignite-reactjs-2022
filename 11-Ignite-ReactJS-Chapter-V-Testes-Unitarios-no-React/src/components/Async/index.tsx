import {useEffect, useState} from 'react'

export function Async(){

    const [isVisible, setIsVisible] = useState(false)

    useEffect(() => {
        setTimeout(() => {
            setIsVisible(!isVisible)
        }, 1000)
        //eslint-disable-next-line
    },[])

    return(
        <div>
            <h1>This is an Async component</h1>
            {isVisible  && <button>Button</button>}
        </div>
    )
}