import { fetchData } from './common.js';

export async function fetchCopyright(setData) {
    const thisYear = new Date().getFullYear();
    const previousYear = thisYear - 1;

    await fetchData({
        uriSuffix: '/copyright',
        localCacheKey: `copyright-key-${thisYear}`,
        previousLocalCacheKey: `copyright-key-${previousYear}`,
        setData,
    });
}

