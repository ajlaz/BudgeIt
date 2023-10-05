import React from 'react'
import axios from 'axios'
import { BarChart, Bar, Cell, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';


export default function MonthyGoal(props) {
    const monthTotal = props.monthTotal;
  return (
    <div>
        <h2>Monthly Spending Goal</h2>
        <BarChart width={500} height={250} data={monthTotal}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="name" />
        <YAxis dataKey="amount" />
        <Tooltip />
        <Legend />
        <Bar dataKey="amount" fill="#8884d8" />
        </BarChart>
    </div>
  )
}
