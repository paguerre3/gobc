import React, { useEffect, useState } from "react";
import Header from './components/Header'
import Container from "./components/Container";
import Footer from './components/Footer'
import { fetchCopyright } from "./api/copyright";

const App = () => {
    const [data, setData] = useState(null);

    useEffect(() => {
        fetchCopyright(setData)
    }, []); // empty array dependency means render only once on mount (page load)

    return (
        <>
            <Header />

            <Container />
            
            { data && <Footer data={data}/> }
        </>
    );
};

export default App;

