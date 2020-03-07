import 'react-dates/initialize'
import 'react-dates/lib/css/_datepicker.css'
import './App.css'
import React, { useState } from 'react'
import moment from 'moment'
import useInterval from '@use-it/interval'
import { AreaChart } from './AreaChart'
import {DatePicker } from './DatePicker'
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
    loadData('http://localhost:5564/getKeys?'+params.toString())
  }, 1000)
  const chartData = data && data.map<any>((tuple: any) => ({
    name: tuple.date,
    keys: tuple.keys || 0,
    clicks: tuple.clicks || 0,
  })) || []
  return (
    <div className="App">
      <h1 style={{marginBottom: 0}}>Statistics</h1>
      <DatePicker 
        startDate={startDate} 
        endDate={endDate} 
        setStartDate={setStartDate}
        setEndDate={setEndDate}
        focusedInput={focusedInput}
        setFocusedInput={setFocusedInput}
      />
      <br />
      <AreaChart data={chartData}/>
    </div>
  )
}

export default App
