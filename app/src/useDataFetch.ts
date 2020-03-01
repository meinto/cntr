import { useState, useEffect } from 'react';

export const useDataFetch = (url: string) => {
  const [data, setData] = useState(null);
  useEffect(() => {
    const loadData = async (url: string) => {
      const response = await fetch(url);
      const json = await response.json();
      setData(json);
    };
    loadData(url);
  }, [url]);
  return data;
};