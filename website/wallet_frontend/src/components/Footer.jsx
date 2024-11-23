import React from 'react'

export default function Footer(props) {
    const {data} = props
    return (
        <footer>
            © {data?.Year} Cami Wallet. All rights reserved.
        </footer>
    )
}
