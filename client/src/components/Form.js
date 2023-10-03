import React, {useState} from 'react'
import axios from 'axios'


export default function Form(props) {
    const [amount, setAmount] = useState(0)
    const [type, setType] = useState("")
    const [date, setDate] = useState("")
    const [error, setError] = useState(false)



    const handleAmount = () => {
        return (e) => {
            setAmount(e.target.value)
        }
    }

    const handleType = () => {
        return (e) => {
            setType(e.target.value)
        }
    }

    const handleDate = () => {
        return (e) => {
            setDate(e.target.value)
        }
    }

    const handleSubmit = () => {
        return (e) => {
            e.preventDefault()
            const data = {
                user_id: parseInt(props.userId, 10),
                amount: parseInt(amount, 10),
                type: type,
                date: date
            }

            if (data.amount === 0 || data.type === "" || data.date === "") {
                setError(true)
                return
            }
            setError(false)




            const res = axios({
                method: 'post',
                type: 'application/json',
                url: 'http://localhost:8080/transactions',
                data: data
            })
            res.then((res) => {
                props.setTransactions([...props.transactions, res.data])

            })
            res.catch(() => {
                console.log("error")
            })



        }
    }
    window.history.pushState({}, null, '/')

  return (
    <div class="transaction-form">
        <input name="amount" type="number" onChange={handleAmount()} placeholder="Amount" />
        <input name="type" type="text" onChange={handleType()} placeholder="Type" />
        <input name="date" type="text"  onChange={handleDate()} placeholder="Date (mm/dd/yyyy)" />
        <button onClick={handleSubmit()}>Submit</button>
        <div class="props">{}</div>
        {
            error ? <p>Missing Fields</p> : <></>
        }
    </div>
  )
}
