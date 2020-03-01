import React, { useState, useEffect } from 'react';
import './App.css';
import { LineChart, XAxis, YAxis, CartesianGrid, Line, Tooltip } from 'recharts';
import useInterval from '@use-it/interval';
import { useDataFetch } from './useDataFetch'

function App() {
  const [data, loadData] = useDataFetch('http://localhost:5564/getKeys?startYear=2020&startMonth=2&startDay=29')
  useInterval(() => {
    loadData();
  }, 1000);
  const chartData = data && data.map<any>((tuple: any) => ({
    name: tuple.date,
    keys: tuple.keys,
    clicks: tuple.clicks,
  })) || [];
  return (
    <div className="App">
      <h1 style={{marginBottom: 0}}>Statistics</h1>
      <strong>(last 10 days)</strong>
      <br />
      <LineChart width={500} height={300} data={chartData}>
        <XAxis dataKey="name"/>
        <YAxis/>
        <Tooltip />
        <CartesianGrid stroke="#eee" strokeDasharray="5 5"/>
        <Line type="monotone" dataKey="keys" stroke="#8884d8" />
        <Line type="monotone" dataKey="clicks" stroke="#82ca9d" />
      </LineChart>
    </div>
  );
}

export default App;
