import React from 'react';
import { format } from 'date-fns'; // Import the format function from date-fns


const Timestamps = (props) => {
    console.log(props.timestamps);

    return (
        <div className='my-5'>
            <h3 className='my-3'>Time stamps</h3>

            {props.timestamps.map((timestamp, index) => {
                const timestampMatch = timestamp.match(/^([\d-]+\s[\d:.]+(?:\.\d+)?)/);

                if (timestampMatch) {

                    const parsedDate = new Date(timestampMatch[1]);

                    if (isNaN(parsedDate)) {
                        return <p key={index}>Invalid timestamp format</p>;
                    }


                    const formattedTimestamp = format(parsedDate, 'eeee, MMMM dd, yyyy h:mm:ss a');
                    return <p key={index}>{formattedTimestamp}</p>;
                }

                return <p key={index}>Invalid timestamp format</p>;
            })}
        </div>
    );
}

export default Timestamps;
