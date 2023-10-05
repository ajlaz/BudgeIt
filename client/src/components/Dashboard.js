import React, {useState, useEffect} from 'react'
import axios from 'axios'
import Form from './Form'
import TransactionList from './TransactionList'
import {LineChart,CartesianGrid, XAxis, YAxis, Tooltip, Legend, Line } from 'recharts'
import MonthyGoal from './MonthyGoal'


export default function Dashboard(props) {
    const [transactions, setTransactions] = useState([])
    const [filteredTransactions, setFilteredTransactions] = useState([{}])
    const [chartData, setChartData] = useState([{}])
    const [monthTotal, setMonthTotal] = useState(0)

    const goalData = [
        {name: 'This Month', amount: monthTotal},
        {name: 'Goal', amount: 1000}
    ]


    //get all transactions for the user
    useEffect(() => {
        axios.get(`http://localhost:8080/transactions/${props.userId}`).then(
            (res) => {
                setTransactions(res.data)
            })
        }, [])

    //Setup chart data and filtered data
    useEffect(() => {
        // convert hte transactions into a format that the line chart can read

        const lineData = transactions.map((t) => {
            return {
                date: t.date,
                amount: t.amount
            }
        })
        const today = new Date()
        const sevenDaysAgo = new Date(today)
        sevenDaysAgo.setDate(sevenDaysAgo.getDate() - 7)
        const filteredLineData = lineData.filter((d) => {
            const date = new Date(d.date)
            return date > sevenDaysAgo
        }).sort((a, b) => {
            const aDate = new Date(a.date)
            const bDate = new Date(b.date)
            return aDate - bDate
        });
        setChartData(filteredLineData)
        const transactionData = transactions.filter((t) => {
            const date = new Date(t.date)
            return date > sevenDaysAgo
        }
        ).sort((a, b) => {
            const aDate = new Date(a.date)
            const bDate = new Date(b.date)
            return bDate-aDate
        })
        setFilteredTransactions(transactionData)
        setMonthTotal(getMonthTotal())


    }, [transactions])

    //get all transactions for this month and return them as a sum
    function getMonthTotal() {
        const today = new Date()
        const month = today.getMonth()
        const year = today.getFullYear()
        let total = 0
        transactions.forEach((t) => {
            const date = new Date(t.date)
            if (date.getMonth() === month && date.getFullYear() === year) {
                total += t.amount
            }
        })
        return total
    }


  return (
    <div className="transaction-body">
        <div className="transactions-left">
            <div className="transaction-chart">
                <h2>Spending over the last 7 days</h2>
                <LineChart width={500} height={300} data={chartData}>
                    <CartesianGrid strokeDasharray="3 3"/>
                    <XAxis dataKey="date"/>
                    <YAxis/>
                    <Tooltip/>
                    <Legend />
                    <Line type="monotone" dataKey="amount" stroke="#8884d8" activeDot={{r: 8}}/>            
                    </LineChart>
            </div>
        </div>
        <div className="transaction-right">
            <TransactionList transactions={filteredTransactions} />
            <div className="transaction-sub-1">
                <MonthyGoal monthTotal={goalData} />
                <Form userId={props.userId} transactions={filteredTransactions} setTransactions={setFilteredTransactions}  />
            </div>
        </div>
    </div>
  )
}
