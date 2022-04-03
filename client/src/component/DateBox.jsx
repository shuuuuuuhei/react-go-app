import { Fragment, useState } from "react";
import DatePicker from "react-datepicker"

import "react-datepicker/dist/react-datepicker.css";

export default function DateBox() {
    const [startDate, setStartDate] = useState(new Date());
    const handleChange = (date) => {
        setStartDate(date)
    }
    return(
        <DatePicker
            dateFormat="yyyy/MM/dd"
            selected={startDate}
            onChange={handleChange}
        />
    )
}