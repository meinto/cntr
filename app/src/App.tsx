import React from 'react';
import './App.css';
import { useDataFetch } from './useDataFetch'
import { LineChart, XAxis, YAxis, CartesianGrid, Line, Tooltip } from 'recharts';

function App() {
  const data = useDataFetch('http://localhost:5564/getKeys?startYear=2020&startMonth=2&startDay=29')
  const chartData = data !== null && data.map<any>((tuple: any) => ({
    name: tuple.date,
    keys: tuple.keys,
    clicks: tuple.clicks,
  }));
  console.log(chartData)
  return (
    <div className="App">
      <h1>Statistics</h1>
      <LineChart width={500} height={300} data={chartData || []}>
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
