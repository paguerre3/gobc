import React, { useEffect, useState } from "react";
import { fetchWallet } from "../../api/wallet";

export default function Wallet() {
  const [data, setData] = useState(null);

  useEffect(() => {
    fetchWallet(setData)
  }, []); // empty array dependency means render only once on mount (page load)

  return (
    <div className='container'>
      <button className="button">Reload Wallet</button>
      <div className='features'>
        <h2>Wallet</h2>
        <h2><div id="walletBalance">0</div></h2>

        <strong>Public Key</strong>
        <textarea id="publicKey" cols="73" rows="2" value={data?.publicKey}/>

        <strong>Private Key</strong>
        <textarea id="privateKey" cols="73" rows="1" value={data?.privateKey}/>

        <strong>Blockchain Address</strong>
        <textarea id="blockChainAddress" cols="73" rows="1" value={data?.blockChainAddress}/>
      </div>
      <div className='features'>
        <h2>Send Money</h2>
        <h2>Amount <input id="sendAmount" type="number"/></h2>

        <strong>Recipient Blockchain Address</strong>
        <textarea id="recipientBlockChainAddress" cols="73" rows="1"/>
      </div>
      <h2><button className="button">Send</button></h2>
    </div>
  )
}
