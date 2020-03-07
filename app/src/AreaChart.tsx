import React from 'react'
import { 
  ResponsiveContainer, 
  XAxis,
  YAxis,
  CartesianGrid, 
  Area,
  Tooltip,
  AreaChart as RCAreaChart,    
} from 'recharts'

type AreaChartProps = {
  data: any[];
}
export const AreaChart = (props: AreaChartProps) => (
  <ResponsiveContainer aspect={2.5}>
    <RCAreaChart data={props.data}>
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
    </RCAreaChart>
  </ResponsiveContainer>
)