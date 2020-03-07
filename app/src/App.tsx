import 'react-dates/initialize';
import 'react-dates/lib/css/_datepicker.css';
import React, { useState } from 'react';
import { DateRangePicker, isInclusivelyBeforeDay } from 'react-dates';
import moment from 'moment'
import './App.css';
import { ResponsiveContainer, XAxis, YAxis, CartesianGrid, Area, Tooltip, AreaChart } from 'recharts';
import useInterval from '@use-it/interval';
import { useDataFetch } from './useDataFetch'

function App() {
  const [data, loadData] = useDataFetch()
  const [startDate, setStartDate] = useState(moment().subtract(10, 'day'))
  const [endDate, setEndDate] = useState(moment())
  const [focusedInput, setFocusedInput] = useState<'startDate' | 'endDate' | null>(null)
  useInterval(() => {
    const params = new URLSearchParams()
    params.append('startYear', startDate.year().toString())
    params.append('startMonth', (startDate.month() + 1).toString())
    params.append('startDay', startDate.date().toString())
    params.append('endYear', endDate.year().toString())
    params.append('endMonth', (endDate.month() + 1).toString())
    params.append('endDay', endDate.date().toString())
    loadData('http://localhost:5564/getKeys?'+params.toString());
  }, 1000);
  const chartData = data && data.map<any>((tuple: any) => ({
    name: tuple.date,
    keys: tuple.keys || 0,
    clicks: tuple.clicks || 0,
  })) || [];
  return (
    <div className="App">
      <h1 style={{marginBottom: 0}}>Statistics</h1>
      <DateRangePicker
        noBorder
        hideKeyboardShortcutsPanel
        startDate={startDate} 
        startDateId="start-date" 
        endDate={endDate} 
        endDateId="end-date" 
        onDatesChange={({ startDate, endDate }) => {
          if (startDate) setStartDate(startDate)
          if (endDate) setEndDate(endDate)
        }} 
        focusedInput={focusedInput} 
        onFocusChange={focusedInput => setFocusedInput(focusedInput)} 
        isOutsideRange={day => !isInclusivelyBeforeDay(day, moment())}
        displayFormat={() => "DD.MM.YYYY"}
      />
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
