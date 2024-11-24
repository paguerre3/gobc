import { fetchData } from './common.js';

export async function fetchWallet(setData) {
    const today = new Date();
    const yesterday = new Date(today);
    yesterday.setDate(today.getDate() - 1);
    const todayStr = today.toISOString().split('T')[0];
    const yesterdayStr = yesterday.toISOString().split('T')[0];

    await fetchData({
        uriSuffix: '/wallet',
        localCacheKey: `wallet-key-${todayStr}`,
        previousLocalCacheKey: `wallet-key-${yesterdayStr}`,
        setData,
    });
}