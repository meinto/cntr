import { useState } from 'react';

export const useDataFetch = (url: string): [[] | null, () => Promise<void>] => {
  const [data, setData] = useState(null);
  const loadData = async () => {
    const response = await fetch(url);
    const json = await response.json();
    if (JSON.stringify(json) !== JSON.stringify(data)) {
      setData(json);
    }
  };
  return [data, loadData];
};