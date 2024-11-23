import React, { useEffect, useState } from "react";
import Header from './components/Header'
import Container from "./components/Container";
import Footer from './components/Footer'


const App = () => {
    const [data, setData] = useState(null);

    useEffect(() => {
        async function fetchCopyright() {
            const copyrightApi = `${import.meta.env.VITE_WALLET_API_URL}/copyright`
        
            // cached data (cleared every year or after local storage refresh)
            const thisYear = new Date().getFullYear()
            const localCopyrightCacheKey = `copyright-key-${thisYear}`
            const cacheCopyrightItem = localStorage.getItem(localCopyrightCacheKey)
            if (cacheCopyrightItem) {
                const cachedCopyrightData = JSON.parse(cacheCopyrightItem)
                setData(cachedCopyrightData)
                console.log('Fetched Copyright Cache DATA:' + localCopyrightCacheKey + '\n', cachedCopyrightData)
                return
            }
            const previousYear = thisYear - 1
            const previousLocalCopyrightCacheKey = `copyright-key-${previousYear}`
            localStorage.removeItem(previousLocalCopyrightCacheKey)
        
            try {
                //console.log('copyright URL:', copyrightApi)
                const res = await fetch (copyrightApi)
                const newCopyrightData = await res.json()
                localStorage.setItem(localCopyrightCacheKey, JSON.stringify(newCopyrightData))
                setData(newCopyrightData)
                console.log(`Fetched "new" Copyright API DATA: ${localCopyrightCacheKey} \n` , newCopyrightData)
            } catch (error) {
                console.error(error)
            }
        }
        fetchCopyright()
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

