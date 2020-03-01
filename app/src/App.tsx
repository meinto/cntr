import React, { useState, useEffect } from 'react';
import './App.css';
import { ResponsiveContainer, XAxis, YAxis, CartesianGrid, Area, Tooltip, AreaChart } from 'recharts';
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
      <ResponsiveContainer aspect={2.5}>
        <AreaChart data={chartData}>
          <defs>
            <linearGradient id="colorKeys" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8}/>
              <stop offset="95%" stopColor="#8884d8" stopOpacity={0}/>
            </linearGradient>
            <linearGradient id="colorClicks" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#82ca9d" stopOpacity={0.8}/>
              <stop offset="95%" stopColor="#82ca9d" stopOpacity={0}/>
            </linearGradient>
          </defs>
          <XAxis dataKey="name"/>
          <YAxis/>
          <Tooltip contentStyle={{color: '#333'}}/>
          <CartesianGrid stroke="#aaa" strokeDasharray="3 3"/>
          <Area type="monotone" dataKey="keys" stroke="#8884d8" fill="url(#colorKeys)" />
          <Area type="monotone" dataKey="clicks" stroke="#82ca9d" fill="url(#colorClicks)" />
        </AreaChart>
      </ResponsiveContainer>
    </div>
  );
}

export default App;
