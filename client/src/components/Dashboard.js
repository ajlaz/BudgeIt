import React, {useState, useEffect} from 'react'
import axios from 'axios'
import Form from './Form'
import TransactionList from './TransactionList'

export default function Dashboard(props) {
    const [transactions, setTransactions] = useState([])

    useEffect(() => {
        axios.get(`http://localhost:8080/transactions/${props.userId}`).then(
            (res) => {
                setTransactions(res.data)
                console.log(transactions)
            })
        }, [])


  return (
    <div>
        <TransactionList transactions={transactions} />
        <Form userId={props.userId} transactions={transactions} setTransactions={setTransactions}  />
    </div>
  )
}
