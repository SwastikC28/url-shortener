import React from 'react'
import { Alert } from 'react-bootstrap'

const Error = (props) => {

    let err = props.error.charAt(0).toUpperCase() + props.error.slice(1);

    return (
        <div className="my-4">
            <Alert key={"danger"} variant={"danger"}>
                {err}
            </Alert>
        </div>
    )
}

export default Error