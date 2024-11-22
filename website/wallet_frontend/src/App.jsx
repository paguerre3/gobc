import React, { useEffect, useState } from "react";

const App = () => {
    const [year, setYear] = useState(null);

    useEffect(() => {
        const copyrightApi = `${import.meta.env.VITE_WALLET_API_URL}/copyright`
        console.log(copyrightApi)
        fetch(copyrightApi)
            .then((res) => {
                console.log(res)
                return res.json()
            })
            .then((data) => { 
                console.log(data)
                setYear(data.year ?? data.Year)
            })
            .catch((err) => console.error(err));
    }, []);

    return (
        <>
            <div className="header">
                <h1>Welcome to Cami Wallet</h1>
                <p>Elegant, Secure, and Stylish Wallets</p>
            </div>

            <div className="container">
                <div className="features">
                    <h2>Why Choose Cami Wallet?</h2>
                    <p>Discover the perfect blend of style and security with our Cami Wallets.</p>
                </div>
                <button className="button">Send</button>
            </div>
            
            <footer>
                © {year} Cami Wallet. All rights reserved.
            </footer>
        </>
    );
};

export default App;

