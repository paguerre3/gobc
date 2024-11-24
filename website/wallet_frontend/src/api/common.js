export async function fetchData({ uriSuffix, localCacheKey, previousLocalCacheKey, setData }) {
    const apiEndpoint = `${import.meta.env.VITE_WALLET_API_URL}${uriSuffix}`;

    // Check local cache
    const cachedData = localStorage.getItem(localCacheKey);
    if (cachedData) {
        const parsedData = JSON.parse(cachedData);
        setData(parsedData);
        console.log(`Fetched Cache DATA: ${localCacheKey}\n`, parsedData);
        return;
    }

    // Remove previous cache
    if (previousLocalCacheKey) {
        localStorage.removeItem(previousLocalCacheKey);
    }

    // Fetch from API
    try {
        const res = await fetch(apiEndpoint);
        const newData = await res.json();
        localStorage.setItem(localCacheKey, JSON.stringify(newData));
        setData(newData);
        console.log(`Fetched "new" API DATA: ${localCacheKey}\n`, newData);
    } catch (error) {
        console.error(`Error fetching data from ${uriSuffix}:`, error);
    }
}
