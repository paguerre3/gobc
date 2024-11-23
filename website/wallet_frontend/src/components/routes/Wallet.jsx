import React from 'react'

export default function Wallet() {
  return (
    <div className='container'>
      <button className="button">Reload Balance</button>
      <div className='features'>
        <p><h2>Wallet</h2></p>
        <h2><div id="walletBalance">0</div></h2>

        <p><strong>Public Key</strong></p>
        <textarea id="publicKey" cols="73" rows="2"></textarea>

        <p><strong>Private Key</strong></p>
        <textarea id="privateKey" cols="73" rows="1"></textarea>

        <p><strong>Blockchain Address</strong></p>
        <textarea id="blockChainAddress" cols="73" rows="1"></textarea>
      </div>
      <div className='features'>
        <p><h2>Send Money</h2></p>
        <h2>Amount <input id="sendAmount" type="number"/></h2>

        <p><strong>Receipient Blockchain Address</strong></p>
        <textarea id="receipientBlockChainAddress" cols="73" rows="1"></textarea>
        <h2><button className="button">Send</button></h2>
      </div>
      
    </div>
  )
}
