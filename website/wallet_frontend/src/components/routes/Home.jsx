import React from 'react'
import { useNavigate } from 'react-router-dom'

export default function Home() {
  const navigate = useNavigate();
  return (
    <div className="container">
        <button className="button" onClick={() => navigate('/wallet')}>Get Started</button>
        <div className="features">
            <div><h2>Secure Storage</h2> Keep your private keys safe with industry-leading encryption.</div>
            <div><h2>Easy Transactions</h2> Send and receive crypto with a seamless user interface.</div>
            <div><h2>Real-Time Tracking</h2> Monitor live prices and track your portfolio performance.</div>
            <div><h2>Decentralized Access</h2> Full control over your wallet—no third parties involved.</div>
        </div>
    </div>
  )
}
