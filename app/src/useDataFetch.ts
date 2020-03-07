import { useState } from 'react';

export const useDataFetch = (): [[] | null, (url: string) => Promise<void>] => {
  const [data, setData] = useState(null);
  const loadData = async (url: string) => {
    const response = await fetch(url);
    const json = await response.json();
    if (JSON.stringify(json) !== JSON.stringify(data)) {
      setData(json);
    }
  };
  return [data, loadData];
};