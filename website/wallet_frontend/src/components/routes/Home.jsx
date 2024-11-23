import React from 'react'
import { BrowserRouter as Router, Routes, Route, useNavigate } from "react-router-dom";

export default function Home() {
  const navigate = useNavigate();
  return (
    <div className="container">
        <button className="button" onClick={() => navigate("/wallet") }>Get Started</button>
        <div className="features">
            <p><h2>Secure Storage</h2> Keep your private keys safe with industry-leading encryption.</p>
            <p><h2>Easy Transactions</h2> Send and receive crypto with a seamless user interface.</p>
            <p><h2>Real-Time Tracking</h2> Monitor live prices and track your portfolio performance.</p>
            <p><h2>Decentralized Access</h2> Full control over your wallet—no third parties involved.</p>
        </div>
    </div>
  )
}
