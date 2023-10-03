import React from 'react'

export default function TransactionList(props) {
    const transactions = props.transactions
  return (
    <div>{transactions.map((t) => {
        return (
            <div>
                <p>{t.amount}</p>
                <p>{t.type}</p>
                <p>{t.date}</p>
            </div>
        )
    })}</div>
  )
}
