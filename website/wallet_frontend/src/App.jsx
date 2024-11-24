import React, { useEffect, useState } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Header from './components/Header'
import Home from "./components/routes/Home";
import Wallet from "./components/routes/Wallet";
import Footer from './components/Footer'
import { fetchCopyright } from "./api/copyright";

const App = () => {
    const [data, setData] = useState(null);

    useEffect(() => {
        fetchCopyright(setData)
    }, []); // empty array dependency means render only once on mount (page load)

    return (
        <Router>
            <Header />

            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/wallet" element={<Wallet />} />
            </Routes>
            
            { data && <Footer data={data}/> }
        </Router>
    );
};

export default App;

