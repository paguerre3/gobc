import React from 'react'

export default function Footer(props) {
    const {data} = props
    return (
        <footer>
            © {data?.Year} Cami Wallet. All rights reserved. Powered by <a href="https://buymeacoffee.com/pabloaguer8" target="_blank" rel="noopener noreferrer">PA</a>.
        </footer>
    )
}
