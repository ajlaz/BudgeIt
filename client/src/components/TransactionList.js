/* eslint-disable react/jsx-key */
import React from 'react'

export default function TransactionList(props) {
    const transactions = props.transactions
  return (
    <div className="transaction-list">
        <h2>Recent Transactions</h2>
        {
        transactions.slice(0, 6).map((t) => {
          return (
              <div className='transaction'>
                  <h3>{t.type}</h3>
                  <p>{`$${t.amount}`}</p>
                  <p>{t.date}</p>
              </div>
          )
      })
        }</div>
  )
}
