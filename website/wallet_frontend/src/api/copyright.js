// fetchCopyright.js
export async function fetchCopyright(setData) {
    const copyrightApi = `${import.meta.env.VITE_WALLET_API_URL}/copyright`;

    const thisYear = new Date().getFullYear();
    const localCopyrightCacheKey = `copyright-key-${thisYear}`;
    const cacheCopyrightItem = localStorage.getItem(localCopyrightCacheKey);

    if (cacheCopyrightItem) {
        const cachedCopyrightData = JSON.parse(cacheCopyrightItem);
        setData(cachedCopyrightData);
        console.log('Fetched Copyright Cache DATA:' + localCopyrightCacheKey + '\n', cachedCopyrightData);
        return;
    }

    const previousYear = thisYear - 1;
    const previousLocalCopyrightCacheKey = `copyright-key-${previousYear}`;
    localStorage.removeItem(previousLocalCopyrightCacheKey);

    try {
        const res = await fetch(copyrightApi);
        const newCopyrightData = await res.json();
        localStorage.setItem(localCopyrightCacheKey, JSON.stringify(newCopyrightData));
        setData(newCopyrightData);
        console.log(`Fetched "new" Copyright API DATA: ${localCopyrightCacheKey} \n`, newCopyrightData);
    } catch (error) {
        console.error(error);
    }
}
